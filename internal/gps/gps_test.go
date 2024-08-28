package gps

import (
	"bytes"
	"desafio-backend/internal/device"
	"desafio-backend/pkg/errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"testing"
	"time"
)

type MockDeviceMain struct {
	device     *device.Device
	deviceErr  errors.Error
	saveDevice *device.Device
	saveErr    errors.Error
}

func TestMain_ParseGps(t *testing.T) {
	var deviceMain device.UseCases
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{
			name:      "null_request",
			input:     "",
			expectErr: true,
		},
		{
			name:      "invalid_request",
			input:     `{"macAddress":"invalidJSON}`,
			expectErr: true,
		},
		{
			name:      "valid_request",
			input:     `{"macAddress":"MAC","timestamp":"2023-03-20T00:00:00Z","latitude":52.5065133,"longitude":13.1445545}`,
			expectErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := &gorm.DB{}
			main := NewMain(db, deviceMain)

			reader := bytes.NewBuffer([]byte(test.input))
			_, err := main.ParseGps(io.NopCloser(reader))

			if test.expectErr {
				assert.NotNil(t, err)
				ae, ok := err.(*errors.AdvancedError)
				if assert.True(t, ok) {
					assert.Equal(t, "Cannot decode data", ae.Title)
				}
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func (d *MockDeviceMain) FindByMacAddress(macAddress string) (*device.Device, errors.Error) {
	return d.device, d.deviceErr
}

func (d *MockDeviceMain) SaveDevice(device device.Device) (device.Device, errors.Error) {
	return *d.saveDevice, d.saveErr
}

func TestMain_SaveGps(t *testing.T) {
	tests := []struct {
		name            string
		in              Request
		setupMockDevice func() *MockDeviceMain
		setupSqlMock    func(mock sqlmock.Sqlmock)
		expectedResp    Response
		expectedErr     errors.Error
	}{
		{
			name: "ValidSaveGps",
			in: Request{
				MacAddress: "99:98:ca:89:be:fe",
				Timestamp:  time.Date(2024, time.August, 28, 9, 26, 46, 547308600, time.Local),
				Latitude:   20.60,
				Longitude:  300.30,
			},
			setupMockDevice: func() *MockDeviceMain {
				return &MockDeviceMain{
					device: &device.Device{},
				}
			},
			setupSqlMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO gps \\(device_id, timestamp, coordinates \\) " +
					"VALUES \\('0', '2024\\-08\\-28T09\\:26\\:46\\-03\\:00', ST_GeogFromText\\('SRID\\=4326;POINT\\(20\\.600000 300\\.300000\\)'\\)\\) RETURNING id;").
					WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow(1))
				mock.ExpectQuery("^SELECT d\\.mac_address, g\\.timestamp, ST_AsGeoJSON\\(g\\.coordinates\\) from public\\.device d, public\\.gps g WHERE g\\.id \\= \\$1 and g\\.device_id \\= d\\.id").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"MacAddress", "Timestamp", "Coordinates"}).
						AddRow("99:98:ca:89:be:fe", time.Date(2024, time.August, 28, 9, 26, 46, 547308600, time.Local), "{\\\"type\\\":\\\"Point\\\",\\\"coordinates\\\":[20.60,300.30]}"))
			},
			expectedResp: Response{
				MacAddress:  "99:98:ca:89:be:fe",
				Timestamp:   time.Date(2024, time.August, 28, 9, 26, 46, 547308600, time.Local),
				Coordinates: "{\\\"type\\\":\\\"Point\\\",\\\"coordinates\\\":[20.60,300.30]}",
			},
			expectedErr: nil,
		},
		{
			name: "InvalidSaveGps_Timestamp",
			in: Request{
				MacAddress: "99:98:ca:89:be:fe",
				Timestamp:  time.Date(2022, time.August, 28, 9, 26, 46, 547308600, time.Local),
				Latitude:   20.60,
				Longitude:  300.30,
			},
			setupMockDevice: func() *MockDeviceMain {
				return &MockDeviceMain{
					device:    &device.Device{},
					deviceErr: errors.NewError("Invalid", "Invalid Timestamp"),
				}
			},
			setupSqlMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO gps \\(device_id, timestamp, coordinates \\) " +
					"VALUES \\('0', '2024\\-08\\-28T09\\:26\\:46\\-03\\:00', ST_GeogFromText\\('SRID\\=4326;POINT\\(20\\.600000 300\\.300000\\)'\\)\\) RETURNING id;").
					WillReturnError(fmt.Errorf("some error"))
			},
			expectedErr: errors.NewError("Invalid", "Invalid Timestamp"),
		},
		{
			name: "InvalidSaveGps_Coordinates",
			in: Request{
				MacAddress: "99:98:ca:89:be:fe",
				Timestamp:  time.Date(2024, time.August, 28, 9, 26, 46, 547308600, time.Local),
				Latitude:   91.60,
				Longitude:  300.30,
			},
			setupMockDevice: func() *MockDeviceMain {
				return &MockDeviceMain{
					device:    &device.Device{},
					deviceErr: errors.NewError("Invalid", "Invalid Coordinates"),
				}
			},
			setupSqlMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO gps \\(device_id, timestamp, coordinates \\) " +
					"VALUES \\('0', '2024\\-08\\-28T09\\:26\\:46\\-03\\:00', ST_GeogFromText\\('SRID\\=4326;POINT\\(20\\.600000 300\\.300000\\)'\\)\\) RETURNING id;").
					WillReturnError(fmt.Errorf("some error"))
			},
			expectedErr: errors.NewError("Invalid", "Invalid Coordinates"),
		},
		{
			name: "InvalidSaveGps_MacAddress",
			in: Request{
				MacAddress: "invalid mac",
				Timestamp:  time.Date(2024, time.August, 28, 9, 26, 46, 547308600, time.Local),
				Latitude:   20.60,
				Longitude:  300.30,
			},
			setupMockDevice: func() *MockDeviceMain {
				return &MockDeviceMain{
					device:    &device.Device{},
					deviceErr: errors.NewError("Invalid", "Invalid MacAddress"),
				}
			},
			setupSqlMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO gps \\(device_id, timestamp, coordinates \\) " +
					"VALUES \\('0', '2024\\-08\\-28T09\\:26\\:46\\-03\\:00', ST_GeogFromText\\('SRID\\=4326;POINT\\(20\\.600000 300\\.300000\\)'\\)\\) RETURNING id;").
					WillReturnError(fmt.Errorf("some error"))
			},
			expectedErr: errors.NewError("Invalid", "Invalid MacAddress"),
		},
		{
			name: "InvalidSaveGps_UnavailableDevice",
			in: Request{
				MacAddress: "99:98:ca:89:be:fe",
				Timestamp:  time.Date(2024, time.August, 28, 9, 26, 46, 547308600, time.Local),
				Latitude:   20.60,
				Longitude:  300.30,
			},
			setupMockDevice: func() *MockDeviceMain {
				return &MockDeviceMain{
					deviceErr: errors.NewError("Invalid", "Device Not Found"),
				}
			},
			setupSqlMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO gps \\(device_id, timestamp, coordinates \\) " +
					"VALUES \\('0', '2024\\-08\\-28T09\\:26\\:46\\-03\\:00', ST_GeogFromText\\('SRID\\=4326;POINT\\(20\\.600000 300\\.300000\\)'\\)\\) RETURNING id;").
					WillReturnError(fmt.Errorf("some error"))
			},
			expectedErr: errors.NewError("Invalid", "Device Not Found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, _ := sqlmock.New()
			db, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.setupSqlMock(mock)

			main := Main{
				db:         db,
				deviceMain: tt.setupMockDevice(),
			}

			resp, err := main.SaveGps(tt.in)

			if tt.expectedErr != nil {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}

func TestMain_ValidateGps(t *testing.T) {
	mainInstance := Main{}

	tests := []struct {
		name string
		gps  Request
		want int
	}{
		{
			name: "EmptyValues",
			gps:  Request{},
			want: 5,
		},
		{
			name: "ValidValues",
			gps: Request{
				MacAddress: "01:23:45:67:89:AB",
				Timestamp:  time.Now(),
				Latitude:   25.43992,
				Longitude:  49.2143992,
			},
			want: 0,
		},
		{
			name: "InvalidMacAddress",
			gps: Request{
				MacAddress: "01:23:45:67:G9:AB",
				Timestamp:  time.Now(),
				Latitude:   25.43992,
				Longitude:  49.2143992,
			},
			want: 1,
		},
		{
			name: "ZeroLatitude",
			gps: Request{
				MacAddress: "01:23:45:67:89:AB",
				Timestamp:  time.Now(),
				Latitude:   0.0,
				Longitude:  49.2143992,
			},
			want: 1,
		},
		{
			name: "ZeroLongitude",
			gps: Request{
				MacAddress: "01:23:45:67:89:AB",
				Timestamp:  time.Now(),
				Latitude:   25.43992,
				Longitude:  0.0,
			},
			want: 1,
		},
		{
			name: "ZeroTimestamp",
			gps: Request{
				MacAddress: "01:23:45:67:89:AB",
				Timestamp:  time.Time{},
				Latitude:   25.43992,
				Longitude:  49.2143992,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errList := mainInstance.ValidateGps(tt.gps)
			if errCount := len(errList.GetErrors()); errCount != tt.want {
				t.Errorf("ValidateGps() Error Count = %v, want %v", errCount, tt.want)
			}
		})
	}
}

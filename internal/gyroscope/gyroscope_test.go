package gyroscope

import (
	"bytes"
	"desafio-backend/internal/device"
	"desafio-backend/pkg/errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockDeviceMain struct {
	device     *device.Device
	deviceErr  errors.Error
	saveDevice *device.Device
	saveErr    errors.Error
}

func TestMain_ParseGyroscope(t *testing.T) {
	var db *gorm.DB

	deviceMain := device.NewMain(db)
	main := NewMain(db, deviceMain)

	tests := []struct {
		name        string
		input       []byte
		wantReq     Request
		wantErrType *errors.AdvancedError
	}{
		{
			name:        "SuccessfulParse",
			input:       []byte(`{"macAddress":"aa:bb:cc:dd:ee:ff","timestamp":"2022-08-05T20:19:49Z","xAxis":0.12,"yAxis":0.43,"zAxis":0.57}`),
			wantReq:     Request{MacAddress: "aa:bb:cc:dd:ee:ff", Timestamp: time.Date(2022, 8, 5, 20, 19, 49, 0, time.UTC), XAxis: 0.12, YAxis: 0.43, ZAxis: 0.57},
			wantErrType: nil,
		},
		{
			name:        "InvalidJson",
			input:       []byte(`{"macAddress": aa:bb:cc:dd:ee:ff,"timestamp":"2022-08-05T20:19:49Z","yAxis":0.43,"zAxis":0.57}`),
			wantReq:     Request{},
			wantErrType: &errors.AdvancedError{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotReq, gotErr := main.ParseGyroscope(ioutil.NopCloser(bytes.NewReader(test.input)))

			if gotErr != nil {
				assert.IsType(t, test.wantErrType, gotErr)
			}
			assert.Equal(t, test.wantReq, gotReq)
		})
	}
}

func (d *MockDeviceMain) FindByMacAddress(macAddress string) (*device.Device, errors.Error) {
	return d.device, d.deviceErr
}

func (d *MockDeviceMain) SaveDevice(device device.Device) (device.Device, errors.Error) {
	return *d.saveDevice, d.saveErr
}

func TestMain_SaveGyroscope(t *testing.T) {
	tests := []struct {
		name            string
		in              Request
		setupMockDevice func() *MockDeviceMain
		setupSqlMock    func(mock sqlmock.Sqlmock)
		expectedResp    Response
		expectedErr     errors.Error
	}{
		{
			name: "ValidSaveGyroscope",
			in: Request{
				MacAddress: "99:98:ca:89:be:fe",
				Timestamp:  time.Date(2024, time.August, 28, 9, 26, 46, 547308600, time.Local),
				XAxis:      20.60,
				YAxis:      300.30,
				ZAxis:      100.00,
			},
			setupMockDevice: func() *MockDeviceMain {
				return &MockDeviceMain{
					device: &device.Device{},
				}
			},
			setupSqlMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO gyroscope \\(device_id, timestamp, x_axis, y_axis, z_axis \\) VALUES \\('0', '2024\\-08\\-28T09\\:26\\:46\\-03\\:00', '20\\.600000', '300\\.300000', '100\\.000000'\\) RETURNING id;").
					WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow(1))
				mock.ExpectQuery("^SELECT d\\.mac_address, g\\.timestamp, g\\.x_axis, g\\.y_axis, g\\.z_axis from public\\.device d, public\\.gyroscope g WHERE g\\.id \\= \\$1 and g\\.device_id \\= d\\.id").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"MacAddress", "Timestamp", "XAxis", "YAxis", "ZAxis"}).
						AddRow("99:98:ca:89:be:fe", time.Date(2024, time.August, 28, 9, 26, 46, 547308600, time.Local), 20.60, 300.30, 100.00))
			},
			expectedResp: Response{
				MacAddress: "99:98:ca:89:be:fe",
				Timestamp:  time.Date(2024, time.August, 28, 9, 26, 46, 547308600, time.Local),
				XAxis:      20.60,
				YAxis:      300.30,
				ZAxis:      100.00,
			},
			expectedErr: nil,
		},
		{
			name: "InvalidSaveGyroscope",
			in: Request{
				MacAddress: "99:98:ca:89:be:fe",
				Timestamp:  time.Time{},
				XAxis:      20.60,
				YAxis:      300.30,
				ZAxis:      100.00,
			},
			setupMockDevice: func() *MockDeviceMain {
				return &MockDeviceMain{
					device:    &device.Device{},
					deviceErr: errors.NewError("Invalid", "Invalid Timestamp"),
				}
			},
			setupSqlMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO gyroscope \\(device_id, timestamp, x_axis, y_axis, z_axis \\) VALUES \\('0', null, '20\\.600000', '300\\.300000', '100\\.000000'\\) RETURNING id;").
					WillReturnError(fmt.Errorf("some error"))
			},
			expectedErr: errors.NewError("Invalid", "Invalid Timestamp"),
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

			resp, err := main.SaveGyroscope(tt.in)

			if tt.expectedErr != nil {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}

func TestMain_ValidateGyroscope(t *testing.T) {
	mainInstance := Main{}

	tests := []struct {
		name string
		gps  Request
		want int
	}{
		{
			name: "EmptyValues",
			gps:  Request{},
			want: 4,
		},
		{
			name: "ValidValues",
			gps: Request{
				MacAddress: "01:23:45:67:89:AB",
				Timestamp:  time.Now(),
				XAxis:      20.60,
				YAxis:      300.30,
				ZAxis:      100.00,
			},
			want: 0,
		},
		{
			name: "InvalidMacAddress",
			gps: Request{
				MacAddress: "MAC",
				Timestamp:  time.Now(),
				XAxis:      20.60,
				YAxis:      300.30,
				ZAxis:      100.00,
			},
			want: 1,
		},
		{
			name: "ZeroTimestamp",
			gps: Request{
				MacAddress: "01:23:45:67:89:AB",
				Timestamp:  time.Time{},
				XAxis:      20.60,
				YAxis:      300.30,
				ZAxis:      100.00,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errList := mainInstance.ValidateGyroscope(tt.gps)
			if errCount := len(errList.GetErrors()); errCount != tt.want {
				t.Errorf("ValidateGyroscope() Error Count = %v, want %v", errCount, tt.want)
			}
		})
	}
}

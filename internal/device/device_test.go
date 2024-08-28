package device

import (
	"desafio-backend/util"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestFindByMacAddress(t *testing.T) {

	type want struct {
		device *Device
		err    bool
	}

	tests := []struct {
		name string
		mac  string
		want want
	}{
		{
			name: "Found a device",
			mac:  "99:98:ca:89:be:fe",
			want: want{
				device: &Device{
					BaseModel: util.BaseModel{
						ID:        1,
						Timestamp: time.Date(2024, time.August, 28, 9, 26, 46, 547308600, time.Local),
					},
					MacAddress: "99:98:ca:89:be:fe",
				},
				err: false,
			},
		},
		{
			name: "Not found a device",
			mac:  "99:98:ca:89:be:ce",
			want: want{
				device: nil,
				err:    false,
			},
		},
		{
			name: "Error",
			mac:  "MAC_3",
			want: want{
				device: nil,
				err:    true,
			},
		},
	}

	mockDB, mock, _ := sqlmock.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			rows := sqlmock.NewRows([]string{"ID", "MacAddress", "Timestamp"})

			if tt.name == "Not found a device" {
				mock.ExpectQuery("^SELECT id, mac_address, created_at as \"timestamp\" from public.device where mac_address = \\$1$").
					WithArgs(tt.mac).WillReturnRows(rows)
			} else if tt.name == "Found a device" {
				rows.AddRow(1, "99:98:ca:89:be:fe", time.Date(2024, time.August, 28, 9, 26, 46, 547308600, time.Local))
				mock.ExpectQuery("^SELECT id, mac_address, created_at as \"timestamp\" from public.device where mac_address = \\$1$").
					WithArgs(tt.mac).WillReturnRows(rows)
			} else {
				mock.ExpectQuery("^SELECT id, mac_address, created_at as \"timestamp\" from public.device where mac_address = \\$1$").
					WithArgs(tt.mac).WillReturnError(fmt.Errorf("some error"))
			}

			main := &Main{db: db}
			dev, err := main.FindByMacAddress(tt.mac)
			equals := assert.Equal(t, tt.want.device, dev)

			if !equals {
				t.Errorf("FindByMacAddress() = %v, want = %v", dev, tt.want.device)
			}

			if (err != nil) != tt.want.err {
				t.Errorf("FindByMacAddress() error = %v, wantErr %v", err, tt.want.err)
				return
			}
		})
	}
}

func TestSaveDevice(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	defer db.Close()

	main := NewMain(gormDB)

	device := Device{
		BaseModel: util.BaseModel{
			ID:        1,
			Timestamp: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		MacAddress: "00:0a:95:9d:68:16",
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO device \\(created_at, mac_address \\) VALUES \\('0001\\-01\\-01T00\\:00\\:00Z', '00\\:0a\\:95\\:9d\\:68\\:16'\\);").
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery("^SELECT id, mac_address, created_at as \"timestamp\" from public.device where mac_address = \\$1$").
			WithArgs(device.MacAddress).
			WillReturnRows(sqlmock.NewRows([]string{"ID", "MacAddress", "Timestamp"}).
				AddRow(1, device.MacAddress, device.Timestamp))

		res, err := main.SaveDevice(device)
		assert.Nil(t, err)
		assert.Equal(t, device, res)
	})

	t.Run("error saving", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO device \\(created_at, mac_address \\) VALUES \\('0001\\-01\\-01T00\\:00\\:00Z', '00\\:0a\\:95\\:9d\\:68\\:16'\\);").
			WillReturnError(fmt.Errorf("some error"))

		_, err := main.SaveDevice(device)
		assert.NotNil(t, err)
	})

	t.Run("error scanning", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO device \\(created_at, mac_address \\) VALUES \\('0001\\-01\\-01T00\\:00\\:00Z', '00\\:0a\\:95\\:9d\\:68\\:16'\\);").
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery("^SELECT id, mac_address, created_at as \"timestamp\" from public.device where mac_address = \\$1$").
			WithArgs(device.MacAddress).
			WillReturnRows(sqlmock.NewRows([]string{"ID", "MacAddress", "Timestamp"}).
				AddRow(1, device.MacAddress, "invalid_data"))

		_, err := main.SaveDevice(device)
		assert.NotNil(t, err)
	})
}

package database

import (
	"database/sql"
	"fmt"
	"github.com/HaroldoFV/desafio/configs"
	"github.com/HaroldoFV/desafio/internal/domain/entity"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type GPSRepositoryTestSuite struct {
	suite.Suite
	DB         *sql.DB
	Repository *GPSRepository
	Config     *configs.Conf
}

func (suite *GPSRepositoryTestSuite) SetupSuite() {
	dir, _ := os.Getwd()
	fmt.Println("currently dir:", dir)

	config, err := configs.LoadConfig(dir)
	if err != nil {
		rootDir := filepath.Join(dir, "..", "..")
		config, err = configs.LoadConfig(rootDir)
		if err != nil {
			fmt.Println("Error loading configurations:", err)
			panic(err)
		}
	}
	suite.Config = config
	fmt.Printf("Configurations loaded: %+v\n", config)

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.TESTDBHost,
		config.TESTDBPort,
		config.TESTDBUser,
		config.TESTDBPassword,
		config.TESTDBName,
	)

	db, err := sql.Open(config.DBDriver, connectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	suite.DB = db
	suite.Repository = NewGPSRepository(db)

	_, err = suite.DB.Exec(`
        CREATE TABLE IF NOT EXISTS gps (
            id UUID PRIMARY KEY,
            latitude DOUBLE PRECISION NOT NULL,
            longitude DOUBLE PRECISION NOT NULL,
            mac_address VARCHAR(17) NOT NULL,
            timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}

func (suite *GPSRepositoryTestSuite) TearDownSuite() {
	_, err := suite.DB.Exec("DROP TABLE IF EXISTS gps")
	if err != nil {
		log.Fatal(err)
	}
	suite.DB.Close()
}

func (suite *GPSRepositoryTestSuite) SetupTest() {
	_, err := suite.DB.Exec("DELETE FROM gps")
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *GPSRepositoryTestSuite) TestCreate() {
	latitude := 40.7128
	longitude := -74.0060
	macAddress := "00:11:22:33:44:55"

	gps, err := entity.NewGPS(latitude, longitude, macAddress)
	assert.NoError(suite.T(), err)

	err = suite.Repository.Create(gps)

	assert.NoError(suite.T(), err)

	var count int
	err = suite.DB.QueryRow("SELECT COUNT(*) FROM gps WHERE id = $1", gps.GetID()).Scan(&count)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, count)

	var storedID, storedMACAddress string
	var storedLatitude, storedLongitude float64
	var storedTimestamp time.Time

	err = suite.DB.QueryRow("SELECT id, latitude, longitude, mac_address, timestamp FROM gps WHERE id = $1", gps.GetID()).Scan(
		&storedID,
		&storedLatitude,
		&storedLongitude,
		&storedMACAddress,
		&storedTimestamp,
	)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), gps.GetID(), storedID)
	assert.Equal(suite.T(), gps.GetLatitude(), storedLatitude)
	assert.Equal(suite.T(), gps.GetLongitude(), storedLongitude)
	assert.Equal(suite.T(), gps.GetMACAddress(), storedMACAddress)
	assert.WithinDuration(suite.T(), gps.GetTimestamp(), storedTimestamp, time.Second)
}

func TestGPSRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(GPSRepositoryTestSuite))
}

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

type PhotoRepositoryTestSuite struct {
	suite.Suite
	DB         *sql.DB
	Repository *PhotoRepository
	Config     *configs.Conf
}

func (suite *PhotoRepositoryTestSuite) SetupSuite() {
	dir, _ := os.Getwd()

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
	suite.Repository = NewPhotoRepository(db)

	_, err = suite.DB.Exec(`
        CREATE TABLE IF NOT EXISTS photos (
            id UUID PRIMARY KEY,
            file_path TEXT NOT NULL,
            mac_address TEXT NOT NULL,
            timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}

func (suite *PhotoRepositoryTestSuite) TearDownSuite() {
	_, err := suite.DB.Exec("DROP TABLE IF EXISTS photos")
	if err != nil {
		log.Fatal(err)
	}
	suite.DB.Close()
}

func (suite *PhotoRepositoryTestSuite) SetupTest() {
	_, err := suite.DB.Exec("DELETE FROM photos")
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *PhotoRepositoryTestSuite) TestCreate() {
	filePath := "/path/to/image.jpg"
	macAddress := "00:11:22:33:44:55"

	photo, err := entity.NewPhoto(filePath, macAddress)
	assert.NoError(suite.T(), err)

	err = suite.Repository.Create(photo)
	assert.NoError(suite.T(), err)

	var count int
	err = suite.DB.QueryRow("SELECT COUNT(*) FROM photos WHERE id = $1", photo.GetID()).Scan(&count)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, count)

	var storedID, storedFilePath, storedMACAddress string
	var storedTimestamp time.Time

	err = suite.DB.QueryRow("SELECT id, file_path, mac_address, timestamp FROM photos WHERE id = $1", photo.GetID()).Scan(
		&storedID,
		&storedFilePath,
		&storedMACAddress,
		&storedTimestamp,
	)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), photo.GetID(), storedID)
	assert.Equal(suite.T(), photo.GetFilePath(), storedFilePath)
	assert.Equal(suite.T(), photo.GetMACAddress(), storedMACAddress)
	assert.WithinDuration(suite.T(), photo.GetTimestamp(), storedTimestamp, time.Second)
}

func TestPhotoRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PhotoRepositoryTestSuite))
}

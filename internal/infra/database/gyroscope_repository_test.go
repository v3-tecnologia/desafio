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

type GyroscopeRepositoryTestSuite struct {
	suite.Suite
	DB         *sql.DB
	Repository *GyroscopeRepository
	Config     *configs.Conf
}

func (suite *GyroscopeRepositoryTestSuite) SetupSuite() {
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
	suite.Repository = NewGyroscopeRepository(db)

	_, err = suite.DB.Exec(`
		CREATE TABLE IF NOT EXISTS gyroscopes (
			id UUID PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			model VARCHAR(50) NOT NULL,
			mac_address VARCHAR(17) NOT NULL,
			x DOUBLE PRECISION DEFAULT 0,
			y DOUBLE PRECISION DEFAULT 0,
			z DOUBLE PRECISION DEFAULT 0,
			timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
		`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}

func (suite *GyroscopeRepositoryTestSuite) TearDownSuite() {
	_, err := suite.DB.Exec("DROP TABLE IF EXISTS gyroscopes")
	if err != nil {
		log.Fatal(err)
	}
	suite.DB.Close()
}

func (suite *GyroscopeRepositoryTestSuite) SetupTest() {
	_, err := suite.DB.Exec("DELETE FROM gyroscopes")
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *GyroscopeRepositoryTestSuite) TestCreate() {
	name := "Test Gyroscope"
	model := "Test Model"
	macAddress := "00:11:22:33:44:55"

	gyroscope, err := entity.NewGyroscope(name, model, 1.0, 2.0, 3.0, macAddress)
	assert.NoError(suite.T(), err)

	err = suite.Repository.Create(gyroscope)

	assert.NoError(suite.T(), err)

	var count int
	err = suite.DB.QueryRow("SELECT COUNT(*) FROM gyroscopes WHERE id = $1", gyroscope.GetID()).Scan(&count)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, count)

	var storedID, storedName, storedModel, storedMACAddress string
	var storedX, storedY, storedZ float64
	var storedTimestamp time.Time

	err = suite.DB.QueryRow("SELECT id, name, model, mac_address, x, y, z, timestamp FROM gyroscopes WHERE id = $1", gyroscope.GetID()).Scan(
		&storedID,
		&storedName,
		&storedModel,
		&storedMACAddress,
		&storedX,
		&storedY,
		&storedZ,
		&storedTimestamp,
	)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), gyroscope.GetID(), storedID)
	assert.Equal(suite.T(), gyroscope.GetName(), storedName)
	assert.Equal(suite.T(), gyroscope.GetModel(), storedModel)
	assert.Equal(suite.T(), gyroscope.GetMACAddress(), storedMACAddress)
	assert.Equal(suite.T(), gyroscope.GetX(), storedX)
	assert.Equal(suite.T(), gyroscope.GetY(), storedY)
	assert.Equal(suite.T(), gyroscope.GetZ(), storedZ)
	assert.WithinDuration(suite.T(), gyroscope.GetTimestamp(), storedTimestamp, time.Second)
}

func TestGyroscopeRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(GyroscopeRepositoryTestSuite))
}

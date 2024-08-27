package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

type Conf struct {
	DBDriver                   string `mapstructure:"DB_DRIVER"`
	DBHost                     string `mapstructure:"DB_HOST"`
	DBPort                     string `mapstructure:"DB_PORT"`
	DBUser                     string `mapstructure:"DB_USER"`
	DBPassword                 string `mapstructure:"DB_PASSWORD"`
	DBName                     string `mapstructure:"DB_NAME"`
	TESTDBHost                 string `mapstructure:"TEST_DB_HOST"`
	TESTDBPort                 string `mapstructure:"TEST_DB_PORT"`
	TESTDBUser                 string `mapstructure:"TEST_DB_USER"`
	TESTDBPassword             string `mapstructure:"TEST_DB_PASSWORD"`
	TESTDBName                 string `mapstructure:"TEST_DB_NAME"`
	WebServerPort              string `mapstructure:"WEB_SERVER_PORT"`
	PhotoStoragePath           string `mapstructure:"PHOTO_STORAGE_PATH"`
	AWSAccessKeyID             string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey         string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AWSRegion                  string `mapstructure:"AWS_REGION"`
	AWSRekognitionCollectionID string `mapstructure:"AWS_REKOGNITION_COLLECTION_ID"`
}

func LoadConfig(path string) (*Conf, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AddConfigPath(filepath.Join(path, ".."))
	viper.AddConfigPath(filepath.Join(path, "..", ".."))
	viper.AddConfigPath("/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo de configuração: %w", err)
	}

	var config Conf
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar configurações: %w", err)
	}

	return &config, nil
}

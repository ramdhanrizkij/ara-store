package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Env  string
		Port string
	}

	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SSLMode  string
	}

	JWT struct {
		Secret string
		Expiry int
	}

	AWS struct {
		AccessKey string
		SecretKey string
		Region    string
		Bucket    string
	}

	Redis struct {
		Host     string
		Port     string
		Password string
	}

	SMTP struct {
		Host     string
		Port     int
		User     string
		Password string
		From     string
	}
}

func Load() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No .env file found, using environment variable")
	}

	cfg := &Config{}

	cfg.App.Env = viper.GetString("APP_ENV")
	cfg.App.Port = viper.GetString("APP_PORT")

	cfg.Database.Host = viper.GetString("DB_HOST")
	cfg.Database.Port = viper.GetString("DB_PORT")
	cfg.Database.User = viper.GetString("DB_USER")
	cfg.Database.Password = viper.GetString("DB_PASSWORD")
	cfg.Database.Name = viper.GetString("DB_NAME")
	cfg.Database.SSLMode = viper.GetString("DB_SSL_MODE")

	cfg.JWT.Secret = viper.GetString("JWT_SECRET")
	cfg.JWT.Expiry = viper.GetInt("JWT_EXPIRY")

	cfg.AWS.AccessKey = viper.GetString("AWS_ACCESS_KEY_ID")
	cfg.AWS.SecretKey = viper.GetString("AWS_SECRET_ACCESS_KEY")
	cfg.AWS.Region = viper.GetString("AWS_REGION")
	cfg.AWS.Bucket = viper.GetString("AWS_S3_BUCKET")

	cfg.Redis.Host = viper.GetString("REDIS_HOST")
	cfg.Redis.Port = viper.GetString("REDIS_PORT")
	cfg.Redis.Password = viper.GetString("REDIS_PASSWORD")

	cfg.SMTP.Host = viper.GetString("SMTP_HOST")
	cfg.SMTP.Port = viper.GetInt("SMTP_PORT")
	cfg.SMTP.User = viper.GetString("SMTP_USER")
	cfg.SMTP.Password = viper.GetString("SMTP_PASSWORD")
	cfg.SMTP.From = viper.GetString("SMTP_FROM")

	return cfg
}

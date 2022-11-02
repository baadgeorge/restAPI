package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func NewDB() (*pgx.Conn, error) {
	config, err := configDB()
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}
	parsedConfig, err := pgx.ParseConnectionString(config)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}
	DB, err := pgx.Connect(parsedConfig)
	if err != nil {
		logrus.Fatal("db connection failed")
		return nil, err
	}

	if err = DB.Ping(context.Background()); err != nil {
		logrus.Fatalf("can't ping db: %s", err)
		return nil, err
	}
	return DB, nil
}

func configDB() (string, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("can't read config: %s", err.Error())
		return "", err
	}

	if err = godotenv.Load(); err != nil {
		logrus.Fatalf("can't load env: %s", err.Error())
		return "", err
	}

	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	dbname := viper.GetString("database.dbname")
	sslMode := viper.GetString("database.sslmode")
	password := os.Getenv("DB_PASSWORD")

	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbname, password, sslMode)

	return config, nil
}

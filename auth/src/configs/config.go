package configs

import (
	"github.com/danborodin/go-logd"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// add log pointer
type Config struct {
	l     *logd.Logger
	Mongo struct {
		Uri      string
		Database string
	}
	//
	Pepper string
}

var Conf *Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file ", err)
	}
}

func New(l *logd.Logger) *Config {
	return &Config{
		l: l,
	}
}

func InitConfig(conf *Config) {
	loadMongoConfig(conf)
	loadPepper(conf)
}

func loadMongoConfig(conf *Config) {
	conf.Mongo.Uri = os.Getenv("MONGO_URI")
	conf.Mongo.Database = os.Getenv("MONGO_DB")
}

func loadPepper(conf *Config) {
	conf.Pepper = os.Getenv("PEPPER")
}

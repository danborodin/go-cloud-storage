package configs

import (
	"log"
	"os"

	"github.com/danborodin/go-logd"
	"github.com/joho/godotenv"
)

// add log pointer
type Config struct {
	l     *logd.Logger
	Host  string
	Mongo struct {
		Uri      string
		Database string
	}
	Gmail struct {
		Host     string
		Port     string
		Pwd      string
		Username string
	}
	//
	Pepper       string
	EmailVerTime string
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
	loadGmail(conf)
	//
	conf.EmailVerTime = os.Getenv("EMAIL_VERIFICATION_TIME")
}

func loadMongoConfig(conf *Config) {
	conf.Mongo.Uri = os.Getenv("MONGO_URI")
	conf.Mongo.Database = os.Getenv("MONGO_DB")
}

func loadPepper(conf *Config) {
	conf.Pepper = os.Getenv("PEPPER")
}

func loadGmail(conf *Config) {
	conf.Gmail.Host = os.Getenv("GMAIL_SMTP_HOST")
	conf.Gmail.Port = os.Getenv("GMAIL_SMTP_PORT")
	conf.Gmail.Pwd = os.Getenv("GMAIL_SMTP_PASSWORD")
	conf.Gmail.Username = os.Getenv("GMAIL_SMTP_USERNAME")
}

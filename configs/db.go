package configs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	MysqlHost                  string `envconfig:"MYSQL_HOST" default:"localhost"`
	MysqlPort                  int    `envconfig:"MYSQL_PORT" default:"3306"`
	MysqlDBName                string `envconfig:"MYSQL_DB_NAME" default:"employee_db"`
	MysqlUsername              string `envconfig:"MYSQL_USERNAME" default:""`
	MysqlPassword              string `envconfig:"MYSQL_PASSWORD" default:""`
	MysqlLogMode               int    `envconfig:"MYSQL_LOG_MODE" default:"1"`
	MysqlParseTime             bool   `envconfig:"MYSQL_PARSE_TIME" default:"true"`
	MysqlCharset               string `envconfig:"MYSQL_CHARSET" default:"utf8mb4"`
	MysqlLoc                   string `envconfig:"MYSQL_LOC" default:"Local"`
	MysqlMaxLifetimeConnection int    `envconfig:"MYSQL_MAX_LIFETIME_CONNECTION" default:"10"`
	MysqlMaxOpenConnection     int    `envconfig:"MYSQL_MAX_OPEN_CONNECTION" default:"50"`
	MysqlMaxIdleConnection     int    `envconfig:"MYSQL_MAX_IDLE_CONNECTION" default:"10"`
}

func (cfg Database) ConnectDB() *gorm.DB {
	// construct connection string
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%+v&loc=%s",
		cfg.MysqlUsername,
		cfg.MysqlPassword,
		cfg.MysqlHost,
		cfg.MysqlPort,
		cfg.MysqlDBName,
		cfg.MysqlCharset,
		cfg.MysqlParseTime,
		cfg.MysqlLoc)

	// open mysql connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(cfg.MysqlLogMode)),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// set configuration pooling connection
	mysqlDb, _ := db.DB()
	mysqlDb.SetMaxOpenConns(cfg.MysqlMaxOpenConnection)
	mysqlDb.SetConnMaxLifetime(time.Duration(cfg.MysqlMaxLifetimeConnection) * time.Minute)
	mysqlDb.SetMaxIdleConns(cfg.MysqlMaxIdleConnection)

	return db
}

func InitDatabaseConfig() Database {
	var cfg Database
	err := godotenv.Overload()
	if err != nil {
		log.Println(err)
	}
	envconfig.MustProcess("", &cfg)
	return cfg
}

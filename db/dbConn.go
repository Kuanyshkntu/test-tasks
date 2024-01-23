package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

type DatabaseParams struct {
	DbUser                 string
	DbPassword             string
	DbHost                 string
	DbPort                 string
	DbName                 string
	MaxPoolConnections     int
	MaxIdlePoolConnections int
	ConnectionTimeout      int
}

type Db struct {
	DbHandler         *gorm.DB
	DbUserName        string
	ConnectionTimeout time.Duration
}

func ConnectDB(pars DatabaseParams) (*Db, error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		pars.DbHost,
		pars.DbPort,
		pars.DbUser,
		pars.DbName,
		pars.DbPassword,
	)

	dbConn, err := gorm.Open(postgres.Open(DBURI), &gorm.Config{Logger: newLogger})

	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxOpenConns(pars.MaxPoolConnections)
	sqlDB.SetMaxIdleConns(pars.MaxIdlePoolConnections)
	//время жизни пула должна быть равно или меньше таймаута запроса
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(pars.ConnectionTimeout))

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &Db{DbHandler: dbConn, DbUserName: pars.DbUser,
		ConnectionTimeout: time.Duration(pars.ConnectionTimeout)}, nil

}

type name struct {
}

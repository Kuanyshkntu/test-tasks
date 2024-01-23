package env

import (
	"fmt"
	"os"
	"strconv"
	"test_tasks/db"
)

var (
	EnvVarNotSpecified = "Переменная окружения «%v» не указана.\n"
)

type DatabaseEnvKeys struct {
	DbUser        string
	DbPassword    string
	DbHost        string
	DbPool        string
	DbPoolIdle    string
	DbConnTimeout string
	DbName        string
	DbPort        string
}

func GetDbEnvVals(dbEnv DatabaseEnvKeys) (*db.DatabaseParams, error) {

	var dbParams db.DatabaseParams

	dbParams.DbUser = os.Getenv(dbEnv.DbUser)
	if dbParams.DbUser == "" {
		return nil, fmt.Errorf("Переменная окружения «%v» не указана.\n", dbEnv.DbUser)
	}

	dbParams.DbPassword = os.Getenv(dbEnv.DbPassword)
	if dbParams.DbPassword == "" {
		return nil, fmt.Errorf("Переменная окружения «%v» не указана.\n", dbEnv.DbPassword)
	}

	dbParams.DbHost = os.Getenv(dbEnv.DbHost)
	if dbParams.DbHost == "" {
		return nil, fmt.Errorf("Переменная окружения «%v» не указана.\n", dbEnv.DbHost)
	}
	dbParams.DbPort = os.Getenv(dbEnv.DbPort)
	if dbParams.DbPort == "" {
		return nil, fmt.Errorf("Переменная окружения «%v» не указана.\n", dbEnv.DbPort)
	}

	dbParams.DbName = os.Getenv(dbEnv.DbName)
	if dbParams.DbPort == "" {
		return nil, fmt.Errorf("Переменная окружения «%v» не указана.\n", dbEnv.DbName)
	}

	if i, b := strconv.Atoi(os.Getenv(dbEnv.DbPool)); b == nil {
		dbParams.MaxPoolConnections = i
	} else {
		return nil, fmt.Errorf("Переменная окружения «%v» не указана.\n", dbEnv.DbPool)
	}

	if i, b := strconv.Atoi(os.Getenv(dbEnv.DbPoolIdle)); b == nil {
		dbParams.MaxIdlePoolConnections = i
	} else {
		return nil, fmt.Errorf("Переменная окружения «%v» не указана.\n", dbEnv.DbPoolIdle)
	}

	if i, b := strconv.Atoi(os.Getenv(dbEnv.DbConnTimeout)); b == nil {
		dbParams.ConnectionTimeout = i
	} else {
		return nil, fmt.Errorf(EnvVarNotSpecified, dbEnv.DbConnTimeout)
	}

	return &dbParams, nil
}

package dataloader

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	envDbHost            string = "DB_HOST"
	envDbPort            string = "DB_PORT"
	envDbName            string = "DB_NAME"
	envDbUsername        string = "DB_USERNAME"
	envDbPassword        string = "DB_PASSWORD"
	envDbMaxOpenConn     string = "DB_MAX_OPEN_CONNS"
	envDbMaxIdleConn     string = "DB_MAX_IDLE_CONNS"
	postgresDriver       string = "postgres"
	stringConnection     string = "host=%s port=%s user=%s password=%s dbname=%s application_name='%s' sslmode=disable"
	defaultApplication   string = "data-loader_%s@%s"
	defaultUnknownUser   string = "unknown-user"
	defaultUnknownHost   string = "unknown-hostname"
	defaultDbMaxOpenConn int    = 10
	defaultDbMaxIdleConn int    = 3

	errorCouldNotCreateDbConn  string = "Could not create db connection: %v"
	errorCouldNotEstablishConn string = "Could not establish connection: %v"
)

func (loader *DataLoader) loadDatabaseConnection() {
	dbConnection := fmt.Sprintf(stringConnection,
		os.Getenv(envDbHost),
		os.Getenv(envDbPort),
		os.Getenv(envDbUsername),
		os.Getenv(envDbPassword),
		os.Getenv(envDbName),
		fmt.Sprintf(defaultApplication, getUsername(), getHostname()),
	)

	db, err := sql.Open(postgresDriver, dbConnection)
	if err != nil {
		log.Fatalf(errorCouldNotCreateDbConn, err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf(errorCouldNotEstablishConn, err)
	}

	db.SetMaxOpenConns(getIntEnv(envDbMaxOpenConn, defaultDbMaxOpenConn))
	db.SetMaxIdleConns(getIntEnv(envDbMaxIdleConn, defaultDbMaxIdleConn))
	loader.database = db
}

func getUsername() string {
	user, err := user.Current()
	if err != nil {
		return defaultUnknownUser
	}

	return user.Username
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return defaultUnknownHost
	}

	return hostname
}

func getIntEnv(env string, defaultValue int) int {
	value := os.Getenv(env)

	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return result
}

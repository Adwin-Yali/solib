package config

import (
	"github.com/J4stEu/solib/internal/app/errors"
	"github.com/J4stEu/solib/internal/app/errors/server_errors"
	"github.com/J4stEu/solib/internal/pkg"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

// Server - server_errors configuration
type Server struct {
	ServerAddr string
	ServerPort uint
	LogLevel   string
}

// DataBase - database configuration
type DataBase struct {
	PostgresIP    string
	PostgresPort  uint
	PostgresDB    string
	PostgresUser  string
	PostgresPass  string
	DataBaseInit  bool
	DataBaseDirty bool
	ForceVersion  uint
}

// Config - application configuration
type Config struct {
	Server   *Server
	DataBase *DataBase
}

func CheckENV() bool {
	_, err := os.LookupEnv("SERVER_ADDR")
	if !err {
		return false
	}
	_, err = os.LookupEnv("SERVER_PORT")
	if !err {
		return false
	}
	_, err = os.LookupEnv("PG_IP")
	if !err {
		return false
	}
	_, err = os.LookupEnv("PG_PORT")
	if !err {
		return false
	}
	_, err = os.LookupEnv("PG_DATABASE")
	if !err {
		return false
	}
	_, err = os.LookupEnv("PG_USER")
	if !err {
		return false
	}
	_, err = os.LookupEnv("PG_PASSWORD")
	if !err {
		return false
	}
	_, err = os.LookupEnv("PG_PASSWORD")
	if !err {
		return false
	}
	_, err = os.LookupEnv("LOG_LEVEL")
	if !err {
		return false
	}
	_, err = os.LookupEnv("DATABASE_INIT")
	if !err {
		return false
	}
	_, err = os.LookupEnv("DATABASE_DIRTY")
	if !err {
		return false
	}
	_, err = os.LookupEnv("FORCE_VERSION")
	if !err {
		return false
	}
	return true
}

func ReadConfiguration(logger *logrus.Logger) *Config {
	config := &Config{&Server{}, &DataBase{}}

	// Server configuration
	// ServerAddr
	serverAddr, err := os.LookupEnv("SERVER_ADDR")
	if !err {
		logger.Fatal(errors.ServerErrorLevel, server_errors.EnvReadError)
	}
	validServerAddr := pkg.IsValidIP(serverAddr)
	if !validServerAddr {
		logger.WithFields(log.Fields{
			"error": "Invalid server IP address.",
		}).Fatal(errors.SetError(errors.ServerErrorLevel, server_errors.EnvSetError))
	}
	// ServerPort
	config.Server.ServerAddr = serverAddr
	serverPort, err := os.LookupEnv("SERVER_PORT")
	if !err {
		logger.Fatal(errors.ServerErrorLevel, server_errors.EnvReadError)
	}
	serverPortUINT, convertErr := strconv.Atoi(serverPort)
	if convertErr != nil {
		logger.WithFields(log.Fields{
			"error": convertErr,
		}).Fatal(errors.SetError(errors.ServerErrorLevel, server_errors.EnvSetError))
	}
	config.Server.ServerPort = uint(serverPortUINT)
	// LogLevel
	logLevel, err := os.LookupEnv("LOG_LEVEL")
	if !err {
		logger.Fatal(errors.ServerErrorLevel, server_errors.EnvReadError)
	}
	config.Server.LogLevel = logLevel

	// Database Configuration
	// PostgresIP
	postgresIP, err := os.LookupEnv("PG_IP")
	if !err {
		logger.Fatal(errors.ServerErrorLevel, server_errors.EnvReadError)
	}
	validPostgresIP := pkg.IsValidIP(postgresIP)
	if !validPostgresIP {
		logger.WithFields(log.Fields{
			"error": "Invalid postgres IP address.",
		}).Fatal(errors.SetError(errors.ServerErrorLevel, server_errors.EnvSetError))
	}
	config.DataBase.PostgresIP = postgresIP
	// PostgresPort
	postgresPort, err := os.LookupEnv("PG_PORT")
	if !err {
		logger.Fatal(errors.ServerErrorLevel, server_errors.EnvReadError)
	}
	postgresPortUINT, convertErr := strconv.Atoi(postgresPort)
	if convertErr != nil {
		logger.WithFields(log.Fields{
			"error": convertErr,
		}).Fatal(errors.SetError(errors.ServerErrorLevel, server_errors.EnvSetError))
	}
	config.DataBase.PostgresPort = uint(postgresPortUINT)
	// PostgresDB
	postgresDB, err := os.LookupEnv("PG_DATABASE")
	if !err {
		logger.Fatal(errors.ServerErrorLevel, server_errors.EnvReadError)
	}
	config.DataBase.PostgresDB = postgresDB
	// PostgresUser
	postgresUser, err := os.LookupEnv("PG_USER")
	if !err {
		logger.Fatal(errors.ServerErrorLevel, server_errors.EnvReadError)
	}
	config.DataBase.PostgresUser = postgresUser
	// PostgresPass
	postgresPass, err := os.LookupEnv("PG_PASSWORD")
	if !err {
		logger.Fatal(errors.ServerErrorLevel, server_errors.EnvReadError)
	}
	config.DataBase.PostgresPass = postgresPass
	// DatabaseReconfigure
	dbReconfigure, err := os.LookupEnv("DATABASE_INIT")
	if !err {
		logger.Fatal(errors.ServerErrorLevel, server_errors.EnvReadError)
	}
	var parseBoolErr error
	config.DataBase.DataBaseInit, parseBoolErr = strconv.ParseBool(dbReconfigure)
	if parseBoolErr != nil {
		logger.WithFields(log.Fields{
			"error": parseBoolErr,
		}).Fatal(errors.SetError(errors.ServerErrorLevel, server_errors.EnvSetError))
	}
	// DataBaseDirty
	dbDirty, err := os.LookupEnv("DATABASE_DIRTY")
	if !err {
		logger.Fatal(errors.ServerErrorLevel, server_errors.EnvReadError)
	}
	config.DataBase.DataBaseDirty, parseBoolErr = strconv.ParseBool(dbDirty)
	if parseBoolErr != nil {
		logger.WithFields(log.Fields{
			"error": parseBoolErr,
		}).Fatal(errors.SetError(errors.ServerErrorLevel, server_errors.EnvSetError))
	}
	// ForceVersion
	forceVersion, err := os.LookupEnv("FORCE_VERSION")
	if !err {
		logger.Fatal(errors.ServerErrorLevel, server_errors.EnvReadError)
	}
	forceVersionUint, convertErr := strconv.Atoi(forceVersion)
	if convertErr != nil {
		logger.WithFields(log.Fields{
			"error": convertErr,
		}).Fatal(errors.SetError(errors.ServerErrorLevel, server_errors.EnvSetError))
	}
	config.DataBase.ForceVersion = uint(forceVersionUint)
	return config
}

func DefaultConfiguration() *Config {
	return &Config{
		Server: &Server{
			ServerAddr: "localhost",
			ServerPort: 8080,
			LogLevel:   "debug",
		},
		DataBase: &DataBase{
			PostgresIP:   "localhost",
			PostgresPort: 5432,
			PostgresDB:   "solib",
			PostgresUser: "postgres",
			PostgresPass: "postgres",
			DataBaseInit: false,
		},
	}
}

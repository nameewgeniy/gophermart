package config

import (
	"flag"
	"os"
)

type CnfContract interface {
	DatabaseDsn() string
	ServerAddr() string
	AccrualAddress() string
	LogLevel() string
}

var Conf CnfContract

func Singleton() {
	if Conf != nil {
		return
	}

	Conf = cnf{
		databaseDsn: "",
		serverAddr:  ":8090",
		logLevel:    "info",
	}
}

type cnf struct {
	databaseDsn    string
	serverAddr     string
	downMigrations bool
	accrualAddress string
	logLevel       string
}

func (c cnf) parseFlag() {
	flag.StringVar(&c.serverAddr, "a", "localhost:8080", "address")
	flag.StringVar(&c.accrualAddress, "r", "localhost:8088", "accrual address")
	flag.StringVar(&c.databaseDsn, "d", "", "database dsn")
	flag.BoolVar(&c.downMigrations, "dm", true, "down migrations after stop")
	flag.Parse()

	envAddr := os.Getenv("RUN_ADDRESS")
	if envAddr != "" {
		c.serverAddr = envAddr
	}

	envDatabaseDsn := os.Getenv("DATABASE_URI")
	if envDatabaseDsn != "" {
		c.databaseDsn = envDatabaseDsn
	}

	envAccrualAddress := os.Getenv("ACCRUAL_SYSTEM_ADDRESS")
	if envAccrualAddress != "" {
		c.accrualAddress = envAccrualAddress
	}
}

func (c cnf) DatabaseDsn() string {
	return c.databaseDsn
}

func (c cnf) ServerAddr() string {
	return c.serverAddr
}

func (c cnf) AccrualAddress() string {
	return c.accrualAddress
}

func (c cnf) LogLevel() string {
	return c.logLevel
}

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
	DownMigrations() bool
	AuthSecretKey() []byte
	AuthAccessTTL() int
	AuthRefreshTTL() int
}

var Conf CnfContract

func Singleton() {
	if Conf != nil {
		return
	}

	c := cnf{
		logLevel: "info",
	}
	c.parse()

	Conf = &c
}

type cnf struct {
	databaseDsn    string
	serverAddr     string
	downMigrations bool
	accrualAddress string
	logLevel       string
	authSecretKey  string
	accessTTL      int
	refreshTTL     int
}

func (c *cnf) parse() {
	flag.StringVar(&c.serverAddr, "a", "localhost:8090", "address")
	flag.StringVar(&c.accrualAddress, "r", "localhost:8088", "accrual address")
	flag.StringVar(&c.databaseDsn, "d", "postgres://user:password@localhost:5452/db?sslmode=disable", "database dsn")
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

	envAuthSecretKey := os.Getenv("AUTH_SECRET_KEY")
	if envAuthSecretKey != "" {
		c.authSecretKey = envAuthSecretKey
	}
}

func (c *cnf) DatabaseDsn() string {
	return c.databaseDsn
}

func (c *cnf) ServerAddr() string {
	return c.serverAddr
}

func (c *cnf) AccrualAddress() string {
	return c.accrualAddress
}

func (c *cnf) LogLevel() string {
	return c.logLevel
}

func (c *cnf) DownMigrations() bool {
	return c.downMigrations
}

func (c *cnf) AuthAccessTTL() int {
	return c.accessTTL
}

func (c *cnf) AuthRefreshTTL() int {
	return c.refreshTTL
}

func (c *cnf) AuthSecretKey() []byte {
	return []byte(c.authSecretKey)
}

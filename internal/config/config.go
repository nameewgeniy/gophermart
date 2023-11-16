package config

type CnfContract interface {
	DatabaseDsn() string
	ServerAddr() string
}

var Conf CnfContract

func Singleton() {
	if Conf != nil {
		return
	}

	Conf = cnf{
		databaseDsn: "",
		serverAddr:  ":8090",
	}
}

type cnf struct {
	databaseDsn string
	serverAddr  string
}

func (c cnf) DatabaseDsn() string {
	return c.databaseDsn
}

func (c cnf) ServerAddr() string {
	return c.serverAddr
}

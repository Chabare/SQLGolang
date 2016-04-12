package cSql

import (
	"log"

	mysql "github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
)

// CreateConnection creates a connection to the database
func CreateConnection(host, port, user, pass, name string) *mysql.Conn {
	d := mysql.New("tcp", "", host+":"+port, user, pass, name)

	err := d.Connect()

	if err != nil {
		log.Fatal(err)
	}

	return &d
}

// CreateConnectionByConfig creates a connection to the database
func CreateConnectionByConfig(c *Config) *mysql.Conn {
	return CreateConnection(c.Host, c.Port, c.User, c.Pass, c.Name)
}

// CreateConnection as function of struct Config
func (c *Config) CreateConnection() *mysql.Conn {
	return CreateConnectionByConfig(c)
}

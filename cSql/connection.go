package cSql

import (
	"log"

	"github.com/ziutek/mymysql/mysql"
)

// CreateConnection creates a connection to the database
func CreateConnection(host, port, user, pass, name string) mysql.Conn {
	d := mysql.New("tcp", "", host+":"+port, user, pass, name)
	err := d.Connect()

	if err != nil {
		log.Fatal(err)
	}

	return d
}
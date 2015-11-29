package sql

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// ConfigFile is a struct for the config file
type ConfigFile struct {
	Name string
}

// Config is a struct where the config values are saved
type Config struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

// open opens the file
func open(name string) *os.File {
	f, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
	}

	return f
}

// getLines gets every line from the file
func (c *ConfigFile) getLines() []string {
	file := open(c.Name)
	defer file.Close()

	data := make([]byte, 4096)
	file.Read(data)
	lines := strings.Split(string(data), "\n")

	// Check for empty byte array
	i, _ := strconv.Atoi(lines[len(lines)-1])
	if i == 0 {
		lines = lines[:len(lines)-1]
	}

	return lines[:]
}

// GetKeys returns a list of the defined keys
func (c *ConfigFile) GetKeys() []string {
	lines := c.getLines()

	for i, ele := range lines {
		lines[i] = ele[:strings.Index(ele, "=")]
	}

	return lines
}

// GetValues returns a list of the defined values
func (c *ConfigFile) GetValues() []string {
	lines := c.getLines()

	for i, ele := range lines {
		lines[i] = ele[strings.Index(ele, "=")+1:]
	}

	return lines
}

// GetValue returns a value from a given key
func (c *ConfigFile) GetValue(key string) {
	file := open(c.Name)

	file.Close()
}

// GetMap returns a map key[value]
func (c *ConfigFile) GetMap() map[string]string {
	config := make(map[string]string)
	keys := c.GetKeys()
	values := c.GetValues()

	for i := range keys {
		config[keys[i]] = values[i]
	}

	return config
}

// GetConfig returns a config struct
func (c *ConfigFile) GetConfig() *Config {
	var cnf Config
	m := c.GetMap()

	cnf.Host = m["Host"]
	cnf.Port = m["Port"]
	cnf.Name = m["Name"]
	cnf.User = m["User"]
	cnf.Pass = m["Pass"]

	return &cnf
}

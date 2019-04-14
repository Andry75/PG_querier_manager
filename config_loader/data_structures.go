package config_loader

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"net"
	"os"
	"strconv"
)

type ConfigError struct {
	Msg string
}

func (e ConfigError) Error() string {
	return fmt.Sprintf("%v", e.Msg)
}

type Config struct {
	Database        Database        `json:"database"`
	WebServer       WebServer       `json:"web_server"`
	InstancesMaster InstancesMaster `json:"instances_master"`
}

type Database struct {
	Vendor   string `json:"vendor"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
	SSLMode  string `json:"ssl_mode"`
}

type WebServer struct {
	Port int `json:"port"`
}

type InstancesMaster struct {
	IpAddress string `json:"ip_address"`
	Port      int    `json:"port"`
}

func (c *Config) load() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	err = gonfig.GetConf(currentDir+"/config.json", c)
	if err != nil {
		fmt.Println(err)
	}
}

func (c Config) GetWebServerPort() string {
	return ":" + strconv.Itoa(c.WebServer.Port)
}

func (c Config) GetInstancesMasterIpAddress() net.IPAddr {
	return net.IPAddr{IP: net.ParseIP(c.InstancesMaster.IpAddress)}
}

func (c Config) GetInstancesMasterPort() string {
	return ":" + strconv.Itoa(c.InstancesMaster.Port)
}

func (c Config) GetConnectionString() string {
	return c.Database.Vendor + "://" + c.Database.User + ":" +
		c.Database.Password + "@" + c.Database.Host + ":" +
		strconv.Itoa(c.Database.Port) + "/" + c.Database.DbName +
		"?sslmode=" + c.Database.SSLMode
}

package config_loader

import (
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"testing"
)

func TestConfigStructure(t *testing.T) {

	Convey(".GetWebServerPort", t, func() {

		Convey("positive", func() {

			Convey("return a correct string if a port exists", func() {
				c := Config{}
				c.WebServer.Port = 8080
				So(c.GetWebServerPort(), ShouldEqual, ":8080")
			})

		})

		Convey("negative", func() {
			c := Config{}
			Convey("return a correct string if port not set", func() {
				So(c.GetWebServerPort(), ShouldEqual, ":0")
			})
		})

	})

	Convey(".GetInstancesMasterIpAddress", t, func() {

		Convey("positive", func() {
			c := Config{}
			c.InstancesMaster.IpAddress = "10.0.0.1"

			Convey("return a correct string if IP address exists", func() {
				So(c.GetInstancesMasterIpAddress(), ShouldResemble,
					net.IPAddr{IP: net.ParseIP("10.0.0.1")})
			})

		})

		Convey("negative", func() {
			c := Config{}
			Convey("return a correct string if ip address not set", func() {
				So(c.GetInstancesMasterIpAddress(), ShouldResemble, net.IPAddr{})
			})

		})

	})

	Convey(".GetInstancesMasterPort", t, func() {

		Convey("positive", func() {

			Convey("return a correct string if a port exists", func() {
				c := Config{}
				c.InstancesMaster.Port = 80
				So(c.GetInstancesMasterPort(), ShouldEqual, ":80")
			})

		})

		Convey("negative", func() {

			Convey("return a correct string if a port not set", func() {
				c := Config{}
				So(c.GetInstancesMasterPort(), ShouldEqual, ":0")
			})

		})

	})

	Convey(".GetConnectionString", t, func() {

		Convey("positive", func() {
			c := Config{Database: Database{
				Vendor:   "postgres",
				Host:     "localhost",
				Port:     5432,
				User:     "test",
				Password: "test_password",
				DbName:   "test_db",
				SSLMode:  "enable",
			}}

			Convey("return a correct string if all data is set", func() {
				expectedConnectionString := "postgres://test:test_password@localhost:5432/test_db?sslmode=enable"
				actualConnectionString, err := c.GetConnectionString()
				So(err, ShouldBeNil)
				So(actualConnectionString, ShouldEqual, expectedConnectionString)
			})

		})

		Convey("negative", func() {

			Convey("assert if Database.Vendor not set", func() {
				c := Config{Database: Database{
					Host:     "localhost",
					Port:     5432,
					User:     "test",
					Password: "test_password",
					DbName:   "test_db",
					SSLMode:  "enable",
				}}

				actualConnectionString, err := c.GetConnectionString()
				So(actualConnectionString, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
				So(err, ShouldBeError)
			})

			Convey("assert if Database.User not set", func() {
				c := Config{Database: Database{
					Vendor:   "postgres",
					Host:     "localhost",
					Port:     5432,
					Password: "test_password",
					DbName:   "test_db",
					SSLMode:  "enable",
				}}

				actualConnectionString, err := c.GetConnectionString()
				So(actualConnectionString, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
				So(err, ShouldBeError)
			})

			Convey("assert if Database.Password not set", func() {
				c := Config{Database: Database{
					Vendor:  "postgres",
					Host:    "localhost",
					Port:    5432,
					User:    "test",
					DbName:  "test_db",
					SSLMode: "enable",
				}}

				actualConnectionString, err := c.GetConnectionString()
				So(actualConnectionString, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
				So(err, ShouldBeError)
			})

			Convey("assert if Database.Host not set", func() {
				c := Config{Database: Database{
					Vendor:   "postgres",
					Port:     5432,
					User:     "test",
					Password: "test_password",
					DbName:   "test_db",
					SSLMode:  "enable",
				}}

				actualConnectionString, err := c.GetConnectionString()
				So(actualConnectionString, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
				So(err, ShouldBeError)
			})

			Convey("assert if Database.Port not set", func() {
				c := Config{Database: Database{
					Vendor:   "postgres",
					Host:     "localhost",
					User:     "test",
					Password: "test_password",
					DbName:   "test_db",
					SSLMode:  "enable",
				}}

				actualConnectionString, err := c.GetConnectionString()
				So(actualConnectionString, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
				So(err, ShouldBeError)
			})

			Convey("assert if Database.DbName not set", func() {
				c := Config{Database: Database{
					Vendor:   "postgres",
					Host:     "localhost",
					Port:     5432,
					User:     "test",
					Password: "test_password",
					SSLMode:  "enable",
				}}

				actualConnectionString, err := c.GetConnectionString()
				So(actualConnectionString, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
				So(err, ShouldBeError)
			})

			Convey("assert if Database.SSLMode not set", func() {
				c := Config{Database: Database{
					Vendor:   "postgres",
					Host:     "localhost",
					Port:     5432,
					User:     "test",
					Password: "test_password",
					DbName:   "test_db",
				}}

				actualConnectionString, err := c.GetConnectionString()
				So(actualConnectionString, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
				So(err, ShouldBeError)
			})

		})

	})

	Convey(".GetConnectionStringWithoutDB", t, func() {

		Convey("positive", func() {

			c := Config{Database: Database{
				Vendor:   "postgres",
				Host:     "localhost",
				Port:     5432,
				User:     "test",
				Password: "test_password",
				DbName:   "test_db",
				SSLMode:  "enable",
			}}

			Convey("return a correct string if all data is set", func() {
				expectedConnectionString := "postgres://test:test_password@localhost:5432/postgres?sslmode=enable"
				actualConnectionString, err := c.GetConnectionStringWithoutDB()
				So(err, ShouldBeNil)
				So(actualConnectionString, ShouldEqual, expectedConnectionString)
			})

			Convey("return a correct string if all data is set except DbName", func() {
				c.Database.DbName = ""
				expectedConnectionString := "postgres://test:test_password@localhost:5432/postgres?sslmode=enable"
				actualConnectionString, err := c.GetConnectionStringWithoutDB()
				So(err, ShouldBeNil)
				So(actualConnectionString, ShouldEqual, expectedConnectionString)
			})

		})

		Convey("negative", func() {

			Convey("assert if Database.Vendor not set", func() {
				c := Config{Database: Database{
					Host:     "localhost",
					Port:     5432,
					User:     "test",
					Password: "test_password",
					DbName:   "test_db",
					SSLMode:  "enable",
				}}

				actualConnectionString, err := c.GetConnectionStringWithoutDB()
				So(actualConnectionString, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
				So(err, ShouldBeError)
			})

			Convey("assert if Database.User not set", func() {
				c := Config{Database: Database{
					Vendor:   "postgres",
					Host:     "localhost",
					Port:     5432,
					Password: "test_password",
					DbName:   "test_db",
					SSLMode:  "enable",
				}}

				actualConnectionString, err := c.GetConnectionStringWithoutDB()
				So(actualConnectionString, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
				So(err, ShouldBeError)
			})

			Convey("assert if Database.Password not set", func() {
				c := Config{Database: Database{
					Vendor:  "postgres",
					Host:    "localhost",
					Port:    5432,
					User:    "test",
					DbName:  "test_db",
					SSLMode: "enable",
				}}

				actualConnectionString, err := c.GetConnectionStringWithoutDB()
				So(actualConnectionString, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
				So(err, ShouldBeError)
			})

			Convey("assert if Database.Host not set", func() {
				c := Config{Database: Database{
					Vendor:   "postgres",
					Port:     5432,
					User:     "test",
					Password: "test_password",
					DbName:   "test_db",
					SSLMode:  "enable",
				}}

				actualConnectionString, err := c.GetConnectionStringWithoutDB()
				So(actualConnectionString, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
				So(err, ShouldBeError)
			})

			Convey("assert if Database.Port not set", func() {
				c := Config{Database: Database{
					Vendor:   "postgres",
					Host:     "localhost",
					User:     "test",
					Password: "test_password",
					DbName:   "test_db",
					SSLMode:  "enable",
				}}

				actualConnectionString, err := c.GetConnectionStringWithoutDB()
				So(actualConnectionString, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
				So(err, ShouldBeError)
			})

			Convey("assert if Database.SSLMode not set", func() {
				c := Config{Database: Database{
					Vendor:   "postgres",
					Host:     "localhost",
					Port:     5432,
					User:     "test",
					Password: "test_password",
					DbName:   "test_db",
				}}

				actualConnectionString, err := c.GetConnectionStringWithoutDB()
				So(actualConnectionString, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
				So(err, ShouldBeError)
			})

		})

	})

}

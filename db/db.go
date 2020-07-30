package db

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // needed for db connection
	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

const (
	//DefaultConnRetryCount ... default retry count for db connection
	DefaultConnRetryCount = 5

	//DefaultConnRetryDelay ... default retry delay for db connection
	DefaultConnRetryDelay = 100

	//DefaultConnAlias ... default alias for db connection
	DefaultConnAlias = "default"

	//DefaultDBDriver ... default driver for db connection
	DefaultDBDriver = "mysql"

	//DefaultConnTimeZone ... default time zone for db connection
	DefaultConnTimeZone = "Asia/Kolkata"
)

//Connection ... default connection struct
type Connection struct {
	Name string
}

//IsDbConnected ...
func IsDbConnected(connectionName string) bool {
	c := new(Connection)
	c.Name = connectionName
	err := c.Connect()
	if err != nil {
		log.Print(err)
		return false
	}
	return true
}

//Connect ... connect to database
func (c *Connection) Connect() error {

	var err error
	uri := "root:chetan@tcp(127.0.0.1:3306)/gl"
	log.Print(uri)

	retryCount := c.getRetryCount()
	retryDelay := c.getRetryDelay()
	alias := c.getAlias()
	driver := c.getDriver()
	timezone := c.getTimeZone()
	maxIdle := 0 //maximum Idle connection orm can have at a particular time (added as it was giving max idle connection reached error)

	for breaker := retryCount; breaker > 0; breaker-- {

		//if it is a retry, wait for sometime before retry
		if breaker < retryCount {
			//sleep for the given delay
			time.Sleep(time.Duration(retryDelay) * time.Millisecond)
		}

		//check if db is already registered, if not register it
		_, err = orm.GetDB(alias)
		if err != nil {
			err = orm.RegisterDataBase(alias, driver, uri, maxIdle)
		}

		//check if db is registered by now, if not continue to retry
		if err != nil {
			log.Print("Connection error:", err, "Breaker:", breaker)
			continue
		}

		//check if connection is still working, if not continue to retry
		err = c.pingCheck(alias)
		if err != nil {
			log.Print("Connection error:", err, "Breaker:", breaker)
			continue
		}

		//All OK !!! -> configure orm/connection with provided input
		orm.DefaultTimeLoc, _ = time.LoadLocation(timezone)
		orm.Debug = c.getDebugFlag()
		break

	}

	return nil
}

func (c *Connection) prefix() string {
	if len(c.Name) > 0 {
		return (c.Name + ".")
	}
	return ""
}

func (c *Connection) getURI() (string, error) {
	k := c.prefix() + "db.connection.uri"
	v := viper.GetString(k)
	if "" == v {
		return "", errors.New("failed to get " + k)
	}
	return v, nil
}

func (c *Connection) getRetryCount() int {
	k := c.prefix() + "db.connection.retrycount"
	v := viper.GetInt(k)
	if 0 == v {
		log.Print("failed to get ", k, ", using ", DefaultConnRetryCount)
		return DefaultConnRetryCount
	}
	return v
}

func (c *Connection) getRetryDelay() int {
	k := c.prefix() + "db.connection.retrydelay"
	v := viper.GetInt(k)
	if 0 == v {
		log.Print("failed to get ", k, ", using ", DefaultConnRetryDelay)
		return DefaultConnRetryDelay
	}
	return v
}

func (c *Connection) getAlias() string {
	k := c.prefix() + "db.connection.alias"
	v := viper.GetString(k)
	if "" == v {
		log.Print("failed to get ", k, ", using ", DefaultConnAlias)
		return DefaultConnAlias
	}
	return v
}

func (c *Connection) getDriver() string {
	k := c.prefix() + "db.connection.driver"
	v := viper.GetString(k)
	if "" == v {
		log.Print("failed to get ", k, " using ", DefaultDBDriver)
		return DefaultDBDriver
	}
	return v
}

func (c *Connection) getTimeZone() string {
	k := c.prefix() + "db.connection.tz"
	v := viper.GetString(k)
	if "" == v {
		log.Print("failed to get ", k, ", using ", DefaultConnTimeZone)
		return DefaultConnTimeZone
	}
	return v
}

func (c *Connection) getDebugFlag() bool {
	k := c.prefix() + "db.connection.debug"
	return viper.GetBool(k)
}

func (c *Connection) pingCheck(alias string) error {
	o := orm.NewOrm()
	o.Using(alias)
	_, err := o.Raw("SELECT 1").Exec()
	return err
}

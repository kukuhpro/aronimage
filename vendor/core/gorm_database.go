package core

import (
	"fmt"
	"sync"

	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var instancesDB *gorm.DB
var once sync.Once

type Db struct {
	db               *gorm.DB
	config           map[string]interface{}
	driverdb         string
	endurlConnection string
}

func (this *Db) SetEndUrlConnection(endUrlConnection string) {
	this.endurlConnection = endUrlConnection
}

func (this *Db) GenerateUrlDatabase() string {
	config := this.config
	usernameAndPassword := fmt.Sprint(config["db_user"]) + ":" + fmt.Sprint(config["db_password"])
	hostName := "tcp(" + fmt.Sprint(config["db_host"]) + ":" + fmt.Sprint(config["db_port"]) + ")"

	urlConnection := usernameAndPassword + "@" + hostName

	return urlConnection
}

func (this *Db) SetConfiguration(config map[string]interface{}) {
	this.config = config
}

func (this *Db) GetConfiguration() map[string]interface{} {
	return this.config
}

func (this *Db) GetEndUrlConnection() string {
	return this.endurlConnection
}

func (this *Db) GetConnection() *gorm.DB {
	return this.db
}

func (this *Db) GetDriver() string {
	return this.driverdb
}

func (this *Db) DefaultConnection() {

	config := this.config
	driver := fmt.Sprint(config["driver"])
	this.driverdb = driver

	usernameAndPassword := fmt.Sprint(config["db_user"]) + ":" + fmt.Sprint(config["db_password"])
	hostName := "tcp(" + fmt.Sprint(config["db_host"]) + ":" + fmt.Sprint(config["db_port"]) + ")"

	log.Println("Connecting to DB Server " + fmt.Sprint(config["db_host"]) + ":" + fmt.Sprint(config["db_port"]) + "...")
	urlConnection := usernameAndPassword + "@" + hostName + "/" + fmt.Sprint(config["db_database"]) + this.endurlConnection

	once.Do(func() {
		db, err := gorm.Open(driver, urlConnection)
		if err != nil {
			panic(err)
		}
		instancesDB = db
	})

	log.Println("DB Server is connected!")

	this.db = instancesDB
}

func Database(configdatabase map[string]interface{}) *Db {
	var db Db

	db.endurlConnection = "?charset=utf8&parseTime=True"
	db.config = configdatabase
	db.DefaultConnection()

	return &db
}

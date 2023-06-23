package infra

import (
	"fmt"
	"math/rand"
	"url_shortening/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	masterDb *gorm.DB
	slaveDbs []*gorm.DB
)

func getMasterDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetConfig().Database.Master.Username,
		config.GetConfig().Database.Master.Password,
		config.GetConfig().Database.Master.Host,
		config.GetConfig().Database.Master.Port,
		config.GetConfig().Database.Master.DatabaseName,
	)
}

func getSlavesDsn() []string {
	var slavesDsn []string
	for _, slave := range config.GetConfig().Database.Slaves {
		slavesDsn = append(slavesDsn, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			slave.Username,
			slave.Password,
			slave.Host,
			slave.Port,
			slave.DatabaseName,
		))
	}
	return slavesDsn
}

func InitDB() error {
	// Initialize the master database
	db, err := gorm.Open(mysql.Open(getMasterDsn()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	masterDb = db

	// Initialize the slave databases
	for _, dsn := range getSlavesDsn() {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to database")
		}
		slaveDbs = append(slaveDbs, db)
	}

	return nil
}

func GetMasterDB() *gorm.DB {
	return masterDb
}

func GetSlaveDB() *gorm.DB {
	if len(slaveDbs) == 0 {
		GetLogger().Fatal("No slave databases available")
		return nil
	}

	// Generate a random index to select a slave database
	index := rand.Intn(len(slaveDbs))

	return slaveDbs[index]
}

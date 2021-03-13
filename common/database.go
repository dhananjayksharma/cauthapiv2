package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
)

type Database struct {
	*gorm.DB
}


var DB *gorm.DB


// Opening a database and save the reference to `Database` struct.
func Init(dbsconn string) *gorm.DB {


	var (
		databasename, hostPort, userPass, userName string
	)

	if dbsconn == "stageenvwrite" {
		userName = viper.GetString("database.stageenvwrite.dbuser")     //"phpmyadmin"
		userPass = viper.GetString("database.stageenvwrite.dbpassword") //"MLIAaCns"
		hostPort = viper.GetString("database.stageenvwrite.hostname")   //"52.66.175.114:3306"
		databasename = viper.GetString("database.stageenvwrite.dbname")       //CMS
	} else if dbsconn == "stageenvread" {
		userName = viper.GetString("database.stageenvread.dbuser")     //"phpmyadmin"
		userPass = viper.GetString("database.stageenvread.dbpassword") //"MLIAaCns"
		hostPort = viper.GetString("database.stageenvread.hostname")   //"52.66.175.114:3306"
		databasename = viper.GetString("database.stageenvread.dbname")       //CMS
	} else if dbsconn == "laptopenv" {
		userName = viper.GetString("database.laptopenv.dbuser")     //"phpmyadmin"
		userPass = viper.GetString("database.laptopenv.dbpassword") //"MLIAaCns"
		hostPort = viper.GetString("database.laptopenv.hostname")   //"52.66.175.114:3306"
		databasename = viper.GetString("database.laptopenv.dbname")       //CMS
	} else {
		userName = viper.GetString("database.liveenv.dbuser")     //"phpmyadmin"
		userPass = viper.GetString("database.liveenv.dbpassword") //"MLIAaCns"
		hostPort = viper.GetString("database.liveenv.hostname")   //"52.66.175.114:3306"
		databasename = viper.GetString("database.liveenv.dbname")       //CMS
	}
	var dataSourceName = userName + ":" + userPass + "@tcp(" + hostPort + ")/"+ databasename +"?charset=utf8&parseTime=True&loc=Local"
	var dbType = "mysql"
	db, err := gorm.Open(dbType, dataSourceName)
	db.SingularTable(true)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	fmt.Println("Database connection done")
	//db.LogMode(true)
	DB = db
	return DB
}



// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}

func DBConn(dbsconn string) *gorm.DB {
	Init(dbsconn)
	dbread := GetDB()
	return dbread
}

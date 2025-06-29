package database

import (
	"encoder/domain"
	"log"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	_"github.com/lib/pq"
	"github.com/jinzhu/gorm"
)


type Database struct {
	Db *gorm.DB
	Dsn string 
	DsnTest string
	DbType string
	DbTypeTest string
	Debug bool
	AutoMigrateDb bool
	Env string 
}

func NewDatabase() *Database {
	return &Database{
	}
}

func NewDatabaseTest() *gorm.DB {

	dbInstance := NewDatabase()	
	dbInstance.Env = "test"
	dbInstance.Debug = true
	dbInstance.AutoMigrateDb = true
	dbInstance.DbType = "sqlite3"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory:"


	connection, err := dbInstance.Connect()
	if err != nil{
		log.Fatalf("Error connecting to database: %v", err)
	}
	return connection
}

func (db *Database) Connect() (*gorm.DB, error) {
	var err error

	if db.Env == "test"{
		db.Db, err = gorm.Open(db.DbTypeTest, db.DsnTest)
	} else {
		db.Db, err = gorm.Open(db.DbType, db.Dsn)
	}

	if err != nil{
		return nil, err
	}

	if db.Debug{
		db.Db.LogMode(true)
	}

	if db.AutoMigrateDb{
		db.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
		db.Db.Model(&domain.Job{}).AddForeignKey("video_id", "videos(id)", "CASCADE", "CASCADE")
	}

	return db.Db, nil
}
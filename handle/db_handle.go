package handle

import (
	"dbresolver/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

type DBClient struct {
	DB  *gorm.DB
	Err error
}

func (client *DBClient) InitDB() error {
	dsn := config.GetDSN()

	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	dbResolverCfg := dbresolver.Config{
		Sources: []gorm.Dialector{postgres.Open(dsn)},
		Policy:  dbresolver.RandomPolicy{},
	}

	readWritePlugin := dbresolver.Register(dbResolverCfg).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(config.Min()).
		SetMaxOpenConns(config.Max())

	err := db.Use(readWritePlugin)
	if err != nil {
		return err
	}
	client.DB = db
	return nil
}

func (client *DBClient) Insert(value interface{}) {
	db := client.DB
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(value).Error; err != nil {
			return err
		}
		return nil
	})

	client.Err = err
}

func (client *DBClient) Find(value interface{}) {
	db := client.DB
	db.Find(value)
}

func (client *DBClient) Retrieved(value interface{}, limit int, offset int) {
	db := client.DB
	db.Limit(limit).Offset(offset).Find(value)
}

package main

import (
	"github.com/3dw1nM0535/galva/store"
	"github.com/3dw1nM0535/galva/store/models"
)

func migrate() error {
	orm, err := store.NewORM()
	defer orm.Store.Close()

	// Drop previously created table
	orm.Store.DropTableIfExists(&models.Land{})
	orm.Store.DropTableIfExists(&models.User{})

	// Migrate data model and create tables based off the models
	err = orm.Store.AutoMigrate(
		&models.Land{},
		&models.User{},
	).Error
	if err != nil {
		return err
	}

	return nil
}

func main() {
	migrate()
}

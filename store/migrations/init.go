package main

import (
	"github.com/3dw1nM0535/galva/store"
	"github.com/3dw1nM0535/galva/store/models"
)

func migrate() error {
	orm, err := store.Factory()
	defer orm.Store.Close()

	err = orm.Store.AutoMigrate(
		&models.Land{},
	).Error
	if err != nil {
		return err
	}

	return nil
}

func main() {
	migrate()
}

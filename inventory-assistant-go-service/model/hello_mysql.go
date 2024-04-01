package model

import "inventory-assistant/tools"

func Migration() error {
	if db, err := tools.SetupDB(); err == nil {
		db.AutoMigrate(&User{}, &Post{})
		return nil
	} else {
		return err
	}

}

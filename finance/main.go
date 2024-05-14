package main

import (
	"Finance/config"
	"Finance/routers"

	"gorm.io/gorm"
)

func main() {

	var db *gorm.DB

	// config.SetupDBConnection(db)

	db = config.Setup(db)

	r := routers.SetupRouter(db)

	r.Run(":3001")
}

package main

import (
	"flag"
	"fmt"
	"os"

	"slashbase.com/backend/config"
	"slashbase.com/backend/db"
	"slashbase.com/backend/models/user"
	"slashbase.com/backend/server"
)

func main() {
	environment := flag.String("e", "local", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.InitGormDB()
	autoMigrate()
	server.Init()
}

func autoMigrate() {
	db.GetDB().AutoMigrate(&user.User{}, &user.UserSession{})
}

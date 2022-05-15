package main

import (
	"os"
	"log"
	"CRUD+Check/pkg/interfaces"
	"CRUD+Check/pkg/models"
	"CRUD+Check/pkg/config"
)

func main() {
	var worker interfaces.IDataWorker
	
	f, err := os.OpenFile("log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("error opening file: %v", err)
    }
    defer f.Close()

    log.SetOutput(f)
	config.Connect()
	sqliteWorker := &models.SQLiteWorker{DB: config.GetDB()}
	worker = sqliteWorker
	worker.DB.AutoMigrate(&models.User{})
	/*user1 := models.User{Mail: "user@example.com", Passhash: "password"}
	user2 := models.User{Mail: "polzovatel@mirea.com", Passhash: "password"}
	user4 := models.User{Mail: "kudzh@kafedra.ru", Passhash: "rektor"}
	user_test1 := models.User{Mail: "polzovatel@mirea.com", Passhash: "password"}
	user_test2 := models.User{Mail: "user@example.com", Passhash: "rektor"}
	user_test3 := models.User{Mail: "user@example.com", Passhash: "password"}
	err = worker.CreateUser(user2)
	if err != nil {
		log.Println("Error creating user", err)
	}
	err = worker.CreateUser(user4)
	if err != nil {
		log.Println("Error creating user", err)
	}
	err = worker.CheckUser(user_test1)
	if err != nil {
		log.Println("Error checking user 1 ", err)
	}
	err = worker.CheckUser(user_test2)
	if err != nil {
		log.Println("Error checking user 2 ", err)
	}
	err = worker.CheckUser(user_test3)
	if err != nil {
		log.Println("Error checking user 3 ", err)
	}
	err = worker.DeleteUser("kudzh@kafedra.ru")
	if err != nil {
		log.Println("Error deleting user", err)
	}*/
	user := models.User{Mail:"mirea@gmail.com", Passhash: "abcdef"}
	err = worker.CheckUser(user)
	if err != nil {
		log.Println("Error updating user", err)
	}
}
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"rpc-tutorial/Item"
	"rpc-tutorial/database"
	"rpc-tutorial/internal/basicCrudServiceServer"
	"rpc-tutorial/rpc/basicCrudService"
)

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("mysql","root:sql@tcp(127.0.0.1:3308)/mypgdb?charset=utf8&parseTime=True")
	if err!=nil{
		fmt.Println("Connection Failed to Open")
	}else{
		fmt.Println("Connection Established")
	}

	database.DBConn.AutoMigrate(Item.Item{})
	fmt.Println("Automigration done")
}

func main() {
	initDatabase()
	defer database.DBConn.Close()
	server := &basicCrudServiceServer.Server{} // implements Haberdasher interface
	twirpHandler := basicCrudService.NewBasicCrudServiceServer(server)

	http.ListenAndServe(":8085", twirpHandler)
}
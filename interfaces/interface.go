// the interface acts as a connecting gateway to application from database, acts as conformer.
// From a phyiscal perspective, consider interfaces to be a port in a plugged device which connects the electricity to the device.

package main

import (
	"fmt"
	"log"

	"github.com/ooonkeet/golang-advanced/interfaces/mysqldb"
)

// connection point for db and application
type dbContract interface{
	Close() //to be implemented by db implementation
	InsertUser(userName string) error //input a string and return an error
	SelectSingleUser(userName string) (string,error) //returns user and an error

} //contract that the db implementation needs to follow or conform to, to plug into our application

type Application struct{
	//for the app
	db dbContract //embed db into the struct
}
func (this Application) Run(){
	// this keyword can access the db property using receiver variable
	userName:="user1"
	err:=this.db.
	InsertUser(userName)
	if err!=nil{
		log.Println(err)
	}
	user,err:=this.db.SelectSingleUser(userName)
	if err!=nil{
		log.Println(err)
	}
	fmt.Println(user)
}
func NewApplication(db dbContract) *Application{
	return &Application{db: db}
} //create a new application - pass db to create a new application, return a pointer to application struct, return an instance of application with the db embedded 
// we pass the db which needs to conform to the contract and then we are going to return the struct to new instance with the db we pass in here embeeded into the instance and then from our main function

func main(){
	dbUser:="user"
	dbPass:="password"
	dbHost:="host"
	dbPort:="port"
	dbName:="name"
	db,err:=mysqldb.New(dbUser,dbPass,dbHost,dbPort,dbName) //going to get the DB back which in this case is mysqldb
	if err!=nil{
		log.Fatal("failed to initiate db connection %v",err)
	}
	defer db.Close()
	// conn point of our application via db
	app:=NewApplication(db) //instantiating the application - plugging in db to the application - connect app and db
	app.Run()
}
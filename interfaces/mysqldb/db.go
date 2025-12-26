package mysqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	dbDriver = "mysql"
)

type Mysql struct {
	db *sql.DB
} //obj returned from this new - call to new function
//return an instance of Mysql DB

func New(dbUser, dbPassword, dbHost, dbPort, dbName string) (*Mysql, error) {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
//db connect start
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Println("mysqldb connection failure: %v", err)
		return nil, err
	}
//db connect end
	err = db.Ping()
	if err != nil {
		log.Println("mysqldb ping failure: %v", err)
		return nil, err
	}

	return &Mysql{db: db}, nil
}

func (this Mysql) Close() { //tying this function to MySQL struct
	// if we create mysql struct it is going to be associated with this close function
	// same for InsertUser function
	// same for SelectSingleUser function
	err := this.db.Close()
	if err != nil {
		log.Println("mysqldb close failure: %v", err)
	}
}

func (this Mysql) InsertUser(userName string) error {
	this.db.Exec("INSERT...")

	return nil
}

func (this Mysql) SelectSingleUser(userName string) (string, error) {
	this.db.Exec("SELECT...")

	return "user", nil
}
package main

import (
	"database/sql"
	"fmt"
	"sql_api_implementation_2/config"
	"sql_api_implementation_2/routers"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
	err error
)

func main() {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// DB, err = sql.Open("postgres", psqlInfo) // untuk memvalidasi argumen-argumen yang diberikan
	// if err != nil {
	// 	panic(err)
	// }
	var PORT = ":8080" 
	DB = config.Connect()
	defer DB.Close()

	err = DB.Ping() //membangun koneksi ke database
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	routers.StartServer().Run(PORT)


	
}

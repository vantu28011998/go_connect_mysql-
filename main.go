package main

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
type Users struct{
	id int
	username string
	password string
}
func main(){
	fmt.Println("Go Mysql connection")
	db,err := sql.Open("mysql","{USERNAME}:{PASSWORD}@tcp(127.0.0.1:3306)/{DATABASE}")
	if err !=nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Successfully connected to mysql")
	insert,err := db.Query("INSERT INTO users VALUES(8, 'helloworld', 'helloworld');")
	if err !=nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successfully inserted to mysql")

	results, err := db.Query("SELECT * FROM USERS")
	if err !=nil {
		panic(err.Error())
	}
	for results.Next(){
		var user Users 
		err :=results.Scan(&user.id,&user.username,&user.password)
		if err !=nil {
			panic(err.Error())
		}
		fmt.Println(user)
	}
	fmt.Println("Successfully selected to mysql")
	
}
package dao

import (
	"database/sql"
	"log"
	"fmt"
	"os/user"
)
// START OMIT
func Get(name string, age int)  {
	var db *sql.DB
	rows, _ := db.Query("SELECT name, num, age " +
		"FROM users WHERE age=? AND name = ?", age, name)
	defer rows.Close()
	for rows.Next() {
		var user users.User
		err := rows.Scan(&user.Name, &user.Num, &user.Age)
// END OMIT
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s is %d\n", name, age)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
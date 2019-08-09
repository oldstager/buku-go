package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/aplikasi")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Sukses!")
	}

	var (
		id    int
		login string
		name  string
	)
	rows, err := db.Query("select id, login, nama from pengguna where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &login, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, login, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

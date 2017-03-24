package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// создаём структуру базы
	// но соединение происходит только при мервом запросе
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/msu-go-11?charset=utf8")
	PanicOnErr(err)

	// проверяем что подключение реально произошло ( делаем запрос )
	err = db.Ping()
	PanicOnErr(err)

	// итерируемся по многим записям
	rows, err := db.Query("SELECT fio FROM students")
	PanicOnErr(err)
	for rows.Next() {
		var fio string
		err = rows.Scan(&fio)
		PanicOnErr(err)
		fmt.Println("rows.Next fio: ", fio)
	}
	rows.Close()

	var fio string
	row := db.QueryRow("SELECT fio FROM students WHERE id = 1")
	err = row.Scan(&fio)
	PanicOnErr(err)
	fmt.Println("db.QueryRow fio: ", fio)

	result, err := db.Exec(
		"INSERT INTO students (fio) VALUES (?)",
		"Ivan Ivanov",
	)
	PanicOnErr(err)

	affected, err := result.RowsAffected()
	PanicOnErr(err)
	lastId, err := result.LastInsertId()
	PanicOnErr(err)

	fmt.Println("RowsAffected", affected, "LastInsertId: ", lastId)

}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

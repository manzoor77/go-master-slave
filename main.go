package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to the master database
	masterDB, err := sql.Open("postgres", "user=test_user password=test_passwd dbname=test_db host=135.181.55.235 port=5438 sslmode=disable")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer masterDB.Close()

	// Insert a record into the master database
	_, err = masterDB.Exec("INSERT INTO test_table (id, name) VALUES ($1, $2)", "value1", "value2")
	if err != nil {
		fmt.Println("err1:", err)
		return
	}

	// Connect to the slave database
	slaveDB, err := sql.Open("postgres", "user=test_user password=test_passwd dbname=test_db host=135.181.55.235 port=5439 sslmode=disable")
	if err != nil {
		fmt.Println("err2:", err)
		return
	}
	defer slaveDB.Close()

	// Read the record from the slave database
	var column1, column2 string
	err = slaveDB.QueryRow("SELECT id, name FROM test_table LIMIT 1").Scan(&column1, &column2)
	if err != nil {
		fmt.Println("err3:", err)
		return
	}

	fmt.Println("Record from the slave database:")
	fmt.Println("id:", column1)
	fmt.Println("name:", column2)
}

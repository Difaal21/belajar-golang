package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	command := "INSERT INTO customer(name) VALUES('al fansha')"

	_, err := db.ExecContext(ctx, command)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var (
			id   int
			name string
		)

		err := rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

		fmt.Println("id", id)
		fmt.Println("name", name)
	}

	defer rows.Close()
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var (
			id          int32
			name, email sql.NullString
			balance     int32
			rating      float64
			birthDate   sql.NullString
			createdAt   time.Time
			married     bool
		)

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)

		if err != nil {
			panic(err)
		}

		fmt.Println("================")
		fmt.Println("id", id)
		fmt.Println("name", name)
		if email.Valid {
			fmt.Println("email", email)
		}
		fmt.Println("balance", balance)
		fmt.Println("rating", rating)
		fmt.Println("birthDate", birthDate)
		fmt.Println("married", married)
		fmt.Println("createdAt", createdAt)
	}

	defer rows.Close()
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// username := "admin'; #"
	username := "admina' OR 1=1; #"
	password := "admin"

	query := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"

	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("Login Success")
	} else {
		fmt.Println("Login Failed")
	}

	defer rows.Close()
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// ==== Script SQL Injection ==== //
	// username := "admin'; #"
	// username := "admina' OR 1=1; #"

	username := "admin"
	password := "admin"

	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"

	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query, username, password)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("Login Success")
	} else {
		fmt.Println("Login Failed")
	}

	defer rows.Close()
}

func TestExecSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	command := "INSERT INTO user(username, password) VALUES(?,?)"

	// ==== Script SQL Injection ==== //
	// username := "superadmin'; #"
	username := "superadmin' OR 1=1; #"

	// username := "superadmin"
	password := "superadmin"

	_, err := db.ExecContext(ctx, command, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")

}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "difaal21@gmail.com"
	comment := "Ini komen kedua"

	command := "INSERT INTO comments(email, comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, command, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment wit id", insertId)
}

// Ketika kita mengeksekusi ExecContext() dan QueryContext() belum tentu menggunakan koneksi yang sama, lebih efisien menggunakan koneksi yang sama dengan cara prepare statement dengan cara Prepare(context, sql)
func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	command := "INSERT INTO comments(email, comment) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, command)
	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "difaal" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini komen prepare statement ke-" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}
}

// DB transaksi, satu gagal semua akan dianggap gagal
func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	//-------- Do Transaction --------//
	command := "INSERT INTO comments(email, comment) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, command)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		email := "difaal" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini komen db transaction ke-" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}

	//-------- End Transaction --------//

	// Commit gunanya untuk memasukkan data ke DB, setiap perintah SQL secara otomatis akan commit (auto commit)
	err = tx.Commit()

	// Rollback berguna untuk membatalkan data, data tidak akan masuk ke DB
	// err = tx.Rollback()

	if err != nil {
		panic(err)
	}
}

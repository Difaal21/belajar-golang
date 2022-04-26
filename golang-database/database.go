package golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_database?parseTime=true")

	if err != nil {
		panic(err)
	}

	// Jumlah minimal koneksi yang dibuat
	db.SetMaxIdleConns(10)
	// Jumlah maksimal koneksi yang dibuat
	db.SetMaxOpenConns(100)
	// Berapa lama koneksi yang sudah tidak digunakan akan dihapus || Setelah lima menit diem aja, yaa udah close aja
	db.SetConnMaxIdleTime(5 * time.Minute)
	// Berapa lama koneksi digunakan || Koneksi apapun yang sudah mencapai 60 menit akan dibuatkan koneksi baru
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

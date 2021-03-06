package models

import (
	"APi_Rest/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	CreateConnection()

}
func CreateConnection() {
	if GetConnection() != nil {
		return
	}
	url := config.GetUrlDatabase()
	if connection, err := sql.Open("mysql", url); err != nil {
		panic(err)
	} else {
		fmt.Println("Conexion exitosa ")
		db = connection
	}

}
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
func CloseConnection() {
	db.Close()
	fmt.Println("Conexion cerrada con exito")
}

/// <username>:<passwrod>@tcp(<host>:<port>)/<database>
func GetConnection() *sql.DB {
	return db
}

func Exec(query string, args ...interface{}) (sql.Result, error) { // recibir n cantidad de parametros con ...interface
	result, err := db.Exec(query, args...) // exec retorna un objeto result ,
	if err != nil {
		log.Println(err)
	}
	return result, err
}
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...) // retorna un objeto rows
	if err != nil {
		log.Println(err)
	}
	return rows, err
}
func InsertData(query string, args ...interface{}) (int64, error) {
	if result, err := Exec(query, args...); err != nil {
		return int64(0), err
	} else {
		id, err := result.LastInsertId()
		return id, err
	}
}

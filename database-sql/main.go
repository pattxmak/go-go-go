package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Cover struct {
	Id   int
	Name string
}

// var db *sql.DB
var db *sqlx.DB

func main() {

	var err error
	// db, err = sql.Open("sqlserver", "sqlserver://sa:p@ssw0rd@13.76.163.73?database=techcoach") // MSSQL
	db, err = sqlx.Open("mysql", "root:p@ssw0rd@tcp(13.76.163.73)/techcoach") // mysql
	if err != nil {
		panic(err)
	}

	// Create
	// cover := Cover{5, "test"}
	// err = AddCover(cover)
	// if err != nil {
	// 	panic(err)
	// }

	// Get
	covers, err := GetCovers()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, cover := range covers {
		fmt.Println(cover)
	}

	// Get by id
	// cover, err = GetCover(1)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(cover)

	// update
	cover := Cover{1, "test update"}
	err = UpdateCover(cover)
	if err != nil {
		panic(err)
	}

	// delete
	err = DeleteCover(2)
	if err != nil {
		panic(err)
	}


}

// Sql
func GetCovers() ([]Cover, error) {

	err := db.Ping()
	if err != nil {
		return nil, err
	}

	// if err = db.Ping(); err != nil {
	// 	panic(err)
	// }

	query := "select id, name from cover"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// read
	covers := []Cover{}
	for rows.Next() {
		cover := Cover{}
		err = rows.Scan(&cover.Id, &cover.Name)
		if err != nil {
			return nil, err
		}
		covers = append(covers, cover)
	}

	return covers, nil
}

func GetCover(id int) (*Cover, error) {

	err := db.Ping()
	if err != nil {
		return nil, err
	}

	// query := "select id, name from cover where id=@id" // MSSQL
	// row := db.QueryRow(query, sql.Named("id", id)) // MSSQL

	query := "select id, name from cover where id=?" // mysql
	row := db.QueryRow(query, id)                    // mysql

	cover := Cover{}
	err = row.Scan(&cover.Id, &cover.Name)
	if err != nil {
		return nil, err
	}
	return &cover, nil

}

func AddCover(cover Cover) error {

	// when work with transaction
	// using begin tran -> can rollback
	// if sure -> using commit for save
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "insert into cover (id, name) values (?, ?)"
	result, err := tx.Exec(query, cover.Id, cover.Name)
	if err != nil {
		return err
	}

	// query := "insert into cover (id, name) values (?, ?)"
	// result, err := db.Exec(query, cover.Id, cover.Name)
	// if err != nil {
	// 	return err
	// }

	affected, err := result.RowsAffected()
	if err != nil {
		// using rollback 
		tx.Rollback()
		return err
	}

	if affected <= 0 {
		return errors.New("cannot insert")
	}

	// if pass
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func UpdateCover(cover Cover) error {

	query := "update cover set name=? where id=?"
	result, err := db.Exec(query, cover.Name, cover.Id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("cannot update")
	}

	return nil
}

func DeleteCover(id int) error {

	query := "delete from cover where id=?"
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("cannot delete")
	}

	return nil
}

// SqlX
func GetCoversX() ([]Cover, error) {

	query := "select id, name from cover"
	covers := []Cover{}

	err := db.Select(&covers, query)
	if err != nil {
		return nil, err
	}

	return covers, nil

}

func GetCoverX(id int) (*Cover, error) {

	query := "select id, name from cover where id=?"
	cover := Cover{}
	err := db.Get(&cover, query, id)
	if err != nil {
		return nil, err
	}

	return &cover, nil

}

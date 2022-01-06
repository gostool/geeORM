package geeORM

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"xorm.io/xorm"
)

func TestDB(t *testing.T) {
	engine, _ := NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}

func TestXorm(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite3", "gee2.db")
	if err != nil {
		t.Fatal(err)
	}
	defer engine.Close()
}
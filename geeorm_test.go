package geeORM

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/database/gdb"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)


func TestDB(t *testing.T) {
	engine, err := NewEngine("sqlite3", "gee.db")
	if err != nil {
		t.Fatal(err)
	}
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

func TestGfOrm(t *testing.T) {
	t.Log(gdb.DefaultGroupName)
}

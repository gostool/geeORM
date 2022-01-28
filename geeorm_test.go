package geeORM

import (
	"geeORM/session"
	"reflect"
	"testing"

	"github.com/gogf/gf/v2/database/gdb"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestDB(t *testing.T) {
	engine, err := NewEngine("sqlite3", "gee.db")
	if err != nil {
		t.Fatal(err)
	}
	defer engine.Close()
	s := engine.NewSession().Model(&User{})
	//_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_ = s.DropTable()
	//_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
	/*
		result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
		count, _ := result.RowsAffected()
		fmt.Printf("Exec success, %d affected\n", count)
	*/
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

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite3", "gee.db")
	if err != nil {
		t.Fatal("failed to open db", err)
	}
	return engine
}

func TestEngine_Transaction(t *testing.T) {
	t.Run("rollback", func(t *testing.T) {
		transactionRollback(t)
	})
	t.Run("commit", func(t *testing.T) {
		transactionCommit(t)
	})
}

func transactionCommit(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return nil, err
	})
	if err == nil || s.HasTable() {
		t.Fatal("failed to rollback")
	}
}

func transactionRollback(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return nil, err
	})
	if err != nil {
		t.Fatal("failed to commit")
	}
	u := &User{}
	_ = s.First(u)
	if u.Name != "Tom" {
		t.Fatal("failed to commit")
	}
}

func TestEngine_Migrate(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text PRIMARY KEY);").Exec()
	_, _ = s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	engine.Migrate(&User{})

	rows, _ := s.Raw("SELECT * FROM User").QueryRows()
	columns, _ := rows.Columns()
	if !reflect.DeepEqual(columns, []string{"Name", "Age"}) {
		t.Fatal("Failed to migrate table User, got columns", columns)
	}
}

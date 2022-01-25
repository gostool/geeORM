package geeORM

import (
	"geeORM/log"
	"geeORM/session"
	"testing"
)

type Account struct {
	ID       int `geeorm:"PRIMARY KEY"`
	Password string
}

func (account *Account) BeforeInsert(s *session.Session) error {
	log.Info("before insert", account)
	account.ID += 1000
	return nil
}

func (account *Account) AfterQuery(s *session.Session) error {
	log.Info("after query", account)
	account.Password = "******"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	t.Helper()
	engine, err := NewEngine("sqlite3", "gee.db")
	if err != nil {
		t.Fatal(err)
	}
	s := engine.NewSession().Model(&Account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(&Account{1, "12345"}, &Account{2, "qwerty"})
	u := &Account{}
	err = s.First(u)
	if err != nil || u.ID != 1001 || u.Password != "******" {
		t.Fatal("Failed to call hooks after query, got", u)
	}
}

package gui_test

import (
	"testing"

	"github.com/walkerxiong/go-chatroom/gui"
	"github.com/walkerxiong/go-chatroom/session"
)

func TestRun(t *testing.T) {
	u := session.User{
		Name: "walker",
	}
	f := func(msg string) error {
		return nil
	}
	app, err := gui.NewApp(&u, f)
	if err != nil {
		t.Fatal(err)
	}
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}

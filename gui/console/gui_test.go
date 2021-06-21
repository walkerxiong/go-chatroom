package console_test

import (
	"testing"

	"github.com/walkerxiong/go-chatroom/gui/console"
	"github.com/walkerxiong/go-chatroom/session"
)

func TestRun(t *testing.T) {
	u := session.User{
		Name: "walker",
	}
	app, err := console.NewApp(&u)
	if err != nil {
		t.Fatal(err)
	}
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}

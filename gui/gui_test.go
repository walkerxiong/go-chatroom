package gui_test

import (
	"testing"

	"github.com/walkerxiong/go-chatroom/gui"
)

func TestRun(t *testing.T) {
	f := func(msg string) error {
		return nil
	}
	app, err := gui.NewApp(f)
	if err != nil {
		t.Fatal(err)
	}
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}

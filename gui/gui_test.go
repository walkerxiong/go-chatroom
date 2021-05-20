package gui_test

import (
	"testing"

	"github.com/walkerxiong/go-chatroom/gui"
)

func TestRun(t *testing.T) {
	app, err := gui.NewApp()
	if err != nil {
		t.Fatal(err)
	}
	app.Run()
}

package gui

import (
	"log"

	"github.com/jroimartin/gocui"
)

type SendMessageFunc func(username string, message string) error

type App struct {
	g           *gocui.Gui
	SendMsgFunc SendMessageFunc
}

func NewApp(f SendMessageFunc) (*App, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, err
	}
	g.Highlight = true
	g.SelFgColor = gocui.ColorRed

	maxX, maxY := g.Size()
	onlineWidget := NewOnlineListWidget(0, 0, maxX/4, 5*maxY/6)
	messageWidget := NewMessageBoxWidget(maxX/4, 0, 3*maxX/4, 5*maxY/6)
	inputWidget := NewInputWidget(0, 5*maxY/6, maxX-1)

	g.SetManager(onlineWidget, messageWidget, inputWidget)
	g.InputEsc = true
	g.Cursor = true
	app := &App{
		g:           g,
		SendMsgFunc: f,
	}
	return app, nil
}

func (app *App) Run() error {
	if err := app.g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := app.g.SetKeybinding(InputView, gocui.KeyEnter, gocui.ModNone, app.sendMsg); err != nil {
		log.Panicln(err)
	}

	if err := app.g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}

// TODO：收到信息
func (app *App) ReceiveMsg(username string, msg string) {

}

// TODO: 用户上线
func (app *App) SetUserOnline(username string) {

}

// TODO: 用户下线
func (app *App) SetUserOffline(username string) {

}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// TODO: implement this
func (app *App) sendMsg(g *gocui.Gui, v *gocui.View) error {
	// get view contents

	// call app sendMsgFunc

	// write to messge box widget

	// clear contents
	return nil
}

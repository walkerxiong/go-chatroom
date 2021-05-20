package gui

import (
	"strings"

	"github.com/jroimartin/gocui"
)

type SendMessageFunc func(message string) error

type App struct {
	g             *gocui.Gui
	SendMsgFunc   SendMessageFunc
	onlineWidget  *OnlineListWidget
	messageWidget *MessageBoxWidget
	inputWidget   *InputWidget
}

func NewApp(f SendMessageFunc) (*App, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, err
	}
	g.Highlight = true
	g.SelFgColor = gocui.ColorRed

	maxX, maxY := g.Size()
	app := &App{
		g:             g,
		SendMsgFunc:   f,
		onlineWidget:  NewOnlineListWidget(0, 0, maxX/4, 5*maxY/6),
		messageWidget: NewMessageBoxWidget(maxX/4, 0, 3*maxX/4, 5*maxY/6),
		inputWidget:   NewInputWidget(0, 5*maxY/6, maxX-1),
	}

	g.SetManager(app.onlineWidget, app.messageWidget, app.inputWidget)
	g.InputEsc = true
	g.Cursor = true

	return app, nil
}

func (app *App) Run() error {
	if err := app.g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := app.g.SetKeybinding(InputView, gocui.KeyEnter, gocui.ModNone, app.sendMsg); err != nil {
		return err
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
	contents := app.inputWidget.view.Buffer()
	// call app sendMsgFunc
	app.SendMsgFunc(contents)
	// write space to assign
	size := app.messageWidget.w - len(contents)
	if size > 0 {
		contents = strings.Repeat(" ", size) + contents
	}
	// write to messge box widget
	app.messageWidget.Msgs = append(app.messageWidget.Msgs, contents)
	// clear contents
	app.inputWidget.view.Clear()
	app.inputWidget.view.SetCursor(0, 0)
	return nil
}

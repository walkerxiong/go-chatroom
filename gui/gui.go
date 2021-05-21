package gui

import (
	"github.com/jroimartin/gocui"
	"github.com/walkerxiong/go-chatroom/session"
)

type SendMessageFunc func(message string) error

type App struct {
	g             *gocui.Gui
	SendMsgFunc   SendMessageFunc
	onlineWidget  *OnlineListWidget
	messageWidget *MessageBoxWidget
	inputWidget   *InputWidget
	user          *session.User
}

func NewApp(u *session.User, f SendMessageFunc) (*App, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, err
	}
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	maxX, maxY := g.Size()
	app := &App{
		g:             g,
		SendMsgFunc:   f,
		onlineWidget:  NewOnlineListWidget(0, 0, maxX/4, 5*maxY/6),
		messageWidget: NewMessageBoxWidget(maxX/4, 0, 3*maxX/4, 5*maxY/6, u),
		inputWidget:   NewInputWidget(0, 5*maxY/6, maxX-1),
		user:          u,
	}
	app.onlineWidget.JoinUser(u.Name)
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

func (app *App) ReceiveMsg(msg session.Message) {
	app.messageWidget.ReceiveMsg(msg)
}

func (app *App) SetUserOnline(username string) {
	app.onlineWidget.JoinUser(username)
}

func (app *App) SetUserOffline(username string) {
	app.onlineWidget.Offline(username)
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (app *App) sendMsg(g *gocui.Gui, v *gocui.View) error {
	// get view contents
	contents := app.inputWidget.view.Buffer()
	if contents == "" {
		return nil
	}
	// call app sendMsgFunc
	app.SendMsgFunc(contents)
	// write to messge box widget
	app.messageWidget.SendMessage(contents)
	// clear contents
	app.inputWidget.view.Clear()
	app.inputWidget.view.SetCursor(0, 0)
	return nil
}

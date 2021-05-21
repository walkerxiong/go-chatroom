package gui

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/walkerxiong/go-chatroom/session"
)

/**
#   online    #   message window  #
#			  #					  #
#			  #					  #
#			  #					  #
#			  #					  #
#			  #					  #
#			  #					  #
###################################
			input box
###################################
**/

const (
	OnlineView  = "online"
	MessageView = "message"
	InputView   = "input"
)

type OnlineListWidget struct {
	x, y  int
	w, h  int
	name  string
	users map[string]struct{}
	view  *gocui.View
}

func NewOnlineListWidget(x, y, w, h int) *OnlineListWidget {
	return &OnlineListWidget{
		x:     x,
		y:     y,
		w:     w,
		h:     h,
		name:  OnlineView,
		users: make(map[string]struct{}),
	}
}

func (online *OnlineListWidget) JoinUser(name string) {
	online.users[name] = struct{}{}
}

func (online *OnlineListWidget) Offline(name string) {
	delete(online.users, name)
}

func (online *OnlineListWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(online.name, online.x, online.y, online.x+online.w, online.y+online.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	v.Clear()
	for user := range online.users {
		fmt.Fprintf(v, "\033[32;1m - %s\033[0m\n", user)
	}
	v.Title = fmt.Sprintf("online users (%d)", len(online.users))
	online.view = v
	return nil
}

type MessageBoxWidget struct {
	x, y     int
	w, h     int
	name     string
	view     *gocui.View
	Msgs     []session.Message
	currUser *session.User
}

func NewMessageBoxWidget(x, y, w, h int, u *session.User) *MessageBoxWidget {
	return &MessageBoxWidget{
		x:        x,
		y:        y,
		w:        w,
		h:        h,
		name:     MessageView,
		currUser: u,
	}
}

func (msg *MessageBoxWidget) SendMessage(content string) {
	msg.Msgs = append(msg.Msgs, session.Message{From: msg.currUser, Content: content})
}

func (msg *MessageBoxWidget) ReceiveMsg(message session.Message) {
	msg.Msgs = append(msg.Msgs, message)
}

// TODO: show latest message or watch all messages by cursor
func (msg *MessageBoxWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(msg.name, msg.x, msg.y, msg.x+msg.w, msg.y+msg.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	v.Clear()
	for _, m := range msg.Msgs {
		// from current session
		// write space to assign
		var content = m.Content
		if m.From == msg.currUser {
			size := msg.w - len(m.Content)
			if size < 0 {
				size = 0
			}
			var name = m.From.Name
			nsize := msg.w - len(name) - 2
			if nsize > 0 {
				name = strings.Repeat(" ", nsize) + name
			}
			fmt.Fprintf(v, "\033[40m%s:\033[0m\n%s\033[43;1m%s\033[0m\n", name, strings.Repeat(" ", size), content)
		} else {
			fmt.Fprintf(v, "\033[40m%s:\033[0m\n\033[42;1m%s\033[0m\n", m.From.Name, content)
		}
	}
	v.Title = "Messages"
	msg.view = v
	return nil
}

type InputWidget struct {
	x, y   int
	w      int
	name   string
	editor InputEditor
	view   *gocui.View
}

func NewInputWidget(x, y, w int) *InputWidget {
	return &InputWidget{
		x:      x,
		y:      y,
		w:      w,
		name:   InputView,
		editor: InputEditor{gocui.DefaultEditor},
	}
}

func (input *InputWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(input.name, input.x, input.y, input.x+input.w, input.y+3)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = true
		v.Wrap = true
		g.SetCurrentView(input.name)
	}
	v.Title = "Input"
	input.view = v
	return nil
}

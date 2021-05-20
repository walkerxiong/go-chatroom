package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
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
	x, y    int
	w, h    int
	name    string
	friends []string
	view    *gocui.View
}

func NewOnlineListWidget(x, y, w, h int) *OnlineListWidget {
	return &OnlineListWidget{
		x:    x,
		y:    y,
		w:    w,
		h:    h,
		name: OnlineView,
	}
}

func (online *OnlineListWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(online.name, online.x, online.y, online.x+online.w, online.y+online.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		for _, name := range online.friends {
			fmt.Fprintln(v, name)
		}
	}
	v.Title = "Online users"
	online.view = v
	return nil
}

type MessageBoxWidget struct {
	x, y int
	w, h int
	name string
	view *gocui.View
	Msgs []string
}

func NewMessageBoxWidget(x, y, w, h int) *MessageBoxWidget {
	return &MessageBoxWidget{
		x:    x,
		y:    y,
		w:    w,
		h:    h,
		name: MessageView,
	}
}

// TODO: define message struct for diffrent color font
func (msg *MessageBoxWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(msg.name, msg.x, msg.y, msg.x+msg.w, msg.y+msg.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	v.Clear()
	for _, m := range msg.Msgs {
		fmt.Fprintf(v, "\033[30;1m%s\033[0m\n", m)
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

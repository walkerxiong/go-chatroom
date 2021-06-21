package console

import "github.com/jroimartin/gocui"

type InputEditor struct {
	editor gocui.Editor
}

func (i InputEditor) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case (ch != 0 || key == gocui.KeySpace) && mod == 0:
		i.editor.Edit(v, key, ch, mod)
		// At the end of the line the default gcui editor adds a whitespace
		// Force him to remove
		ox, _ := v.Cursor()
		if ox > 1 && ox >= len(v.Buffer())-2 {
			v.EditDelete(false)
		}
		return
	case key == gocui.KeyEnter:
		return
	case key == gocui.KeyArrowRight:
		ox, _ := v.Cursor()
		if ox >= len(v.Buffer())-1 {
			return
		}
	case key == gocui.KeyHome || key == gocui.KeyArrowUp:
		v.SetCursor(0, 0)
		v.SetOrigin(0, 0)
		return
	case key == gocui.KeyEnd || key == gocui.KeyArrowDown:
		width, _ := v.Size()
		lineWidth := len(v.Buffer()) - 1
		if lineWidth > width {
			v.SetOrigin(lineWidth-width, 0)
			lineWidth = width - 1
		}
		v.SetCursor(lineWidth, 0)
		return
	}
	i.editor.Edit(v, key, ch, mod)
}

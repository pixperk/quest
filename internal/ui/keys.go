package ui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up              key.Binding
	Down            key.Binding
	Left            key.Binding
	Right           key.Binding
	Enter           key.Binding
	Tab             key.Binding
	Quit            key.Binding
	Send            key.Binding
	Help            key.Binding
	NextTab         key.Binding
	PrevTab         key.Binding
	AddHeader       key.Binding
	ClearHeaders    key.Binding
	SaveRequest     key.Binding
	LoadRequest     key.Binding
	NextResponseTab key.Binding
	PrevResponseTab key.Binding
	NextFocus       key.Binding
	PrevFocus       key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.NextTab, k.Send, k.SaveRequest, k.LoadRequest, k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right},
		{k.Tab, k.NextTab, k.PrevTab},
		{k.NextFocus, k.PrevFocus},
		{k.NextResponseTab, k.PrevResponseTab},
		{k.Send, k.AddHeader, k.ClearHeaders, k.Enter},
		{k.SaveRequest, k.LoadRequest},
		{k.Help, k.Quit},
	}
}

var DefaultKeys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "confirm"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next field"),
	),
	NextTab: key.NewBinding(
		key.WithKeys("ctrl+right", "ctrl+l"),
		key.WithHelp("ctrl+→", "next tab"),
	),
	PrevTab: key.NewBinding(
		key.WithKeys("ctrl+left", "ctrl+h"),
		key.WithHelp("ctrl+←", "prev tab"),
	),
	Send: key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "send request"),
	),
	AddHeader: key.NewBinding(
		key.WithKeys("ctrl+a"),
		key.WithHelp("ctrl+a", "add header"),
	),
	ClearHeaders: key.NewBinding(
		key.WithKeys("ctrl+x"),
		key.WithHelp("ctrl+x", "clear headers"),
	),
	SaveRequest: key.NewBinding(
		key.WithKeys("ctrl+w"),
		key.WithHelp("ctrl+w", "save request"),
	),
	LoadRequest: key.NewBinding(
		key.WithKeys("ctrl+r"),
		key.WithHelp("ctrl+r", "load request"),
	),
	NextResponseTab: key.NewBinding(
		key.WithKeys("shift+right", "shift+l"),
		key.WithHelp("shift+→", "next response tab"),
	),
	PrevResponseTab: key.NewBinding(
		key.WithKeys("shift+left", "shift+h"),
		key.WithHelp("shift+←", "prev response tab"),
	),
	NextFocus: key.NewBinding(
		key.WithKeys("alt+right", "alt+l"),
		key.WithHelp("alt+→", "next focus"),
	),
	PrevFocus: key.NewBinding(
		key.WithKeys("alt+left", "alt+h"),
		key.WithHelp("alt+←", "prev focus"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

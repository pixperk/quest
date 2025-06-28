package ui

import (
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/pixperk/quest/internal/styles"
)

type HTTPMethod struct {
	Name string
	Desc string
}

func (m HTTPMethod) FilterValue() string {
	return m.Name
}

func (m HTTPMethod) Title() string {
	return styles.StyledMethod(m.Name)
}

func (m HTTPMethod) Description() string {
	return m.Desc
}

type MethodDelegate struct{}

func NewMethodDelegate() MethodDelegate {
	return MethodDelegate{}
}

func (d MethodDelegate) Height() int {
	return 1
}

func (d MethodDelegate) Spacing() int {
	return 0
}

func (d MethodDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d MethodDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	method := item.(HTTPMethod)

	var style lipgloss.Style
	if index == m.Index() {
		// Selected item style
		style = lipgloss.NewStyle().
			Foreground(styles.White).
			Background(styles.HotPink).
			Padding(0, 1).
			Bold(true)
	} else {
		// Normal item style
		style = lipgloss.NewStyle().
			Foreground(styles.MethodColors[method.Name]).
			Padding(0, 1)
	}

	w.Write([]byte(style.Render(method.Name)))
}

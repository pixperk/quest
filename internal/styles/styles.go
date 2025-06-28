package styles

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	HotPink   = lipgloss.Color("#FF06B7")
	DarkGray  = lipgloss.Color("#767676")
	LightGray = lipgloss.Color("#C4C4C4")
	Green     = lipgloss.Color("#04B575")
	Red       = lipgloss.Color("#FF4757")
	Blue      = lipgloss.Color("#3742FA")
	Yellow    = lipgloss.Color("#FFA502")
	Purple    = lipgloss.Color("#9C88FF")
	Orange    = lipgloss.Color("#FF7675")
	White     = lipgloss.Color("#FFFFFF")
	Black     = lipgloss.Color("#000000")
)

var (
	TitleStyle = lipgloss.NewStyle().
			Foreground(HotPink).
			Bold(true).
			Padding(0, 1)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(DarkGray).
			Italic(true)

	StatusStyle = lipgloss.NewStyle().
			Foreground(Green).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(Red).
			Bold(true)

	InfoStyle = lipgloss.NewStyle().
			Foreground(Blue)

	JsonStyle = lipgloss.NewStyle().
			Foreground(LightGray)

	HelpStyle = lipgloss.NewStyle().
			Foreground(DarkGray).
			Italic(true)
)

var (
	HeaderStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(HotPink).
			Padding(0, 1)

	FocusedStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(HotPink).
			Padding(0, 1)

	BlurredStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(DarkGray).
			Padding(0, 1)

	ResponseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(Blue).
			Padding(1, 2)
)

var (
	TabStyle = lipgloss.NewStyle().
			Padding(0, 1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(DarkGray)

	ActiveTabStyle = lipgloss.NewStyle().
			Padding(0, 1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(HotPink).
			Background(HotPink).
			Foreground(White).
			Bold(true)
)

var MethodColors = map[string]lipgloss.Color{
	"GET":     Green,
	"POST":    Blue,
	"PUT":     Yellow,
	"DELETE":  Red,
	"PATCH":   Purple,
	"HEAD":    LightGray,
	"OPTIONS": DarkGray,
}

func StatusCodeColor(code int) lipgloss.Color {
	switch {
	case code >= 200 && code < 300:
		return Green
	case code >= 300 && code < 400:
		return Yellow
	case code >= 400 && code < 500:
		return Orange
	case code >= 500:
		return Red
	default:
		return DarkGray
	}
}

func StyledMethod(method string) string {
	color, exists := MethodColors[method]
	if !exists {
		color = DarkGray
	}
	return lipgloss.NewStyle().Foreground(color).Bold(true).Render(method)
}

func StyledStatusCode(code int) string {
	color := StatusCodeColor(code)
	return lipgloss.NewStyle().Foreground(color).Bold(true).Render(fmt.Sprintf("%d", code))
}

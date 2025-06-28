package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"

	"github.com/pixperk/quest/internal/http"
	"github.com/pixperk/quest/internal/syntax"
)

type Tab int

const (
	URLTab Tab = iota
	HeadersTab
	BodyTab
	ResponseTab
	LoadRequestTab
)

type ResponseSubTab int

const (
	ResponseBodySubTab ResponseSubTab = iota
	ResponseHeadersSubTab
)

type SavedRequest struct {
	Name    string            `json:"name"`
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

func (r SavedRequest) FilterValue() string {
	return r.Name
}

func (r SavedRequest) Title() string {
	return r.Name
}

func (r SavedRequest) Description() string {
	return fmt.Sprintf("%s • %s", r.Method, r.URL)
}

type Model struct {
	width  int
	height int

	urlInput         textinput.Model
	methodList       list.Model
	requestList      list.Model
	headerKey        textinput.Model
	headerValue      textinput.Model
	bodyTextarea     textarea.Model
	responseViewport viewport.Model
	headersViewport  viewport.Model
	help             help.Model
	spinner          spinner.Model
	highlighter      *syntax.Highlighter

	activeTab           Tab
	responseSubTab      ResponseSubTab
	focused             int
	loading             bool
	response            string
	responseContentType string
	statusCode          int
	responseTime        time.Duration
	responseHeaders     map[string]string
	requestHeaders      map[string]string
	httpClient          *http.Client
	showingLoadDialog   bool
	savedRequests       []SavedRequest

	keys KeyMap
}

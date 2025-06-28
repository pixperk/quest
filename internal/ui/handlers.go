package ui

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/pixperk/quest/internal/http"
	"github.com/pixperk/quest/internal/styles"
	"github.com/pixperk/quest/internal/syntax"
)

func NewModel() Model {
	urlInput := textinput.New()
	urlInput.Placeholder = "https://api.example.com/endpoint"
	urlInput.Focus()
	urlInput.CharLimit = 500
	urlInput.Width = 50

	headerKey := textinput.New()
	headerKey.Placeholder = "Header Name (e.g., Authorization)"
	headerKey.Width = 25

	headerValue := textinput.New()
	headerValue.Placeholder = "Header Value (e.g., Bearer token123)"
	headerValue.Width = 35

	bodyTextarea := textarea.New()
	bodyTextarea.Placeholder = "Request body (JSON, XML, text, etc.)\n\nExample:\n{\n  \"name\": \"John Doe\",\n  \"email\": \"john@example.com\"\n}"
	bodyTextarea.SetWidth(60)
	bodyTextarea.SetHeight(10)

	methods := []list.Item{
		HTTPMethod{Name: "GET", Desc: "Retrieve data"},
		HTTPMethod{Name: "POST", Desc: "Create new resource"},
		HTTPMethod{Name: "PUT", Desc: "Update/replace resource"},
		HTTPMethod{Name: "DELETE", Desc: "Remove resource"},
		HTTPMethod{Name: "PATCH", Desc: "Partial update"},
		HTTPMethod{Name: "HEAD", Desc: "Headers only"},
		HTTPMethod{Name: "OPTIONS", Desc: "Available methods"},
	}

	methodList := list.New(methods, NewMethodDelegate(), 15, 7)
	methodList.Title = "HTTP Method"
	methodList.SetShowStatusBar(false)
	methodList.SetFilteringEnabled(false)
	methodList.SetShowHelp(false)

	requestList := list.New([]list.Item{}, list.NewDefaultDelegate(), 60, 15)
	requestList.Title = "Saved Requests"
	requestList.SetShowStatusBar(false)
	requestList.SetFilteringEnabled(true)
	requestList.SetShowHelp(false)

	viewport := viewport.New(60, 15)

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(styles.HotPink)

	help := help.New()
	help.ShowAll = false

	return Model{
		urlInput:          urlInput,
		methodList:        methodList,
		requestList:       requestList,
		headerKey:         headerKey,
		headerValue:       headerValue,
		bodyTextarea:      bodyTextarea,
		responseViewport:  viewport,
		headersViewport:   viewport,
		help:              help,
		spinner:           s,
		highlighter:       syntax.NewHighlighter(),
		activeTab:         URLTab,
		responseSubTab:    ResponseBodySubTab,
		focused:           0,
		keys:              DefaultKeys,
		requestHeaders:    make(map[string]string),
		responseHeaders:   make(map[string]string),
		savedRequests:     make([]SavedRequest, 0),
		httpClient:        http.NewClient(),
		showingLoadDialog: false,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,
		textarea.Blink,
		m.spinner.Tick,
	)
}

func (m *Model) updateSizes() {
	m.urlInput.Width = m.width - 30
	m.methodList.SetSize(15, 10)
	m.requestList.SetSize(m.width-10, m.height-15)
	m.headerKey.Width = (m.width - 35) / 2
	m.headerValue.Width = (m.width - 35) / 2
	m.bodyTextarea.SetWidth(m.width - 10)
	m.bodyTextarea.SetHeight(m.height - 25)
	m.responseViewport.Width = m.width - 6
	m.responseViewport.Height = m.height - 25
	m.headersViewport.Width = m.width - 6
	m.headersViewport.Height = m.height - 25
}

func (m *Model) updateFocus() {
	m.urlInput.Blur()
	m.headerKey.Blur()
	m.headerValue.Blur()
	m.bodyTextarea.Blur()

	switch m.activeTab {
	case URLTab:
		if m.focused == 0 {
			m.urlInput.Focus()
		}
	case HeadersTab:
		if m.focused == 0 {
			m.headerKey.Focus()
		} else {
			m.headerValue.Focus()
		}
	case BodyTab:
		m.bodyTextarea.Focus()
	}
}

func (m Model) getSelectedMethod() string {
	if selected := m.methodList.SelectedItem(); selected != nil {
		return selected.(HTTPMethod).Name
	}
	return "GET"
}

func (m Model) sendRequest() (Model, tea.Cmd) {
	m.loading = true
	m.activeTab = ResponseTab

	req := http.Request{
		Method:  m.getSelectedMethod(),
		URL:     m.urlInput.Value(),
		Headers: m.requestHeaders,
		Body:    m.bodyTextarea.Value(),
	}

	return m, tea.Batch(
		m.spinner.Tick,
		func() tea.Msg {
			resp := m.httpClient.SendRequest(req)
			return ResponseMessage{
				StatusCode:   resp.StatusCode,
				Headers:      resp.Headers,
				Body:         resp.Body,
				ContentType:  resp.ContentType,
				ResponseTime: resp.ResponseTime,
				Error:        resp.Error,
			}
		},
	)
}

func (m Model) saveCurrentRequest() (Model, tea.Cmd) {
	if m.urlInput.Value() == "" {
		return m, nil
	}

	request := SavedRequest{
		Name:    fmt.Sprintf("%s %s", m.getSelectedMethod(), m.urlInput.Value()),
		Method:  m.getSelectedMethod(),
		URL:     m.urlInput.Value(),
		Headers: make(map[string]string),
		Body:    m.bodyTextarea.Value(),
	}

	for k, v := range m.requestHeaders {
		request.Headers[k] = v
	}

	requests := m.loadRequestsFromFile()
	requests = append(requests, request)

	if err := m.saveRequestsToFile(requests); err != nil {
		return m, nil
	}

	return m, nil
}

func (m Model) loadRequestsFromFile() []SavedRequest {
	data, err := os.ReadFile(".quest")
	if err != nil {
		return []SavedRequest{}
	}

	var requests []SavedRequest
	if err := json.Unmarshal(data, &requests); err != nil {
		return []SavedRequest{}
	}

	return requests
}

func (m Model) saveRequestsToFile(requests []SavedRequest) error {
	data, err := json.MarshalIndent(requests, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(".quest", data, 0644)
}

func (m Model) showLoadRequestDialog() (Model, tea.Cmd) {
	savedRequests := m.loadRequestsFromFile()

	items := make([]list.Item, len(savedRequests))
	for i, req := range savedRequests {
		items[i] = req
	}

	m.requestList.SetItems(items)
	m.savedRequests = savedRequests
	m.activeTab = LoadRequestTab
	m.showingLoadDialog = true
	m.focused = 0

	return m, nil
}

func (m Model) loadSelectedRequest(request SavedRequest) (Model, tea.Cmd) {
	m.urlInput.SetValue(request.URL)
	m.bodyTextarea.SetValue(request.Body)

	for i, item := range m.methodList.Items() {
		if method, ok := item.(HTTPMethod); ok && method.Name == request.Method {
			m.methodList.Select(i)
			break
		}
	}

	m.requestHeaders = make(map[string]string)
	for k, v := range request.Headers {
		m.requestHeaders[k] = v
	}

	m.activeTab = URLTab
	m.showingLoadDialog = false
	m.focused = 0
	m.updateFocus()

	return m, nil
}

package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/pixperk/quest/internal/styles"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.updateSizes()
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Send):
			if !m.loading && m.urlInput.Value() != "" {
				return m.sendRequest()
			}

		case key.Matches(msg, m.keys.NextTab):
			if m.activeTab == LoadRequestTab {
				m.activeTab = URLTab
				m.showingLoadDialog = false
			} else {
				m.activeTab = Tab((int(m.activeTab) + 1) % 4)
			}
			m.focused = 0
			m.updateFocus()

		case key.Matches(msg, m.keys.PrevTab):
			if m.activeTab == LoadRequestTab {
				m.activeTab = ResponseTab
				m.showingLoadDialog = false
			} else {
				m.activeTab = Tab((int(m.activeTab) + 3) % 4)
			}
			m.focused = 0
			m.updateFocus()

		case key.Matches(msg, m.keys.NextFocus):
			switch m.activeTab {
			case URLTab:
				m.focused = (m.focused + 1) % 2
			case HeadersTab:
				m.focused = (m.focused + 1) % 2
			}
			m.updateFocus()

		case key.Matches(msg, m.keys.PrevFocus):
			switch m.activeTab {
			case URLTab:
				m.focused = (m.focused + 1) % 2
			case HeadersTab:
				m.focused = (m.focused + 1) % 2
			}
			m.updateFocus()

		case key.Matches(msg, m.keys.Tab):
			switch m.activeTab {
			case URLTab:
				m.focused = (m.focused + 1) % 2
			case HeadersTab:
				m.focused = (m.focused + 1) % 2
			}
			m.updateFocus()

		// Handle arrow keys for method selection when focused
		case msg.String() == "left":
			if m.activeTab == URLTab && m.focused == 1 {
				currentIndex := 0
				for i, item := range m.methodList.Items() {
					if method, ok := item.(HTTPMethod); ok && method.Name == m.getSelectedMethod() {
						currentIndex = i
						break
					}
				}
				newIndex := (currentIndex - 1 + len(m.methodList.Items())) % len(m.methodList.Items())
				m.methodList.Select(newIndex)
			}

		case msg.String() == "right":
			if m.activeTab == URLTab && m.focused == 1 {
				currentIndex := 0
				for i, item := range m.methodList.Items() {
					if method, ok := item.(HTTPMethod); ok && method.Name == m.getSelectedMethod() {
						currentIndex = i
						break
					}
				}
				newIndex := (currentIndex + 1) % len(m.methodList.Items())
				m.methodList.Select(newIndex)
			}

		case key.Matches(msg, m.keys.AddHeader):
			if m.activeTab == HeadersTab && m.headerKey.Value() != "" && m.headerValue.Value() != "" {
				m.requestHeaders[m.headerKey.Value()] = m.headerValue.Value()
				m.headerKey.SetValue("")
				m.headerValue.SetValue("")
				m.headerKey.Focus()
				m.headerValue.Blur()
				m.focused = 0
			}

		case key.Matches(msg, m.keys.ClearHeaders):
			if m.activeTab == HeadersTab {
				m.requestHeaders = make(map[string]string)
			}

		case key.Matches(msg, m.keys.NextResponseTab):
			if m.activeTab == ResponseTab {
				m.responseSubTab = ResponseSubTab((int(m.responseSubTab) + 1) % 2)
			}

		case key.Matches(msg, m.keys.PrevResponseTab):
			if m.activeTab == ResponseTab {
				m.responseSubTab = ResponseSubTab((int(m.responseSubTab) + 1) % 2)
			}

		case key.Matches(msg, m.keys.SaveRequest):
			return m.saveCurrentRequest()

		case key.Matches(msg, m.keys.LoadRequest):
			return m.showLoadRequestDialog()

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case msg.String() == "esc":
			if m.activeTab == LoadRequestTab {
				m.activeTab = URLTab
				m.showingLoadDialog = false
				m.focused = 0
				m.updateFocus()
			}
		}

	case ResponseMessage:
		m.loading = false
		m.statusCode = msg.StatusCode
		m.responseTime = msg.ResponseTime
		m.responseHeaders = msg.Headers

		if msg.Error != nil {
			m.response = styles.ErrorStyle.Render("Error: " + msg.Error.Error())
		} else {
			m.response = msg.Body
		}

		m.responseViewport.SetContent(m.response)
		m.activeTab = ResponseTab
		return m, nil

	case spinner.TickMsg:
		if m.loading {
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	switch m.activeTab {
	case URLTab:
		if m.focused == 0 {
			m.urlInput, cmd = m.urlInput.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			// Don't pass left/right arrows to the method list when focused
			if keyMsg, ok := msg.(tea.KeyMsg); ok && (keyMsg.String() == "left" || keyMsg.String() == "right") {
				// Already handled above
			} else {
				m.methodList, cmd = m.methodList.Update(msg)
				cmds = append(cmds, cmd)
			}
		}
	case HeadersTab:
		if m.focused == 0 {
			m.headerKey, cmd = m.headerKey.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			m.headerValue, cmd = m.headerValue.Update(msg)
			cmds = append(cmds, cmd)
		}
	case BodyTab:
		m.bodyTextarea, cmd = m.bodyTextarea.Update(msg)
		cmds = append(cmds, cmd)
	case ResponseTab:
		if m.responseSubTab == ResponseBodySubTab {
			m.responseViewport, cmd = m.responseViewport.Update(msg)
		} else {
			m.headersViewport, cmd = m.headersViewport.Update(msg)
		}
		cmds = append(cmds, cmd)
	case LoadRequestTab:
		m.requestList, cmd = m.requestList.Update(msg)
		cmds = append(cmds, cmd)
		if keyMsg, ok := msg.(tea.KeyMsg); ok && key.Matches(keyMsg, m.keys.Enter) {
			if selected := m.requestList.SelectedItem(); selected != nil {
				return m.loadSelectedRequest(selected.(SavedRequest))
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return styles.InfoStyle.Render("Initializing Quest...")
	}

	header := m.renderHeader()
	tabs := m.renderTabs()

	var content string
	switch m.activeTab {
	case URLTab:
		content = m.renderURLTab()
	case HeadersTab:
		content = m.renderHeadersTab()
	case BodyTab:
		content = m.renderBodyTab()
	case ResponseTab:
		content = m.renderResponseTab()
	case LoadRequestTab:
		content = m.renderLoadRequestTab()
	}

	statusBar := m.renderStatusBar()
	helpView := m.help.View(m.keys)

	main := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		"",
		tabs,
		"",
		content,
		"",
		statusBar,
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		main,
		"",
		helpView,
	)
}

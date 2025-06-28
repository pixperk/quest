package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/pixperk/quest/internal/styles"
)

func (m Model) renderHeader() string {
	title := styles.TitleStyle.Render("Quest - Terminal HTTP Client")
	subtitle := styles.SubtitleStyle.Render("Beautiful API testing in your terminal")

	return lipgloss.JoinVertical(lipgloss.Left, title, subtitle)
}

// renderTabs renders the tab navigation
func (m Model) renderTabs() string {
	var tabs []string

	tabNames := []string{"URL", "Headers", "Body", "Response"}
	for i, name := range tabNames {
		if Tab(i) == m.activeTab {
			tabs = append(tabs, styles.ActiveTabStyle.Render(name))
		} else {
			tabs = append(tabs, styles.TabStyle.Render(name))
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, tabs...)
}

func (m Model) renderURLTab() string {
	// URL section
	urlSection := styles.HeaderStyle.Render("Request URL") + "\n"
	if m.focused == 0 {
		urlSection += styles.FocusedStyle.Render(m.urlInput.View())
	} else {
		urlSection += styles.BlurredStyle.Render(m.urlInput.View())
	}

	// HTTP methods in horizontal layout
	methodSection := styles.HeaderStyle.Render("HTTP Method") + "\n"

	var methodButtons []string
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"}
	selectedMethod := m.getSelectedMethod()

	for _, method := range methods {
		if method == selectedMethod && m.focused == 1 {
			methodButtons = append(methodButtons, styles.ActiveTabStyle.Render(" "+method+" "))
		} else if method == selectedMethod {
			methodButtons = append(methodButtons, styles.TabStyle.Render(" "+method+" "))
		} else {
			methodButtons = append(methodButtons, styles.BlurredStyle.Render(" "+method+" "))
		}
	}

	methodRow := lipgloss.JoinHorizontal(lipgloss.Left, methodButtons...)

	// Focus indicator
	var focusHelp string
	if m.focused == 1 {
		focusHelp = styles.HelpStyle.Render("Use ←/→ to select method")
	} else {
		focusHelp = styles.HelpStyle.Render("Alt+↓ or Tab to select method • Ctrl+R to load saved requests")
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		urlSection,
		"",
		methodSection,
		methodRow,
		"",
		focusHelp,
	)
}

// renderHeadersTab renders the headers management tab
func (m Model) renderHeadersTab() string {
	headerSection := styles.HeaderStyle.Render("Request Headers") + "\n"

	var keyInput, valueInput string
	if m.focused == 0 {
		keyInput = styles.FocusedStyle.Render(m.headerKey.View())
		valueInput = styles.BlurredStyle.Render(m.headerValue.View())
	} else {
		keyInput = styles.BlurredStyle.Render(m.headerKey.View())
		valueInput = styles.FocusedStyle.Render(m.headerValue.View())
	}

	inputs := lipgloss.JoinHorizontal(lipgloss.Left, keyInput, "  ", valueInput)
	addInstruction := styles.HelpStyle.Render("Ctrl+A: Add Header • Ctrl+X: Clear All • Tab: Switch Fields")

	var headersList []string
	for key, value := range m.requestHeaders {
		headerItem := styles.InfoStyle.Render(key) + ": " + styles.JsonStyle.Render(value)
		headersList = append(headersList, headerItem)
	}

	var existingHeaders string
	if len(headersList) > 0 {
		existingHeaders = "\n\n" + styles.HeaderStyle.Render("Active Headers:") + "\n" +
			strings.Join(headersList, "\n")
	} else {
		existingHeaders = "\n\n" + styles.HelpStyle.Render("No custom headers added yet")
	}

	return headerSection + inputs + "\n\n" + addInstruction + existingHeaders
}

// renderBodyTab renders the request body tab
func (m Model) renderBodyTab() string {
	bodySection := styles.HeaderStyle.Render("Request Body") + "\n"
	bodySection += styles.HelpStyle.Render("Enter your request body (JSON, XML, plain text, etc.)") + "\n\n"
	bodySection += styles.FocusedStyle.Render(m.bodyTextarea.View())

	return bodySection
}

// renderResponseTab renders the response display tab
func (m Model) renderResponseTab() string {
	responseSection := styles.HeaderStyle.Render("Response") + "\n"

	if m.loading {
		return responseSection + m.spinner.View() + " " +
			styles.InfoStyle.Render("Sending request...")
	}

	if m.response == "" {
		return responseSection + styles.HelpStyle.Render("No response yet. Send a request to see the response here.")
	}

	// Response sub-tabs
	responseTabs := m.renderResponseSubTabs() + "\n\n"

	// Content based on active response sub-tab
	var content string
	switch m.responseSubTab {
	case ResponseBodySubTab:
		content = m.renderResponseBody()
	case ResponseHeadersSubTab:
		content = m.renderResponseHeaders()
	}

	return responseSection + responseTabs + content
}

// renderResponseSubTabs renders the sub-tabs within the response tab
func (m Model) renderResponseSubTabs() string {
	var tabs []string

	subTabNames := []string{"Body", "Headers"}
	for i, name := range subTabNames {
		if ResponseSubTab(i) == m.responseSubTab {
			tabs = append(tabs, styles.ActiveTabStyle.Render(name))
		} else {
			tabs = append(tabs, styles.TabStyle.Render(name))
		}
	}

	helpText := styles.HelpStyle.Render("Use Shift+←/→ to switch between response tabs")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(lipgloss.Left, tabs...),
		helpText,
	)
}

// renderResponseBody renders the response body content
func (m Model) renderResponseBody() string {
	return m.responseViewport.View()
}

// renderResponseHeaders renders the response headers content
func (m Model) renderResponseHeaders() string {
	if len(m.responseHeaders) == 0 {
		content := styles.HelpStyle.Render("No response headers available")
		m.headersViewport.SetContent(content)
	} else {
		var headerLines []string
		for key, value := range m.responseHeaders {
			headerLine := styles.InfoStyle.Render(key) + ": " + styles.JsonStyle.Render(value)
			headerLines = append(headerLines, headerLine)
		}
		content := strings.Join(headerLines, "\n")
		m.headersViewport.SetContent(content)
	}

	return m.headersViewport.View()
}

// renderStatusBar renders the status bar with response information
func (m Model) renderStatusBar() string {
	if m.statusCode == 0 {
		return ""
	}

	// Status code with color
	statusColor := styles.StatusCodeColor(m.statusCode)
	statusText := lipgloss.NewStyle().
		Foreground(statusColor).
		Bold(true).
		Render(fmt.Sprintf("Status: %d", m.statusCode))

	// Response time
	responseTimeText := styles.InfoStyle.Render(fmt.Sprintf("Time: %v", m.responseTime))

	// Method indicator
	methodText := styles.InfoStyle.Render(fmt.Sprintf("Method: %s", m.getSelectedMethod()))

	// URL indicator
	url := m.urlInput.Value()
	if len(url) > 50 {
		url = url[:47] + "..."
	}
	urlText := styles.HelpStyle.Render(fmt.Sprintf("URL: %s", url))

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		statusText,
		"  ",
		responseTimeText,
		"  ",
		methodText,
		"  ",
		urlText,
		"  ",
		styles.HelpStyle.Render("Ctrl+W: Save • Ctrl+R: Load"),
	)
}

// renderLoadRequestTab renders the saved requests loading dialog
func (m Model) renderLoadRequestTab() string {
	title := styles.HeaderStyle.Render("Load Saved Request")
	subtitle := styles.HelpStyle.Render("↑/↓: Navigate • Enter: Select • Esc: Cancel • /: Search")

	// Create a nice border around the list
	listStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(styles.Blue).
		Padding(1).
		Width(m.width - 10).
		Height(m.height - 15)

	listView := listStyle.Render(m.requestList.View())

	if len(m.savedRequests) == 0 {
		emptyStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(styles.Yellow).
			Padding(2).
			Width(m.width - 10).
			Align(lipgloss.Center)

		emptyMsg := emptyStyle.Render(
			"No saved requests found\n\n" +
				"Save a request with Ctrl+W first\n" +
				"Then use Ctrl+R to load it here",
		)

		return lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			"",
			subtitle,
			"",
			emptyMsg,
		)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		subtitle,
		"",
		listView,
	)
}

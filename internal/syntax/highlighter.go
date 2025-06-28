package syntax

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/pixperk/quest/internal/styles"
)

// Highlighter provides syntax highlighting for different content types
type Highlighter struct {
	jsonKeyStyle    lipgloss.Style
	jsonStringStyle lipgloss.Style
	jsonNumberStyle lipgloss.Style
	jsonBoolStyle   lipgloss.Style
	jsonNullStyle   lipgloss.Style
	htmlTagStyle    lipgloss.Style
	htmlAttrStyle   lipgloss.Style
	xmlTagStyle     lipgloss.Style
}

// NewHighlighter creates a new syntax highlighter
func NewHighlighter() *Highlighter {
	return &Highlighter{
		jsonKeyStyle:    lipgloss.NewStyle().Foreground(styles.Blue).Bold(true),
		jsonStringStyle: lipgloss.NewStyle().Foreground(styles.Green),
		jsonNumberStyle: lipgloss.NewStyle().Foreground(styles.Purple),
		jsonBoolStyle:   lipgloss.NewStyle().Foreground(styles.Orange),
		jsonNullStyle:   lipgloss.NewStyle().Foreground(styles.DarkGray),
		htmlTagStyle:    lipgloss.NewStyle().Foreground(styles.HotPink),
		htmlAttrStyle:   lipgloss.NewStyle().Foreground(styles.Blue),
		xmlTagStyle:     lipgloss.NewStyle().Foreground(styles.Purple),
	}
}

// Highlight applies syntax highlighting based on content type
func (h *Highlighter) Highlight(content, contentType string) string {
	if content == "" {
		return content
	}

	// Detect content type if not provided
	if contentType == "" {
		contentType = h.detectContentType(content)
	}

	switch {
	case strings.Contains(contentType, "json"):
		return h.highlightJSON(content)
	case strings.Contains(contentType, "html"):
		return h.highlightHTML(content)
	case strings.Contains(contentType, "xml"):
		return h.highlightXML(content)
	default:
		return content
	}
}

// detectContentType attempts to detect the content type
func (h *Highlighter) detectContentType(content string) string {
	trimmed := strings.TrimSpace(content)

	if (strings.HasPrefix(trimmed, "{") && strings.HasSuffix(trimmed, "}")) ||
		(strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]")) {
		// Try to parse as JSON
		var js json.RawMessage
		if json.Unmarshal([]byte(content), &js) == nil {
			return "application/json"
		}
	}

	if strings.HasPrefix(trimmed, "<") && strings.HasSuffix(trimmed, ">") {
		if strings.Contains(strings.ToLower(content), "<html") {
			return "text/html"
		}
		return "application/xml"
	}

	return "text/plain"
}

// highlightJSON applies JSON syntax highlighting
func (h *Highlighter) highlightJSON(content string) string {
	// First, try to format the JSON
	var obj interface{}
	if err := json.Unmarshal([]byte(content), &obj); err == nil {
		if formatted, err := json.MarshalIndent(obj, "", "  "); err == nil {
			content = string(formatted)
		}
	}

	// Apply syntax highlighting with regex patterns
	patterns := []struct {
		regex *regexp.Regexp
		style lipgloss.Style
	}{
		// JSON keys (quoted strings followed by colon)
		{regexp.MustCompile(`"([^"\\]|\\.)*":`), h.jsonKeyStyle},
		// JSON string values
		{regexp.MustCompile(`:\s*"([^"\\]|\\.)*"`), h.jsonStringStyle},
		// JSON numbers
		{regexp.MustCompile(`:\s*-?\d+\.?\d*([eE][+-]?\d+)?`), h.jsonNumberStyle},
		// JSON booleans
		{regexp.MustCompile(`:\s*(true|false)`), h.jsonBoolStyle},
		// JSON null
		{regexp.MustCompile(`:\s*null`), h.jsonNullStyle},
	}

	result := content
	for _, pattern := range patterns {
		result = pattern.regex.ReplaceAllStringFunc(result, func(match string) string {
			if strings.Contains(match, ":") {
				// For key-value pairs, only style the value part
				parts := strings.SplitN(match, ":", 2)
				if len(parts) == 2 {
					if strings.Contains(match, `"`) && !strings.Contains(parts[1], `"`) {
						// This is a key, style it differently
						return h.jsonKeyStyle.Render(parts[0]) + ":" + parts[1]
					}
					return parts[0] + ":" + pattern.style.Render(parts[1])
				}
			}
			return pattern.style.Render(match)
		})
	}

	return result
}

// highlightHTML applies basic HTML syntax highlighting
func (h *Highlighter) highlightHTML(content string) string {
	// HTML tags
	tagRegex := regexp.MustCompile(`<[^>]+>`)
	content = tagRegex.ReplaceAllStringFunc(content, func(match string) string {
		return h.htmlTagStyle.Render(match)
	})

	// HTML attributes (simplified)
	attrRegex := regexp.MustCompile(`(\w+)=("[^"]*"|'[^']*')`)
	content = attrRegex.ReplaceAllStringFunc(content, func(match string) string {
		return h.htmlAttrStyle.Render(match)
	})

	return content
}

// highlightXML applies basic XML syntax highlighting
func (h *Highlighter) highlightXML(content string) string {
	// XML tags
	tagRegex := regexp.MustCompile(`<[^>]+>`)
	content = tagRegex.ReplaceAllStringFunc(content, func(match string) string {
		return h.xmlTagStyle.Render(match)
	})

	return content
}

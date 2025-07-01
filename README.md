# 🚀 Quest - Terminal HTTP Client

A beautiful, modular terminal-based HTTP client built with Go, Bubble Tea, and Charm CLI. Think Postman, but in your terminal with a gorgeous TUI interface.

![Quest Screenshot](https://img.shields.io/badge/Built%20with-Go-00ADD8?style=for-the-badge&logo=go)
![Quest Screenshot](https://img.shields.io/badge/TUI-Bubble%20Tea-FF69B4?style=for-the-badge)
![Quest Screenshot](https://img.shields.io/badge/Style-Lipgloss-9C88FF?style=for-the-badge)

## ✨ Features

- 🎨 **Beautiful TUI Interface** - Powered by Charm's Bubble Tea and Lipgloss
- 🌐 **Full HTTP Support** - GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS
- 📊 **Real-time Response** - View responses with syntax highlighting
- ⚡ **Performance Metrics** - Response time and status code display
- 🎯 **Easy Navigation** - Keyboard-driven interface with tabs
- 📱 **Responsive Design** - Adapts to your terminal size
- 🔧 **Modular Architecture** - Clean, maintainable codebase
- 🎪 **Custom Headers** - Add and manage request headers
- 📝 **Request Body Support** - JSON, XML, plain text support
- 🌈 **Color-coded Methods** - Visual distinction for HTTP methods
- 💾 **Clean Error Handling** - Graceful error messages
- 📂 **Request Saving** - Save requests to .quest files
- 🔄 **Response Sub-tabs** - Separate views for headers and body
- 📜 **Scrollable Method List** - Better navigation through HTTP methods

## 🚀 Installation

### Prerequisites
- Go 1.23+ installed on your system
- A terminal that supports ANSI colors

### Simple install (via Go driectly)
To install **quest** simply run:
```bash
go install github.com/pixperk/quest@latest
```
This will install the binary into your GOBIN (linux/macOS: $GOBIN, windows: %GOBIN%) which in most cases already is in your path.

### Build from Source

1. Clone the repository:
```bash
git clone https://github.com/pixperk/quest.git
cd quest
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the application:
```bash
go build -o quest .
```

4. Run Quest:
```bash
./quest        # On Linux/macOS
quest.exe      # On Windows
```

### Quick Install (one-liner)

> [!INFO]  
> This will download the source code and build quest, not add it to your PATH

```bash
git clone https://github.com/pixperk/quest.git && cd quest && go build -o quest . && ./quest
```

## 🎮 Usage

### Interface Overview

Quest features a tabbed interface with four main sections:

1. **URL Tab** - Enter your API endpoint and select HTTP method (🌐 🚀)
2. **Headers Tab** - Add custom request headers (📋)  
3. **Body Tab** - Enter request body (for POST/PUT/PATCH) (📝)
4. **Response Tab** - View formatted response with headers (📊)

### Keyboard Shortcuts

#### Navigation
- **Ctrl+→** / **Ctrl+L** - Next tab
- **Ctrl+←** / **Ctrl+H** - Previous tab
- **Tab** / **Alt+→** / **Alt+L** - Next field within current tab
- **Alt+←** / **Alt+H** - Previous field within current tab
- **↑/↓** or **j/k** - Navigate through HTTP methods
- **Enter** - Select HTTP method

#### Actions
- **Ctrl+S** - Send the HTTP request
- **Ctrl+A** - Add header (in Headers tab)
- **Ctrl+X** - Clear all headers (in Headers tab)
- **Ctrl+W** - Save current request to .quest file
- **Ctrl+R** - Load saved request
- **Shift+←/→** - Switch between response sub-tabs (Headers/Body)
- **Esc** - Cancel load dialog
- **/** - Search saved requests (when in load dialog)
- **?** - Toggle help menu
- **q** or **Ctrl+C** - Quit the application

### Making Your First Request

1. **Enter URL**: Type your API endpoint (e.g., `https://jsonplaceholder.typicode.com/posts/1`)
2. **Select Method**: Press Tab to focus on methods, use ↑/↓ to select (default: GET)
3. **Add Headers** (optional): Navigate to Headers tab, enter key-value pairs
4. **Add Body** (optional): For POST/PUT/PATCH, navigate to Body tab and enter JSON/XML
5. **Send Request**: Press Ctrl+S to send the request
6. **View Response**: Automatically switches to Response tab with formatted output

### Example Requests

#### Simple GET Request
```
URL: https://api.github.com/users/octocat
Method: GET
```

#### POST with JSON Body
```
URL: https://jsonplaceholder.typicode.com/posts
Method: POST
Headers:
  Content-Type: application/json
Body:
{
  "title": "My Post",
  "body": "This is my post content",
  "userId": 1
}
```

#### Authenticated Request
```
URL: https://api.example.com/protected
Method: GET
Headers:
  Authorization: Bearer your-token-here
  Accept: application/json
```

## 🎨 Interface Walkthrough

### URL Tab
- **URL Input**: Enter your complete API endpoint with protocol (http/https)
- **Method Selection**: Choose from 7 HTTP methods with color coding (now scrollable with ↑/↓):
  - 🟢 **GET** - Retrieve data
  - 🔵 **POST** - Create new resource
  - 🟡 **PUT** - Update/replace resource
  - 🔴 **DELETE** - Remove resource
  - 🟣 **PATCH** - Partial update
  - ⚪ **HEAD** - Headers only


### Headers Tab
- Add custom headers like `Authorization`, `Content-Type`, etc.
- Press **Ctrl+A** to add a header after entering key and value
- View all added headers below the input fields
- Press **Ctrl+X** to clear all headers

### Body Tab
- Large text area for request body
- Supports JSON, XML, plain text, or any format
- Automatic Content-Type header for requests with body

### Response Tab
- **Response Sub-tabs**: Switch between Headers and Body views with Shift+←/→
- **Headers Sub-tab**: Clean display of all response headers
- **Body Sub-tab**: Formatted response body with JSON auto-formatting
- **Status Code** with color coding (green=2xx, yellow=3xx, orange=4xx, red=5xx)
- **Response Time** measurement

### Request Saving
- Press **Ctrl+W** to save the current request to a `.quest` file
- Requests are saved with URL, method, headers, and body
- Saved requests persist between sessions


### Built With
- **Go 1.23+** - Modern Go with latest features
- **Bubble Tea** - TUI framework for rich terminal applications  
- **Lipgloss** - Styling and layout for terminal UIs
- **Bubbles** - Common TUI components (textinput, viewport, etc.)

### Design Principles
- **Modular**: Clean separation of concerns
- **Responsive**: Adapts to different terminal sizes
- **Accessible**: Keyboard-driven interface
- **Beautiful**: Carefully crafted visual design
- **Fast**: Efficient rendering and HTTP client

## 🎯 Roadmap

### Upcoming Features
- [ ] Request history and recall
- [ ] Save/load request collections
- [ ] Environment variables and templating
- [ ] Multiple authentication methods (Basic, OAuth, API Key)
- [ ] Response export (JSON, text files)
- [ ] Custom themes and color schemes
- [ ] Request/response tabs for multiple concurrent requests
- [ ] cURL command export
- [ ] Request duration graphs
- [ ] Configuration file support

### Advanced Features
- [ ] GraphQL support
- [ ] WebSocket connections
- [ ] Request scripting with JavaScript
- [ ] Mock server integration
- [ ] CI/CD integration
- [ ] Team collaboration features

## 🤝 Contributing

Contributions are welcome! Here's how to get started:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests if applicable
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

### Development Setup
```bash
git clone https://github.com/pixperk/quest.git
cd quest
go mod tidy
go run main.go
```

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Charm CLI](https://charm.sh/) for the amazing TUI components
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) for the TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) for beautiful styling
- [Bubbles](https://github.com/charmbracelet/bubbles) for UI components

---

**Made with ❤️ and Go**

*Quest - Making API testing beautiful, one terminal at a time.*

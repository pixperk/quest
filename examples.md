# Quest Configuration Examples

## Environment Variables
```bash
# Set default timeout
export QUEST_TIMEOUT=30s

# Set default headers
export QUEST_DEFAULT_HEADERS="Authorization:Bearer token123,Accept:application/json"

# Enable debug mode
export DEBUG=1
```

## Sample Requests

### Loading Saved Requests
```
1. Press Ctrl+R to open the saved requests dialog
2. Use ↑/↓ to navigate through saved requests
3. Press Enter to load a selected request
4. Press Esc to cancel and go back
5. Use / to search through requests
```

### JSON API Request
```
URL: https://jsonplaceholder.typicode.com/posts/1
Method: GET
Headers:
  Accept: application/json
  User-Agent: Quest/1.0
```

### Create Resource
```
URL: https://jsonplaceholder.typicode.com/posts
Method: POST
Headers:
  Content-Type: application/json
Body:
{
  "title": "Quest Test Post",
  "body": "This is a test post created with Quest",
  "userId": 1
}
```

### Authentication Examples
```
# Bearer Token
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

# Basic Auth (base64 encoded username:password)
Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=

# API Key
X-API-Key: your-api-key-here
```

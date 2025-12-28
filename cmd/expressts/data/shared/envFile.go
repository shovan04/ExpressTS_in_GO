package shared

func GetENVFileContent() []byte {
	return []byte(`PORT=8000
HOST=localhost
BASE_API_PATH=/api/v1
BACKEND_URI=http://localhost:8000
COOKIE_SECRET=cookie_secret
`)
}

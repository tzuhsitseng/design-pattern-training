package main

import (
	"fmt"
	"net/http"
)

func main() {
	n := newNginx()
	appStatusURL := "/app/status"
	createUserURL := "/app/user"

	statusCode, respBody := n.handleRequest(appStatusURL, http.MethodGet)
	fmt.Println(appStatusURL, http.MethodGet, statusCode, respBody)

	statusCode, respBody = n.handleRequest(appStatusURL, http.MethodGet)
	fmt.Println(appStatusURL, http.MethodGet, statusCode, respBody)

	statusCode, respBody = n.handleRequest(appStatusURL, http.MethodGet)
	fmt.Println(appStatusURL, http.MethodGet, statusCode, respBody)

	statusCode, respBody = n.handleRequest(appStatusURL, http.MethodGet)
	fmt.Println(appStatusURL, http.MethodGet, statusCode, respBody)

	statusCode, respBody = n.handleRequest(createUserURL, http.MethodGet)
	fmt.Println(createUserURL, http.MethodGet, statusCode, respBody)

	statusCode, respBody = n.handleRequest(createUserURL, http.MethodPost)
	fmt.Println(createUserURL, http.MethodPost, statusCode, respBody)
}

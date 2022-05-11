package main

type application struct {
}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "OK"
	}
	if url == "/app/user" && method == "POST" {
		return 201, "CREATED"
	}

	return 404, "NOT FOUND"
}

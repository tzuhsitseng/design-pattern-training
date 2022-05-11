package main

type nginx struct {
	application *application
	rateLimiter map[string]int
	maxRequest  int
}

func (n *nginx) handleRequest(url, method string) (int, string) {
	if cnt := n.rateLimiter[url]; cnt >= n.maxRequest {
		return 403, "NOT ALLOWED"
	}

	n.rateLimiter[url]++
	return n.application.handleRequest(url, method)
}

func newNginx() *nginx {
	return &nginx{
		application: &application{},
		rateLimiter: map[string]int{},
		maxRequest:  3,
	}
}

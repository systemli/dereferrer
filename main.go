package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	listenAddr := ":8080"
	if os.Getenv("LISTEN_ADDR") != "" {
		listenAddr = os.Getenv("LISTEN_ADDR")
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(listenAddr, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.RawQuery != "" {
		decodedUrl, err := url.PathUnescape(r.URL.RawQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !isUrl(decodedUrl) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 Bad Request"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`<html lang="en">
		<head>
			<meta charset="utf-8">
			<meta http-equiv="refresh" content="0; URL=%s"/>
			<title>Redirecting to %s</title>
		</head>
		<body>
		<div align="center">
			<p>Redirecting to <a href="%s">%s</a></p>
		</div>
		</body>
		</html>`, decodedUrl, decodedUrl, decodedUrl, decodedUrl)))
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

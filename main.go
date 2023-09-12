package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	logLevel := "info"
	if os.Getenv("LOG_LEVEL") != "" {
		logLevel = os.Getenv("LOG_LEVEL")
	}

	atomic := zap.NewAtomicLevel()
	level, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		log.Fatal(err)
	}
	atomic.SetLevel(level)
	logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(os.Stdout),
		atomic,
	))
}

func main() {
	listenAddr := ":8080"
	if os.Getenv("LISTEN_ADDR") != "" {
		listenAddr = os.Getenv("LISTEN_ADDR")
	}

	logger.Info("Starting server", zap.String("listenAddr", listenAddr))
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.RawQuery != "" {
		decodedUrl, err := url.PathUnescape(r.URL.RawQuery)
		if err != nil {
			logger.Error("Error decoding URL", zap.Error(err), zap.String("query", r.URL.RawQuery))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !isUrl(decodedUrl) {
			logger.Error("Invalid URL", zap.String("query", r.URL.RawQuery), zap.String("decoded_url", decodedUrl))
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("400 Bad Request"))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf(`<html lang="en">
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
	_, _ = w.Write([]byte("404 Not Found"))
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

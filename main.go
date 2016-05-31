package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var port = flag.Int("port", 3000, "listen port")

func main() {
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", md5HashHandler)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", *port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func md5HashHandler(w http.ResponseWriter, r *http.Request) {
	h := md5.New()
	io.Copy(h, r.Body)
	println(hex.EncodeToString(h.Sum(nil)))
}

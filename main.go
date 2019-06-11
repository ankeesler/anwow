package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"
)

func wow(s string) string {
	b := []byte(s)
	bNew := make([]rune, len(b))

	for i := 0; i < len(b); i++ {
		if i%2 == 1 {
			bNew[i] = unicode.ToUpper(rune(b[i]))
		} else {
			bNew[i] = rune(b[i])
		}
	}

	sNew := string(bNew)

	log.Printf("wow: '%s' -> '%s'", s, sNew)

	return sNew
}

func handle(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("\"%s\"\n", err.Error())))
		return
	}
	defer r.Body.Close()

	path := r.URL.Path
	bodyString := string(bodyBytes)
	log.Printf("%s %s '%s'", r.Method, path, bodyString)

	if path == "/" {
		w.Write([]byte(fmt.Sprintf("\"%s\"\n", wow(bodyString))))
	} else {
		w.Write([]byte(fmt.Sprintf("\"%s\"\n", wow(strings.TrimSpace(strings.ReplaceAll(path, "/", " "))))))
	}
}

func main() {
	log.SetOutput(os.Stdout)

	port := "8080"
	if p, ok := os.LookupEnv("PORT"); ok {
		port = p
	}

	log.Println(
		http.ListenAndServe(
			fmt.Sprintf(":%s", port),
			http.HandlerFunc(handle),
		),
	)
}

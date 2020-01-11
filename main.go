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
	defer r.Body.Close()
	defer ioutil.ReadAll(r.Body)

	path := r.URL.Path
	contentType := r.Header.Get("Accept")
	log.Printf("%s %s Content-Type: %s", r.Method, path, contentType)

	text := wow(
		strings.TrimSpace(
			strings.ReplaceAll(path, "/", " "),
		),
	)

	if r.URL.Query().Get("clap") == "1" {
		text = strings.ReplaceAll(text, " ", " &#x1f44f ")
	}

	if strings.Contains(contentType, "text/html") {
		w.Header().Set("Content-Type", "text/html")
		w.Write(
			[]byte(
				fmt.Sprintf(
					`<div style="font-size:50px;margin:30px" align="center">%s</div>`,
					text,
				),
			),
		)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(
			[]byte(
				fmt.Sprintf(
					"\"%s\"\n",
					text,
				),
			),
		)
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

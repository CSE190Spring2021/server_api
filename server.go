// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"strconv"
	"bytes"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/gendata", gendataHandler)

	// for local testing
	//log.Fatal(http.ListenAndServe("localhost:8000", nil))
	// for web server
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

//!-handler

//!+randSeq
// this random string generation was inspired by:
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go/22892986#22892986
// https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
var characters = []rune("abcedfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*")

func randSeq(n int) string {
	var b bytes.Buffer
	for i := 0; i < n; i++ {
		b.WriteString(string(characters[rand.Intn(len(characters))]))
	}
	return b.String()
}

//!-randSeq

//!+gendataHandler
func gendataHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	queriesArray, received := r.URL.Query()["numBytes"]

	numBytesString := queriesArray[0]
	numBytes, err:= strconv.Atoi(numBytesString)
	if !received || len(queriesArray[0]) < 1 || err != nil {
		numBytes = 1
	}
	fmt.Fprintf(w, "%s", randSeq(numBytes))
}

//!-gendataHandler

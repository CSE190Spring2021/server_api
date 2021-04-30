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
	"strings"
)


// Struct for URL and trackers
type urlTrackerStruct struct {
	URL string
	Trackers [] string
}

// Parse the HTTP request to get out the URL and the trackers
func parseURLAndTrackers (httpRequest string) urlTrackerStruct {
	// To populate and return
	completedStruct := urlTrackerStruct{}
	var extractedURL string
	var extractedTrackers [] string
	// Start by parsing up to the first 
	forwardSlash := "/"
	for strings.Contains(httpRequest, forwardSlash) {
		slashIndex := strings.Index(httpRequest, forwardSlash)
		extractedURL = httpRequest[slashIndex:]
		break
	}
	// Parse the trackers, how do they come in? After/until what character?


	completedStruct.URL = extractedURL
	completedStruct.Trackers = extractedTrackers
	return completedStruct
}
/*
// Extract content type from mime.types file
func (hs *HttpServer) parseURI(requestHeader *HttpRequestHeader, conn net.Conn, uri string) (result string) {
	// pars the .html part off
	dotChar := "."
	// Parse file extension
	for strings.Contains(uri, dotChar) {
		dotIndex := strings.Index(uri, dotChar)
		result = uri[dotIndex:]
		break
	}
	// Here result = .extension, look it up in mime map
	result = hs.MIMEMap[result]
	return result
}
*/



func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/gendata", gendataHandler)
	http.HandleFunc("/url", urlHandler)


	// for local testing
	//log.Fatal(http.ListenAndServe("localhost:8000", nil))
	// for web server

	// Listens on port 0.0.0.0:8000 and checks
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	//log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
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

// We get sent URL and a list of trackers, 
// What type of data are we getting sent? 
// .r.url.query numbytes.
// http://localhost:8000/=10
// dcIJ!V1Lef

// Take the HTTP request from the extension
// Parse it to get the URL, trackers
// Session cookie
// Struct to store the url and trackers.
// Have a list of trackers, append to it 

func urlHandler(w http.ResponseWriter, r *http.Request) {

	rand.Seed(time.Now().UnixNano())
	queriesArray, received := r.URL.Query()["url"]

	numBytesString := queriesArray[0]
	numBytes, err:= strconv.Atoi(numBytesString)
	if !received || len(queriesArray[0]) < 1 || err != nil {
		numBytes = 1
	}
	fmt.Fprintf(w, "%s", randSeq(numBytes))

	

}

// 
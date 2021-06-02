package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
	"os/exec"
)

type Message struct {
	Addr string `json:"addr"`
	SafeStatus string `json:"safestatus"`
}

// curl localhost:8000 -d '{"name":"Hello"}'
func handler(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Message  
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return 
	}

	fmt.Println(msg.Addr)
	cmd := exec.Command("python3", "-W", "ignore", "./Model/model_script.py", "-u " + msg.Addr)
	out, err := cmd.CombinedOutput()
	if err!=nil {
		fmt.Println(err)
	}
	msg.SafeStatus = string(out)
	fmt.Println(out)
	
	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
	fmt.Println(msg.Addr)
}


func main() {

	http.HandleFunc("/", handler)
	er := http.ListenAndServe(":8080", nil)
	if er != nil {
		fmt.Println(er)
	}

}

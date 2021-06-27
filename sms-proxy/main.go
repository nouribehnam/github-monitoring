package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type test_struct struct {
	Message string
	Title   string
}

func sendNotify(rw http.ResponseWriter, request *http.Request) {
	var t test_struct
	decoder := json.NewDecoder(request.Body)

	numbers := fmt.Sprintf("%s", request.URL.Query().Get("to"))
	fmt.Println(numbers)

	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	client := http.Client{}

	req, err := http.NewRequest("GET", "http://api.payamak-panel.com/post/sendsms.ashx", nil)
	req.Header = http.Header{
		"Host":            []string{"api.payamak-panel.com"},
		"User-Agent":      []string{"golang"},
		"Accept":          []string{"*/*"},
		"Accept-Encoding": []string{"gzip, deflate, br"},
		"Connection":      []string{"keep-alive"},
	}
	q := req.URL.Query()
	q.Add("username", "09121400548")
	q.Add("password", "password")
	q.Add("to", numbers)
	q.Add("from", "30008666003003")
	messages := t.Title + "\n" + t.Message
	q.Add("text", messages)
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	_, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {
	http.HandleFunc("/send/notif", sendNotify)
	http.ListenAndServe(":9999", nil)
}


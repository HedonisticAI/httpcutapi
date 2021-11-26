package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789"

var in_mem map[string]string

func t_rand() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 10)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func main() {

	if len(os.Args) == 2 && (os.Args[1] == "in_mem" || os.Args[1] == "post") {

		if os.Args[1] == "in_mem" {
			in_mem = make(map[string]string)
		}
		if os.Args[1] == "post" {

		}
		http.HandleFunc("/post/", handlePOST)
		http.HandleFunc("/get/", handleGET)
		log.Fatal(http.ListenAndServe("localhost:8080", nil))

	} else {
		fmt.Println("wrong args, please use 'in_mem' for in-memory mode, or 'post' for postgres")
	}
}
func handlePOST(w http.ResponseWriter, req *http.Request) {
	var bf string
	query := req.URL.Query()
	param := query.Get("url")
	if isValidUrl(param) {
		bf = t_rand()
		if os.Args[1] == "in_mem" {
			if findInMap(reverseMap(in_mem), param) {
				bf = reverseMap(in_mem)[param]
			} else {
				in_mem[bf] = param
				fmt.Print(param + " " + in_mem[bf] + "\n")
			}
		}

		fmt.Fprint(w, bf+" ", param)
	} else {

		fmt.Fprint(w, "bad params, example of good params - http://localhost:8080/post/?url=https:/developer.mozilla.org")
	}

}
func check_short(token string) bool {
	if len(token) != 10 {
		return false
	}
	for i := range token {
		if strings.IndexByte(letter, token[i]) == -1 {
			return false
		}
	}
	return true
}

func isValidUrl(token string) bool {
	_, err := url.ParseRequestURI(token)
	return err == nil

}

func handleGET(w http.ResponseWriter, req *http.Request) {
	var bf string
	query := req.URL.Query()
	param := query.Get("url")
	if check_short(param) {
		if os.Args[1] == "in_mem" {
			bf = in_mem[param]
			if bf == "" {
				bf = "not found"
			}
			fmt.Fprintf(w, bf)
		}
	} else {
		fmt.Fprint(w, "bad params, example of good params http://localhost:8080/get/?url=tJceNCDARG")
	}
}

func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

func findInMap(m map[string]string, s string) bool {
	for v := range m {
		if v == s {
			return true
		}
	}
	return false
}

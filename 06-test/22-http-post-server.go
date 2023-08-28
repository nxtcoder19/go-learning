package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getData(w http.ResponseWriter, req *http.Request) {
	fmt.Println("get data called")
	io.WriteString(w, "This is my website")
}

func main() {
	http.HandleFunc("/", getData)
	http.HandleFunc("/post", getPost)

	err := http.ListenAndServe(":8090", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

type Article struct {
	Title   string `json:"article-title"`
	Content string `json:"article-content"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getPost(w http.ResponseWriter, r *http.Request) {
	// enableCors(&w)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	reqBody, _ := io.ReadAll(r.Body)
	var post Article
	json.Unmarshal(reqBody, &post)

	json.NewEncoder(w).Encode(post)

	newData, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("post called")
		fmt.Println(string(newData))
	}
}

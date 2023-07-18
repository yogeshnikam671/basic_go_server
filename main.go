package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type HelloHandler struct {}

type HelloRequest struct {
  Name string `json:"name"`  // This tilda thing is important while using json.unmarshal
}

func (helloHandler HelloHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) { 
  if(r.URL.Path != "/hello" || r.Method != "POST") {
    http.Error(w, "404 Not found", http.StatusNotFound)
    return
  }

  body, err := io.ReadAll(r.Body)
  if(err != nil) {
    http.Error(w, "400 Bad Request", http.StatusBadRequest)
  }
  
  var helloRequest HelloRequest
  if err := json.Unmarshal(body, &helloRequest); err != nil {
    http.Error(w, "400 Bad Request", http.StatusBadRequest)
    return
  }

  response := fmt.Sprintf("hello %s!!", helloRequest.Name)
  if _, err := w.Write([]byte(response)); err != nil {
    log.Fatal(err);
  }
}

func main() {
  fileServer := http.FileServer(http.Dir("./static"))
  http.Handle("/", fileServer) // Serves the index.html file when '/' is accessed.

  http.Handle("/hello", HelloHandler{})
  

  // starts the server.
  fmt.Println("Listening on port 8080...");
  if err := http.ListenAndServe("localhost:8080", nil); err != nil {
    log.Fatal(err);
  }
}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
	fmt.Println(r.RequestURI)
}

func main() {
	Greet(os.Stdout, "mary\n")
	log.Fatal(http.ListenAndServe(":8001", http.HandlerFunc(GreetHandler)))
}

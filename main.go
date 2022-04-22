package main

import (
	_ "embed"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/apex/gateway/v2"
)

//go:embed thought-jar.txt
var thoughtBytes []byte

//go:embed jar.txt
var jarBytes []byte

func main() {
	if isWeb() {
		serveWeb()
	} else {
		print()
	}
}

// Determines if the user wants this to operates as a web server
// or a one time print.
func isWeb() bool {
	if os.Getenv("WEB") != "" {
		return true
	}

	return len(os.Args) == 2 && os.Args[1] == "web"
}

// Serves web requests asking for a thought.
func serveWeb() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, thought())
	})

	if isLocal() {
		fmt.Println("👂 on port 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		log.Fatal(gateway.ListenAndServe(":8080", nil))
	}
}

// Prints a thought to stdout.
func print() {
	fmt.Println(thought())
}

// Gets a single thought from the jar.
func thought() string {
	image := string(jarBytes)
	text := string(thoughtBytes)

	thoughts := strings.Split(text, "\n--\n")

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	randomIndex := rand.Intn(len(thoughts))

	return image + thoughts[randomIndex]
}

func isLocal() bool {
	return os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == ""
}

package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//go:embed thought-jar.txt
var thoughtBytes []byte

//go:embed jar.txt
var jarBytes []byte

func main() {
	text := string(thoughtBytes)

	thoughts := strings.Split(text, "\n--\n")

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	randomIndex := rand.Intn(len(thoughts))

	message := string(jarBytes) + thoughts[randomIndex]

	fmt.Println(message)
}

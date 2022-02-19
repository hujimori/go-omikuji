package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var rs1Lettesr = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var omikuji = []string{"大吉", "吉", "凶", "小吉", "大凶"}

func getOmikuji() string {
	return omikuji[rand.Intn(len(omikuji))]
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getOmikuji())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	PORT        = "/8080"
)

var (
	baseURL = "http://short.est/"
)

type Response struct {
	ShortURL    string `json:"short_url,omitempty"`
	OriginalURL string `json:"original_url,omitempty"`
	Message     string `json:"message,omitempty"`
}

// To respond with json format
func JsonResponder(w http.ResponseWriter, res Response, httpStatusCode int) {
	jsonResp, err := json.Marshal(res)
	if err != nil {
		http.Error(w, fmt.Sprintf("error building the response, %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write(jsonResp)
}

// to generate random id and append to baseurl
func GenerateShortLink(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	shortURL := fmt.Sprintf("%s%s", baseURL, string(b))
	return shortURL
}

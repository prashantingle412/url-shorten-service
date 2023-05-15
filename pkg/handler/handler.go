package handler

import (
	"encoding/json"
	"net/http"
	"url-shorten-service/pkg/utils"
)

type URL struct {
	URL string `json:"url"`
}

var (
	urlMap = make(map[string]string)
)

func EnocdeURL(w http.ResponseWriter, r *http.Request) {
	var response utils.Response
	if r.Method == http.MethodPost {
		var url URL
		json.NewDecoder(r.Body).Decode(&url)
		_, ok := urlMap[url.URL]
		if !ok {
			shortLink := utils.GenerateShortLink(10)
			urlMap[shortLink] = url.URL
			url.URL = shortLink
		}
		response.ShortURL = url.URL
		response.Message = "URL encoded successfully"
		utils.JsonResponder(w, response, http.StatusCreated)
	} else {
		response.Message = "Method not allowed"
		utils.JsonResponder(w, response, http.StatusMethodNotAllowed)
	}
}

func DecodeURL(w http.ResponseWriter, r *http.Request) {
	var url URL
	var response utils.Response
	if r.Method == http.MethodPost {
		json.NewDecoder(r.Body).Decode(&url)
		v, ok := urlMap[url.URL]
		if ok {
			response.Message = "URL decoded successfully"
			response.OriginalURL = v
			utils.JsonResponder(w, response, http.StatusOK)
		} else {
			response.Message = "URL not found"
			utils.JsonResponder(w, response, http.StatusNotFound)
		}
	} else {
		response.Message = "Method not allowed"
		utils.JsonResponder(w, response, http.StatusMethodNotAllowed)
	}
}

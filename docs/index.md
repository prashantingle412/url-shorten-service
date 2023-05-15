# Techdocs for url shortner Service
* URL shortner service is developed with golang, to develop rest server used inuild package of Golang "net/http".
* Used "testing" inbuilt package for writing unit test.
## REST Endpoints
   ## Encode
   * Enocde api is available on [https://localhost:8080/shortner/encode](https://localhost:8080/shortner/encode) endpoint.
   * Encode api is developed to encode the long or orignial url to the shorten form and save into non persistence stoarage.
   * Used Map for in-memory data stoarage. 
   * It accept http POST method with JSON request body and will return JSON response body with encoded URL, message and http status code.
   * follwoing code snippet show logic for generating random id and appending in base url.
```go
    func GenerateShortLink(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	shortURL := fmt.Sprintf("%s%s", baseURL, string(b))
	return shortURL
}
```
   ### Request body : 
   ```json
      {
       "url":"https://codesubmit.io/library/react"
      }
   ```   
  ### Response body : 
```json
   {
    "short_url": "http://short.est/XVlBzgbaiC",
    "message": "URL encoded successfully"
   }
```
   ## Decode
   * Deocde api is available on https://localhost:8080/shortner/decode endpoint.
   * decode api is developed to decode the short url to the original form of url and send in json format. 
   * It accept http POST method with JSON request body and will return JSON response body with decoded URL, message and http status code.
   ### Follwoing code snippet show the logic for decoding the url to original form.
```go
		v, ok := urlMap[url.URL]
		if ok {
			response.Message = "URL decoded successfully"
			response.OriginalURL = v
			utils.JsonResponder(w, response, http.StatusOK)
		} else {
			response.Message = "URL not found"
			utils.JsonResponder(w, response, http.StatusNotFound)
		}
```      
   ### Request body : 
   ```json
      {
       "url":"http://short.est/XVlBzgbaiC"
      }
   ```   
  ### Response body : 
```json
{
    "original_url": "https://codesubmit.io/library/react",
    "message": "URL decoded successfully"
}
```
### following is the structure for my service.
* pkg - placed all my internal packages here.
* main -   will do my all initiazation part like to initialze the logger,db,and configuration.
* Router - this package is reponsible for routing the http reqest.
* Hanlder - This package i used to store all http handler 
* Service - This packge used for all the bussiness logic of app
* Repository - this package used for DB operations.
* Utils - is my utility package which contains all necessory helper function and constant.
* Config - keep all the configuration required for app like port, host, address but for this app i used constant  
```
url-shoten-service
│   main.go
│   go.mod
│       
└───pkg
│   │   
│   │   
│   │
│   └───handler
│   │       hanlder.go
│   │       handler_test.go
│   │    
│   └───router
│   │     router.go
│   │    
│   │    
│   │   
│   └───service
│   │     service.go
│   │  
│   │
│   └───repository
│   │    repository.go
│   │
│   │
│   └───utils
│         utils.go
└───docs
    │   index.md
```    

### To run app run command
   go run main.go or go build

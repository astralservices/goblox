// Inherited struct and functionality provided from this file
// All API structs will use this.

package network

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var requestTypeToString = []string{"GET", "POST"}

type INetworkRequest interface {
	SendRequest(url string, data map[string]interface{}) (string, error)
	GetCookieArray() []*http.Cookie
	GetCookieString() string
	GetCookie(string) *http.Cookie
	GetHeaders() map[string][]string
	GetResponse() *http.Response
	SetHeaders(map[string][]string)
	SetContentType(HTTPContentType)
	SetRequestType(HTTPRequestType)
	SetResponse(*http.Response)
	AddHeader(string, []string)
	AddCookie(*http.Cookie)
	New()
}

type NetworkRequest struct {
	cookieString string
	cookieArray  []*http.Cookie
	headers      map[string][]string
	requestType  HTTPRequestType
	contentType  HTTPContentType
	response     *http.Response
}

func (ref *NetworkRequest) New() {
	ref.cookieArray = []*http.Cookie{}
	ref.headers = map[string][]string{}
	ref.requestType = GET
}

func (ref *NetworkRequest) GetCookieArray() []*http.Cookie {
	return ref.cookieArray
}

func (ref *NetworkRequest) GetCookieString() string {
	return ref.cookieString
}

func (ref *NetworkRequest) GetCookie(name string) *http.Cookie {
	for _, cookie := range ref.GetCookieArray() {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

func (ref *NetworkRequest) GetHeaders() map[string][]string {
	return ref.headers
}

func (ref *NetworkRequest) GetResponse() *http.Response {
	return ref.response
}

func (ref *NetworkRequest) SetHeaders(header map[string][]string) {
	ref.headers = header
}
func (ref *NetworkRequest) SetContentType(content HTTPContentType) {
	ref.contentType = content
}

func (ref *NetworkRequest) SetRequestType(request HTTPRequestType) {
	ref.requestType = request
}

func (ref *NetworkRequest) SetResponse(res *http.Response) {
	ref.response = res
}

func (ref *NetworkRequest) AddHeader(key string, value []string) {
	ref.headers[key] = value
}

func (ref *NetworkRequest) AddCookie(cookie *http.Cookie) {
	ref.cookieArray = append(ref.cookieArray, cookie)
}

func (ref *NetworkRequest) SendRequest(url string, data map[string]interface{}) (string, error) {
	if data == nil {
		data = map[string]interface{}{}
	}
	marshal, jsonErr := json.Marshal(data)
	if jsonErr != nil {
		return "", jsonErr
	}
	client := http.Client{}
	req, reqError := http.NewRequest(requestTypeToString[ref.requestType], url, bytes.NewBuffer([]byte(marshal)))
	if reqError != nil {
		return "", reqError
	}

	// Set headers
	req.Header = ref.headers

	// Set cookies
	for _, cookie := range ref.GetCookieArray() {
		req.AddCookie(cookie)
	}

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	log.Println("URL: " + url)
	for _, cookie := range response.Cookies() {
		log.Println("Adding Cookie: " + cookie.Name)
		ref.AddCookie(cookie)
	}
	read, _ := ioutil.ReadAll(response.Body)
	ref.response = response
	response.Body.Close()
	return string(read), nil
}

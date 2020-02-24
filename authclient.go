package sumsubgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type AuthClient struct {
	host     string
	login    string
	password string
}

const (
	UriAuthLogin = "/resources/auth/login"
)

func CreateAuthClient(host string, login string, password string) *AuthClient {
	return &AuthClient{host: host, login: login, password: password}
}

func (self *AuthClient) PostCreateTokenAuth() *string {
	url := self.createUrl(UriAuthLogin)
	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte{}))
	self.addHeaderContentJson(request)
	request.SetBasicAuth(self.login, self.password)

	response := self.doRequest(request)

	data, _ := ioutil.ReadAll(response.Body)
	tokenDto := TokenDto{}
	_ = json.Unmarshal(data, &tokenDto)
	if !tokenDto.IsOk() {
		panic("todo")
	}

	return &tokenDto.Payload
}

func (self *AuthClient) createUrl(uri string) string {
	return self.host + uri
}

func (self *AuthClient) addHeaderContentJson(request *http.Request) {
	request.Header.Set(HeaderContentType, HeaderContentTypeAppJson)
}

func (self *AuthClient) doRequest(request *http.Request) *http.Response {
	client := &http.Client{Timeout: 60 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		panic(err)
	}
	//todo проверка статуса кода

	return response
}

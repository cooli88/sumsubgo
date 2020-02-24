package sumsubgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type BaseClient struct {
	host       string
	token      *AuthToken
	authClient *AuthClient
}

func CreateBaseClient(host string, login string, password string) *BaseClient {
	return &BaseClient{host: host, token: &AuthToken{}, authClient: CreateAuthClient(host, login, password)}
}

const (
	HeaderAuth               = "Authorization"
	HeaderContentType        = "Content-Type"
	HeaderContentTypeAppJson = "application/json"
)

func (self *BaseClient) createUrl(uri string) string {
	return self.host + uri
}

func (self *BaseClient) getToken() *string {
	if self.token.IsValid() {
		return self.token.token
	}
	self.token.setToken(self.authClient.PostCreateTokenAuth())
	return self.token.token
}

func (self *BaseClient) Post(uri string, body interface{}) *[]byte {
	return self.doCommonRequest(http.MethodPost, uri, body)
}

func (self *BaseClient) Patch(uri string, body interface{}) *[]byte {
	return self.doCommonRequest(http.MethodPatch, uri, body)
}

func (self *BaseClient) Get(uri string) *[]byte {
	return self.doCommonRequest(http.MethodGet, uri, []byte{})
}

func (self *BaseClient) GetReturnResponse(uri string) *http.Response {
	return self.doCommonRequestRaw(http.MethodGet, uri, []byte{})
}

func (self *BaseClient) Upload(uri string, values *map[string]io.Reader, uploadedFileClient *UploadedFileClient) *http.Response {
	var b bytes.Buffer
	w := self.multipatrTobytes(values, &b)

	url := self.createUrl(uri)

	request, err := http.NewRequest(http.MethodPost, url, &b)
	if err != nil {
		return nil
	}

	self.addHeaderContentType(request, w.FormDataContentType())
	self.addHeaderAuth(request, self.getToken())
	response := self.doRequest(request)

	// Check the response
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", response.Status)
	}
	return response
}

func (self *BaseClient) multipatrTobytes(values *map[string]io.Reader, b *bytes.Buffer) *multipart.Writer {

	var err error
	w := multipart.NewWriter(b)
	for key, r := range *values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				panic(err)
			}
		} else {
			if fw, err = w.CreateFormField(key); err != nil {
				panic(err)
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			panic(err)
		}

	}
	w.Close()

	return w
}

func (self *BaseClient) addBaseHeaders(request *http.Request) {
	self.addHeaderAuth(request, self.getToken())
	self.addHeaderContentJson(request)
}

func (self *BaseClient) addHeaderAuth(request *http.Request, token *string) {
	request.Header.Set(HeaderAuth, "Bearer "+*token)
}

func (self *BaseClient) addHeaderContentJson(request *http.Request) {
	self.addHeaderContentType(request, HeaderContentTypeAppJson)
}

func (self *BaseClient) addHeaderContentType(request *http.Request, contentType string) {
	request.Header.Set(HeaderContentType, contentType)
}

func (self *BaseClient) doCommonRequest(method string, uri string, body interface{}) *[]byte {
	response := self.doCommonRequestRaw(method, uri, body)
	data, _ := ioutil.ReadAll(response.Body)

	return &data
}

func (self *BaseClient) doCommonRequestRaw(method string, uri string, body interface{}) *http.Response {
	jsonValue, _ := json.Marshal(&body)
	url := self.createUrl(uri)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))

	self.addBaseHeaders(request)
	return self.doRequest(request)
}

func (self *BaseClient) doRequest(request *http.Request) *http.Response {
	client := &http.Client{Timeout: 60 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		panic(err)
	}
	//todo проверка статуса кода

	return response
}

package sumsubgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	UriCreateSourceOfFunds string = "/resources/accounts/-/applicants/%s/info"

	HeaderSumsubExtUserId = "X-Data-External-User-Id"
)

type SourceOfIncomeClient struct {
	host string
}

func CreateSourceOfIncomeClient(host string) *SourceOfIncomeClient {
	return &SourceOfIncomeClient{host: host}
}

//testtestuEUoXeYbEJehqGWrVJLgEJwfVFKiDhLu
func (self *SourceOfIncomeClient) CreateSourceOfFunds(applicantId string, kycProcessId string, token string, proofOfIncome ProofOfIncome) *SumsubApplicantInfo {
	uri := fmt.Sprintf(UriCreateSourceOfFunds, applicantId)
	data := self.PatchSourceOfFunds(uri, token, kycProcessId, proofOfIncome)
	fmt.Println(string(*data))
	applicantInfo := SumsubApplicantInfo{}
	_ = json.Unmarshal(*data, &applicantInfo)

	return &applicantInfo
}

func (self *SourceOfIncomeClient) PatchSourceOfFunds(uri string, token string, kycProcessId string, body interface{}) *[]byte {

	jsonValue, _ := json.Marshal(&body)
	url := self.createUrl(uri)
	request, _ := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonValue))

	self.addHeaderAuth(request, &token)
	self.addHeaderContentJson(request)
	self.addHeaderExternalUserId(request, kycProcessId)
	response := self.doRequest(request)
	data, _ := ioutil.ReadAll(response.Body)

	return &data
}

func (self *SourceOfIncomeClient) doRequest(request *http.Request) *http.Response {
	client := &http.Client{Timeout: 60 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		panic(err)
	}
	//todo проверка статуса кода

	return response
}

func (self *SourceOfIncomeClient) createUrl(uri string) string {
	return self.host + uri
}

func (self *SourceOfIncomeClient) addHeaderAuth(request *http.Request, token *string) {
	request.Header.Set(HeaderAuth, "Bearer "+*token)
}
func (self *SourceOfIncomeClient) addHeaderContentJson(request *http.Request) {
	request.Header.Set(HeaderContentType, HeaderContentTypeAppJson)
}

func (self *SourceOfIncomeClient) addHeaderExternalUserId(request *http.Request, kycProcessId string) {
	request.Header.Set(HeaderSumsubExtUserId, kycProcessId)
}

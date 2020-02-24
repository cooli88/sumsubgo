package sumsubcl

import (
	"encoding/json"
	"fmt"
	"io"
)

type SumsubClient struct {
	client               *BaseClient
	sumsubServiceFactory *SumsubServiceFactory
}

const (
	UriGetApplicantState      = "/resources/applicants/%s/state"
	UriGetAdverseMedia        = "/resources/inspections/%s/checks"
	UriCreateApplicant        = "/resources/applicants"
	UriUpdateApplicant        = "/resources/applicants/%s/info"
	UriGetApplicant           = "/resources/applicants/%s"
	UriStartProccesing        = "/resources/applicants/%s/status/pending?%s"
	UriUpdateRequireDocs      = "/resources/applicants/%s/requiredIdDocs"
	UriGetDocument            = "/resources/inspections/%s/resources/%s"
	UriUploadFile             = "/resources/applicants/%s/info/idDoc"
	UriCreateAccessTokenUser  = "/resources/accessTokens?userId=%s"
	UriCreateLoginAccessToken = "/resources/accounts/-/loginAccessToken?token=%s&X-Data-External-User-Id=%s"
)

func NewSumsubClient(BClient *BaseClient) *SumsubClient {
	return &SumsubClient{client: BClient}
}

//CreateApplicant https://developers.sumsub.com/api-reference/#creating-an-applicant
func (self *SumsubClient) CreateApplicant(applicantRequestBody CreateApplicantRequest) *ApplicantResponse {
	data := self.client.Post(UriCreateApplicant, applicantRequestBody)
	applicantResponse := ApplicantResponse{}
	_ = json.Unmarshal(*data, &applicantResponse)

	return &applicantResponse
}

//GetState https://developers.sumsub.com/api-reference/#getting-applicant-status-api
func (self *SumsubClient) GetState(applicantId string) *StatusDto {
	uri := fmt.Sprintf(UriGetApplicantState, applicantId)
	data := self.client.Get(uri)
	statusDto := StatusDto{}
	_ = json.Unmarshal(*data, &statusDto)

	return &statusDto
}

//Update https://developers.sumsub.com/api-reference/#changing-applicant-data
func (self *SumsubClient) Update(updateApplicantRequestBody UpdateApplicantRequest, applicantId string) *IdDocResponse {
	uri := fmt.Sprintf(UriUpdateApplicant, applicantId)
	data := self.client.Patch(uri, updateApplicantRequestBody)
	idDocResponse := IdDocResponse{}
	_ = json.Unmarshal(*data, &idDocResponse)

	return &idDocResponse
}

//GetCheckAdverseMedia
func (self *SumsubClient) GetCheckAdverseMedia(inspectionId string) *MainResponse {
	uri := fmt.Sprintf(UriGetAdverseMedia, inspectionId)
	data := self.client.Get(uri)
	mainResponse := MainResponse{}
	_ = json.Unmarshal(*data, &mainResponse)

	return &mainResponse
}

//GetApplicant https://developers.sumsub.com/api-reference/#getting-applicant-data
func (self *SumsubClient) GetApplicant(applicantId string) *GetApplicantsResponse {
	uri := fmt.Sprintf(UriGetApplicant, applicantId)
	data := self.client.Get(uri)
	getApplicantsResponse := GetApplicantsResponse{}
	_ = json.Unmarshal(*data, &getApplicantsResponse)

	return &getApplicantsResponse
}

//StartProcessing https://developers.sumsub.com/api-reference/#requesting-an-applicant-re-check
func (self *SumsubClient) StartProcessing(applicantId string, reason string) *CompletedResponse {
	uri := fmt.Sprintf(UriStartProccesing, applicantId, reason)
	data := self.client.Post(uri, nil)
	completedResponse := CompletedResponse{}
	_ = json.Unmarshal(*data, &completedResponse)

	return &completedResponse
}

//UpdateRequireDocs https://developers.sumsub.com/api-reference/#changing-required-document-set
func (self *SumsubClient) UpdateRequireDocs(applicantId string, requireDocs RequireDocSet) *RequireDocSet {
	uri := fmt.Sprintf(UriUpdateRequireDocs, applicantId)
	data := self.client.Post(uri, requireDocs)
	requireDocSet := RequireDocSet{}
	_ = json.Unmarshal(*data, &requireDocSet)

	return &requireDocSet
}

//GetFile https://developers.sumsub.com/api-reference/#getting-document-images
func (self *SumsubClient) GetFile(inspectionId string, imageId string) *DownloadedFile {
	uri := fmt.Sprintf(UriGetDocument, inspectionId, imageId)
	response := self.client.GetReturnResponse(uri)
	return self.sumsubServiceFactory.createDownloadFileFromResponse(response)
}

//UploadFile https://developers.sumsub.com/api-reference/#adding-an-id-document
func (self *SumsubClient) UploadFile(applicantId string, values *map[string]io.Reader, uploadedFileClient *UploadedFileClient) *UploadedFile {
	uri := fmt.Sprintf(UriUploadFile, applicantId)
	response := self.client.Upload(uri, values, uploadedFileClient)
	return self.sumsubServiceFactory.createUploadedFileFromResponse(response)
}

//CreateAccessTokenForUser
func (self *SumsubClient) CreateAccessTokenForUser(kycProcessId string) *AccessToken {
	uri := fmt.Sprintf(UriCreateAccessTokenUser, kycProcessId)
	data := self.client.Post(uri, nil)
	accessToken := AccessToken{}
	_ = json.Unmarshal(*data, &accessToken)

	return &accessToken
}

//CreateAccessLoginForUser
func (self *SumsubClient) CreateAccessLoginForUser(kycProcessId string, token string) *TokenDto {
	uri := fmt.Sprintf(UriCreateLoginAccessToken, token, kycProcessId)
	data := self.client.Post(uri, nil)
	tokenDto := TokenDto{}
	_ = json.Unmarshal(*data, &tokenDto)

	return &tokenDto
}

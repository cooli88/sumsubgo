package sumsubcl

import (
	"encoding/json"
	"fmt"
	"io"
	"sumsubcl/dto"
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
func (self *SumsubClient) CreateApplicant(applicantRequestBody dto.CreateApplicantRequest) *dto.ApplicantResponse {
	data := self.client.Post(UriCreateApplicant, applicantRequestBody)
	applicantResponse := dto.ApplicantResponse{}
	_ = json.Unmarshal(*data, &applicantResponse)

	return &applicantResponse
}

//GetState https://developers.sumsub.com/api-reference/#getting-applicant-status-api
func (self *SumsubClient) GetState(applicantId string) *dto.StatusDto {
	uri := fmt.Sprintf(UriGetApplicantState, applicantId)
	data := self.client.Get(uri)
	statusDto := dto.StatusDto{}
	_ = json.Unmarshal(*data, &statusDto)

	return &statusDto
}

//Update https://developers.sumsub.com/api-reference/#changing-applicant-data
func (self *SumsubClient) Update(updateApplicantRequestBody dto.UpdateApplicantRequest, applicantId string) *dto.IdDocResponse {
	uri := fmt.Sprintf(UriUpdateApplicant, applicantId)
	data := self.client.Patch(uri, updateApplicantRequestBody)
	idDocResponse := dto.IdDocResponse{}
	_ = json.Unmarshal(*data, &idDocResponse)

	return &idDocResponse
}

//GetCheckAdverseMedia
func (self *SumsubClient) GetCheckAdverseMedia(inspectionId string) *dto.MainResponse {
	uri := fmt.Sprintf(UriGetAdverseMedia, inspectionId)
	data := self.client.Get(uri)
	mainResponse := dto.MainResponse{}
	_ = json.Unmarshal(*data, &mainResponse)

	return &mainResponse
}

//GetApplicant https://developers.sumsub.com/api-reference/#getting-applicant-data
func (self *SumsubClient) GetApplicant(applicantId string) *dto.GetApplicantsResponse {
	uri := fmt.Sprintf(UriGetApplicant, applicantId)
	data := self.client.Get(uri)
	getApplicantsResponse := dto.GetApplicantsResponse{}
	_ = json.Unmarshal(*data, &getApplicantsResponse)

	return &getApplicantsResponse
}

//StartProcessing https://developers.sumsub.com/api-reference/#requesting-an-applicant-re-check
func (self *SumsubClient) StartProcessing(applicantId string, reason string) *dto.CompletedResponse {
	uri := fmt.Sprintf(UriStartProccesing, applicantId, reason)
	data := self.client.Post(uri, nil)
	completedResponse := dto.CompletedResponse{}
	_ = json.Unmarshal(*data, &completedResponse)

	return &completedResponse
}

//UpdateRequireDocs https://developers.sumsub.com/api-reference/#changing-required-document-set
func (self *SumsubClient) UpdateRequireDocs(applicantId string, requireDocs dto.RequireDocSet) *dto.RequireDocSet {
	uri := fmt.Sprintf(UriUpdateRequireDocs, applicantId)
	data := self.client.Post(uri, requireDocs)
	requireDocSet := dto.RequireDocSet{}
	_ = json.Unmarshal(*data, &requireDocSet)

	return &requireDocSet
}

//GetFile https://developers.sumsub.com/api-reference/#getting-document-images
func (self *SumsubClient) GetFile(inspectionId string, imageId string) *dto.DownloadedFile {
	uri := fmt.Sprintf(UriGetDocument, inspectionId, imageId)
	response := self.client.GetReturnResponse(uri)
	return self.sumsubServiceFactory.createDownloadFileFromResponse(response)
}

//UploadFile https://developers.sumsub.com/api-reference/#adding-an-id-document
func (self *SumsubClient) UploadFile(applicantId string, values *map[string]io.Reader, uploadedFileClient *dto.UploadedFileClient) *dto.UploadedFile {
	uri := fmt.Sprintf(UriUploadFile, applicantId)
	response := self.client.Upload(uri, values, uploadedFileClient)
	return self.sumsubServiceFactory.createUploadedFileFromResponse(response)
}

//CreateAccessTokenForUser
func (self *SumsubClient) CreateAccessTokenForUser(kycProcessId string) *dto.AccessToken {
	uri := fmt.Sprintf(UriCreateAccessTokenUser, kycProcessId)
	data := self.client.Post(uri, nil)
	accessToken := dto.AccessToken{}
	_ = json.Unmarshal(*data, &accessToken)

	return &accessToken
}

//CreateAccessLoginForUser
func (self *SumsubClient) CreateAccessLoginForUser(kycProcessId string, token string) *dto.TokenDto {
	uri := fmt.Sprintf(UriCreateLoginAccessToken, token, kycProcessId)
	data := self.client.Post(uri, nil)
	tokenDto := dto.TokenDto{}
	_ = json.Unmarshal(*data, &tokenDto)

	return &tokenDto
}

package sumsub

import (
	"sumsubcl/common"
	"sumsubcl/dto"
	"sumsubcl/summodel"
)

type SumsubService struct {
	sumsubClient         *SumsubClient
	sourceOfIncomeClient *SourceOfIncomeClient
	sumsubServiceFactory *SumsubServiceFactory
}

var sumsubService *SumsubService

func GetSumsubService(sumsubConfig common.SumsubConfigI) *SumsubService {
	if sumsubService == nil{
		sumsubService = CreateSumsubService(sumsubConfig)
	}

	return sumsubService
}

func CreateSumsubService(sumsubConfig common.SumsubConfigI) *SumsubService {
	baseClient := CreateBaseClient(sumsubConfig.GetHost(), sumsubConfig.GetLogin(), sumsubConfig.GetPassword())
	sumsubClient := NewSumsubClient(baseClient)

	return &SumsubService{sumsubClient: sumsubClient, sourceOfIncomeClient:CreateSourceOfIncomeClient(sumsubConfig.GetHost())}
}

func (self *SumsubService) CreateApplicant(countryCodeISO3 string, applicantProperties summodel.ApplicantPropertiesI, kycProcessId string) *dto.ApplicantResponse {
	applicantRequestBody := self.sumsubServiceFactory.CreateApplicantRequestBody(countryCodeISO3, applicantProperties, kycProcessId, 1)
	return self.sumsubClient.CreateApplicant(applicantRequestBody)
}

func (self *SumsubService) GetState(applicantId string) *dto.StatusDto {
	return self.sumsubClient.GetState(applicantId)
}

func (self *SumsubService) Update(applicantProperties summodel.ApplicantPropertiesI, applicantId string) *dto.IdDocResponse {
	updateApplicantRequestBody := self.sumsubServiceFactory.CreateUpdateApplicantRequest(applicantProperties)
	return self.sumsubClient.Update(updateApplicantRequestBody, applicantId)
}

func (self *SumsubService) GetCheckAdverseMedia(inspectionId string) *dto.MainResponse {
	return self.sumsubClient.GetCheckAdverseMedia(inspectionId)
}

func (self *SumsubService) GetApplicant(applicantId string) *dto.GetApplicantsResponse {
	return self.sumsubClient.GetApplicant(applicantId)
}

func (self *SumsubService) StartProcessing(applicantId string, reason string) *dto.CompletedResponse {
	return self.sumsubClient.StartProcessing(applicantId, reason)
}

func (self *SumsubService) UpdateRequireDocsByTier(applicantId string, tier dto.TierLevel) *dto.RequireDocSet {
	requireDocs := self.sumsubServiceFactory.GetRequireDocsByTier(tier)
	return self.sumsubClient.UpdateRequireDocs(applicantId, requireDocs)
}

func (self *SumsubService) GetFile(inspectionId string, imageId string) *dto.DownloadedFile {
	return self.sumsubClient.GetFile(inspectionId, imageId)
}

func (self *SumsubService) UploadFile(applicantId string, uploadedFileClient *dto.UploadedFileClient, country string) *dto.UploadedFile {
	values := self.sumsubServiceFactory.CreateBodyUploadFile(uploadedFileClient, country)
	return self.sumsubClient.UploadFile(applicantId, values, uploadedFileClient)
}

func (self *SumsubService) CreateSourceOfFunds(applicantId string, kycProcessId string, sourceOfIncome summodel.SourceOfIncomeI) *dto.SumsubApplicantInfo {
	accessToken := self.sumsubClient.CreateAccessTokenForUser(kycProcessId)
	token := self.sumsubClient.CreateAccessLoginForUser(kycProcessId, accessToken.Token)
	proofOfIncome := self.sumsubServiceFactory.CreateProofOfFundsDto(sourceOfIncome)
	return self.sourceOfIncomeClient.CreateSourceOfFunds(applicantId, kycProcessId, token.Payload, proofOfIncome)
}

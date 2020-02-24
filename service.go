package sumsubcl

type SumsubService struct {
	sumsubClient         *SumsubClient
	sourceOfIncomeClient *SourceOfIncomeClient
	sumsubServiceFactory *SumsubServiceFactory
}

var sumsubService *SumsubService

func GetSumsubService(sumsubConfig SumsubConfigI) *SumsubService {
	if sumsubService == nil {
		sumsubService = CreateSumsubService(sumsubConfig)
	}

	return sumsubService
}

func CreateSumsubService(sumsubConfig SumsubConfigI) *SumsubService {
	baseClient := CreateBaseClient(sumsubConfig.GetHost(), sumsubConfig.GetLogin(), sumsubConfig.GetPassword())
	sumsubClient := NewSumsubClient(baseClient)

	return &SumsubService{sumsubClient: sumsubClient, sourceOfIncomeClient: CreateSourceOfIncomeClient(sumsubConfig.GetHost())}
}

func (self *SumsubService) CreateApplicant(countryCodeISO3 string, applicantProperties ApplicantPropertiesI, kycProcessId string) *ApplicantResponse {
	applicantRequestBody := self.sumsubServiceFactory.CreateApplicantRequestBody(countryCodeISO3, applicantProperties, kycProcessId, 1)
	return self.sumsubClient.CreateApplicant(applicantRequestBody)
}

func (self *SumsubService) GetState(applicantId string) *StatusDto {
	return self.sumsubClient.GetState(applicantId)
}

func (self *SumsubService) Update(applicantProperties ApplicantPropertiesI, applicantId string) *IdDocResponse {
	updateApplicantRequestBody := self.sumsubServiceFactory.CreateUpdateApplicantRequest(applicantProperties)
	return self.sumsubClient.Update(updateApplicantRequestBody, applicantId)
}

func (self *SumsubService) GetCheckAdverseMedia(inspectionId string) *MainResponse {
	return self.sumsubClient.GetCheckAdverseMedia(inspectionId)
}

func (self *SumsubService) GetApplicant(applicantId string) *GetApplicantsResponse {
	return self.sumsubClient.GetApplicant(applicantId)
}

func (self *SumsubService) StartProcessing(applicantId string, reason string) *CompletedResponse {
	return self.sumsubClient.StartProcessing(applicantId, reason)
}

func (self *SumsubService) UpdateRequireDocsByTier(applicantId string, tier TierLevel) *RequireDocSet {
	requireDocs := self.sumsubServiceFactory.GetRequireDocsByTier(tier)
	return self.sumsubClient.UpdateRequireDocs(applicantId, requireDocs)
}

func (self *SumsubService) GetFile(inspectionId string, imageId string) *DownloadedFile {
	return self.sumsubClient.GetFile(inspectionId, imageId)
}

func (self *SumsubService) UploadFile(applicantId string, uploadedFileClient *UploadedFileClient, country string) *UploadedFile {
	values := self.sumsubServiceFactory.CreateBodyUploadFile(uploadedFileClient, country)
	return self.sumsubClient.UploadFile(applicantId, values, uploadedFileClient)
}

func (self *SumsubService) CreateSourceOfFunds(applicantId string, kycProcessId string, sourceOfIncome SourceOfIncomeI) *SumsubApplicantInfo {
	accessToken := self.sumsubClient.CreateAccessTokenForUser(kycProcessId)
	token := self.sumsubClient.CreateAccessLoginForUser(kycProcessId, accessToken.Token)
	proofOfIncome := self.sumsubServiceFactory.CreateProofOfFundsDto(sourceOfIncome)
	return self.sourceOfIncomeClient.CreateSourceOfFunds(applicantId, kycProcessId, token.Payload, proofOfIncome)
}

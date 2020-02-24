package sumsubgo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func CreateSumsubServiceFactory() *SumsubServiceFactory {
	return &SumsubServiceFactory{}
}

type SumsubServiceFactory struct {
}

const (
	FieldFirstName          string = "firstName"
	FieldLastName           string = "lastName"
	FieldDateOfBirth        string = "dob"
	KeyDocSets              string = "docSets"
	KeyMetadataSelfiePhrase string = "selfiePhrase"
	HeaderXImageId          string = "X-Image-Id"
)

/**
 * DocSetDto for first tier
 *
 * @return DocSetDto
 */
func (self *SumsubServiceFactory) CreateUpdateApplicantRequest(applicantProperties ApplicantPropertiesI) UpdateApplicantRequest {
	return UpdateApplicantRequest{FirstName: applicantProperties.GetFirstName(), LastName: applicantProperties.GetLastName(), Dob: applicantProperties.GetDateOfBirth()}
}

func (self *SumsubServiceFactory) CreateBodyUploadFile(uploadedFileClient *UploadedFileClient, country string) *map[string]io.Reader {
	metadata := self.createMetadata(uploadedFileClient, country)
	metadataJson, err := json.Marshal(metadata)

	file, err := uploadedFileClient.File.Open()
	if err != nil {
		panic(err)
	}
	values := map[string]io.Reader{
		"content":  file, // lets assume its this file
		"metadata": bytes.NewReader(metadataJson),
	}

	return &values
}

func (self *SumsubServiceFactory) createMetadata(uploadedFileClient *UploadedFileClient, country string) map[string]string {
	metadata := map[string]string{
		"idDocType": *uploadedFileClient.DocType,
		"country":   country,
	}
	if uploadedFileClient.DocSubType != nil {
		metadata["idDocSubType"] = *uploadedFileClient.DocSubType
	}
	return metadata
}

/**
 * DTO for create applicant
 *
 * @param countryCodeISO3     iso-3166 country code {"RUS", "USA", "GBR"}
 * @param applicantProperties
 * @param kycProcessId        in sumsubcl externalUserId
 * @param tierLevel
 * @return
 */
func (self *SumsubServiceFactory) CreateApplicantRequestBody(countryCodeISO3 string, applicantProperties ApplicantPropertiesI, kycProcessId string, tierLevel TierLevel) CreateApplicantRequest {
	createApplicantRequest := CreateApplicantRequest{}
	createApplicantRequest.ExternalUserId = kycProcessId
	createApplicantRequest.Metadata = self.createMetadataDtoArray(applicantProperties)
	createApplicantRequest.Info = self.createApplicantInfo(applicantProperties, countryCodeISO3)
	createApplicantRequest.RequiredIdDocs = self.GetRequireDocsByTier(tierLevel)
	return createApplicantRequest
}

/**
 * create slice MetadataDto, key-value
 *
 * @param applicantProperties
 * @return
 */
func (self *SumsubServiceFactory) createMetadataDtoArray(applicantProperties ApplicantPropertiesI) []Metadata {
	metadataDto := Metadata{KeyMetadataSelfiePhrase, applicantProperties.GetSelfiePhrase()}
	return []Metadata{metadataDto}
}

/**
 * Содздание ApplicantInfoDto
 *
 * @param applicantProperties
 * @param country
 * @return
 */
func (self *SumsubServiceFactory) createApplicantInfo(applicantProperties ApplicantPropertiesI, country string) ApplicantInfo {
	applicantInfo := ApplicantInfo{
		Country:   country,
		FirstName: applicantProperties.GetFirstName(),
		LastName:  applicantProperties.GetLastName(),
		Dob:       applicantProperties.GetDateOfBirth(),
		Phone:     applicantProperties.GetPhone(),
	}
	return applicantInfo
}

/**
 * Создаем Map с необходимыми Документами, используется при создании аппликанта, также при повышении требуемых доков
 * Какие доки нужны для проверки на каждом тире
 *
 * @param tier
 * @return
 */
func (self *SumsubServiceFactory) GetRequireDocsByTier(tier TierLevel) RequireDocSet {

	requiredIdDocs := self.createDocSetByTier(tier)
	requiredIdDocs = append(requiredIdDocs, self.createDocSetForFirstTier())

	if tier > TierLevel1 {
		requiredIdDocs = append(requiredIdDocs, self.createDocSetForIdentity())
		requiredIdDocs = append(requiredIdDocs, self.createDocSetForSelfie())
	}

	if tier > TierLevel2 {
		requiredIdDocs = append(requiredIdDocs, self.createDocSetForInvestability())
	}

	if tier > TierLevel3 {
		requiredIdDocs = append(requiredIdDocs, self.createDocSetForProofOfResidence())
	}

	return RequireDocSet{DocSets: requiredIdDocs}
}

/**
 * Создаем массив пустой DocSetDto, взависимости от tier кол-во элементов в массиве
 *
 * @param tier
 * @return
 */
func (self *SumsubServiceFactory) createDocSetByTier(tier TierLevel) []DocSet {
	//todo сделать нормальное создане сейчас будет аллокейт какждый раз
	//var countElements int8
	//switch (tier) {
	//case 2:
	//	countElements = 3;
	//	break;
	//case 3:
	//	countElements = 4;
	//	break;
	//case 4:
	//	countElements = 5;
	//	break;
	//default:
	//	countElements = 1;
	//	break;
	//}

	a := []DocSet{}

	return a
}

/**
 * DocSetDto для первого тира
 *
 * @return DocSetDto
 */
func (self *SumsubServiceFactory) createDocSetForFirstTier() DocSet {
	docSetDtoForFirstTier := DocSet{IdDocSetType: ApplicantData,
		Fields: []DocSetField{
			self.createDocSetField(FieldFirstName),
			self.createDocSetField(FieldLastName),
			self.createDocSetField(FieldDateOfBirth),
		}}
	return docSetDtoForFirstTier
}

/**
 * @param name
 * @return DocSetField
 */
func (self *SumsubServiceFactory) createDocSetField(name string) DocSetField {
	field := DocSetField{Name: name, Required: true}
	return field
}

/**
* DocSetField для второго тира Identity
 */
func (self *SumsubServiceFactory) createDocSetForIdentity() DocSet {
	return DocSet{IdDocSetType: Identity, Types: self.getDocTypeSetForIdentity(), SubTypes: self.getDocSubTypeSetForIdentity()}
}

/**
* Указываем какие документы принимаем для Identity
 */
func (self *SumsubServiceFactory) getDocTypeSetForIdentity() []SumsubDocType {
	return []SumsubDocType{Passport, IdCard, Drivers}
}

/**
* Указываем какие подтипы документы принимаем для Identity
* <p>
* Примечание 2-сторонние доки только ID_CARD и DRIVER_LICENCE, они всегда 2 сторонние
 */
func (self *SumsubServiceFactory) getDocSubTypeSetForIdentity() []SumsubDocSubType {
	return []SumsubDocSubType{FrontSide, BackSide}
}

//
/**
* DocSetField для второго тира Selfie
 */
func (self *SumsubServiceFactory) createDocSetForSelfie() DocSet {
	return DocSet{IdDocSetType: SelfieSetType, Types: self.getDocTypeSetForSelfie()}
}

/**
* Указываем какие документы принимаем для Selfie
 */
func (self *SumsubServiceFactory) getDocTypeSetForSelfie() []SumsubDocType {
	return []SumsubDocType{Selfie}
}

/**
* DocSetField для 3го тира Investability
 */
func (self *SumsubServiceFactory) createDocSetForInvestability() DocSet {
	return DocSet{IdDocSetType: Investability, Types: self.getDocTypeSetForInvestability()}
}

/**
* Указываем какие документы принимаем для Investability
 */
func (self *SumsubServiceFactory) getDocTypeSetForInvestability() []SumsubDocType {
	return []SumsubDocType{IncomeSource}
}

/**
* DocSetField для 4го тира PROOF_OF_RESIDENCE, может быть принят на 3ем тире, на 4ом обезательно
 */
func (self *SumsubServiceFactory) createDocSetForProofOfResidence() DocSet {
	return DocSet{IdDocSetType: ProofOfResidence, Types: self.getDocTypeSetForProofOfResidenc()}
}

/**
* Указываем какие документы принимаем для PROOF_OF_RESIDENCE
 */
func (self *SumsubServiceFactory) getDocTypeSetForProofOfResidenc() []SumsubDocType {
	return []SumsubDocType{UtilityBill}
}

//
///**
//* Создание из ответа Сумсаба CheckApplicantDto
//*
//* @param response
//* @return
//*/
//func (self *SumsubServiceFactory) CheckApplicantDto createCheckApplicantDtoFroMainDto(MainDto response) {
//CheckApplicantDto checkApplicantDto = new CheckApplicantDto();
//
//if (response.getPersonWatchlist() != null) {
//checkApplicantsetAnswer(response.getPersonWatchlist().getAnswer());
//checkApplicantsetHits(response.getPersonWatchlist().getWatchlistInfo().getHits());
//}
//
//return checkApplicantDto;
//}

/**
* Создание тела запроса для ProofOfFunds
*
* @param applicantProperties
* @return
 */
func (self *SumsubServiceFactory) CreateProofOfFundsDto(sourceOfIncome SourceOfIncomeI) ProofOfIncome {
	return ProofOfIncome{
		InvestabilityDto{
			sourceOfIncome.GetInvestabilityType(),
			sourceOfIncome.GetTransactionAmount(),
			sourceOfIncome.GetSourceOfIncome(),
			sourceOfIncome.GetInvestmentKnowledge(),
			RangeOfMoney{
				sourceOfIncome.GetAnnualIncome().GetFrom(),
				sourceOfIncome.GetAnnualIncome().GetTo(),
				sourceOfIncome.GetAnnualIncome().GetCurrency()},
			RangeOfMoney{
				sourceOfIncome.GetNetWorth().GetFrom(),
				sourceOfIncome.GetNetWorth().GetTo(),
				sourceOfIncome.GetNetWorth().GetCurrency()},
		},
	}
}

//
/**
* Создание UploadedFileDto из ResponseEntity
*
* @param responseEntity
* @return
 */
func (self *SumsubServiceFactory) createUploadedFileFromResponse(response *http.Response) *UploadedFile {
	return &UploadedFile{response.Header.Get(HeaderXImageId)}
}

/**
* Создание DownloadFileDto из ResponseEntity
*
* @param response
* @return
 */
func (self *SumsubServiceFactory) createDownloadFileFromResponse(response *http.Response) *DownloadedFile {
	return &DownloadedFile{
		Body:          response.Body,
		ContentType:   response.Header.Get(HeaderContentType),
		ContentLength: response.ContentLength,
	}
}

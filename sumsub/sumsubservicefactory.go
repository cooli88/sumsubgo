package sumsub

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sumsubcl/dto"
	"sumsubcl/summodel"
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
func (self *SumsubServiceFactory) CreateUpdateApplicantRequest(applicantProperties summodel.ApplicantPropertiesI) dto.UpdateApplicantRequest {
	return dto.UpdateApplicantRequest{FirstName: applicantProperties.GetFirstName(), LastName: applicantProperties.GetLastName(), Dob: applicantProperties.GetDateOfBirth(),}
}

func (self *SumsubServiceFactory) CreateBodyUploadFile(uploadedFileClient *dto.UploadedFileClient, country string) *map[string]io.Reader {
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

func (self *SumsubServiceFactory) createMetadata(uploadedFileClient *dto.UploadedFileClient, country string) map[string]string {
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
func (self *SumsubServiceFactory) CreateApplicantRequestBody(countryCodeISO3 string, applicantProperties summodel.ApplicantPropertiesI, kycProcessId string, tierLevel dto.TierLevel) dto.CreateApplicantRequest {
	createApplicantRequest := dto.CreateApplicantRequest{}
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
func (self *SumsubServiceFactory) createMetadataDtoArray(applicantProperties summodel.ApplicantPropertiesI) []dto.Metadata {
	metadataDto := dto.Metadata{KeyMetadataSelfiePhrase, applicantProperties.GetSelfiePhrase()}
	return []dto.Metadata{metadataDto}
}

/**
 * Содздание ApplicantInfoDto
 *
 * @param applicantProperties
 * @param country
 * @return
 */
func (self *SumsubServiceFactory) createApplicantInfo(applicantProperties summodel.ApplicantPropertiesI, country string) dto.ApplicantInfo {
	applicantInfo := dto.ApplicantInfo{
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
func (self *SumsubServiceFactory) GetRequireDocsByTier(tier dto.TierLevel) dto.RequireDocSet {

	requiredIdDocs := self.createDocSetByTier(tier);
	requiredIdDocs = append(requiredIdDocs, self.createDocSetForFirstTier())

	if tier > dto.TierLevel1 {
		requiredIdDocs = append(requiredIdDocs, self.createDocSetForIdentity())
		requiredIdDocs = append(requiredIdDocs, self.createDocSetForSelfie())
	}

	if tier > dto.TierLevel2 {
		requiredIdDocs = append(requiredIdDocs, self.createDocSetForInvestability())
	}

	if tier > dto.TierLevel3 {
		requiredIdDocs = append(requiredIdDocs, self.createDocSetForProofOfResidence())
	}

	return dto.RequireDocSet{DocSets: requiredIdDocs}
}

/**
 * Создаем массив пустой DocSetDto, взависимости от tier кол-во элементов в массиве
 *
 * @param tier
 * @return
 */
func (self *SumsubServiceFactory) createDocSetByTier(tier dto.TierLevel) []dto.DocSet {
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

	a := []dto.DocSet{}

	return a
}

/**
 * DocSetDto для первого тира
 *
 * @return DocSetDto
 */
func (self *SumsubServiceFactory) createDocSetForFirstTier() dto.DocSet {
	docSetDtoForFirstTier := dto.DocSet{IdDocSetType: dto.ApplicantData,
		Fields: []dto.DocSetField{
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
func (self *SumsubServiceFactory) createDocSetField(name string) dto.DocSetField {
	field := dto.DocSetField{Name: name, Required: true}
	return field
}

/**
* DocSetField для второго тира Identity
*/
func (self *SumsubServiceFactory) createDocSetForIdentity() dto.DocSet {
	return dto.DocSet{IdDocSetType: dto.Identity, Types: self.getDocTypeSetForIdentity(), SubTypes: self.getDocSubTypeSetForIdentity()}
}

/**
* Указываем какие документы принимаем для Identity
*/
func (self *SumsubServiceFactory) getDocTypeSetForIdentity() []dto.SumsubDocType {
	return []dto.SumsubDocType{dto.Passport, dto.IdCard, dto.Drivers,}
}

/**
* Указываем какие подтипы документы принимаем для Identity
* <p>
* Примечание 2-сторонние доки только ID_CARD и DRIVER_LICENCE, они всегда 2 сторонние
*/
func (self *SumsubServiceFactory) getDocSubTypeSetForIdentity() []dto.SumsubDocSubType {
	return []dto.SumsubDocSubType{dto.FrontSide, dto.BackSide}
}

//
/**
* DocSetField для второго тира Selfie
*/
func (self *SumsubServiceFactory) createDocSetForSelfie() dto.DocSet {
	return dto.DocSet{IdDocSetType: dto.SelfieSetType, Types: self.getDocTypeSetForSelfie()}
}

/**
* Указываем какие документы принимаем для Selfie
*/
func (self *SumsubServiceFactory) getDocTypeSetForSelfie() []dto.SumsubDocType {
	return []dto.SumsubDocType{dto.Selfie}
}

/**
* DocSetField для 3го тира Investability
*/
func (self *SumsubServiceFactory) createDocSetForInvestability() dto.DocSet {
	return dto.DocSet{IdDocSetType: dto.Investability, Types: self.getDocTypeSetForInvestability()}
}

/**
* Указываем какие документы принимаем для Investability
*/
func (self *SumsubServiceFactory) getDocTypeSetForInvestability() []dto.SumsubDocType {
	return []dto.SumsubDocType{dto.IncomeSource}
}

/**
* DocSetField для 4го тира PROOF_OF_RESIDENCE, может быть принят на 3ем тире, на 4ом обезательно
*/
func (self *SumsubServiceFactory) createDocSetForProofOfResidence() dto.DocSet {
	return dto.DocSet{IdDocSetType: dto.ProofOfResidence, Types: self.getDocTypeSetForProofOfResidenc()}
}

/**
* Указываем какие документы принимаем для PROOF_OF_RESIDENCE
*/
func (self *SumsubServiceFactory) getDocTypeSetForProofOfResidenc() []dto.SumsubDocType {
	return []dto.SumsubDocType{dto.UtilityBill}
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
//checkApplicantDto.setAnswer(response.getPersonWatchlist().getAnswer());
//checkApplicantDto.setHits(response.getPersonWatchlist().getWatchlistInfo().getHits());
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
func (self *SumsubServiceFactory) CreateProofOfFundsDto(sourceOfIncome summodel.SourceOfIncomeI) dto.ProofOfIncome {
	return dto.ProofOfIncome{
		dto.InvestabilityDto{
			sourceOfIncome.GetInvestabilityType(),
			sourceOfIncome.GetTransactionAmount(),
			sourceOfIncome.GetSourceOfIncome(),
			sourceOfIncome.GetInvestmentKnowledge(),
			dto.RangeOfMoney{
				sourceOfIncome.GetAnnualIncome().GetFrom(),
				sourceOfIncome.GetAnnualIncome().GetTo(),
				sourceOfIncome.GetAnnualIncome().GetCurrency()},
			dto.RangeOfMoney{
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
func (self *SumsubServiceFactory) createUploadedFileFromResponse(response *http.Response) *dto.UploadedFile {
	return &dto.UploadedFile{response.Header.Get(HeaderXImageId)}
}

/**
* Создание DownloadFileDto из ResponseEntity
*
* @param response
* @return
*/
func (self *SumsubServiceFactory) createDownloadFileFromResponse(response *http.Response) *dto.DownloadedFile {
	return &dto.DownloadedFile{
		Body:          response.Body,
		ContentType:   response.Header.Get(HeaderContentType),
		ContentLength: response.ContentLength,
	}
}

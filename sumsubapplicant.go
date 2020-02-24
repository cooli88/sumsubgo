package sumsubcl

type (
	///------------------create data------------------///

	ApplicantInfo struct {
		Country   string `json:"country"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Dob       string `json:"dob"`
		Phone     string `json:"phone"`
	}

	CreateApplicantRequest struct {
		ExternalUserId string        `json:"externalUserId"`
		Metadata       []Metadata    `json:"metadata"`
		Info           ApplicantInfo `json:"info"`
		RequiredIdDocs RequireDocSet `json:"requiredIdDocs"`
	}

	RequireDocSet struct {
		DocSets []DocSet `json:"docSets"`
	}

	DocSet struct {
		IdDocSetType SumsubIdDocSetType `json:"idDocSetType"`
		Types        []SumsubDocType    `json:"types"`
		SubTypes     []SumsubDocSubType `json:"subTypes"`
		Fields       []DocSetField      `json:"fields"`
	}

	DocSetField struct {
		Name     string `json:"name"`
		Required bool   `json:"required"`
	}

	Metadata struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	///------------------create data------------------///

	///------------------response data------------------///
	ApplicantResponse struct {
		Id                string              `json:"id"`
		CreatedAt         string              `json:"createdAt"`
		ClientId          string              `json:"clientId"`
		InspectionId      string              `json:"inspectionId"`
		JobId             string              `json:"jobId"`
		ExternalUserId    string              `json:"externalUserId"`
		Email             string              `json:"email"`
		Status            string              `json:"status"`
		ApplicantPlatform string              `json:"applicantPlatform"`
		Info              SumsubApplicantInfo `json:"info"`
		Env               string              `json:"env"`
		Notes             string              `json:"notes"`
	}

	SumsubApplicantInfo struct {
		FirstName      string           `json:"firstName"`
		MiddleName     string           `json:"middleName"`
		FirstNameEn    string           `json:"firstNameEn"`
		LastNameEn     string           `json:"lastNameEn"`
		LastName       string           `json:"lastName"`
		Dob            string           `json:"dob"`
		PlaceOfBirth   string           `json:"placeOfBirth"`
		Country        string           `json:"country"`
		CountryOfBirth string           `json:"countryOfBirth"`
		Phone          string           `json:"phone"`
		Gender         string           `json:"gender"`
		StateOfBirth   string           `json:"stateOfBirth"`
		Addresses      []Address        `json:"addresses"`
		IdDocs         []IdDocResponse  `json:"idDocs"`
		Investability  InvestabilityDto `json:"investability"`
	}

	Address struct {
		SubStreet string `json:"subStreet"`
		Street    string `json:"street"`
		State     string `json:"state"`
		Town      string `json:"town"`
		PostCode  string `json:"postCode"`
		Country   string `json:"country"`
	}

	IdDocResponse struct {
		IdDocType       string `json:"idDocType"`
		Country         string `json:"country"`
		FirstNam        string `json:"firstNam"`
		MiddleName      string `json:"middleName"`
		LastName        string `json:"lastName"`
		IssuedDate      string `json:"issuedDate"`
		ValidUntil      string `json:"validUntil"`
		FirstIssuedDate string `json:"firstIssuedDate"`
		Number          string `json:"number"`
		Dob             string `json:"dob"`
	}

	///------------------response data------------------///

	UpdateApplicantRequest struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Dob       string `json:"dob"`
	}

	GetApplicantsResponse struct {
		List ItemsSumsubApplicant `json:"list"`
	}

	ItemsSumsubApplicant struct {
		Items []ApplicantResponse `json:"items"`
	}

	///------------------source of funds data------------------///

	ProofOfIncome struct {
		Investability InvestabilityDto `json:"investability"`
	}

	InvestabilityDto struct {
		InvestabilityType        InvestabilityTypeEnum `json:"investabilityType"`
		TransactionAmount        int64                 `json:"transactionAmount"`
		IncomeSourceType         SumsubSourceOfIncome  `json:"incomeSourceType"`
		InvestmentKnowledgeLevel InvestmentKnowledge   `json:"investmentKnowledgeLevel"`
		AnnualIncome             RangeOfMoney          `json:"annualIncome"`
		NetWorth                 RangeOfMoney          `json:"netWorth"`
	}

	RangeOfMoney struct {
		From     *int64 `json:"from"`
		To       int64  `json:"to"`
		Currency string `json:"currency"`
	}

	///------------------source of funds data------------------///
)

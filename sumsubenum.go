package sumsubgo

type SumsubDocSubType string
type SumsubDocType string
type SumsubIdDocSetType string
type TierLevel int8
type InvestabilityTypeEnum string
type SumsubSourceOfIncome string
type InvestmentKnowledge string
type ReviewAnswerState string
type ReviewStatusState string
type ReviewRejectType string

const (
	FrontSide SumsubDocSubType = "FRONT_SIDE"
	BackSide  SumsubDocSubType = "BACK_SIDE"

	Passport                       SumsubDocType = "PASSPORT"
	IdCard                         SumsubDocType = "ID_CARD"
	Drivers                        SumsubDocType = "DRIVERS"
	Selfie                         SumsubDocType = "SELFIE"
	VideoSelfie                    SumsubDocType = "VIDEO_SELFIE"
	IncomeSource                   SumsubDocType = "INCOME_SOURCE"
	UtilityBill                    SumsubDocType = "UTILITY_BILL"
	BankCard                       SumsubDocType = "BANK_CARD"
	BankStatement                  SumsubDocType = "BANK_STATEMENT"
	ProfileImage                   SumsubDocType = "PROFILE_IMAGE"
	IdDocPhoto                     SumsubDocType = "ID_DOC_PHOTO"
	Agreement                      SumsubDocType = "AGREEMENT"
	Contract                       SumsubDocType = "CONTRACT"
	ResidencePermit                SumsubDocType = "RESIDENCE_PERMIT"
	DriversTranslation             SumsubDocType = "DRIVERS_TRANSLATION"
	InvestorDoc                    SumsubDocType = "INVESTOR_DOC"
	VehicleRegistrationCertificate SumsubDocType = "VEHICLE_REGISTRATION_CERTIFICATE"
	EmploymentCertificate          SumsubDocType = "EMPLOYMENT_CERTIFICATE"
	Snils                          SumsubDocType = "SNILS"
	OtherDocType                   SumsubDocType = "OTHER"

	ApplicantData    SumsubIdDocSetType = "APPLICANT_DATA"
	Identity         SumsubIdDocSetType = "IDENTITY"
	Identity2        SumsubIdDocSetType = "IDENTITY2"
	SelfieSetType    SumsubIdDocSetType = "SELFIE"
	SelfieSetType2   SumsubIdDocSetType = "SELFIE2"
	Investability    SumsubIdDocSetType = "INVESTABILITY"
	ProofOfResidence SumsubIdDocSetType = "PROOF_OF_RESIDENCE"
	PaymentMethods   SumsubIdDocSetType = "PAYMENT_METHODS"

	InvestabilityTypePerson InvestabilityTypeEnum = "person"

	SumsubSourceOfIncomeSalary      SumsubSourceOfIncome = "salary"
	SumsubSourceOfIncomeRent        SumsubSourceOfIncome = "rent"
	SumsubSourceOfIncomeInvestments SumsubSourceOfIncome = "investments"
	SumsubSourceOfIncomeOther       SumsubSourceOfIncome = "other"

	SumsubInvestmentKnowledgePoor          InvestmentKnowledge = "poor"
	SumsubInvestmentKnowledgeLimited       InvestmentKnowledge = "limited"
	SumsubInvestmentKnowledgeGood          InvestmentKnowledge = "good"
	SumsubInvestmentKnowledgeSophisticated InvestmentKnowledge = "sophisticated"

	TierLevel0 TierLevel = 0
	TierLevel1 TierLevel = 1
	TierLevel2 TierLevel = 2
	TierLevel3 TierLevel = 3
	TierLevel4 TierLevel = 4

	Init                ReviewStatusState = "init"
	Pending             ReviewStatusState = "pending"
	Queued              ReviewStatusState = "queued"
	Completed           ReviewStatusState = "completed"
	CompletedSent       ReviewStatusState = "completedSent"
	CompletedSetFailure ReviewStatusState = "completedSentFailure"

	Red   ReviewAnswerState = "RED"
	Green ReviewAnswerState = "GREEN"

	Final    ReviewRejectType = "FINAL"
	Retry    ReviewRejectType = "RETRY"
	External ReviewRejectType = "EXTERNAL"
)

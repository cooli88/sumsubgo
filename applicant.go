package sumsubgo

type ApplicantPropertiesI interface {
	GetSelfiePhrase() string
	GetFirstName() string
	GetLastName() string
	GetDateOfBirth() string
	GetPhone() string
}

type SourceOfIncomeI interface {
	GetInvestabilityType() InvestabilityTypeEnum
	GetTransactionAmount() int64
	GetSourceOfIncome() SumsubSourceOfIncome
	GetInvestmentKnowledge() InvestmentKnowledge
	GetAnnualIncome() RangeOfMoneyI
	GetNetWorth() RangeOfMoneyI
}

type RangeOfMoneyI interface {
	GetFrom() *int64
	GetTo() int64
	GetCurrency() string
}

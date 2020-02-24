package summodel

import "sumsubcl/dto"

type ApplicantPropertiesI interface {
	GetSelfiePhrase() string
	GetFirstName() string
	GetLastName() string
	GetDateOfBirth() string
	GetPhone() string
}

type SourceOfIncomeI interface {
	GetInvestabilityType() dto.InvestabilityTypeEnum
	GetTransactionAmount() int64
	GetSourceOfIncome() dto.SumsubSourceOfIncome
	GetInvestmentKnowledge() dto.InvestmentKnowledge
	GetAnnualIncome() RangeOfMoneyI
	GetNetWorth() RangeOfMoneyI
}

type RangeOfMoneyI interface {
	GetFrom() *int64
	GetTo() int64
	GetCurrency() string
}

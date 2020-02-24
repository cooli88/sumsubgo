
```go
package main

import (
	"fmt"
	"github.com/cooli88/sumsubgo"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

const (
	ApplicantId  = "1**********************8"
	InspectionId = "1**********************8"
	SumsubHost   = "https://test-api.sumsub.com"
	SumsubLogin  = "LOGIN"
	SumsubPass   = "PASS"
)

type (
	MainConfig struct {
		Sumsub SumsubConfig `yaml:"sumsub"`
	}

	SumsubConfig struct {
		Host     string `yaml:"host"`
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
	}
)

func init() {
	mainConfig = createMainConfig()
}

const configFilePath = "config.yml"

var mainConfig *MainConfig

func GetMainConfig() *MainConfig {
	return mainConfig
}

func createMainConfig() *MainConfig {
	return &MainConfig{
		SumsubConfig{SumsubHost, SumsubLogin, SumsubPass},
	}
}

func (self SumsubConfig) GetHost() string {
	return self.Host
}

func (self SumsubConfig) GetLogin() string {
	return self.Login
}

func (self SumsubConfig) GetPassword() string {
	return self.Password
}

type (
	ApplicantProperties struct {
		FirstName             string
		LastName              string
		DateOfBirth           string
		SelfiePhrase          string
		PhoneCountryCode      string
		PhoneSubscriberNumber string
	}

	SourceOfIncome struct {
		InvestabilityType   sumsubgo.InvestabilityTypeEnum
		TransactionAmount   int64
		SourceOfIncome      sumsubgo.SumsubSourceOfIncome
		InvestmentKnowledge sumsubgo.InvestmentKnowledge
		AnnualIncome        RangeOfMoney
		NetWorth            RangeOfMoney
	}

	RangeOfMoney struct {
		From     *int64
		To       int64
		Currency string
	}
)

func (self ApplicantProperties) GetSelfiePhrase() string {
	return self.SelfiePhrase
}

func (self ApplicantProperties) GetFirstName() string {
	return self.FirstName
}

func (self ApplicantProperties) GetLastName() string {
	return self.LastName
}

func (self ApplicantProperties) GetDateOfBirth() string {
	return self.DateOfBirth
}

func (self ApplicantProperties) GetPhone() string {
	return self.PhoneCountryCode + self.PhoneSubscriberNumber
}

func (self SourceOfIncome) GetInvestabilityType() sumsubgo.InvestabilityTypeEnum {
	return self.InvestabilityType
}

func (self SourceOfIncome) GetTransactionAmount() int64 {
	return self.TransactionAmount
}

func (self SourceOfIncome) GetSourceOfIncome() sumsubgo.SumsubSourceOfIncome {
	return self.SourceOfIncome
}

func (self SourceOfIncome) GetInvestmentKnowledge() sumsubgo.InvestmentKnowledge {
	return self.InvestmentKnowledge
}

func (self SourceOfIncome) GetAnnualIncome() sumsubgo.RangeOfMoneyI {
	return self.AnnualIncome
}

func (self SourceOfIncome) GetNetWorth() sumsubgo.RangeOfMoneyI {
	return self.NetWorth
}

func (self RangeOfMoney) GetFrom() *int64 {
	return self.From
}

func (self RangeOfMoney) GetTo() int64 {
	return self.To
}

func (self RangeOfMoney) GetCurrency() string {
	return self.Currency
}

var todoController *TodoController

func init() {
	rand.Seed(time.Now().UnixNano())
	todoController = createTodoController()
}

func GetTodoContoller() *TodoController {
	return todoController
}

func createTodoController() *TodoController {
	return &TodoController{sumsubService: sumsubgo.CreateSumsubService(GetMainConfig().Sumsub)}
}

type TodoController struct {
	sumsubService *sumsubgo.SumsubService
}

func (self *TodoController) Create(c *gin.Context) {

	applicantProperties := ApplicantProperties{FirstName: "ferd", LastName: "tred", DateOfBirth: "1985-01-01", SelfiePhrase: "I am ferd", PhoneCountryCode: "7", PhoneSubscriberNumber: "0002221111"}
	createdApplicant := self.sumsubService.CreateApplicant("RUS", applicantProperties, "testtest"+RandStringBytes(32))
	fmt.Println(createdApplicant)
	applicantProperties.LastName = "lotest"
	applicantProperties.FirstName = "test23"
	applicantProperties.DateOfBirth = "1980-04-04"
	updatedApplicant := self.sumsubService.Update(applicantProperties, createdApplicant.Id)
	fmt.Println(updatedApplicant)
	startedProcessing := self.sumsubService.StartProcessing(createdApplicant.Id, "reason=WL")
	fmt.Println(startedProcessing)
	res := self.sumsubService.UpdateRequireDocsByTier(createdApplicant.Id, 2)
	fmt.Println(res)
	self.sumsubService.UpdateRequireDocsByTier(createdApplicant.Id, 3)
	self.sumsubService.UpdateRequireDocsByTier(createdApplicant.Id, 4)
	getApplicant := self.sumsubService.GetApplicant(createdApplicant.Id)
	fmt.Println(getApplicant)
	getState := self.sumsubService.GetState(createdApplicant.Id)
	fmt.Println(getState)
	//result := self.sumsubService.GetCheckAdverseMedia("1**********************8")
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": getApplicant})
}

func (self *TodoController) FetchAll(c *gin.Context) {
	imageId := "2********4"

	downloadedFile := self.sumsubService.GetFile(InspectionId, imageId)
	c.DataFromReader(http.StatusOK, downloadedFile.ContentLength, downloadedFile.ContentType, downloadedFile.Body, downloadedFile.GetExtaHeaders("test"))
}

func (self *TodoController) FetchSingle(c *gin.Context) {
	id := c.Param("id")
	downloadedFile := self.sumsubService.GetFile(InspectionId, id)
	c.DataFromReader(http.StatusOK, downloadedFile.ContentLength, downloadedFile.ContentType, downloadedFile.Body, downloadedFile.GetExtaHeaders("test"))
}

func (self *TodoController) Update(c *gin.Context) {
	todoID := c.Param("id")
	externaUserId := "testtestbSQtQxAXjLKgMCgtqNzGRhxQqiMaCABc"

	sourceOfIncome := SourceOfIncome{
		InvestabilityType:   sumsubgo.InvestabilityTypePerson,
		TransactionAmount:   1000,
		SourceOfIncome:      sumsubgo.SumsubSourceOfIncomeSalary,
		InvestmentKnowledge: sumsubgo.SumsubInvestmentKnowledgePoor,
		AnnualIncome:        RangeOfMoney{To: 200000, Currency: "USD"},
		NetWorth:            RangeOfMoney{To: 500000, Currency: "USD"},
	}

	result := self.sumsubService.CreateSourceOfFunds(ApplicantId, externaUserId, sourceOfIncome)
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": fmt.Sprintf("updated id: %+s", todoID)})
}

func (self *TodoController) Delete(c *gin.Context) {
	todoID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": fmt.Sprintf("deleted id: %+s", todoID)})
}

func (self *TodoController) Upload(c *gin.Context) {
	form, _ := c.MultipartForm()

	uploadedFileClient := sumsubgo.UploadedFileClient{}

	if len(form.File["file"]) > 0 {
		uploadedFileClient.File = form.File["file"][0]
	}

	if len(form.Value["docType"]) > 0 {
		uploadedFileClient.DocType = &form.Value["docType"][0]
	}

	if len(form.Value["docSubType"]) > 0 {
		uploadedFileClient.DocSubType = &form.Value["docSubType"][0]
	}

	if uploadedFileClient.File == nil {
		panic("alll GOOOD")
	}

	uploadedFile := self.sumsubService.UploadFile(ApplicantId, &uploadedFileClient, "RUS")

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": *uploadedFile})
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func StartServer() {
	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		todoController := GetTodoContoller()
		v1.POST("/", todoController.Create)
		v1.GET("/", todoController.FetchAll)
		v1.GET("/:id", todoController.FetchSingle)
		v1.PUT("/:id", todoController.Update)
		v1.DELETE("/:id", todoController.Delete)
		v1.POST("/upload", todoController.Upload)
	}
	router.Run()
}

func main() {

	StartServer()

}
```
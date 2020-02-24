package sumsubgo

import (
	"io"
	"mime/multipart"
)

type (
	CompletedResponse struct {
		Ok bool `json:"ok"`
	}

	UploadedFileClient struct {
		File       *multipart.FileHeader
		DocType    *string
		DocSubType *string
	}

	UploadedFile struct {
		Id string `json:"id"`
	}

	FileToSumsub struct {
		File       *multipart.FileHeader
		DocType    *string
		DocSubType *string
	}

	DownloadedFile struct {
		Body          io.ReadCloser
		ContentType   string
		ContentLength int64
	}

	TokenDto struct {
		Status  *string `json:"status"`
		Payload string  `json:"payload"`
	}

	AccessToken struct {
		UserId string `json:"userId"`
		Token  string `json:"token"`
	}
)

func (self *TokenDto) IsOk() bool {
	return self.Status != nil && *self.Status == "ok"
}

func (self *DownloadedFile) GetExtaHeaders(fileName string) map[string]string {
	return map[string]string{
		"Content-Disposition": `attachment; filename="` + fileName + `"`,
	}
}

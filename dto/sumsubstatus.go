package dto

type (
	DocumentStatusDto struct {
		IdDocType    string          `json:"idDocType"`
		IdDocSubType string          `json:"idDocSubType"`
		Country      string          `json:"country"`
		ImageId      int64           `json:"imageId"`
		ReviewResult ReviewResultDto `json:"reviewResult"`
		AddedDate    string          `json:"addedDate"`
	}

	ReviewResultDto struct {
		ModerationComment string             `json:"moderationComment"`
		ClientComment     string             `json:"clientComment"`
		ReviewAnswer      *ReviewAnswerState `json:"reviewAnswer"`
		RejectLabels      []string           `json:"rejectLabels"`
		Label             string             `json:"label"`
		ReviewRejectType  *ReviewRejectType  `json:"reviewRejectType"`
		CustomTouch       bool               `json:"customTouch"`
	}

	StatusDetailDto struct {
		Id                     string            `json:"id"`
		InspectionId           string            `json:"inspectionId"`
		ApplicantId            string            `json:"applicantId"`
		JobId                  string            `json:"jobId"`
		ElapsedSincePendingMs  int64             `json:"elapsedSincePendingMs"`
		ElapsedSinceQueuedMs   int64             `json:"elapsedSinceQueuedMs"`
		CreateDate             string            `json:"createDate"`
		StartDate              string            `json:"startDate"`
		ReviewDate             string            `json:"reviewDate"`
		ReviewStatus           ReviewStatusState `json:"reviewStatus"`
		Priority               int64             `json:"priority"`
		ReviewResult           ReviewResultDto   `json:"reviewResult"`
		NotificationFailureCnt int64             `json:"notificationFailureCnt"`
	}

	StatusDto struct {
		Id             string              `json:"id"`
		Status         StatusDetailDto     `json:"status"`
		DocumentStatus []DocumentStatusDto `json:"documentStatus"`
	}
)

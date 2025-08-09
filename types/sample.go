package types

type Sample struct {
	ID        int       `json:"id"`
	DATA      string    `json:"data"`
}

type CreateSampleReq struct {
	Data string `json:"data" binding:"required,min=20"`
}
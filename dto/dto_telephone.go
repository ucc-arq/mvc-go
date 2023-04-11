package dto

type TelephoneDto struct {
	Code   string `json:"code"`
	Number string `json:"number"`
	UserId int    `json:"user_id,omitempty"`
}

type TelephonesDto []TelephoneDto

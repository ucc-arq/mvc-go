package dto

type UserDetailDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`

	Street string `json:"street_name"`
	Number int    `json:"street_number"`

	TelephonesDto TelephonesDto `json:"telephones,omitempty"`
}

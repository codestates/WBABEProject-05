package protocol

type RequestPostUser struct {
	Name        string `json:"name"`
	NicName     string `json:"nic_name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
}

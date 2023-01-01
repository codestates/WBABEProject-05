package response

import (
	"github.com/codestates/WBABEProject-05/model/entity"
)

type ResponseUser struct {
	ID          string `json:"user_id,omitempty"`
	Name        string `json:"name,omitempty"`
	NicName     string `json:"nic_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Role        string `json:"role,omitempty"`
}

func NewResponseUserFromUser(user *entity.User) *ResponseUser {
	return &ResponseUser{
		ID:          user.ID.Hex(),
		Name:        user.Name,
		NicName:     user.NicName,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
	}
}

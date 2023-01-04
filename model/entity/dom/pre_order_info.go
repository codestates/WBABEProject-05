package dom

type PreOrderInfo struct {
	Address     *Address `bson:"pre_order_address,omitempty"`
	PhoneNumber string   `bson:"pre_order_phone_number,omitempty"`
}

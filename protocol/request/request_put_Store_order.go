package request

type RequestPutStoreOrder struct {
	ID string `json:"_id,omitempty"`
	// Status 접수중/접수취소/추가접수/접수중/조리중/배달중/배달완료
	Status string `json:"status,omitempty"`
}

package models

type CreateItemRequest struct {
	SellerID   *int       `json:"sellerID"`
	Name       *string    `json:"name"`
	Price      *int       `json:"price"`
	Statistics Statistics `json:"statistics"`
}

type Statistics struct {
	Likes     *int `json:"likes"`
	ViewCount *int `json:"viewCount"`
	Contacts  *int `json:"contacts"`
}

func IntPtr(i int) *int {
	return &i
}

func StringPtr(s string) *string {
	return &s
}

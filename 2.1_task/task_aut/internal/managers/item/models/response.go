package models

type CreateItemResponse struct {
	Status string `json:"status"`
	Result Result `json:"result"`
}

type Result struct {
	Message  string                 `json:"message"`
	Messages map[string]interface{} `json:"messages"`
}

type ItemResponse struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Price      int        `json:"price"`
	SellerID   int        `json:"sellerId"`
	CreatedAt  string     `json:"createdAt"`
	Statistics Statistics `json:"statistics"`
}

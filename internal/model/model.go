package model

type Order struct {
	OrderUID          string      `json:"order_uid" validate:"required"`
	TrackNumber       string      `json:"track_number" validate:"required"`
	Entry             string      `json:"entry" validate:"required"`
	Delivery          Delivery    `json:"delivery" validate:"required,dive"`
	Payment           Payment     `json:"payment" validate:"required,dive"`
	Items             []OrderItem `json:"items" validate:"required,dive"`
	Locale            string      `json:"locale" validate:"required"`
	InternalSignature string      `json:"internal_signature"`
	CustomerID        string      `json:"customer_id" validate:"required"`
	DeliveryService   string      `json:"delivery_service" validate:"required"`
	ShardKey          string      `json:"shardkey" validate:"required"`
	SmID              int         `json:"sm_id" validate:"checkint"`
	DateCreated       string      `json:"date_created" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	OofShard          string      `json:"oof_shard" validate:"required"`
}

type Delivery struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Zip     string `json:"zip" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}

type Payment struct {
	Transaction  string `json:"transaction" validate:"required"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency" validate:"required"`
	Provider     string `json:"provider" validate:"required"`
	Amount       int    `json:"amount" validate:"checkint"`
	PaymentDT    int64  `json:"payment_dt" validate:"checkint"`
	Bank         string `json:"bank" validate:"required"`
	DeliveryCost int    `json:"delivery_cost" validate:"checkint"`
	GoodsTotal   int    `json:"goods_total" validate:"checkint"`
	CustomFee    int    `json:"custom_fee" validate:"checkint"`
}

type OrderItem struct {
	ChrtID      int    `json:"chrt_id" validate:"checkint"`
	TrackNumber string `json:"track_number" validate:"required"`
	Price       int    `json:"price" validate:"checkint"`
	Rid         string `json:"rid" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Sale        int    `json:"sale" validate:"checkint"`
	Size        string `json:"size" validate:"required"`
	TotalPrice  int    `json:"total_price" validate:"checkint"`
	NmID        int    `json:"nm_id" validate:"checkint"`
	Brand       string `json:"brand" validate:"required"`
	Status      int    `json:"status" validate:"checkint"`
}

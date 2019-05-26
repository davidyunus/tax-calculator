package tax

import "time"

// Tax represent tax struct
type Tax struct {
	TaxID      int        `json:"taxId"`
	Name       string     `json:"name"`
	TaxCode    int        `json:"taxCode"`
	Type       string     `json:"type"`
	Refundable string     `json:"refundable"`
	Price      float32    `json:"price"`
	Tax        float32    `json:"tax"`
	Amount     float32    `json:"amount"`
	CreatedAt  string     `json:"createdAt"`
	UpdatedAt  string     `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}

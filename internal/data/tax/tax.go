package tax

import "time"

// Tax ...
type Tax struct {
	TaxID      int        `json:"taxId"`
	Name       string     `json:"name"`
	TaxCode    int        `json:"taxCode"`
	Type       string     `json:"type"`
	Refundable string     `json:"refundable"`
	Price      float32    `json:"price"`
	Tax        float32    `json:"tax"`
	Amount     float32    `json:"amount"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}

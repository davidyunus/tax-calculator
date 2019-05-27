package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/davidyunus/tax-calculator/internal/httpserver/response"

	"github.com/davidyunus/tax-calculator/internal/tax"
)

// TaxController ...
type TaxController struct {
	taxService *tax.Service
}

// NewTaxController ...
func NewTaxController(taxService *tax.Service) *TaxController {
	return &TaxController{
		taxService: taxService,
	}
}

// TaxRequest represent tax request from user
type TaxRequest struct {
	Name    string `json:"name"`
	TaxCode int    `json:"taxCode"`
	Price   int    `json:"price"`
}

// TaxResponse represent tax response
type TaxResponse struct {
	TaxID      int     `json:"taxId"`
	Name       string  `json:"name"`
	TaxCode    int     `json:"taxCode"`
	Type       string  `json:"type"`
	Refundable string  `json:"refundable"`
	Price      float32 `json:"price"`
	Tax        float32 `json:"tax"`
	Amount     float32 `json:"amount"`
}

func makeTaxResponse(tax *tax.Tax) *TaxResponse {
	return &TaxResponse{
		TaxID:      tax.TaxID,
		Name:       tax.Name,
		TaxCode:    tax.TaxCode,
		Type:       tax.Type,
		Refundable: tax.Refundable,
		Price:      tax.Price,
		Tax:        tax.Tax,
		Amount:     tax.Amount,
	}
}

func makeTaxResponseList(tax []*tax.Tax) []*TaxResponse {
	taxResp := []*TaxResponse{}
	for i := range tax {
		tax := &TaxResponse{
			TaxID:      tax[i].TaxID,
			Name:       tax[i].Name,
			TaxCode:    tax[i].TaxCode,
			Type:       tax[i].Type,
			Refundable: tax[i].Refundable,
			Price:      tax[i].Price,
			Tax:        tax[i].Tax,
			Amount:     tax[i].Amount,
		}
		taxResp = append(taxResp, tax)
	}
	return taxResp
}

// Create ...
func (tc *TaxController) Create(w http.ResponseWriter, r *http.Request) {

	var tax TaxRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tax)

	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tc.taxService.Create(context.Background(), tax.Name, tax.TaxCode, tax.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusCreated, nil)
}

// FindAll ...
func (tc *TaxController) FindAll(w http.ResponseWriter, r *http.Request) {
	tax, err := tc.taxService.FindAll(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, tax)
}

// FindByQuery ...
func (tc *TaxController) FindByQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	tax, err := tc.taxService.FindByQuery(context.Background(), query)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Fail to find tax data by query", err)
		return
	}
	response.JSON(w, http.StatusOK, tax)
}

// FindByKeys ...
func (tc *TaxController) FindByKeys(w http.ResponseWriter, r *http.Request) {
	tID := chi.URLParam(r, "taxId")
	taxID, err := strconv.Atoi(tID)

	tax, err := tc.taxService.FindByKeys(context.Background(), taxID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}
	response.JSON(w, http.StatusOK, tax)
}

// Update ...
func (tc *TaxController) Update(w http.ResponseWriter, r *http.Request) {
	tx := chi.URLParam(r, "taxId")
	taxID, err := strconv.Atoi(tx)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Fail to parse taxID", err)
		return
	}

	var tax *TaxRequest
	err = json.NewDecoder(r.Body).Decode(&tax)
	defer r.Body.Close()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal Server Error", err)
	}

	err = tc.taxService.Update(context.Background(), taxID, tax.Name, tax.TaxCode, tax.Price)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}
	response.JSON(w, http.StatusOK, err)
}

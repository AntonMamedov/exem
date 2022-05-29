package httpdelivery

import (
	"encoding/json"
	"net/http"
	"tester/models"
)

func (d Delivery) mul(w http.ResponseWriter, r *http.Request) {
	mulReq := models.MulRequest{}
	err := json.NewDecoder(r.Body).Decode(&mulReq)
	if err != nil {
		d.newErrorResponse(w, r, http.StatusUnprocessableEntity, err)
		return
	}
	resp, err := d.services.Mul(r.Context(), mulReq)
	if err != nil {
		d.newErrorResponse(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	d.newResponse(w, r, http.StatusCreated, resp)
}
package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	kt "github.com/chandrafortuna/assesment/domain/kitara"
)

type KitaraHandler struct {
	service kt.Service
}

func NewKitaraHandler(s kt.Service) *KitaraHandler {
	return &KitaraHandler{
		service: s,
	}
}

func (h *KitaraHandler) Init(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req kt.ProductVariant
	if err := decoder.Decode(&req); err != nil {
		log.Println("Error:", err)
		ReplyError(&w, http.StatusBadRequest, "Create Product failed : ", err)
		return
	}

	res, err := h.service.InitData(req)
	if err != nil {
		ReplyError(&w, http.StatusBadRequest, "Create Product failed : ", err)
		return
	}

	ReplySuccess(&w, &res)
}

func (h *KitaraHandler) Reserve(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req kt.ReserveProductReq
	if err := decoder.Decode(&req); err != nil {
		log.Println("Error:", err)
		ReplyError(&w, http.StatusBadRequest, "Reserve failed : ", err)
		return
	}

	err := req.Validate()
	if err != nil {
		ReplyError(&w, http.StatusBadRequest, "Reserve failed : ", err)
		return
	}

	res, err := h.service.Reserve(req)
	if err != nil {
		log.Println("Error:", err)
		ReplyError(&w, http.StatusBadRequest, "Reserve failed : ", err)
		return
	}

	log.Println(fmt.Sprintf("Order Status: %v, Qty: %v", res.Success, res.Data.Qty))
	log.Println("Error:", err)
	ReplySuccess(&w, &res)
}

func (h *KitaraHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetAll()
	if err != nil {
		log.Println("Error:", err)
		ReplyError(&w, http.StatusInternalServerError, "Get Product Variant Failed:", err)
		return
	}

	ReplySuccess(&w, res)
}

func (h *KitaraHandler) Clear(w http.ResponseWriter, r *http.Request) {
	status, err := h.service.Clear()
	res := kt.Response{
		Success: status,
		Message: err.Error(),
	}
	ReplySuccess(&w, res)
}

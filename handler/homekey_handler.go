package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	hk "github.com/chandrafortuna/assesment/domain/homekey"
)

type HomekeyHandler struct {
	service hk.Service
}

func NewHomekeyHandler(s hk.Service) *HomekeyHandler {
	return &HomekeyHandler{
		service: s,
	}
}

func (h *HomekeyHandler) GetSolution(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req hk.MatrixPoint
	if err := decoder.Decode(&req); err != nil {
		log.Println("Error:", err)
		ReplyError(&w, http.StatusBadRequest, "Get Solution failed : ", err)
		return
	}

	err := req.Validate()
	if err != nil {
		log.Println("Error:", err)
		ReplyError(&w, http.StatusBadRequest, "Reserve failed : ", err)
		return
	}

	res, err := h.service.Solution(req)
	if err != nil {
		log.Println("Error:", err)
		ReplyError(&w, http.StatusBadRequest, "Get Solution failed : ", err)
		return
	}

	log.Println(fmt.Sprintf("Success Total Dot: %v", res.TotalStep))
	ReplySuccess(&w, &res)
}

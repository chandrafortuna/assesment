package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	ball "github.com/chandrafortuna/assesment/domain/ball"
)

type BallHandler struct {
	service ball.Service
}

func NewBallHandler(s ball.Service) *BallHandler {
	return &BallHandler{
		service: s,
	}
}

func (h *BallHandler) Init(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req ball.InitRequest
	if err := decoder.Decode(&req); err != nil {
		log.Println("Error:", err)
		ReplyError(&w, http.StatusBadRequest, "Init failed : ", err)
		return
	}

	err := req.Validate()
	if err != nil {
		ReplyError(&w, http.StatusBadRequest, "Init failed : ", err)
		return
	}

	res, err := h.service.Init(req.NumContainer, req.MaxLoad)
	if err != nil {
		ReplyError(&w, http.StatusInternalServerError, "failed to write message: ", err)
		return
	}

	ReplySuccess(&w, &res)
}

func (h *BallHandler) AddBallToContainer(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.AddBallToContainer()
	if err != nil {
		log.Println("Error:", err)
		ReplyError(&w, http.StatusInternalServerError, "Add Ball To Container Failed:", err)
		return
	}

	log.Println(fmt.Sprintf("Add Ball To Container success container ID:%v isVerified:%v ", res.ID, res.IsVerified()))
	ReplySuccess(&w, res)
}

func (h *BallHandler) GetAllContainers(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetAllContainer()
	if err != nil {
		ReplyError(&w, http.StatusInternalServerError, "Get Container Failed:", err)
		return
	}

	ReplySuccess(&w, res)
}

func (h *BallHandler) GetVerifiedContainer(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetVerifiedContainer()
	if err != nil {
		ReplyError(&w, http.StatusInternalServerError, "Get Container Failed:", err)
		return
	}

	ReplySuccess(&w, res)
}

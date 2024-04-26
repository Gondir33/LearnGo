package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"petstore/internal/infrastructure/responder"
	"petstore/internal/models"
	service "petstore/internal/modules/store/service"
	"strconv"

	"github.com/go-chi/chi"
)

type Storere interface {
	CreateOrderHandler(w http.ResponseWriter, r *http.Request)
	FindOrderHandler(w http.ResponseWriter, r *http.Request)
	DeleteOrderHandler(w http.ResponseWriter, r *http.Request)
	GetInventoryOrderHandler(w http.ResponseWriter, r *http.Request)
}

type Store struct {
	store   service.Storere
	respond responder.Responder
}

func NewStoreHandler(store service.Storere, respond responder.Responder) Storere {
	return &Store{
		store:   store,
		respond: respond,
	}
}

//	@Summary	Place an order for a pet
//	@Tags		store
//	@Accept		json
//	@Produce	json
//	@Param		request	body	CreateOrderRequest	true	"Order Info"
//	@Success	200
//	@Router		/store/order [post]
func (s *Store) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateOrderRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		s.respond.ErrorInternal(w, err)
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		s.respond.ErrorBadRequest(w, err)
		return
	}
	responseBody, err := s.store.Create(models.Order(reqBody))
	if err != nil {
		s.respond.ErrorInternal(w, err)
		return
	}
	s.respond.OutputJSON(w, responseBody)
}

//	@Summary		Find purchase order by ID
//	@Description	For valid response try integer IDs with value >= 1 and <= 10. Other values will generated exceptions
//	@Tags			store
//	@Accept			json
//	@Produce		json
//	@Param			orderId	path		int	true	"1"
//	@Success		200		{object}	models.Order
//	@Router			/store/order/{orderId} [get]
func (s *Store) FindOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "orderId")
	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		s.respond.ErrorBadRequest(w, err)
		return
	}
	if orderIdInt <= 1 && orderIdInt >= 10 {
		s.respond.ErrorBadRequest(w, errors.New("for valid response try integer IDs with value >= 1 and <= 10"))
		return
	}
	responseBody, err := s.store.Find(orderIdInt)
	if err != nil {
		s.respond.ErrorInternal(w, err)
		return
	}
	s.respond.OutputJSON(w, responseBody)
}

//	@Summary	Delete purchase order by ID
//	@Tags		store
//	@Accept		json
//	@Produce	json
//	@Param		orderId	path	int	true	"1"
//	@Success	200
//	@Router		/store/order/{orderId} [delete]
func (s *Store) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "orderId")
	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		s.respond.ErrorBadRequest(w, err)
		return
	}
	err = s.store.Delete(orderIdInt)
	if err != nil {
		s.respond.ErrorInternal(w, err)
		return
	}
	s.respond.OutputJSON(w, "")
}

//	@Summary	Returns pet inventory by status
//	@Tags		store
//	@Security	ApiKeyAuth
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	map[string]int
//	@Router		/store/inventory [get]
func (s *Store) GetInventoryOrderHandler(w http.ResponseWriter, r *http.Request) {
	responseBody, err := s.store.GetInventory()
	if err != nil {
		s.respond.ErrorInternal(w, err)
	}
	s.respond.OutputJSON(w, responseBody)
}

type CreateOrderRequest models.Order

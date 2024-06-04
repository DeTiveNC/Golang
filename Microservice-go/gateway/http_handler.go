package main

import (
	"errors"
	common "github.com/DeTiveNC/commons"
	pb "github.com/DeTiveNC/commons/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

type Handler struct {
	// gateway
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *Handler {
	return &Handler{client}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	// gateway
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *Handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleCreateOrder")

	customerID := r.PathValue("customerID")
	var items []*pb.ItemsWithQuatity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, "Failed to read items")
		return
	}

	if err := validateItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	o, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
	rStatus := status.Convert(err)
	if rStatus != nil {
		if rStatus.Code() == codes.InvalidArgument {
			common.WriteError(w, http.StatusNotFound, rStatus.Message())
			return
		}

		common.WriteError(w, http.StatusInternalServerError, "Failed to create order")
		return
	}

	common.WriteJSON(w, http.StatusOK, o)
}

func validateItems(items []*pb.ItemsWithQuatity) error {
	if len(items) == 0 {
		return common.ErrNoItems
	}

	for _, item := range items {
		if item.ID == "" {
			return errors.New("item ID is required")
		}
		if item.Quantity <= 0 {
			return errors.New("item quantity must be greater than 0")
		}
	}

	return nil
}

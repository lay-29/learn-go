package handler

import (
	"encoding/json"
	"learn-go/No8/middleware"
	"learn-go/No8/store"
	"net/http"
	"strconv"
)

type CreateOrderReq struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value(middleware.UserIDKey).(string)
	userID, _ := strconv.ParseInt(userIDStr, 10, 64)

	var req CreateOrderReq
	json.NewDecoder(r.Body).Decode(&req)

	o := store.CreatOrder(userID, req.Title, req.Price)
	json.NewEncoder(w).Encode(o)
}

func MyOrders(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value(middleware.UserIDKey).(string)
	userID, _ := strconv.ParseInt(userIDStr, 10, 64)

	list := store.GetOrderByUser(userID)
	json.NewEncoder(w).Encode(list)
}

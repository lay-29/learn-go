package handler

import (
	"encoding/json"
	"learn-go/No8/store"
	"net/http"
	"strconv"
)

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Request Body Decoder error", http.StatusBadRequest)
		return
	}

	if store.GetUserByName(req.Username) != nil {
		http.Error(w, "user exists", http.StatusBadRequest)
		return
	}
	u := store.CreatUser(req.Username, req.Password)
	json.NewEncoder(w).Encode(u)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req RegisterReq
	json.NewDecoder(r.Body).Decode(&req)

	u := store.GetUserByName(req.Username)
	if u == nil || u.Password != req.Password {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	// demo token：直接用 userID
	token := strconv.FormatInt(u.ID, 10)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

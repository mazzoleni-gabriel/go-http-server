package handler

import (
	"encoding/json"
	"github.com/mazzoleni-gabriel/go-http-server/user"
	"github.com/mazzoleni-gabriel/go-http-server/user/service"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//Decode json into user
	var u user.User
	json.NewDecoder(req.Body).Decode(&u)

	u = service.RenameUser(u)

	// Decode user into json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(u)
	_, _ = w.Write(b)
}

package accounts

import (
  "log"
  "net/http"
  "github.com/alexkarpovich/mowy-api-go/lib"
  "github.com/alexkarpovich/mowy-api-go/database/users"
)

func (h *AccountHandler) View(w http.ResponseWriter, r *http.Request) {
  user, _ := r.Context().Value("user").(*users.User)
  log.Printf("LOGGED_IN_USER: %s", user)

  lib.SendJson(w, user, http.StatusOK)
}

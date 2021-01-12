package handlers

import (
	"go-rest2/models"
	"log"
	"net/http"
	"os"
)

type AdminHandler struct {
	admin *models.Admin
}

func NewAdmin() *AdminHandler {
	password := os.Getenv("ADMIN_PASSWORD")
	if password == "" {
		log.Fatal("Password not set")
	}
	return &AdminHandler{
		admin: &models.Admin{Password: password},
	}
}

func (a *AdminHandler) GetAdmin(w http.ResponseWriter, r *http.Request) {
	un, pass, ok := r.BasicAuth()
	if !ok || un != "admin" || pass != a.admin.Password {
		http.Error(w, "ERROR 401 - unauthorized access", http.StatusUnauthorized)
		return
	}
	w.Write([]byte("Welcome admin"))
}

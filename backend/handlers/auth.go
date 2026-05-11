package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"admin-api/database"
	"admin-api/middleware"
	"admin-api/models"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string

func SetJWTSecret(s string) {
	jwtSecret = s
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "request tidak valid", http.StatusBadRequest)
		return
	}

	var user models.User
	err := database.DB.Get(&user,
		"SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL",
		req.Email,
	)
	if err != nil {
		http.Error(w, "email atau password salah", http.StatusUnauthorized)
		return
	}

	if !user.CheckPassword(req.Password) {
		http.Error(w, "email atau password salah", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "gagal generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loginResponse{Token: tokenStr, User: user})
}

func Me(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)

	var user models.User
	err := database.DB.Get(&user,
		"SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL",
		userID,
	)
	if err != nil {
		http.Error(w, "user tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

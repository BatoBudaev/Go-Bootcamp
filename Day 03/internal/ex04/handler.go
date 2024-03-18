package ex04

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/BatoBudaev/Go-Bootcamp/internal/elastic"
	"github.com/BatoBudaev/Go-Bootcamp/internal/ex03"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
	"time"
)

var secretKey []byte

func GenerateRandomKey(length int) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func init() {
	var err error
	secretKey, err = GenerateRandomKey(32)
	if err != nil {
		log.Fatalf("Не удалось сгенерировать секретный ключ: %v", err)
	}
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	credentials := &Credentials{
		Username: "User",
		Password: "password",
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"token": tokenString}
	json.NewEncoder(w).Encode(response)
}

func RecommendHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Требуется авторизация", http.StatusUnauthorized)
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Неподдерживаемый метод подписи: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		http.Error(w, "Неверные аутентификационные данные", http.StatusUnauthorized)
		return
	}

	es := elastic.InitClient()
	store := &elastic.ElasticsearchStore{Client: es}

	handler := ex03.HandleRecommendationRequest(store)
	handler(w, r)
}

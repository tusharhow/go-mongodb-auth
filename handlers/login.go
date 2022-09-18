package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	db "github.com/tusharhow/go-api/db"
	model "github.com/tusharhow/go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)


func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user model.User
	var dbUser model.User
	json.NewDecoder(request.Body).Decode(&user)
	collection := db.MGI.Db.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		log.Println(passErr)
		response.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}

	token, err := GenerateToken()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	response.Write([]byte(`{"token":"` + token + `"}`))
}

func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(db.SECRET_KEY)
	if err != nil {
		log.Println("Error in generating token")
		return "", err
	}
	return tokenString, nil
}

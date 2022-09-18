package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	db "github.com/tusharhow/go-api/db"
	model "github.com/tusharhow/go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)


func Signup(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var user model.User

	json.NewDecoder(request.Body).Decode(&user)
	user.Password = getHash([]byte(user.Password))
	collection := db.MGI.Db.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user.ID = primitive.NewObjectID().Hex()
	validate := validateUser(user)
	if validate == true {
		response.Write([]byte(`{"response":"User Already Exists!"}`))
		return
	}
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(result)

}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func validateUser(user model.User) bool {
	collection := db.MGI.Db.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var dbUser model.User
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)
	if err != nil {
		return false
	}
	println("User already exists")
	return true
}

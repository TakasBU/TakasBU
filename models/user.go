package models

import (
	"log"

	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name             string    `gorm:"type:varchar(255);not null"`
	Email            string    `gorm:"uniqueIndex;not null"`
	Password         string    `gorm:"not null"`
	Role             string    `gorm:"type:varchar(255);not null"`
	Provider         string    `gorm:"not null"`
	Photo            string    `gorm:"not null"`
	VerificationCode string
	Verified         bool      `gorm:"not null"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
}

type SignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	Photo    string `json:"photo" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var SECRET_KEY = []byte("gosecretkey")

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

func CreateUser(db *gorm.DB, User *User) (err error) {

	err = db.Table("users").Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserById(db *gorm.DB, User *User, id int) (err error) {
	err = db.Table("users").Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Table("users").Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Table("users").Save(User)
	return nil
}

func DeleteUser(db *gorm.DB, User *User, id int) (err error) {
	db.Table("users").Where("id = ?", id).Delete(User)
	return nil
}

/*
package main

import (
	"context"
	"time"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)


var SECRET_KEY = []byte("gosecretkey")

type User struct{
	FirstName string `json:"firstname" bson:"firstname"`
	LastName string `json:"lastname" bson:"lastname"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

var client *mongo.Client

func getHash(pwd []byte) string {
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    return string(hash)
}

func GenerateJWT()(string,error){
	token:= jwt.New(jwt.SigningMethodHS256)
	tokenString, err :=  token.SignedString(SECRET_KEY)
	if err !=nil{
		log.Println("Error in JWT token generation")
		return "",err
	}
	return tokenString, nil
}

func userSignup(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-Type","application/json")
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	user.Password = getHash([]byte(user.Password))
	collection := client.Database("GODB").Collection("user")
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	result,_ := collection.InsertOne(ctx,user)
	json.NewEncoder(response).Encode(result)
}


func userLogin(response http.ResponseWriter, request *http.Request){
  response.Header().Set("Content-Type","application/json")
  var user User
  var dbUser User
  json.NewDecoder(request.Body).Decode(&user)
  collection:= client.Database("GODB").Collection("user")
  ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
  err:= collection.FindOne(ctx, bson.M{"email":user.Email}).Decode(&dbUser)

  if err!=nil{
	  response.WriteHeader(http.StatusInternalServerError)
	  response.Write([]byte(`{"message":"`+err.Error()+`"}`))
	  return
  }
  userPass:= []byte(user.Password)
  dbPass:= []byte(dbUser.Password)

  passErr:= bcrypt.CompareHashAndPassword(dbPass, userPass)

  if passErr != nil{
	  log.Println(passErr)
	  response.Write([]byte(`{"response":"Wrong Password!"}`))
	  return
  }
  jwtToken, err := GenerateJWT()
  if err != nil{
	response.WriteHeader(http.StatusInternalServerError)
	response.Write([]byte(`{"message":"`+err.Error()+`"}`))
	return
  }
  response.Write([]byte(`{"token":"`+jwtToken+`"}`))

}


func main(){
	log.Println("Starting the application")

	router:= mux.NewRouter()
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	client,_= mongo.Connect(ctx,options.Client().ApplyURI("mongodb://localhost:27017"))

	router.HandleFunc("/api/user/login",userLogin).Methods("POST")
	router.HandleFunc("/api/user/signup",userSignup).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}
*/

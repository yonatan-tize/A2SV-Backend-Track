package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"task_manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client
var userCollection *mongo.Collection
var SECRET_KEY = []byte("MY-Secret-Key")

type Claims struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Role     string `json:"role"`
    jwt.StandardClaims
}

func init() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("database not connected")
	}
	fmt.Println("connected to database")
	userCollection = client.Database("User_registration").Collection("users")
	taskCollection = client.Database("TaskManagers").Collection("tasks")

}

var validate = validator.New()

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func verifyPassword(userPassword string, foundPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(foundPassword), []byte(userPassword))
	return err == nil
}

func generateTokens(user models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours

	claims := &Claims{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	//generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret key
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateAccount(user models.User) (models.User, error) {
	//validate if the input is compatible with the struct
	err := validate.Struct(user)
	if err != nil {
		return models.User{}, err
	}

	// check if the user exist
	var existingUser models.User
	err = userCollection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return models.User{}, errors.New("username already exists")
	}

	// hash the user password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = hashedPassword

	// count users in the database. if no user make the user admin
	count, _ := userCollection.CountDocuments(context.TODO(), bson.M{})
	if count == 0 {
		// Promote the first user to admin
		user.Role = "ADMIN"
	}

	if user.Role == "" {
		user.Role = "USER"
	}

	user.ID = primitive.NewObjectID()

	// Set timestamps
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = userCollection.InsertOne(context.Background(), user)
	return user, err
}

func AuthenticateUser(userName string, password string) (models.User, string, error) {
	//find the user name
	var foundUser models.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": userName}).Decode(&foundUser)
	if err != nil {
		return models.User{}, "", err
	}

	// check if the password is the same
	isValidPassword := verifyPassword(password, foundUser.Password)
	if !isValidPassword {
		return models.User{}, "", errors.New("incorrect password")
	}
	token, err := generateTokens(foundUser)
	if err != nil {
		return models.User{}, "", err
	}

	return foundUser, token, nil
}

func UpdateUserRoll(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"role": "ADMIN",
		},
	}
	result, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments // No user found with the given ID
	}
	return nil
}

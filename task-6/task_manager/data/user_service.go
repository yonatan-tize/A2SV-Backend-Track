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

// init initializes the MongoDB client and establishes a connection to the database.
// It sets up the necessary client options and connects to the MongoDB server running on localhost:27017.
// If the connection is successful, it prints a message indicating that it is connected to the database.
// It also sets the userCollection and taskCollection variables to reference the respective collections in the database.
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

// hashPassword hashes the plain password to human unreadable format
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

//checks the users password with the one in the database
func verifyPassword(userPassword string, foundPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(foundPassword), []byte(userPassword))
	return err == nil
}

// generateTokens generates a JWT token for the given user.
// It takes a user model as input and returns the generated token string and an error, if any.
// The token is valid for 24 hours.
// The token contains the user's ID, username, role, and expiration time.
// It uses the HS256 signing method and signs the token with the SECRET_KEY.
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


// CreateAccount creates a new user account in the system.
// It takes a user model as input and returns the created user model and an error, if any.
// The function performs the following steps:
// 1. Validates if the input user is compatible with the user model struct.
// 2. Checks if the user already exists in the database by querying the username.
// 3. Hashes the user's password for security.
// 4. If there are no existing users in the database, the first user is promoted to an admin role.
// 5. If the user's role is not specified, it is set to "USER" by default.
// 6. Generates a new unique ID for the user.
// 7. Sets the created and updated timestamps for the user.
// 8. Inserts the user into the database.
// The function returns the created user model and any error that occurred during the process.
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
	}else {
		user.Role = "USER"
	}

	user.ID = primitive.NewObjectID()

	// Set timestamps
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = userCollection.InsertOne(context.Background(), user)
	return user, err
}

// AuthenticateUser authenticates a user by their username and password.
// It searches for the user with the given username in the user collection and verifies the password.
// If the username or password is incorrect, it returns an empty user model, an empty token, and an error.
// If the authentication is successful, it returns the authenticated user model, a token, and no error.
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

// UpdateUserRoll updates the role of a user with the given ID to "ADMIN".
// It takes the ID of the user as a parameter and returns an error if any occurred.
// If no user is found with the given ID, it returns mongo.ErrNoDocuments.
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

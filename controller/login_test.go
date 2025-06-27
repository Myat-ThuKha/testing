package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing-api/database"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func seedTestUser(t *testing.T, db *database.DB) {
	pass := "mypassword"
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	_, err := db.Client.Database(db.Name).Collection("users").InsertOne(context.TODO(), bson.M{
		"_id":             "u1",
		"username":        "testuser",
		"hashed_password": string(hashed),
	})
	assert.NoError(t, err)
}

func clearUsers(t *testing.T, db *database.DB) {
	_ = db.Client.Database(db.Name).Collection("users").Drop(context.TODO())
}

func TestLogin_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, err := database.ConnectMongo()
	assert.NoError(t, err)
	clearUsers(t, db)
	seedTestUser(t, db)

	router := gin.Default()
	router.POST("/login", LoginHandler(db))

	body := `{"username":"testuser","password":"mypassword"}`
	req := httptest.NewRequest("POST", "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "access_token")
}

func TestLogin_MissingUsername(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, err := database.ConnectMongo()
	assert.NoError(t, err)

	clearUsers(t, db)
	seedTestUser(t, db)

	router := gin.Default()
	router.POST("/login", LoginHandler(db))

	body := `{"password":"mypassword"}`
	req := httptest.NewRequest("POST", "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid input")
}

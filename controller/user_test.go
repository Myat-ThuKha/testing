package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing-api/database"
	"testing-api/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUserController_Create_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, err := database.ConnectMongo()
	assert.NoError(t, err)

	router := gin.Default()
	router.POST("/user", Create(db))

	// Input request
	reqBody := `{
		"companyId": "company123",
		"username": "newuser",
		"password": "pass123"
		"email": "newuser@example.com",
		"fullName": "New User"
	}`

	hashed_password, err := bcrypt.GenerateFromPassword([]byte("pass123"), bcrypt.DefaultCost)
	assert.NoError(t, err)

	// Expected user returned from CreateUser (simulate DB result)
	expectedUser := &model.CreatedUser{
		Id:             "user123",
		Username:       "newuser",
		Email:          "newuser@example.com",
		FullName:       "New User",
		CompanyId:      "company123",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		HashedPassword: string(hashed_password),
		// HashedPassword: "hashed_password", // Uncomment if you want to include hashed
		// add other fields if needed
	}

	// Create HTTP request
	req := httptest.NewRequest("POST", "/users", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Record response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)

	// Optionally decode response body
	var resp map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, resp)
	// assert.Equal(t, expectedUser.Username, resp["username"])
	// assert.Equal(t, expectedUser.Email, resp["email"])
}

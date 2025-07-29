package controller

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing-api/database"
	"testing-api/model"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Create(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashed_password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		collection := db.Client.Database(db.Name).Collection("users")

		_, err = collection.InsertOne(context.Background(), bson.M{
			"company_id":      req.CompanyId,
			"username":        req.Username,
			"hashed_password": string(hashed_password),
			"email":           req.Email,
			"full_name":       req.FullName,
			"created_at":      time.Now(),
			"updated_at":      time.Now(),
		})
		if err != nil {
			if mongo.IsDuplicateKeyError(err) {
				if strings.Contains(err.Error(), "username_1") {
					c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("user with username %s already exists", req.Username)})
				}
				if strings.Contains(err.Error(), "email_1") {
					c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("user with email %s already exists", req.Email)})
				}
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error creating user with username %s: %v", req.Username, err)})
		}

		c.JSON(http.StatusCreated, "User created successfully")
	}
}

package handlers

import (
	"net/http"
	"time"
	"todo-list-api/db"
	"todo-list-api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var Jwtkey = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0X2tleSIsInNjb3BlIjplbmQwfQ.dhfurhfu4hfurhfu4hfurhf4uhfu4hf4uhfu4hf4uhf4uhfurhf4uhfurhf4uhfurhf4uhf")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	user.Password = string(hashedPassword)
	_, err = db.DB.Exec("insert into user (username , password) values (?,?)", user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User register successfuly"})
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var storedUser models.User
	err := db.DB.QueryRow("select id , username , password from user where username = ?", user.Username).Scan(&storedUser.Id, &storedUser.Username, &storedUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Credentialse"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Credentialse"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(Jwtkey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

package utils

import (
	"errors"
	"os"
	"time"

	"ynbauth/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateAccessToken creates a short-lived access token.
func GenerateAccessToken(email string) (string, error) {
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // Access token expires in 15 minutes
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// GenerateRefreshToken creates a long-lived refresh token.
func GenerateRefreshToken(email string) (string, error) {
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // Refresh token expires in 7 days
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ParseToken verifies a token and returns its claims.
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")

}

// ObtainToken handles login and generates both access and refresh tokens.
func ObtainToken(c *gin.Context) {
	
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Find the user by email
	user, err := models.FindUser(req.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare the provided password with the hashed password from the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate Access and refresh tokens
	accessToken, err := GenerateAccessToken(req.Email)
	if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate access token"})
			return
	}

	refreshToken, err := GenerateRefreshToken(req.Email)
	if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate refresh token"})
			return
	}

	c.JSON(200, gin.H{"access":  accessToken,"refresh": refreshToken})
	return
}

// RefreshToken generates a new access token using a valid refresh token.
func RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	claims, err := ParseToken(req.RefreshToken)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	// Generate a new access token
	accessToken, err := GenerateAccessToken(claims.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate access token"})
		return
	}

	c.JSON(200, gin.H{"access": accessToken})
}

// Verify Access Token
func VerifyToken(c *gin.Context) {
	var req struct {
		Token string `json:"access"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	_, err := ParseToken(req.Token)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid or expired token"})
		return
	}

	c.JSON(200, gin.H{"message": "Token is valid"})
}

package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tpSpace/go-jwt/initializers"
	"github.com/tpSpace/go-jwt/models"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("Authenticating...")
	// Get cookie from request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Decode and verify the JWT
	// Parse takes the token string and a function for looking up the key. The latter is especially

	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check the expiration date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println("Token expired")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with token sub 
		var user models.User
		initializers.DB.Where("ID = ?", claims["sub"]).First(&user)
		
		if user.ID == 0 {
			fmt.Println("User not found")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach request to context
		c.Set("user", user)

		// Call the next middleware
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	
}
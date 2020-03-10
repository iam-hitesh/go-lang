package Services


import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"../Models"
)


type Claims struct {
	Email string `json:"email"`
	IsActive bool `json:"IsActive"`
	Role int `json:"Role"`
	IsAdmin bool `json:"IsAdmin"`
	jwt.StandardClaims
}


type Header struct {
	Token   int    `header:"token"`
	Authorization string `header:"authorization"`
}


func GenerateToken(user *Models.User) string {
	// We will kept it for next 5 Minutes
	expirationTime := time.Now().Add(500 * time.Minute)

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &Claims{
		Email: user.Email,
		IsActive: user.IsActive,
		Role: user.Role,
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return tokenString
}


func ParseAndValidateToken(tokenString string) map[string]interface{} {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		User := map[string]interface{}{
			"Email":          claims["email"].(string),
			"IsActive":       claims["IsActive"].(bool),
			"Role":           int(claims["Role"].(float64)),
			"IsAdmin":        claims["IsAdmin"].(bool),
		}

		return User
	}

	fmt.Println(err)
	return nil
}


func TokenAuthMiddleware() gin.HandlerFunc {
	return func(req *gin.Context) {
		h := Header{}

		if err := req.ShouldBindHeader(&h); err != nil {
			req.JSON(200, err)
		}

		tokenSlice := strings.Split(h.Authorization, " ")

		// Token should be in `Bearer token_string` format
		if len(tokenSlice) != 2 {
			req.Abort()

			req.JSON(http.StatusUnauthorized, gin.H{
				"message": "Improper token string format",
			})

			return
		}

		jwtToken := tokenSlice[1]
		user := ParseAndValidateToken(jwtToken)

		if user == nil {
			req.Abort()
			//req.Writer.WriteHeader(http.StatusUnauthorized)

			req.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})

			return
		}

		req.Set("currentUser", user)
		req.Next()
	}
}
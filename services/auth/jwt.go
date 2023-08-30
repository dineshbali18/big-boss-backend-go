package auth

// import (
// 	"big-boss-7/config"
// 	commonErrors "big-boss-7/helpers/errors"
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"super-app/helpers/utils"

// 	"big-boss-7/services/crypto"

// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/gommon/log"
// )

// type Claims struct {
// 	Auth string
// 	jwt.RegisteredClaims
// 	UserId      string
// 	DeviceID    string
// 	PhoneNumber string
// }

// // GenerateJWT generates a JWT signed by HS256 method. This JWT is then used across all requests pertaining to the user.
// // For the full list of URLs which use this JWT, refer to the list of URLs present in JWTSkipper
// func GenerateJWT(userID, phoneNumber, deviceID string) (string, error) {
// 	// Create new token object
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	// Add claims
// 	claims := token.Claims.(jwt.MapClaims)

// 	if userID != "" {
// 		claims["userId"] = userID
// 	}

// 	if phoneNumber != "" {
// 		claims["phoneNumber"] = phoneNumber
// 	}

// 	if deviceID != "" {
// 		claims["deviceId"] = deviceID
// 	}

// 	// Generate the final token string
// 	webToken, err := token.SignedString([]byte(config.AuthConfig.SecretKey))
// 	if err != nil {
// 		return "", err
// 	}

// 	// Encrypt the token
// 	encryptedToken, err := crypto.ServerEncrypt([]byte(webToken))
// 	if err != nil {
// 		log.Printf("encryption failed: %v", err)
// 		return "", err
// 	}

// 	return encryptedToken, nil
// }

// // JWTSkipper is used to define the routes which will be checked for JWT authentication
// func JWTURLchecker(context echo.Context) bool {
// 	// Define the URLs that require JWT authentication
// 	jwtURLs := []string{
// 		"/v1/superApps/users/address",
// 		"/v1/superApps/register/device",
// 		"/v1/superApps/users/cart/increment",
// 		"/v1/superApps/users/cart/decrement",
// 		"/v1/superApps/users/carts",
// 		"/v1/superApps/users/orders",
// 		"/v1/superApps/users/orders/:orderId",
// 		// Add more URLs here as needed
// 	}

// 	// Get the request URL path
// 	requestURL := context.Request().URL.Path

// 	log.Printf(requestURL)

// 	// Check if the request URL requires JWT authentication
// 	if utils.ContainsPartialMatch(jwtURLs, requestURL) {
// 		return false
// 	}

// 	// JWT authentication is not required for the request URL
// 	return true
// }

// // JWTErrorHandler is used to action if JWT authentication fails
// // On failure, show a generic error "not authorized" message to the user and log the actual error
// func JWTErrorHandler(err error) error {
// 	return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
// }

// // Middleware function to authenticate the endpoint with JWT
// func JWTAuthenticator(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(context echo.Context) error {

// 		if JWTURLchecker(context) {
// 			return next(context)
// 		}

// 		// Extract the JWT token from the request header
// 		authHeader := context.Request().Header.Get("jwt")
// 		if authHeader == "" {
// 			return echo.NewHTTPError(http.StatusUnauthorized, "Missing JWT token")
// 		}

// 		// Split the Authorization header into two parts: "Bearer" and the token
// 		tokenParts := strings.Split(authHeader, " ")
// 		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
// 			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid JWT token format")
// 		}
// 		token := tokenParts[1]

// 		// Decrypting the Token server-side
// 		decryptedToken, err := crypto.ServerDecrypt(string(token))
// 		if err != nil {
// 			log.Error("Unable to decrypt JWT token:", err)
// 			return echo.NewHTTPError(commonErrors.InvalidJWTToken.Status, commonErrors.InvalidJWTToken)
// 		}
// 		if token == "" {
// 			return echo.NewHTTPError(http.StatusUnauthorized, "Missing JWT token")
// 		}

// 		// Validate and parse the JWT token
// 		claims := &Claims{}
// 		jwtToken, err := jwt.ParseWithClaims(decryptedToken, claims, func(token *jwt.Token) (interface{}, error) {
// 			// Provide the secret key used for signing the JWT token
// 			return []byte(config.AuthConfig.SecretKey), nil
// 		})

// 		if err != nil {
// 			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid JWT token")
// 		}

// 		// Check if the JWT token is valid
// 		if !jwtToken.Valid {
// 			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid JWT token")
// 		}

// 		// Get the custom claims from token
// 		customClaims, ok := jwtToken.Claims.(*Claims)
// 		if !ok {
// 			return echo.NewHTTPError(commonErrors.InvalidJWTToken.Status, commonErrors.InvalidJWTToken)
// 		}

// 		// Extract necessary values from custom claims
// 		UserID := customClaims.UserId
// 		deviceID := customClaims.DeviceID
// 		phoneNumber := customClaims.PhoneNumber
// 		if UserID != "" {
// 			userIDStr := string(UserID)
// 			log.Printf(userIDStr)
// 			userID, err := strconv.ParseUint(userIDStr, 10, 64)
// 			if err != nil {
// 				log.Printf("Failed to convert UserID to uint: %v", err)
// 				return err
// 			}
// 			userId := uint(userID)

// 			// Set necessary values in the context for later use
// 			context.Set("X-User-ID", userId)
// 		}

// 		// Set necessary values in the context for later use
// 		context.Set("X-Device-ID", deviceID)
// 		// context.Set("X-User-ID", userId)
// 		context.Set("X-Phone-Number", phoneNumber)

// 		// Call the next handler
// 		return next(context)
// 	}
// }

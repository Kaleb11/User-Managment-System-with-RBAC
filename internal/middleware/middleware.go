package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func memberRoleHandler(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("ACCESS_SECRET")), nil
			})
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "props", claims)
				// Access context values in handlers like this
				role := claims["role"]
				if role == "member" {
					fmt.Println("User role : ", role)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					fmt.Println("User role : ", role)
					fmt.Println("Invalid role")
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized"))
				}

			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	}
}

// username := c.FormValue("username")
// password := c.FormValue("password")
// // Check in your db if the user exists or not
// if username == "jon" && password == "password" {
//     // Create token
//     token := jwt.New(jwt.SigningMethodHS256)
//     // Set claims
//     // This is the information which frontend can use
//     // The backend can also decode the token and get admin etc.
//     claims := token.Claims.(jwt.MapClaims)
//     claims["role"] = "admin"
//     claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
//     // Generate encoded token and send it as response.
//     // The signing string should be secret (a generated UUID          works too)
//     t, err := token.SignedString([]byte("secret"))
//     if err != nil {
//         return err
//     }
//     return c.JSON(http.StatusOK, map[string]string{
//         "token": t,
//     })

//member token
//role
//api return
// rolee, err = model.User.Role
// role, err := session.GetString(r, "role")
// if err != nil {
// 	writeError(http.StatusInternalServerError, "ERROR", w, err)
// 	return
// }
// writeSuccess(fmt.Sprintf("User with Role: %s", role), w)

func Authmiddleware(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("ACCESS_SECRET")), nil
			})
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "props", claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value("props").(jwt.MapClaims)

				role := claims["role"]
				if role == "admin" {
					fmt.Println("User role : ", role)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					fmt.Println("User role : ", role)
					fmt.Println("Invalid role")
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized"))
				}

			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	}
}

func adminHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//writeSuccess("I'm an Admin!", w)
	})
}

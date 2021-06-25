package router

import (
	"Auth/internal/middleware"
	"Auth/internal/user/handler"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NewRouter return all router
func NewRouter() *httprouter.Router {
	// http.Handle("/ping", ss.Authmiddleware(handler.AddUsers))
	router := httprouter.New()
	router.POST("/user", middleware.Authmiddleware(handler.AddUsers()))
	router.GET("/user", handler.GetUsers)
	//router.GET("/user/:id", handler.GetUsersByID)
	router.PUT("/user", middleware.Authmiddleware(handler.UpdateUsers()))
	router.DELETE("/user/:id", handler.DeleteUsers)
	router.PATCH("/user/photoupload/:id", handler.Photoupload)
	router.POST("/signup", handler.Signup)
	router.POST("/signin", handler.Signin)
	router.GET("/user/:username", handler.GetUsersByusername)
	router.NotFound = http.FileServer(http.Dir("C:\\Users\\Tilefamily\\Documents\\AuthHexa\\cmd\\rest\\router\\views"))
	//router.Handle("/", http.FileServer(http.Dir("C:\\Users\\Tilefamily\\Documents\\AuthHexa\\cmd\\rest\\router\\views")))
	//router.ServeFiles("views/*", http.Dir("/"))

	return router

}

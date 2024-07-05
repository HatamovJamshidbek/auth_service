package api

import (
	_ "github.com/gin-contrib/cors"
	"root/api/handler"
	"root/api/middleware"

	"github.com/gin-gonic/gin"
)

func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.MiddleWare())
	u := r.Group("/user")
	u.POST("/registr", h.RegisterUser)
	u.PUT("/update/:id", h.UpdateUser)
	u.DELETE("/delete/:id", h.DeleteUser)
	u.GET("/getall", h.GetAllUser)
	u.GET("/getbyid/:id", h.GetbyIdUser)
	u.POST("/login", h.LoginUser)
	return r
}

package handler

import (
	"errors"
	"fmt"
	"intern-bcc/internal/service"
	"intern-bcc/pkg/middleware"
	"intern-bcc/pkg/response"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	r.router.POST("/register", r.Register)
	routerGroup := r.router.Group("api/v1")
	r.router.Use(r.middleware.Timeout())

	//r.middleware.OnlyAdmin,
	routerGroup.GET("login-user", r.middleware.AuthenticateUser, r.middleware.OnlyAdmin, testGetLoginUser)
	routerGroup.GET("time-out", testTimeout)
	routerGroup.POST("/login", r.Login)
	routerGroup.GET("/get-user/:name", r.GetUserByName)
}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}

func testTimeout(ctx *gin.Context) {
	time.Sleep(3 * time.Second)

	response.Success(ctx, http.StatusOK, "success", nil)
}

func testGetLoginUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		response.Error(ctx, http.StatusUnauthorized, "test get login user", errors.New(" "))
		return
	}

	response.Success(ctx, http.StatusOK, "success", user)
}

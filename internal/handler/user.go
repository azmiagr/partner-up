package handler

import (
	// "intern-bcc/internal/service"
	"intern-bcc/model"
	"intern-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) Register(ctx *gin.Context) {
	param := model.UserRegister{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = r.service.User.Register(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to register new user", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "successfully register new user", nil)
}

func (r *Rest) Login(ctx *gin.Context) {
	param := model.UserLogin{}
	err := ctx.ShouldBindJSON(&param)

	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}
	result, err := r.service.User.Login(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to login user", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully login to system", result)

}

func (r *Rest) GetUserByName(ctx *gin.Context) {
	name := ctx.Param("name")

	user, err := r.service.User.GetUserByName(name)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get user", err)
		return
	}

	response.Success(ctx, http.StatusOK, "user found", user)
}

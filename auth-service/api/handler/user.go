package handler

import (
	"root/api/token"

	pb "root/genproto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(ctx *gin.Context) {
	arr := pb.User{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.User.CreateUser(ctx, &arr)
	if err != nil {
		panic(err)
	}
	t := token.GenereteJWTToken(&arr)
	ctx.JSON(200, t)
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	arr := pb.User{}
	arr.Id = ctx.Param("id")
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.User.UpdateUser(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	id := pb.IdRequest{Id: ctx.Param("id")}
	_, err := h.User.DeleteUser(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

func (h *Handler) GetAllUser(ctx *gin.Context) {
	user := &pb.User{}
	user.Email = ctx.Query("email")
	user.UserName = ctx.Query("user_name")

	res, err := h.User.GetAllUser(ctx, user)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

func (h *Handler) GetbyIdUser(ctx *gin.Context) {
	id := pb.IdRequest{Id: ctx.Param("id")}
	res, err := h.User.GetByIdUser(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

func (h *Handler) LoginUser(ctx *gin.Context) {
	user := &pb.User{}
	err := ctx.ShouldBindJSON(user)
	if err != nil {
		panic(err)
	}
	res, err := h.User.LoginUser(ctx, user)
	if err != nil {
		panic(err)
	}
	t := token.GenereteJWTToken(res)
	ctx.JSON(200, t)
}

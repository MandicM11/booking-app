package handler

import (
	"booking-app/server/user-service/dto"
	"booking-app/server/user-service/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignupAdmin(c *gin.Context) {
	var userDto dto.RegisterUserDTO
	if err := c.BindJSON(&userDto); err != nil {
		c.JSON(400, "Can't unmarshal request body, error: "+err.Error())
		return
	}
	createdUser, err := service.SignupAdmin(dto.RegisterUserDTOToUser(userDto))
	if err != nil {
		c.JSON(400, "Can't register user, error:"+err.Error())
		return
	}
	c.JSON(200, createdUser)
}

func SignupGuest(ctx *gin.Context) {
	var userDto dto.RegisterUserDTO
	if err := ctx.BindJSON(&userDto); err != nil {
		ctx.JSON(400, "Can't unmarshal request, error: "+err.Error())
		return
	}
	createdUser, err := service.SignupGuest(dto.RegisterUserDTOToUser(userDto))
	if err != nil {
		ctx.JSON(400, "Can't register user, error: "+err.Error())
		return
	}
	ctx.JSON(200, createdUser)
}

func Login(ctx *gin.Context) {
	loginInfo := dto.LoginDTO{}
	if err := ctx.ShouldBindJSON(&loginInfo); err != nil {
		ctx.JSON(400, "Can't unmarshal request body, error: "+err.Error())
		return
	}
	token, err := service.LoginUser(loginInfo)
	if err != nil {
		ctx.JSON(400, "Can't login user, error: "+err.Error())
		return
	}
	ctx.JSON(200, token)
}

func GetAll(ctx *gin.Context) {
	ctx.JSON(200, service.GetAll())
}

func GetById(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, "Can't convert to ObjectId, error: "+err.Error())
		return
	}
	user, err := service.GetById(id)
	if err != nil {
		ctx.JSON(400, "Can't get user by id, error: "+err.Error())
		return
	}
	ctx.JSON(200, user)
}

func GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := service.GetByEmail(email)
	if err != nil {
		ctx.JSON(400, "Can't get user by email, error: "+err.Error())
		return
	}
	ctx.JSON(200, user)
}

func UpdateUser(ctx *gin.Context) {
	var userDto dto.UpdateUserDTO
	if err := ctx.BindJSON(&userDto); err != nil {
		ctx.JSON(400, "Can't unmarshal request, error: "+err.Error())
		return
	}
	updatedUser, err := service.UpdateUser(dto.UpdateUserDTOtoUser(userDto))
	if err != nil {
		ctx.JSON(400, "Can't update user, error: "+err.Error())
		return
	}
	ctx.JSON(200, updatedUser)
}

// func DeleteUser(ctx *gin.Context) {
// 	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
// 	if err != nil {
// 		ctx.JSON(400, "Can't convert to ObjectId, error: "+err.Error())
// 		return
// 	}
// 	err = service.DeleteUser(id)
// 	if err != nil {
// 		ctx.JSON(400, "Can't delete user, error: "+err.Error())
// 		return
// 	}
// 	ctx.JSON(200, "User deleted successfully")
// }

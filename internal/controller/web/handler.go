package web

import (
	"api/internal/entity"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase entity.UserUseCase
}

func NewRouter(c *gin.Engine, us entity.UserUseCase) {
	handler := &UserHandler{
		UserUseCase: us,
	}
	c.GET("/:id", handler.GetByUID)
	c.GET("/users", handler.GetAll)
	c.GET("/", handler.Get)
	c.GET("/friends/:id", handler.GetAllFriends)

	c.POST("/create", handler.Create)
	c.POST("/make_friends", handler.MakeFriends)

	c.DELETE("/:id", handler.Delete)
	c.PUT("/:id", handler.Update)

}

func (u *UserHandler) GetByUID(c *gin.Context) {

	id := c.Param("id")
	user, err := u.UserUseCase.GetByUUID(c, id)
	if err != nil {
		c.JSON(http.StatusBadGateway, entity.ErrInternalServerError)
	}

	friends, err := u.UserUseCase.GetAllFriends(c, id)
	var responseFriends string
	for _, friend := range friends {
		responseFriends += fmt.Sprintf("Name: %v\n", friend.Name)
	}

	response := fmt.Sprintf("Name: %v\nAge: %v\nFriends:\n %+v\n", user.Name, user.Age, responseFriends)
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(response))
}

func (u *UserHandler) Get(c *gin.Context) {
	fmt.Println("Hello world")
	c.JSON(http.StatusOK, "Hello world")
}

func (u *UserHandler) Create(c *gin.Context) {
	var params entity.User

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Create handler", params.Name, params.Age)
	if _, err := u.UserUseCase.Set(c, &params); err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are created user"})
	return

}

func (u *UserHandler) GetAll(c *gin.Context) {
	response := ""
	users, err := u.UserUseCase.GetAll(c)

	if err != nil {
		c.JSON(http.StatusBadGateway, entity.ErrInternalServerError)
	}

	for _, user := range users {
		response += fmt.Sprintf("Name: %v\nAge: %v\n\n", user.Name, user.Age)
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(response))
}

func (u *UserHandler) Delete(c *gin.Context) {

	id := c.Param("id")

	if err := u.UserUseCase.Delete(c, id); err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "user was deleted"})
	return
}

func (u *UserHandler) Update(c *gin.Context) {

	id := c.Param("id")

	var params entity.User
	fmt.Println(params)
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(params)
	err := u.UserUseCase.Update(c, id, params.Age)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrBadParamInput)
		return
	}

	c.JSON(http.StatusOK, "user update age")
}

func (u *UserHandler) MakeFriends(c *gin.Context) {

	var params entity.FriendPool

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(params.SourceID, params.TargetID)

	err := u.UserUseCase.MakeFriends(c, params.SourceID, params.TargetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "make friends")
}

func (u *UserHandler) GetAllFriends(c *gin.Context) {
	uid := c.Param("id")

	friends, err := u.UserUseCase.GetAllFriends(c, uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var response string
	for i, _ := range friends {
		response += fmt.Sprintf("Name: %v\n Age: %v\n", friends[i].Name, friends[i].Age)
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(response))

}

func (u *UserHandler) toStringFriends(user *entity.User) string {
	return fmt.Sprintf("%v %v\n", user.Name, user.Age)
}

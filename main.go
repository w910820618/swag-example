// @title User-List API
// @version 1.0
// @description This is a sample User-list server.
// @license.name Github Repository
// @license.url https://github.com/w910820618/swag-example
// @BasePath /api/v1

package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"

	// for generate swagger ui
	_ "swag-example/docs"
)

func init() {
	flag.Set("logtostderr", "true") // 输出到标准错误流
	flag.Set("v", "3")              // 输出级别为 3
	flag.Parse()
	glog.Info("Program started.")
}

func MWHandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}

		statusCode := c.Writer.Status()
		if statusCode == http.StatusOK {
			statusCode = http.StatusInternalServerError
		}

		c.AbortWithStatusJSON(statusCode, gin.H{
			"message": err.Error(),
		})
	}
}

func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://127.0.0.1"
		},
		MaxAge: 12 * time.Hour,
	}))
	r.Use(MWHandleErrors())
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/user")
		{
			users.GET(":id", getUser)
			users.GET("", listUser)
			users.POST("", addUser)
			users.DELETE(":id", deleteUser)
			users.PUT(":id", updateUser)
		}
	}
	// must access /swagger/index.html
	// access /swagger will get 404
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	_ = r.Run()
}

// @Summary Create a user
// @Description Create a user
// @Tags user
// @Accept  json
// @Produce  json
// @Param   User    body User true "User"
// @Success 200 {string} string	"ok"
// @Router /user [post]
func addUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(err)
		return
	}
	db.Create(&user)
	c.JSON(http.StatusOK, user)
}

// @Summary Get a user
// @Description Get a user
// @Tags user
// @Accept  json
// @Produce  json
// @Param   id     path string true "id"
// @Success 200 {string} string	"ok"
// @Router /user/{id} [get]
func getUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	db.First(&user, id)
	c.JSON(http.StatusOK, user)
}

// @Summary Get all users
// @Description Get all users
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /user [get]
func listUser(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// @Summary Update a user
// @Description Update a user
// @Tags user
// @Accept  json
// @Produce  json
// @Param   User   body User true "User"
// @Param   id     path string true "id"
// @Success 200 {string} string	"ok"
// @Router /user/{id} [put]
func updateUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	db.First(&user, id)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(err)
		return
	}
	db.Model(&user).Update(&user)
	c.JSON(http.StatusOK, user)
}

// @Summary Delete a user
// @Description Delete a user
// @Tags user
// @Accept  json
// @Produce  json
// @Param   id     path string true "id"
// @Success 200 {string} string	"ok"
// @Router /user/{id} [delete]
func deleteUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	db.First(&user, id)
	db.Delete(&user)
	c.JSON(http.StatusOK, user)
}

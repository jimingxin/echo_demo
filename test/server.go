package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"github.com/labstack/echo/middleware"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

type (
	User struct {
		ID int `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int] *User{}
	seq = 1

)

func loginUser(c echo.Context) error  {
	token := jwt.New(jwt.SigningMethodES256)
	fmt.Println("token-->",token)
	return nil
}

func createUser(c echo.Context) error  {
	u := &User{
		ID: seq,
	}

	if err := c.Bind(u); err != nil {
		return err
	}

	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)

}

func getUser(c echo.Context) error  {
	id,_ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK,users[id])
}

func updateUser(c echo.Context) error {
	u := new(User)
	if err:= c.Bind(u); err != nil  {
		return err
	}

	id,_ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK,users[id])
}

func deleteUser(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))
	delete(users,id)
	return c.NoContent(http.StatusNoContent)
}


func main()  {
	e := echo.New()

	// 中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 路由
	e.POST("/users",createUser)
	e.GET("users/:id",getUser)
	e.PUT("/users/:id",updateUser)
	e.DELETE("/users/:id",deleteUser)
	e.Logger.Fatal(e.Start(":8087"))
}

package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"time"
	"net/http"
	"github.com/labstack/echo/middleware"
)

type jwtCustomClaims struct {
	Name string `json:"name"`
	Admin bool `json:"admin"`
	jwt.StandardClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon" && password =="shh!" {
		chaims := &jwtCustomClaims{
			"Jon Snow",
			true,
			jwt.StandardClaims{
				ExpiresAt:time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// 获取token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256,chaims)

		t,err := token.SignedString([]byte("secret"))
		if err != nil{
			 return err
		}

		return c.JSON(http.StatusOK,echo.Map{
			"token":t,
		})
	}
	return echo.ErrUnauthorized
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK,"Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK,"Welcom" + name + "!")
}

func main()  {

	e:= echo.New()

	// 中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login",login)

	e.GET("/",accessible)

	r:= e.Group("/restricted")
	config := middleware.JWTConfig{
		Claims:&jwtCustomClaims{},
		SigningKey:[]byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("",restricted)

	e.Logger.Fatal(e.Start(":8087"))
}
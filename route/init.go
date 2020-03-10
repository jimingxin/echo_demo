package route

import (
	"github.com/labstack/echo"
	"echo_demo/controller"
	"github.com/labstack/echo/middleware"
)

func init()  {
	e:= echo.New()

	//中间件基本认证
	//e.Use(middleware.BasicAuth(func(username,password string, c echo.Context) (bool, error) {
	//	if username == "jmx" && password == "jmx"{
	//		return true,nil
	//	}
	//	return false, nil
	//}))

	// JWT中间件
	//e.Use(middleware.JWT([]byte("secret")))

	// log中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")
	e.GET("/allClass",controller.GetAllClass)
	e.GET("/getIdClass",controller.GetQueryClass)
	e.POST("/insertClass",controller.InsertClass)
	e.POST("/updateClass",controller.UpdateClass)
	e.Logger.Fatal(e.Start(":8087"))
}


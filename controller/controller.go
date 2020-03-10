package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"echo_demo/model"
	"echo_demo/util"
	"fmt"
	"encoding/json"
	"strconv"
)

func GetAllClass(c echo.Context) error {
	results := model.GetClassAll()

	if len(results) == 0 {
		return model.Err
	} else {
		return c.JSON(http.StatusOK, util.NewResult(util.Success, "获取数据成功", results))
	}
}


func GetQueryClass(c echo.Context) error {

	id_query := c.QueryParam("id")

	id,_ := strconv.Atoi(id_query)

	results := model.GetQueryClass(id)

	if len(results) == 0 {
		return model.Err
	} else {
		return c.JSON(http.StatusOK, util.NewResult(util.Success, "获取数据成功", results))
	}
}

/*
	进行数据的插入
*/
func InsertClass(c echo.Context) error {

	//name := c.QueryParam("name")
	//desc := c.QueryParam("desc")
	// post form请求
	name := c.FormValue("name")
	desc := c.FormValue("desc")
	fmt.Printf("name %v : desc %v", name, desc)
	mods, err := model.InsertClass(name, desc)

	if err != nil {
		return c.JSON(http.StatusNotModified, util.NewResult(util.Success, "添加数据失败",mods))
	} else {
		return c.JSON(http.StatusOK, util.NewResult(util.Success, "添加数据成功",mods))
	}
}

/*
	进行数据的修改
*/
func UpdateClass(c echo.Context) error {

	// json方式获取body对象
	json_map := make(map[string]interface{})

	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	//body, err := ioutil.ReadAll(c.Request().Body)
	//
	//if err != nil {
	//	return c.JSON(http.StatusExpectationFailed, util.NewResult(101, "请求参数获取出错"))
	//}
	//
	//fmt.Printf("json获取的body数据:%v",body)
	//if err = json.Unmarshal(body, &json_map); err != nil {
	//	return c.JSON(http.StatusExpectationFailed, util.NewResult(101, err.Error()))
	//}

	fmt.Printf("json获取数据:%v \n", json_map)
	name, _ := json_map["name"].(string)
	desc, _ := json_map["desc"].(string)
	f_id, _ := json_map["id"].(float64)

	i_id := int(f_id)


	//name := c.FormValue("name")
	//desc := c.FormValue("desc")
	//i_id, err:= strconv.Atoi(c.FormValue("id"))

	if err != nil {
		return c.JSON(http.StatusExpectationFailed, util.NewResult(101, "id 不能为空"))
	}
	fmt.Printf("修改操作的数据：name %v : desc %v ：id %v", name, desc, i_id)
	//id , _ := strconv.Atoi(i_id)
	err = model.UpdateClass(name, desc, i_id)

	if err != nil {
		return c.JSON(http.StatusNotModified, util.NewResult(101, err.Error()))
	} else {
		return c.JSON(http.StatusOK, util.NewResult(util.Success, "修改数据成功"))
	}
}

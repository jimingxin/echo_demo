package model

import (
	"fmt"
)

type Class struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Desc string `json:"desc" form:"desc"`
}

func GetClassAll() []Class {
	mods := make([]Class, 0)
	// 进行查询
	Err = db.Select(&mods, `select * from class`)

	return mods
}

/*
	根据条件获取数据
*/
func GetQueryClass(id int) []Class {
	mods := make([]Class, 0)
	// 进行查询
	Err = db.Select(&mods, `select * from class where id >= ?`,id)

	return mods
}


/*
	进行数据的插入
*/
func InsertClass(name,desc string) ([]Class,error) {
	result, Err := db.Exec("INSERT INTO class(`name`,`desc`) VALUES(?, ?)", name, desc)

	id,_ := result.LastInsertId();
	fmt.Printf("reslut-id:%d",id)

	mods := GetQueryClassUsingId(int(id));
	// 进行查询

	return  mods,Err
}

func UpdateClass(name,desc string, id int) error {
	_, Err = db.Exec("update class set `name` = ?,`desc` = ? where id = ?",name,desc,id)
	return Err
}

/*
	根据条件获取数据
*/
func GetQueryClassUsingId(id int) []Class {
	mods := make([]Class, 0)
	// 进行查询
	Err = db.Select(&mods, `select * from class where id = ?`,id)

	return mods
}
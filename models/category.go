package models

type ParamCategories struct {
	Name string `json:"name" db:"name" binding:"required"`
}

type ParamCategoryId struct {
	Id int `json:"id" binding:"required"`
}

type ParamEditCategory struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

//type CategoryName struct {
//	Name string `db:"name"`
//}

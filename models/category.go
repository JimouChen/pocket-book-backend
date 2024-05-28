package models

type ParamCategories struct {
	Name string `json:"name" db:"name" binding:"required"`
	//UserId string `json:"UserId"  binding:"required"`
}

type ParamCategoryId struct {
	Id int `json:"id" binding:"required"`
}

type ParamEditCategory struct {
	Id   int    `json:"id" db:"id" binding:"required"`
	Name string `json:"name" db:"name" binding:"required"`
}

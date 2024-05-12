package models

type ParamCategories struct {
	Name string `json:"name" binding:"required"`
}

type ParamCategoryId struct {
	Id int `json:"id" binding:"required"`
}

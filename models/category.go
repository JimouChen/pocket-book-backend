package models

type ParamCategories struct {
	Name string `json:"name" binding:"required"`
}

type ParamCategoryId struct {
	Id string `json:"id" binding:"required"`
}

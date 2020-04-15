package services

import "ThesisProject/models"

func AddCategory(Category models.Category) {
	Db.Save(&Category)
}

func RemoveCategory(Category models.Category) {
	Db.Delete(&Category)
}

func UpdateCategory(Category models.Category) {
	Db.Save(&Category)
}

func GetAllCategory() []models.Category {
	var Categories []models.Category
	Db.Preload("Product").Find(&Categories)
	return Categories
}

func GetAllCategoryName() []string {
	var categories []models.Category
	Db.Select("category_name").Find(&categories)
	var categoryName []string
	for i := 0; i < len(categories); i++ {
		categoryName = append(categoryName, categories[i].CategoryName)
	}
	return categoryName
}

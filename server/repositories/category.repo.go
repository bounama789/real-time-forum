package repositories

import (
	"fmt"
	db "forum/database"
	q "forum/database/query"
	"forum/models"
)

type CatRepo struct {
	BaseRepo
}

func (r *CatRepo) init() {
	r.DB = db.DB
	r.TableName = db.CATEGORIES_TABLE
}

func (r *CatRepo) SaveCategory(category models.Category) error {
	err := r.DB.Insert(r.TableName, category)
	if err != nil {
		return err
	}
	return nil
}

func (r *CatRepo) DeleteCategory(categoryId string) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"category_id": categoryId})
	if err != nil {
		return err
	}
	return nil
}

func (r *CatRepo) UpdateCategory(Category models.Category) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"category_id": Category.CategoryId})
	if err != nil {
		return err
	}
	return nil
}

func (r *CatRepo) GetCategory(categoryId string) (category models.Category, err error) {
	row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"category_id": categoryId})
	if err != nil {
		return category, err
	}
	err = row.Scan(&category.CategoryId, &category.Name)
	if err != nil {
		return category, fmt.Errorf("no value found")
	}
	return category, nil
}

func (r *CatRepo) GetCategories() (categories []models.Category, err error) {
	var category models.Category
	rows, err := r.DB.GetAllFrom(r.TableName, nil,"")
	if err != nil {
		return categories, err
	}

	for rows.Next() {
		rows.Scan(&category.CategoryId, &category.Name,&category.Color)
		categories = append(categories, category)

	}

	return categories, err
}

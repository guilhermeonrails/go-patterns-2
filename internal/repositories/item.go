package repositories

import (
	"myapi/internal/config"
	"myapi/internal/models"
)

type ItemRepository struct{}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{}
}

func (r *ItemRepository) ListAll() ([]models.Item, error) {
	var items []models.Item
	if err := config.DB.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ItemRepository) GetByID(id int) (*models.Item, error) {
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) GetByCode(code string) (*models.Item, error) {
	var item models.Item
	if err := config.DB.Where("codigo = ?", code).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) Create(item *models.Item) (*models.Item, error) {
	if err := config.DB.Create(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (r *ItemRepository) Update(item *models.Item) error {
	return config.DB.Save(item).Error
}

func (r *ItemRepository) Delete(id int) error {
	return config.DB.Delete(&models.Item{}, id).Error
}

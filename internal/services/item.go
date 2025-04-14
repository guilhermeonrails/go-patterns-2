package services

import (
	"errors"
	"myapi/internal/models"
)

// service para criar um item
func CreateItem(item *models.Item) (*models.Item, error) {
	if item.Nome == "" {
		return nil, errors.New("nome do item não pode ser vazio")
	}
	return item, nil
}

package repositories

import (
	"errors"

	"github.com/satti999/PractiseProject/models"
)

var items = []models.Item{}

type ItemRepository interface {
	GetAll() ([]models.Item, error)
	GetByID(int) (models.Item, error)
	Create(models.Item) (models.Item, error)
	Update(int, models.Item) (models.Item, error)
	Delete(int) error
}

type itemRepository struct{}

func NewItemRepository() ItemRepository {
	return &itemRepository{}
}

func (r *itemRepository) GetAll() ([]models.Item, error) {
	return items, nil
}

func (r *itemRepository) GetByID(id int) (models.Item, error) {
	for _, item := range items {
		if item.ID == id {
			return item, nil
		}
	}
	return models.Item{}, errors.New("item not found")
}

func (r *itemRepository) Create(item models.Item) (models.Item, error) {
	item.ID = len(items) + 1
	items = append(items, item)
	return item, nil
}

func (r *itemRepository) Update(id int, item models.Item) (models.Item, error) {
	for i, existingItem := range items {
		if existingItem.ID == id {
			item.ID = id
			items[i] = item
			return item, nil
		}
	}
	return models.Item{}, errors.New("item not found")
}

func (r *itemRepository) Delete(id int) error {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

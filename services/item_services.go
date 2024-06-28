package services

import (
	"github.com/satti999/PractiseProject/models"
	"github.com/satti999/PractiseProject/repositories"
)

type ItemService interface {
	GetAllItems() ([]models.Item, error)
	GetItemByID(int) (models.Item, error)
	CreateItem(models.Item) (models.Item, error)
	UpdateItem(int, models.Item) (models.Item, error)
	DeleteItem(int) error
}

type itemService struct {
	repo repositories.ItemRepository
}

func NewItemService() ItemService {
	return &itemService{
		repo: repositories.NewItemRepository(),
	}
}

func (s *itemService) GetAllItems() ([]models.Item, error) {
	return s.repo.GetAll()
}

func (s *itemService) GetItemByID(id int) (models.Item, error) {
	return s.repo.GetByID(id)
}

func (s *itemService) CreateItem(item models.Item) (models.Item, error) {
	return s.repo.Create(item)
}

func (s *itemService) UpdateItem(id int, item models.Item) (models.Item, error) {
	return s.repo.Update(id, item)
}

func (s *itemService) DeleteItem(id int) error {
	return s.repo.Delete(id)
}

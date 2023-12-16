package services

import (
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

type computer struct {
	repo ports.StorageComputerService
}

// NewComputer return instance of computer service
func NewComputer(r ports.StorageComputerService) *computer {
	return &computer{
		repo: r,
	}
}
func (c *computer) ListComputers() []*models.CPU {
	result := c.repo.SelectComputers()
	return result
}

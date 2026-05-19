package store

import (
	"testing"

	"github.com/mrckurz/CI-CD-MCM/internal/model"
)

func TestCreateAndGet(t *testing.T) {
	s := NewMemoryStore()

	product := model.Product{
		Name:  "Widget",
		Price: 9.99,
	}

	created := s.Create(product)

	found, err := s.GetByID(created.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if found.Name != "Widget" {
		t.Errorf("expected Widget, got %s", found.Name)
	}
}

func TestGetAllEmpty(t *testing.T) {
	s := NewMemoryStore()

	products := s.GetAll()

	if len(products) != 0 {
		t.Errorf("expected 0 products, got %d", len(products))
	}
}

func TestDeleteNonExistent(t *testing.T) {
	s := NewMemoryStore()

	err := s.Delete(999)

	if err != ErrNotFound {
		t.Error("expected ErrNotFound when deleting non-existent product")
	}
}

func TestUpdateProduct(t *testing.T) {
	s := NewMemoryStore()

	product := model.Product{
		Name:  "Widget",
		Price: 9.99,
	}

	created := s.Create(product)
	created.Name = "Updated Widget"

	updated, err := s.Update(created.ID, created)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if updated.Name != "Updated Widget" {
		t.Errorf("expected updated name")
	}
}

func TestDeleteExisting(t *testing.T) {
	s := NewMemoryStore()

	product := model.Product{
		Name:  "Widget",
		Price: 9.99,
	}

	created := s.Create(product)

	err := s.Delete(created.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetByIDInvalid(t *testing.T) {
	s := NewMemoryStore()

	_, err := s.GetByID(999)

	if err != ErrNotFound {
		t.Error("expected ErrNotFound")
	}
}

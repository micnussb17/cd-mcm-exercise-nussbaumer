package store

import (
	"testing"

	"github.com/mrckurz/CI-CD-MCM/internal/model"
)

func TestCreateAndGet(t *testing.T) {
	s := NewMemoryStore()

	created := s.Create(model.Product{
		Name:  "Laptop",
		Price: 999.99,
	})

	if created.ID == 0 {
		t.Fatal("expected created product to have a non-zero ID")
	}

	got, err := s.GetByID(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got.ID != created.ID {
		t.Errorf("expected ID %d, got %d", created.ID, got.ID)
	}
	if got.Name != created.Name {
		t.Errorf("expected name %q, got %q", created.Name, got.Name)
	}
	if got.Price != created.Price {
		t.Errorf("expected price %v, got %v", created.Price, got.Price)
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

	created := s.Create(model.Product{
		Name:  "Mouse",
		Price: 25.0,
	})

	updated, err := s.Update(created.ID, model.Product{
		Name:  "Gaming Mouse",
		Price: 39.99,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if updated.ID != created.ID {
		t.Errorf("expected ID %d, got %d", created.ID, updated.ID)
	}
	if updated.Name != "Gaming Mouse" {
		t.Errorf("expected name %q, got %q", "Gaming Mouse", updated.Name)
	}
	if updated.Price != 39.99 {
		t.Errorf("expected price %v, got %v", 39.99, updated.Price)
	}

	got, err := s.GetByID(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got.Name != "Gaming Mouse" {
		t.Errorf("expected stored name %q, got %q", "Gaming Mouse", got.Name)
	}
	if got.Price != 39.99 {
		t.Errorf("expected stored price %v, got %v", 39.99, got.Price)
	}
}

func TestDeleteProduct(t *testing.T) {
	s := NewMemoryStore()

	created := s.Create(model.Product{
		Name:  "Keyboard",
		Price: 49.99,
	})

	err := s.Delete(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = s.GetByID(created.ID)
	if err != ErrNotFound {
		t.Fatalf("expected ErrNotFound after delete, got %v", err)
	}
}

func TestGetByIDNotFound(t *testing.T) {
	s := NewMemoryStore()

	tests := []struct {
		name string
		id   int
	}{
		{name: "zero ID", id: 0},
		{name: "positive missing ID", id: 123},
		{name: "negative ID", id: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetByID(tt.id)
			if err != ErrNotFound {
				t.Fatalf("expected ErrNotFound for ID %d, got %v", tt.id, err)
			}
		})
	}
}
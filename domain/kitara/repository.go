package kitara

import (
	"errors"
	"sync"
)

type Repository interface {
	Save(pv *ProductVariant) (*ProductVariant, error)
	GetByID(productVariantId int) (*ProductVariant, error)
	InitData(variants []*ProductVariant) ([]*ProductVariant, error)
	Clear() error
	GetAll() ([]*ProductVariant, error)
}

var ErrProductVariantNotFound = errors.New("Product Variant Not Found")

type MemoryRepository struct {
	productVariants []*ProductVariant
	m               *sync.Mutex
}

func (r *MemoryRepository) InitData(variants []*ProductVariant) ([]*ProductVariant, error) {
	r.m.Lock()
	defer r.m.Unlock()
	for _, pv := range variants {
		r.productVariants = append(r.productVariants, NewProductVariant(pv.ID, pv.Sku, pv.Size, pv.Color, pv.Qty))
	}
	return r.productVariants, nil
}

func (r *MemoryRepository) GetByID(productVariantId int) (*ProductVariant, error) {
	r.m.Lock()
	defer r.m.Unlock()
	for _, pv := range r.productVariants {
		if pv.ID == productVariantId {
			return NewProductVariant(pv.ID, pv.Sku, pv.Size, pv.Color, pv.Qty), nil
		}
	}

	return nil, ErrProductVariantNotFound
}

func (r *MemoryRepository) GetAll() ([]*ProductVariant, error) {
	r.m.Lock()
	defer r.m.Unlock()

	containers := make([]*ProductVariant, len(r.productVariants))
	i := 0
	for _, pv := range r.productVariants {
		containers[i] = NewProductVariant(pv.ID, pv.Sku, pv.Size, pv.Color, pv.Qty)
		i++
	}
	return containers, nil

}

func (r *MemoryRepository) Save(pv *ProductVariant) (*ProductVariant, error) {
	r.m.Lock()
	defer r.m.Unlock()
	i := 0
	index := 0
	flag := false
	for _, variant := range r.productVariants {
		if variant.ID == pv.ID {
			index = i
			flag = true
		}
		i++
	}
	if !flag {
		r.productVariants = append(r.productVariants, NewProductVariant(pv.ID, pv.Sku, pv.Size, pv.Color, pv.Qty))
		return pv, nil
	}

	r.productVariants[index] = NewProductVariant(pv.ID, pv.Sku, pv.Size, pv.Color, pv.Qty)
	return r.productVariants[index], nil
}

func (r *MemoryRepository) Clear() error {
	r.m.Lock()
	defer r.m.Unlock()
	r.productVariants = []*ProductVariant{}
	return nil
}

func NewRepository() (r Repository) {
	r = &MemoryRepository{
		productVariants: []*ProductVariant{},
		m:               &sync.Mutex{},
	}
	return
}

package ball

import "errors"

//BallRepository interface
type Repository interface {
	SaveContainer(c *Container) (*Container, error)
	GetAllContainer() ([]*Container, error)
	GetVerifiedContainer() *Container
	AddContainers(containers []*Container) ([]*Container, error)
}

var ErrContainerNotFound = errors.New("Container Not Found")

type MemoryRepository struct {
	containers []*Container
}

func (r MemoryRepository) SaveContainer(c *Container) (*Container, error) {

	for _, container := range r.containers {
		if container.ID == c.ID {
			container = c
			return container, nil
		}
	}
	return nil, ErrContainerNotFound
}

func (r *MemoryRepository) GetAllContainer() ([]*Container, error) {
	return r.containers, nil
}

func (r *MemoryRepository) AddContainers(containers []*Container) ([]*Container, error) {
	r.containers = containers
	return r.containers, nil
}

func (r *MemoryRepository) GetVerifiedContainer() *Container {
	for _, container := range r.containers {
		if container.IsVerified() {
			return container
		}
	}

	return nil
}

func NewRepository(m []*Container) (r Repository) {
	r = &MemoryRepository{m}
	return
}

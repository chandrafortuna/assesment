package ball

import (
	"errors"
	"sync"
)

//BallRepository interface
type Repository interface {
	SaveContainer(c *Container) (*Container, error)
	GetAllContainer() ([]*Container, error)
	VerifiedContainer() *Container
	AddContainers(containers []*Container) ([]*Container, error)
	ClearContainer() error
	GetContainerByID(containerId int) (*Container, error)
}

var ErrContainerNotFound = errors.New("Container Not Found")

type MemoryRepository struct {
	containers []*Container
	verified   *Container
	m          *sync.Mutex
}

func (r MemoryRepository) SaveContainer(c *Container) (*Container, error) {

	r.m.Lock()
	defer r.m.Unlock()
	i := 0
	index := 0
	flag := false
	for _, container := range r.containers {
		if container.ID == c.ID {
			index = i
			flag = true
		}
		i++
	}
	if !flag {
		return nil, ErrContainerNotFound
	}

	r.containers[index] = NewContainer(c.ID, c.Max, c.Balls)
	return r.containers[index], nil
}

func (r *MemoryRepository) GetAllContainer() ([]*Container, error) {
	r.m.Lock()
	defer r.m.Unlock()

	containers := make([]*Container, len(r.containers))
	i := 0
	for _, c := range r.containers {
		containers[i] = NewContainer(c.ID, c.Max, c.Balls)
		i++
	}
	return containers, nil

}

func (r *MemoryRepository) ClearContainer() error {
	r.containers = []*Container{}
	r.verified = nil
	return nil
}
func (r *MemoryRepository) AddContainers(containers []*Container) ([]*Container, error) {
	r.m.Lock()
	defer r.m.Unlock()
	for _, c := range containers {
		r.containers = append(r.containers, NewContainer(c.ID, c.Max, c.Balls))
	}
	return r.containers, nil
}

func (r *MemoryRepository) VerifiedContainer() *Container {
	r.m.Lock()
	defer r.m.Unlock()
	if r.verified != nil {
		return r.verified
	}
	for _, container := range r.containers {
		if container.IsVerified() {
			r.verified = NewContainer(container.ID, container.Max, container.Balls)
			return r.verified
		}
	}

	return nil
}

func (r *MemoryRepository) GetContainerByID(containerId int) (*Container, error) {
	r.m.Lock()
	defer r.m.Unlock()
	for _, container := range r.containers {
		if container.ID == containerId {
			return NewContainer(container.ID, container.Max, container.Balls), nil
		}
	}

	return nil, ErrContainerNotFound
}

func NewRepository() (r Repository) {
	r = &MemoryRepository{
		containers: []*Container{
			NewContainer(1, 10, []*Ball{}),
			NewContainer(2, 10, []*Ball{}),
			NewContainer(3, 10, []*Ball{}),
		},
		verified: nil,
		m:        &sync.Mutex{},
	}
	return
}

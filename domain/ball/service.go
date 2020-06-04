package ball

import (
	"errors"
	"fmt"
	"math/rand"
)

type Service struct {
	repo Repository
}

var ErrAddBallFailed = errors.New("Add Ball Failed")
var ErrContainerIsEmpty = errors.New("Container Is Empty")

func NewService(r Repository) Service {
	return Service{
		repo: r,
	}
}

func (s *Service) Init(containerNum int, maxLoad int) ([]*Container, error) {

	err := s.repo.ClearContainer()
	if err != nil {
		return []*Container{}, err
	}

	var containers []*Container
	for i := 1; i <= containerNum; i++ {
		containers = append(containers, &Container{
			ID:    i,
			Max:   maxLoad,
			Balls: []*Ball{},
		})
	}
	savedContainers, err := s.repo.AddContainers(containers)
	if err != nil {
		return []*Container{}, err
	}

	return savedContainers, nil
}

func (s *Service) GetAllContainer() ([]*Container, error) {
	containers, err := s.repo.GetAllContainer()
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func (s *Service) GetVerifiedContainer() (*Container, error) {
	container := s.repo.VerifiedContainer()
	return container, nil
}

func (s *Service) AddBallToContainer() (*Container, error) {
	containers, err := s.repo.GetAllContainer()
	if err != nil {
		return nil, err
	}

	if len(containers) <= 0 {
		return nil, ErrContainerIsEmpty
	}

	containerLen := len(containers)
	randIndex := rand.Intn(containerLen)
	containerId := containers[randIndex].ID

	verifiedContainer := s.repo.VerifiedContainer()
	if verifiedContainer != nil {
		return nil, errors.New(fmt.Sprintf("Container ID:%d Is Full", verifiedContainer.ID))
	}

	selectedContainer, err := s.repo.GetContainerByID(containerId)
	if err != nil {
		return nil, err
	}

	selectedContainer.Balls = append(selectedContainer.Balls, NewBall())

	savedContainer, err := s.repo.SaveContainer(selectedContainer)
	if err != nil {
		return nil, err
	}

	return savedContainer, nil
}

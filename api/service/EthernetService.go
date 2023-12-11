package service

import (
	"errors"
	"smartgw/api/domain"
	"smartgw/api/repository"
)

type (
	EthernetService interface {
		Add(ethernet *domain.Ethernet) error
		Update(ethernet *domain.Ethernet) error
		Delete(name string) error
		Find(name string) (domain.Ethernet, error)
		FindAll() ([]domain.Ethernet, error)
	}

	ethernetService struct {
		ethernetRepository repository.EthernetRepository
	}
)

var _ EthernetService = (*ethernetService)(nil)

func NewEthernetService(ethernetRepository repository.EthernetRepository) EthernetService {
	return &ethernetService{
		ethernetRepository: ethernetRepository,
	}
}

func (e ethernetService) Add(ethernet *domain.Ethernet) error {
	if _, err := e.ethernetRepository.Find(ethernet.Name); err == nil {
		return errors.New("此ip已存在")
	}

	return e.ethernetRepository.Save(ethernet)
}

func (e ethernetService) Update(ethernet *domain.Ethernet) error {
	return e.ethernetRepository.Save(ethernet)
}

func (e ethernetService) Delete(name string) error {
	return e.ethernetRepository.Delete(name)
}

func (e ethernetService) Find(name string) (domain.Ethernet, error) {
	return e.ethernetRepository.Find(name)
}

func (e ethernetService) FindAll() ([]domain.Ethernet, error) {
	return e.ethernetRepository.FindAll()
}

package service

import (
	"errors"
	"greatwall/api/domain"
	"greatwall/api/repository"
	"greatwall/lib/task"
)

type (
	CollectTaskService interface {
		Add(collectTask *domain.CollectTask) error
		Update(collectTask *domain.CollectTask) error
		Delete(name string) error
		Find(name string) (domain.CollectTask, error)
		FindAll() ([]domain.CollectTask, error)
	}

	collectTaskService struct {
		collectTaskRepository repository.CollectTaskRepository
		collectTaskServer     *task.CollectTaskServer
	}
)

var _ CollectTaskService = (*collectTaskService)(nil)

func NewCollectTaskService(collectTaskRepository repository.CollectTaskRepository, collectTaskServer *task.CollectTaskServer) CollectTaskService {
	return &collectTaskService{
		collectTaskRepository: collectTaskRepository,
		collectTaskServer:     collectTaskServer,
	}
}

func (cts *collectTaskService) Add(collectTask *domain.CollectTask) error {
	if _, err := cts.collectTaskRepository.Find(collectTask.Name); err == nil {
		return errors.New("采集任务已存在")
	}
	err := cts.collectTaskRepository.Save(collectTask)
	if err != nil {
		return err
	}
	cts.collectTaskServer.Add(collectTask)
	return err
}

func (cts *collectTaskService) Update(collectTask *domain.CollectTask) error {
	err := cts.collectTaskRepository.Save(collectTask)
	if err != nil {
		return err
	}
	cts.collectTaskServer.Update(collectTask)
	return err
}

func (cts *collectTaskService) Delete(name string) error {
	err := cts.collectTaskRepository.Delete(name)
	if err != nil {
		return err
	}
	cts.collectTaskServer.Delete(name)
	return err
}

func (cts *collectTaskService) Find(name string) (domain.CollectTask, error) {
	return cts.collectTaskRepository.Find(name)
}

func (cts *collectTaskService) FindAll() ([]domain.CollectTask, error) {
	return cts.collectTaskRepository.FindAll()
}

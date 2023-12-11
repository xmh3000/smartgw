package service

import (
	"errors"
	"greatwall/api/domain"
	"greatwall/api/repository"
	"greatwall/lib/task"
)

type (
	ReportTaskService interface {
		Add(reportTask *domain.ReportTask) error
		Update(reportTask *domain.ReportTask) error
		Delete(name string) error
		Find(name string) (domain.ReportTask, error)
		FindAll() ([]domain.ReportTask, error)
	}

	reportTaskService struct {
		reportTaskRepository repository.ReportTaskRepository
		reportTaskServer     *task.ReportTaskServer
	}
)

var _ ReportTaskService = (*reportTaskService)(nil)

func NewReportTaskService(reportTaskRepository repository.ReportTaskRepository, reportTaskServer *task.ReportTaskServer) ReportTaskService {
	return &reportTaskService{
		reportTaskRepository: reportTaskRepository,
		reportTaskServer:     reportTaskServer,
	}
}

// Add 新增任务
func (s *reportTaskService) Add(reportTask *domain.ReportTask) error {
	if _, err := s.reportTaskRepository.Find(reportTask.Name); err == nil {
		return errors.New("上报任务已经存在")
	}
	err := s.reportTaskRepository.Save(reportTask)
	if err != nil {
		return err
	}
	s.reportTaskServer.Add(reportTask)
	return err
}

func (s *reportTaskService) Update(reportTask *domain.ReportTask) error {
	err := s.reportTaskRepository.Save(reportTask)
	if err != nil {
		return err
	}
	s.reportTaskServer.Update(reportTask)
	return err
}

func (s *reportTaskService) Delete(name string) error {
	err := s.reportTaskRepository.Delete(name)
	if err != nil {
		return err
	}
	s.reportTaskServer.Delete(name)
	return err
}

func (s *reportTaskService) Find(name string) (domain.ReportTask, error) {
	return s.reportTaskRepository.Find(name)
}

func (s *reportTaskService) FindAll() ([]domain.ReportTask, error) {
	return s.reportTaskRepository.FindAll()
}

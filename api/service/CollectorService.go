package service

import (
	"errors"
	"smartgw/api/domain"
	"smartgw/api/repository"
	"smartgw/lib/collect"
)

type (
	CollectorService interface {
		Add(collector *domain.Collector) error
		Update(collector *domain.Collector) error
		Delete(name string) error
		Find(name string) (domain.Collector, error)
		FindAll() ([]domain.Collector, error)
	}

	collectorService struct {
		collectorRepository   repository.CollectorRepository
		deviceRepository      repository.DeviceRepository
		collectTaskRepository repository.CollectTaskRepository
		collectServer         *collect.CollectorServer
	}
)

var _ CollectorService = (*collectorService)(nil)

func NewCollectorService(
	collectorRepository repository.CollectorRepository,
	deviceRepository repository.DeviceRepository,
	collectTaskRepository repository.CollectTaskRepository,
	collectServer *collect.CollectorServer) CollectorService {
	return &collectorService{
		collectorRepository:   collectorRepository,
		deviceRepository:      deviceRepository,
		collectTaskRepository: collectTaskRepository,
		collectServer:         collectServer,
	}
}

func (cs *collectorService) Add(collector *domain.Collector) error {
	if _, err := cs.collectorRepository.Find(collector.Name); err == nil {
		return errors.New("采集接口已存在")
	}

	err := cs.collectorRepository.Save(collector)
	if err != nil {
		return err
	}
	cs.collectServer.Add(*collector)
	return err
}

func (cs *collectorService) Update(collector *domain.Collector) error {
	err := cs.collectorRepository.Save(collector)
	if err != nil {
		return err
	}
	cs.collectServer.Update(*collector)

	// 同步更新设备信息
	cs.deviceRepository.CollectorChanged(*collector)
	//cs.collectTaskRepository.CollectorChanged(collector)
	return err
}

func (cs *collectorService) Delete(name string) error {
	// 检测设备是否引用了采集接口
	collector, err := cs.Find(name)
	if err == nil {
		if cs.deviceRepository.CollectorIsUsed(&collector) {
			return errors.New("已经有设备引用了采集接口，不能删除！")
		}

		//if cs.collectTaskRepository.CollectorIsUsed(&collector) {
		//	return errors.New("已经有采集任务引用了采集接口，不能删除！")
		//}
	}

	err = cs.collectorRepository.Delete(name)
	if err != nil {
		return err
	}
	cs.collectServer.Delete(name)
	return err
}

func (cs *collectorService) Find(name string) (domain.Collector, error) {
	return cs.collectorRepository.Find(name)
}

func (cs *collectorService) FindAll() ([]domain.Collector, error) {
	return cs.collectorRepository.FindAll()
}

package repository

import (
	"influ-dojo/api/domain/apperr"
	domainModel "influ-dojo/api/domain/model"
	"influ-dojo/api/domain/repository"
	dataModel "influ-dojo/api/infrastructure/persistence/model"

	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

type dailyWork struct {
	GormRepository
}

func NewDailyWork(db *gorm.DB) repository.Work {
	return &dailyWork{GormRepository{db}}
}

func (repo *dailyWork) Load() ([]*domainModel.Work, error) {
	mdls := make([]*dataModel.DailyWork, 0)
	if err := repo.DB.Find(&mdls).Error; err != nil {
		return nil, xerrors.Errorf("failed to load daily works: %w", err)
	}

	entities := make([]*domainModel.Work, 0)
	for _, mdl := range mdls {
		entities = append(entities, mdl.MakeEntity())
	}

	return entities, nil
}

func (repo *dailyWork) LoadOrderByRanking() ([]*domainModel.Work, error) {
	mdls := make([]*dataModel.DailyWork, 0)
	if err := repo.DB.Order("point desc").Find(&mdls).Error; err != nil {
		return nil, xerrors.Errorf("failed to load work ranking: %w", err)
	}

	entities := make([]*domainModel.Work, 0)
	for _, mdl := range mdls {
		entities = append(entities, mdl.MakeEntity())
	}

	return entities, nil
}

func (repo *dailyWork) LoadByID(id string) (*domainModel.Work, error) {
	mdl := new(dataModel.DailyWork)
	if err := repo.DB.Where("user_id = ?", id).First(mdl).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, apperr.ErrRecordNotFound
		}

		return nil, err
	}

	return mdl.MakeEntity(), nil
}

func (repo *dailyWork) Save(entity *domainModel.Work) error {
	mdl := dataModel.NewDailyWork(entity)

	return repo.store(repo.DB, mdl)
}

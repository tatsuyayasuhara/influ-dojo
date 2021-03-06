package repository

import (
	"influ-dojo/api/domain/apperr"
	domainModel "influ-dojo/api/domain/model"
	"influ-dojo/api/domain/repository"
	dataModel "influ-dojo/api/infrastructure/persistence/model"

	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

type monthlyWork struct {
	GormRepository
}

func NewMonthlyWork(db *gorm.DB) repository.Work {
	return &monthlyWork{GormRepository{db}}
}

func (repo *monthlyWork) Load() ([]*domainModel.Work, error) {
	mdls := make([]*dataModel.MonthlyWork, 0)
	if err := repo.DB.Find(&mdls).Error; err != nil {
		return nil, xerrors.Errorf("failed to load daily works: %w", err)
	}

	entities := make([]*domainModel.Work, 0)
	for _, mdl := range mdls {
		entities = append(entities, mdl.MakeEntity())
	}

	return entities, nil
}

func (repo *monthlyWork) LoadOrderByRanking() ([]*domainModel.Work, error) {
	mdls := make([]*dataModel.MonthlyWork, 0)
	if err := repo.DB.Order("point desc").Find(&mdls).Error; err != nil {
		return nil, xerrors.Errorf("failed to load work ranking: %w", err)
	}

	entities := make([]*domainModel.Work, 0)
	for _, mdl := range mdls {
		entities = append(entities, mdl.MakeEntity())
	}

	return entities, nil
}

func (repo *monthlyWork) LoadByID(id string) (*domainModel.Work, error) {
	mdl := new(dataModel.MonthlyWork)
	if err := repo.DB.Where("user_id = ?", id).First(mdl).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, apperr.ErrRecordNotFound
		}

		return nil, err
	}

	return mdl.MakeEntity(), nil
}

func (repo *monthlyWork) Save(entity *domainModel.Work) error {
	mdl := dataModel.NewMonthlyWork(entity)

	return repo.store(repo.DB, mdl)
}

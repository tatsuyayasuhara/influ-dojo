package repository

import (
	"errors"
	domainModel "influ-dojo/api/domain/model"
	"influ-dojo/api/domain/repository"
	dataModel "influ-dojo/api/infrastructure/persistence/model"

	"github.com/jinzhu/gorm"
)

type monthlyResult gormRepository

func NewMonthlyResult(db *gorm.DB) repository.MonthlyResult {
	return &monthlyResult{db}
}

func (result *monthlyResult) LoadTop3() ([]*domainModel.Result, error) {
	mdls := make([]*dataModel.MonthlyResult, 0)
	if err := result.DB.Order("point desc").Limit(3).Find(&mdls).Error; err != nil {
		return nil, err
	}

	entities := make([]*domainModel.Result, len(mdls))
	for _, mdl := range mdls {
		count := 0
		if mdl.IncreaseFollowersCount == nil {
			mdl.IncreaseFollowersCount = &count
		}
		if mdl.Point == nil {
			mdl.Point = &count
		}

		entities = append(entities, &domainModel.Result{
			ScreenName:             mdl.ScreenName,
			FollowersCount:         mdl.FollowersCount,
			IncreaseFollowersCount: *mdl.IncreaseFollowersCount,
			Point:                  *mdl.Point,
		})
	}

	return entities, nil
}

func (result *monthlyResult) LoadByScreenName(screenName string) (*domainModel.Result, error) {
	mdl := new(dataModel.MonthlyResult)
	if err := result.DB.Where("screen_name = ?", screenName).First(mdl).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("not found")
		}

		return nil, err
	}

	count := 0
	if mdl.IncreaseFollowersCount == nil {
		mdl.IncreaseFollowersCount = &count
	}
	if mdl.Point == nil {
		mdl.Point = &count
	}

	return &domainModel.Result{
		ScreenName:             mdl.ScreenName,
		FollowersCount:         mdl.FollowersCount,
		IncreaseFollowersCount: *mdl.IncreaseFollowersCount,
		Point:                  *mdl.Point,
	}, nil
}

func (result *monthlyResult) Save(entity *domainModel.Result) error {
	mdl := &dataModel.MonthlyResult{
		ScreenName:             entity.ScreenName,
		FollowersCount:         entity.FollowersCount,
		IncreaseFollowersCount: &entity.IncreaseFollowersCount,
		Point:                  &entity.Point,
	}

	return result.DB.Where("screen_name = ?", mdl.ScreenName).Assign(*mdl).FirstOrCreate(mdl).Error
}
package model

import "influ-dojo/api/domain/model"

type WeeklyWork struct {
	UserID                 string `gorm:"primary_key"`
	TweetsCount            int
	IncreaseTweetsCount    *int
	MyTweetsCount          *int
	RepliesCount           *int
	FavoritesCount         int
	IncreaseFavoritesCount *int
	Point                  *float64
	Ranking                int
	LastRanking            int
	Model
}

func (mdl *WeeklyWork) IsNew() bool {
	return len(mdl.UserID) == 0
}

func (mdl *WeeklyWork) AttachID() error {
	return nil
}

func (mdl *WeeklyWork) MakeEntity() *model.Work {
	zero := 0
	if mdl.IncreaseTweetsCount == nil {
		mdl.IncreaseTweetsCount = &zero
	}
	if mdl.MyTweetsCount == nil {
		mdl.MyTweetsCount = &zero
	}
	if mdl.RepliesCount == nil {
		mdl.RepliesCount = &zero
	}
	if mdl.IncreaseFavoritesCount == nil {
		mdl.IncreaseFavoritesCount = &zero
	}
	zeroZero := 0.0
	if mdl.Point == nil {
		mdl.Point = &zeroZero
	}

	return &model.Work{
		UserID:                 mdl.UserID,
		TweetsCount:            mdl.TweetsCount,
		IncreaseTweetsCount:    *mdl.IncreaseTweetsCount,
		MyTweetsCount:          *mdl.MyTweetsCount,
		RepliesCount:           *mdl.RepliesCount,
		FavoritesCount:         mdl.FavoritesCount,
		IncreaseFavoritesCount: *mdl.IncreaseFavoritesCount,
		Point:                  *mdl.Point,
		Ranking:                mdl.Ranking,
		LastRanking:            mdl.LastRanking,
	}
}

func NewWeeklyWork(entity *model.Work) *WeeklyWork {
	return &WeeklyWork{
		UserID:                 entity.UserID,
		TweetsCount:            entity.TweetsCount,
		IncreaseTweetsCount:    &entity.IncreaseTweetsCount,
		MyTweetsCount:          &entity.MyTweetsCount,
		RepliesCount:           &entity.RepliesCount,
		FavoritesCount:         entity.FavoritesCount,
		IncreaseFavoritesCount: &entity.IncreaseFavoritesCount,
		Point:                  &entity.Point,
		Ranking:                entity.Ranking,
		LastRanking:            entity.LastRanking,
	}
}

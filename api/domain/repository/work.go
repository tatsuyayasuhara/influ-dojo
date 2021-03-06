//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=../../mock/$GOPACKAGE/mock_$GOFILE

package repository

import "influ-dojo/api/domain/model"

type Work interface {
	Load() ([]*model.Work, error)
	LoadOrderByRanking() ([]*model.Work, error)
	LoadByID(id string) (*model.Work, error)
	Save(work *model.Work) error
}

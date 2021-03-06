//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=../../mock/$GOPACKAGE/mock_$GOFILE

package repository

import "influ-dojo/api/usecase/dto"

type Ranking interface {
	LoadAll() (*dto.RankingAll, error)
	Store(all *dto.RankingAll) error
}

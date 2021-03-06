//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=../../mock/$GOPACKAGE/mock_$GOFILE

package client

import "influ-dojo/api/domain/model"

type Follower interface {
	CountFollowers() (int, error)
	GetFollowers() ([]*model.Follower, error)
}

package api

import (
	domainClient "influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	infraClient "influ-dojo/api/infrastructure/client"
	persistence "influ-dojo/api/infrastructure/persistence/repository"

	"github.com/jinzhu/gorm"
)

type Dependency struct {
	FollowerClient    domainClient.Follower
	BotClient         domainClient.Bot
	UserRepo          repository.User
	DailyWorkRepo     repository.Work
	DailyResultRepo   repository.Result
	WeeklyWorkRepo    repository.Work
	WeeklyResultRepo  repository.Result
	MonthlyWorkRepo   repository.Work
	MonthlyResultRepo repository.Result
}

func Inject(cfg *Config, db *gorm.DB) (*Dependency, error) {
	at := cfg.Twitter.AccessToken
	ats := cfg.Twitter.AccessTokenSecret
	ck := cfg.Twitter.ConsumerKey
	cs := cfg.Twitter.ConsumerSecret

	return &Dependency{
		FollowerClient:    infraClient.NewFollower(at, ats, ck, cs),
		BotClient:         infraClient.NewBot(at, ats, ck, cs),
		UserRepo:          persistence.NewUser(db),
		DailyWorkRepo:     persistence.NewDailyWork(db),
		DailyResultRepo:   persistence.NewDailyResult(db),
		WeeklyWorkRepo:    persistence.NewWeeklyWork(db),
		WeeklyResultRepo:  persistence.NewWeeklyResult(db),
		MonthlyWorkRepo:   persistence.NewMonthlyWork(db),
		MonthlyResultRepo: persistence.NewMonthlyResult(db),
	}, nil
}

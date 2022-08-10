package spider

//go:generate mockgen -destination spider_gomock.go -package spider github.com/kingofzihua/learn-go/testing/gomock/spider Spider

type Spider interface {
	GetBody() string
}

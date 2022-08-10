package gomock

import (
	"github.com/kingofzihua/learn-go/testing/gomock/spider"
)

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}

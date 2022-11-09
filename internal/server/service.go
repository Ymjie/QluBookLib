package server

import (
	"cklib/internal/config"
	"cklib/internal/cron"
	"cklib/internal/cron/job"
	"cklib/internal/user"
	"cklib/pkg/logger"
	"cklib/pkg/notice/http"
	"fmt"
)

func New(config *config.Config) *Service {
	return &Service{config: config}
}
func (s *Service) Start() {
	spec := s.config.Cron
	loglv := s.config.LogLevel
	t := s.config.Threads
	for _, v := range s.config.User {
		Mlog := logger.New(nil, int64(loglv), 0)
		Mlog.SetProfix(fmt.Sprintf("[%s]", v.Username))
		newUser := user.NewUser(v.Username, v.Password, v.BookList)
		jobone := job.NewJob(newUser, http.NewNt(v.InfoAPI, "{%msg%}"), t, Mlog)
		cron.Newcron(spec, jobone.Start, Mlog)
	}
	select {}
}

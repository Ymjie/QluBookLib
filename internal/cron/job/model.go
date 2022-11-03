package job

import (
	"cklib/internal/user"
	"cklib/pkg/logger"
	"cklib/pkg/notice"
)

type Job struct {
	Nt   notice.Notifier
	Mlog *logger.MyLogger
	u    *user.User
	t    int
}

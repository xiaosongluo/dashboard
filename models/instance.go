package models

import (
	"github.com/xiaosongluo/dashboard/db"
	"github.com/xiaosongluo/dashboard/conf"
)

var(
	Database db.DB
	Cfg      *conf.Config
)

package models

import (
	"github.com/xiaosongluo/dashboard/conf"
	"github.com/xiaosongluo/dashboard/db"
)

var (
	Database db.DB        //Database
	Cfg      *conf.Config //Cfg
)

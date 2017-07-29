package models

import (
	"github.com/xiaosongluo/dashboard/db"
	"github.com/xiaosongluo/dashboard/conf"
)

var (
	Database db.DB        //Database
	Cfg      *conf.Config //Cfg
)

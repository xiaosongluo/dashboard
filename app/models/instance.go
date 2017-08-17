package models

import (
	"github.com/xiaosongluo/dashboard/app/storage"
	"github.com/xiaosongluo/dashboard/app/utils/config"
)

var (
	Storage storage.Storage //Database
	Config  *config.Config  //Cfg
)

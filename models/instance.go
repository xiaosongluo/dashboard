package models

import (
	"github.com/xiaosongluo/dashboard/config"
	"github.com/xiaosongluo/dashboard/storage"
)

var (
	Storage storage.Storage //Database
	Config  *config.Config  //Cfg
)

package main

import (
	"weblogindemo/dao"
	"weblogindemo/routes"
	"weblogindemo/service"
	"weblogindemo/utils"
)

func main() {

	if dao, err := dao.NewDao(); err != nil {
		panic(err)
	} else {
		service.Dao = dao
	}

	r := routes.InitRouter()
	r.Run(utils.HttpPort)
}

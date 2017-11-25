package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/iorellana18/Team-Udes-Back/utils"
)

var (
	rdbms, user, pass, ipPg, portPg, database, sslmode string
)

func PostgresSetup() {
	rdbms = utils.Config.Postgres.Rdbms
	user = utils.Config.Postgres.User
	pass = utils.Config.Postgres.Pass
	ipPg = utils.Config.Postgres.Ip
	portPg = utils.Config.Postgres.Port
	database = utils.Config.Postgres.Name
	sslmode = utils.Config.Postgres.Sslmode
}

func Database() *gorm.DB {
	db, e := gorm.Open(rdbms, rdbms+"://"+user+":"+pass+"@"+ipPg+":"+portPg+"/"+database+"?sslmode="+sslmode)
	utils.Check(e)
	return db
}

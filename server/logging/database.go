package logging

import (
	"context"
	"database/sql"
	"fmt"

	"Server/util"

	// used at sql.Open(->"mysql"<-, fmt.Sprintf
	_ "github.com/go-sql-driver/mysql"
)

/*
	Create and open SQL Connection
*/
func SQLInit() {
	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", util.GetConfig().Database.DBUser, util.GetConfig().Database.DBPassword, util.GetConfig().Database.DBHost, util.GetConfig().Database.DBDatabase))
	if err != nil {
		util.Err(util.SQL, err, true, "Error creating connection")
		panic(err)
	}
	ctx := context.Background()
	err = DB.PingContext(ctx)
	if err != nil {
		util.Err(util.SQL, err, true, "Error creating connection")
		panic(err)
	}
	util.Log(util.SQL, "Connection established")
}

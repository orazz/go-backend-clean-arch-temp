package bootstrap

import "database/sql"

type Application struct {
	Env   *Env
	SqlDb *sql.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.SqlDb = newMysqlDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	app.SqlDb.Close()
}

package app

func (app *Application) RunSubscribers() {
	CreateLicense(app)
	GetLicense(app)
}

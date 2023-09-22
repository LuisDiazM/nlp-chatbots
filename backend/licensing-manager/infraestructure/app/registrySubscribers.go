package app

import "context"

func (app *Application) RunSubscribers(ctx context.Context) {
	CreateLicense(app, ctx)
	GetLicense(app, ctx)
	IncrementLicenseUsage(app, ctx)
}

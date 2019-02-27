package testhorizon

func initTestCoreInfo(app *App) {
	app.UpdateTestCoreInfo()
}

func init() {
	appInit.Add("testCoreInfo", initTestCoreInfo, "app-context", "log")
}

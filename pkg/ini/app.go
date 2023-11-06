package ini

func (ctx *Context) InitApp() {
	ctx.dummyDbInit()
	// initialize services
	//ctx.Wait4Ever(60000) // dummy service logs every 10 minute
}

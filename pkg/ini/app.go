package ini

import "github.com/sourcehaven/mypass-godbridge/pkg/services"

func init() {
	// app environment initialization
	configInit()
	loggerInit()
	dummyDbInit()

	// initialize services
	services.Wait4Ever(60000) // dummy service logs every 1 minute
}

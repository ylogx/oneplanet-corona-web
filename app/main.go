package main

import (
	"cloud.google.com/go/profiler"
	"gitlab.com/oneplanet/corona-web/app/controllers"
	"gitlab.com/oneplanet/corona-web/app/utils/log"
	"os"
)

func setupGceProfiler() {
	// Profiler initialization, best done as early as possible.
	if err := profiler.Start(profiler.Config{
		Service:        "oneplanet-corona-web-app",
		ServiceVersion: "1.0.0",
		// ProjectID must be set if not running on GCP.
		// ProjectID: "oneplanet-corona-web",
	}); err != nil {
		log.Error("Error in creating the profiler")
	}
}

func main() {
	log.Setup()
	setupGceProfiler()
	//models.Init()
	r := controllers.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_ = r.Run(":" + port)
}

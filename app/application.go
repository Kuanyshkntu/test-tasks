package app

import (
	"go.uber.org/zap"
	"log"
	"test_tasks/controllers"
	"test_tasks/db"
	"test_tasks/env"
	"test_tasks/repository"
	"test_tasks/service"
)

func StartApplication() {
	coreDbEnv, err := env.GetDbEnvVals(env.DatabaseEnvKeys{
		DbUser:        "db_user",
		DbPassword:    "db_pass",
		DbHost:        "db_host",
		DbPool:        "db_pool",
		DbPoolIdle:    "db_pool_idle",
		DbConnTimeout: "db_timeout",
		DbName:        "db_name",
		DbPort:        "db_port",
	})
	if err != nil {
		log.Fatal(err)
	}

	tgwCoreDbPg, err := db.ConnectDB(*coreDbEnv)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(tgwCoreDbPg.DbHandler)

	svc := service.NewService(repo)

	handler := controllers.NewHandler(svc)

	router, err := handler.GetRouter()
	if err != nil {
		log.Fatal("Error initializing router", zap.Error(err))
	}

	err = router.Run(":8089")
	if err != nil {
		log.Fatal("Error running http", zap.Error(err))
	}
}

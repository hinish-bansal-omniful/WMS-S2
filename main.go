package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
	appinit "wms/init"
	"wms/router"

	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/db/sql/migration"
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/shutdown"
)

const (
	modeWorker     = "worker"
	modeHttp       = "http"
	modeMigration  = "migration"
	upMigration    = "up"
	downMigration  = "down"
	forceMigration = "force"
)

func main() {
	fmt.Println("hello")
	err := config.Init(time.Second * 10)
	if err != nil {
		log.Panicf("Error while initialising config, err: %v", err)
		panic(err)
	}
	ctx, err := config.TODOContext()
	if err != nil {
		log.Panicf("Error while getting context from config, err: %v", err)
		panic(err)
	}
	//initoialise connection
	appinit.Initialize(ctx)

	//run server
	runHttpServer(ctx)

}
func runHttpServer(ctx context.Context) {

	server := http.InitializeServer(config.GetString(ctx, "server.port"), 10*time.Second, 10*time.Second, 70*time.Second)

	// Initialize middlewares and routes
	//err := router.Initialize(ctx, server)
	//if err != nil {
	//	log.Errorf(err.Error())
	//	panic(err)
	//}
	//
	err := router.InternalRoutes(ctx, server)
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}

	log.Infof("Starting server on port" + config.GetString(ctx, "server.port"))

	err = server.StartServer("WMS-S2")
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}

	<-shutdown.GetWaitChannel()

	var mode, migrationType, number string
	flag.StringVar(
		&mode,
		"mode",
		modeHttp,
		"Pass the flag to run in different modes (worker or default)",
	)

	flag.StringVar(
		&migrationType,
		"migrationType",
		upMigration,
		"Pass the flag to run migration in different modes (worker or default)",
	)

	flag.StringVar(
		&number,
		"migrationNumber",
		"0",
		"Pass the flag to force migration to that version(number)",
	)
	flag.Parse()

	//if config.GetBool(ctx, "migration.flag") {
	//	runMigration(ctx, migrationType, number)
	//}

	switch strings.ToLower(mode) {
	case modeHttp:
		runHttpServer(ctx)
	// case modeWorker:
	// 	runWorker(ctx)
	case modeMigration:
		runMigration(ctx, migrationType, number)
	default:
		runHttpServer(ctx)
	}
}

func runMigration(ctx context.Context, migrationType string, number string) {
	database := config.GetString(ctx, "postgresql.database")
	mysqlWriteHost := config.GetString(ctx, "postgresql.master.host")
	mysqlWritePort := config.GetString(ctx, "postgresql.master.port")
	mysqlWritePassword := config.GetString(ctx, "postgresql.master.password")
	mysqlWriterUsername := config.GetString(ctx, "postgresql.master.username")

	m, err := migration.InitializeMigrate("file://deployment/migration", "postgres://"+mysqlWriteHost+":"+mysqlWritePort+"/"+database+"?user="+mysqlWriterUsername+"&password="+mysqlWritePassword+"&sslmode=disable")
	if err != nil {
		panic(err)
	}

	switch migrationType {
	case upMigration:
		err = m.Up()
		if err != nil {
			panic(err)
		}
		break
	case downMigration:
		err = m.Down()
		if err != nil {
			panic(err)
		}
		break
	case forceMigration:
		version, parseErr := strconv.Atoi(number)
		if parseErr != nil {
			panic(parseErr)
		}

		err = m.ForceVersion(version)
		if err != nil {
			return
		}
		break
	default:
		err = m.Up()
		if err != nil {
			panic(err)
		}
		break
	}
}

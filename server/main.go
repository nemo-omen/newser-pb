package main

import (
	"log"
	"os"
	"server/handler"
	"server/service"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	// uncomment once you have at least one .go migration file in the "migrations" directory
	// _ "yourpackage/migrations"
)

var (
	searchHandler    *handler.SearchHandler
	searchService    *service.SearchService
	subscribeService *service.SubscribeService
)

func main() {
	app := pocketbase.New()
	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	// Run auto migrations if isRunCmd is true
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isGoRun,
	})

	registerHandlers(app)
	routes(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func routes(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST("/search", searchHandler.Search)
		return nil
	})
}

func registerHandlers(app *pocketbase.PocketBase) {
	searchService = service.NewSearchService()
	subscribeService = service.NewSubscribeService(app.Dao())
	searchHandler = handler.NewSearchHandler(searchService)
}

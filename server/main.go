package main

import (
	"log"
	"os"
	"server/handler"
	"server/service"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	// uncomment once you have at least one .go migration file in the "migrations" directory
	// _ "yourpackage/migrations"
)

var (
	searchHandler    *handler.SearchHandler
	subscribeHandler *handler.SubscribeHandler
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

	app.OnModelAfterCreate("users").Add(func(e *core.ModelEvent) error {
		log.Printf("User created: %v", e.Model)
		createUserCollections(app, e.Model)
		return nil
	})

	registerHandlers(app)
	routes(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func createUserCollections(app *pocketbase.PocketBase, user models.Model) error {
	collection, err := app.Dao().FindCollectionByNameOrId("collections")
	if err != nil {
		return err
	}

	return app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
		readRecord := models.NewRecord(collection)
		savedRecord := models.NewRecord(collection)

		readRecord.Set("user", user.GetId())
		savedRecord.Set("user", user.GetId())
		readRecord.Set("title", "read")
		savedRecord.Set("title", "saved")

		if err := app.Dao().SaveRecord(readRecord); err != nil {
			return err
		}

		if err := app.Dao().SaveRecord(savedRecord); err != nil {
			return err
		}

		return nil
	})
}

func routes(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST("/search", searchHandler.Search)
		e.Router.POST("/subscribe", subscribeHandler.Subscribe)
		return nil
	})
}

func registerHandlers(app *pocketbase.PocketBase) {
	searchService = service.NewSearchService()
	subscribeService = service.NewSubscribeService(app.Dao())
	searchHandler = handler.NewSearchHandler(searchService)
	subscribeHandler = handler.NewSubscribeHandler(subscribeService)
}

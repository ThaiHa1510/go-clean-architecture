package main

import(
	"embed"
	"log"
	"time"
	"fmt"
	"net/http"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/go-clean-archituecture/routes"
	"github.com/go-clean-archituecture/middlewares"
	"github.com/go-clean-archituecture/config"
)
// Embed a single file
//go:embed static/index.html
var f embed.FS

// Embed a directory
//go:embed static/*
var embedDirStatic embed.FS
func main(){
	fmt.Println("Start server....")
	config := config.GetConfig()
	app := fiber.New()
	app.Use("/assets", filesystem.New(filesystem.Config{
		Root: http.FS(embedDirStatic),
		PathPrefix:"assets",
		Browse: true,
	}))
	defer func (){
		if errorMsg := recover(); errorMsg != nil{
			_, shutDownContext := context.WithTimeout(context.Background(), 5*time.Second)
			defer shutDownContext()
			app.ShutdownWithContext(shutDownContext)
		}
	}()
	log.Info("Stared server....")
	// static 
	app.Static("/static","./static")
	routes.LoadApiRoutes(&app)
	middlewares.SetupMiddleware(&app)
	app.Listen(fmt.Printf(":%d", config.App.Port))
}

func LoadDatabase(cfg *config.Config) *database.database{
	db, err := databse.NewMysqlDatabase(cfg)
	if err != nil{
		panic(err.Error())
		return nil
	}
	return db
}

func emitServerStartEvent(){
	go func(){
		eventStart,_ := event.New(ProcessStart , ProcessStartEvent{startBy:"admin"})
		agg := &ServerStartedAggregator{}
		agg.AggregateCluster = eventsourcing.New(agg, agg.Transition, utils.UUIDGenerator)
		agg.Apply(eventStart)
	}()
}

type ServerStartedAggregator struct {
	*event.AggregateCluster
	ProcessId int
	ProcessStatus ProcessStatus
	
}

type ProcessStatus string

const (
	ProcessStart ProcessStatus = "start"
	ProcessStoped ProcessStatus = "stoped"
	ProcessString ProcessStatus = "staring"
	ProcessStoping ProcessStatus = "stoping"
 )
 
type ProcessStartEvent struct {
	startBy string
}

type ProcessStopEvent struct {
	stopReason string
}
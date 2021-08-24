package assets

import "fmt"

func GetMultiplexer(projectName string) (string, string) {
	return fmt.Sprintf(`package multiplexer

import (
	"%[1]s/source/config"

	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// params is the input parameter struct for the handler module.
type params struct {
	fx.In

	Config     config.Config
	Logger     *zap.Logger
	Lifecycle  fx.Lifecycle
	ShutDowner fx.Shutdowner
}

//Module is one main modules
var Module = fx.Options(
	fx.Provide(NewHandlerServer),
	fx.Invoke(registerServer),
)

//NewHandlerServer create new mux and server
func NewHandlerServer(p params) (router *gin.Engine, server *http.Server, err error) {

	gin.SetMode(gin.ReleaseMode)

	rootPath, err := os.Getwd()

	ginLoggerFile, err := os.Create(rootPath + "/delivery-gin.log")

	if err != nil {
		p.Logger.Fatal("error when creating log file for gin " + err.Error())
		return nil, nil, err
	}

	gin.DefaultWriter = io.MultiWriter(ginLoggerFile)

	router = gin.Default()
	router.Use(gin.Recovery())
	router.Use(cros)

	port := p.Config.GetString("SERVER_PORT")

	server = &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return router, server, nil

}

type registerServerParams struct {
	fx.In

	Logger     *zap.Logger
	Lifecycle  fx.Lifecycle
	Shutdowner fx.Shutdowner
	Server     *http.Server
}

func cros(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	// c.Header("Content-Type", "application/json; charset=utf-8")

	// Second, we handle the OPTIONS problem
	if c.Request.Method == "OPTIONS" {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)

	} else {

		c.Next()

	}
}

func registerServer(p registerServerParams) {

	p.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {

				go p.Server.ListenAndServe()

				c := make(chan os.Signal, 1)
				signal.Notify(c, os.Interrupt)

				// Block until a signal is received.
				go func() {
					s := <-c

					p.Logger.Info(
						"Got signal.",
						zap.String("signal", s.String()),
					)

					if err := p.Shutdowner.Shutdown(); err != nil {
						p.Logger.Error("Could not shutdown.", zap.Error(err))
						os.Exit(1)
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				p.Logger.Info("Shutting down server.")
				return p.Server.Shutdown(ctx)
			},
		},
	)
}
`, projectName), "./" + projectName + "/source/multiplexer/multiplexer.go"
}

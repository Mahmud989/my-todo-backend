package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"my-todo/controllers"
	"my-todo/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func setupRoutes() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: utils.DebugLogger.Writer(),
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%v | %3d | %13v | %15s | %-7s %s\n%s",
				param.TimeStamp.Format("2006/01/02 - 15:04:05"),
				param.StatusCode,
				param.Latency,
				param.ClientIP,
				param.Method,
				param.Path,
				param.ErrorMessage,
			)
		},
	}))
	v1 := router.Group("/api/v1")
	{
		v1.POST("/create", controllers.CreateTodo)
		v1.GET("/all", controllers.GetTodos)
	}
	return router
}

// Run will run the HTTP Server
func Run(router http.Handler) {
	// Set up a channel to listen to for interrupt signals
	var runChan = make(chan os.Signal, 1)

	// Set up a context to allow for graceful server shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*4, //ms
	)
	defer cancel()

	// Define server options
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	// Alert the user that the server is starting
	log.Printf("Server is starting on %s\n", server.Addr)

	// Run the server on a new goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				// Normal interrupt operation, ignore
			} else {
				utils.ErrorLogger.Fatalf("Server failed to start due to err: %v", err)
			}
		}
	}()

	// Block on this channel listeninf for those previously defined syscalls assign
	// to variable so we can let the user know why the server is shutting down
	interrupt := <-runChan

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user

	utils.ErrorLogger.Printf("Server is shutting down due to %+v\n", interrupt)

	if err := server.Shutdown(ctx); err != nil {
		utils.ErrorLogger.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	} else {
		utils.InfoLogger.Println("Server stopped")
	}
}
func main() {
	//define log
	initLogs()
	Run(setupRoutes())

}
func initLogs() {
	utils.InfoLogger = log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	utils.WarningLogger = log.New(log.Writer(), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	utils.ErrorLogger = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	utils.DebugLogger = log.New(log.Writer(), "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	initLoggerToPath("/logs/logger-info.log", utils.InfoLogger)
	initLoggerToPath("/logs/logger-warning.log", utils.WarningLogger)
	initLoggerToPath("/logs/logger-error.log", utils.ErrorLogger)
	initLoggerToPath("/logs/logger-debug.log", utils.DebugLogger)
}
func initLoggerToPath(path string, logger *log.Logger) *log.Logger {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		log.Println(err.Error())
	}
	return logger
}

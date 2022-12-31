package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"live-easy-backend/sdk/log"
	"live-easy-backend/src/usecase"
)

var once = sync.Once{}

type rest struct {
	http *gin.Engine
	log  *log.Logger
	uc   *usecase.Usecase
}

func Init(uc *usecase.Usecase, log *log.Logger) *rest {
	r := &rest{}
	once.Do(func() {
		gin.SetMode(gin.ReleaseMode) // TODO: Move to config later

		r.http = gin.New()
		r.log = log
		r.uc = uc

		r.RegisterMiddlewareAndRoutes()
	})

	return r
}

func (r *rest) RegisterMiddlewareAndRoutes() {
	// Global middleware
	r.http.Use(r.CorsMiddleware())
	r.http.Use(gin.Recovery())
	r.http.Use(r.SetTimeout)
	r.http.Use(r.AddFieldsToContext)

	// Auth routes
	r.http.POST("api/v1/auth/register", r.Register)
	r.http.POST("api/v1/auth/login", r.Login)
	r.http.POST("api/v1/auth/login/google", r.LoginWithGoogle)

	// Protected Routes
	v1 := r.http.Group("api/v1", r.Authorization())

	// User routes
	v1.Group("user")
	{
		v1.GET("user/profile", r.GetUserProfile)
	}

	// Medicine routes
	v1.Group("medicine")
	{
		v1.POST("medicine", r.CreateMedicine)
		v1.GET("medicine/:id", r.GetMedicine)
		v1.GET("medicine", r.GetListMedicines)
		v1.PUT("medicine/:id", r.UpdateMedicine)
		v1.DELETE("medicine/:id", r.DeleteMedicine)
	}
}

func (r *rest) Run() {
	/*
		Create context that listens for the interrupt signal from the OS.
		This will allow us to gracefully shutdown the server.
	*/
	c := context.Background()
	ctx, stop := signal.NotifyContext(c, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	port := ":8080"
	if os.Getenv("APP_PORT") != "" {
		port = ":" + os.Getenv("APP_PORT")
	}
	server := &http.Server{
		Addr:              port,
		Handler:           r.http,
		ReadHeaderTimeout: 2 * time.Second,
	}

	// Run the server in a goroutine so that it doesn't block the graceful shutdown handling below

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			r.log.Error(ctx, err.Error())
		}
	}()

	r.log.Info(context.Background(), "Server is running on port "+os.Getenv("APP_PORT"))

	// Block until we receive our signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	r.log.Info(context.Background(), "Shutting down server...")

	// Create a deadline to wait for.
	quitCtx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()
	if err := server.Shutdown(quitCtx); err != nil {
		r.log.Fatal(quitCtx, fmt.Sprintf("Server Shutdown error: %s", err.Error()))
	}

	r.log.Info(context.Background(), "Server gracefully stopped")
}

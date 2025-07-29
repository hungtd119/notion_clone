package cmd

import (
	"context"
	"fmt"
	"net/http"
	"notion/src/routes"
	"os"
	"os/signal"
	"sync"

	"notion/src/config"
	"notion/src/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func server() *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.GetInstance()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// Setup router
			r := gin.Default()

			r.Use(middleware.Recover)
			r.Use(middleware.Cors)

			if cfg.GetEnv() == "production" {
				gin.SetMode(gin.ReleaseMode)
			}

			// Setup routes
			routes.Bootstrap(r)

			var wg sync.WaitGroup
			wg.Add(1)
			go runGin(ctx, r, cfg, &wg)
			wg.Wait()

			zap.S().Info("Shutdown server")
		},
	}
}

func runGin(ctx context.Context, r *gin.Engine, cfg config.IConfig, wg *sync.WaitGroup) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.GetPort()),
		Handler: r,
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, os.Kill)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			zap.S().Error(err)
		}
	}()

	<-sigint
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Error(err)
	}

	wg.Done()
	zap.S().Info("Server shutdown")
}

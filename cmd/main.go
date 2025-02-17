package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"s-coder-snippet-sharder/api"
	"s-coder-snippet-sharder/internal/app/account"
	"s-coder-snippet-sharder/internal/app/snippet"
	"s-coder-snippet-sharder/internal/db/repo"
	db "s-coder-snippet-sharder/internal/db/sqlc"
	"s-coder-snippet-sharder/pkg/session"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var loggerLevels = map[string]slog.Level{
	"info":  slog.LevelInfo,
	"debug": slog.LevelDebug,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

var opts struct {
	http struct {
		port   int
		secure bool
	}
	db struct {
		migrationsDir string
	}
	logger struct {
		style string
		level string
	}
}

func main() {
	flag.IntVar(&opts.http.port, "http-port", 8000, "The HTTP port to connect to")
	flag.BoolVar(&opts.http.secure, "https", false, "Use HTTPS instead of HTTP")
	flag.StringVar(&opts.logger.level, "logger-level", "debug", "Set level of logging")
	flag.StringVar(&opts.logger.style, "logger-style", "text", "Set style of logging (json, console, text)")
	flag.Parse()

	logLevel, ok := loggerLevels[opts.logger.level]
	if !ok {
		log.Fatalf("Invalid logger level: %s", opts.logger.level)
	}

	var handler slog.Handler
	switch opts.logger.style {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})

	case "text":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})

	//case "dev":
	//	handler = NewDevHandler(os.Stdout, logLevel)

	default:
		log.Fatalf("invalid log style supplied: %s", opts.logger.style)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//todo: config for features toggle

	var longTasks sync.WaitGroup
	envPort, err := strconv.ParseInt(os.Getenv("PORT"), 0, 64)
	if err != nil {
		log.Fatal("Invalid port found in env: ", os.Getenv("PORT"))
	}

	//todo: connect database
	//postgresDB, err := postgres.New(ctx, cfg)
	//if err != nil {
	//	log.Fatalf("Error creating Postgres DB: %v", err)
	//}

	conn := db.Conn()
	defer conn.Close()

	snippetRepo := repo.NewSnippetRepo(conn)
	accountRepo := repo.NewAccountRepo(conn)

	snippetService, err := snippet.NewService(ctx, snippetRepo)
	if err != nil {
		log.Fatalf("Error creating snippet service: %v", err)
	}

	accountService, err := account.NewService(ctx, accountRepo)
	if err != nil {
		log.Fatalf("Error creating account service: %v", err)
	}

	//run session
	sessionutil.InitSessionStore()

	// add register auth service
	accountService.RegisterAuthService()

	repos := api.Repo{
		Snippet: snippetRepo,
		Account: accountRepo,
	}

	services := &api.Services{
		SnippetAPI: snippetService,
		AccountAPI: accountService,
	}

	apiInstance := api.New(
		logger,
		&longTasks,
		services,
		repos,
		api.WithPort(int(envPort)),
		api.WithHTTPS(opts.http.secure),
	)

	go apiInstance.StartServer(ctx)
	longTasks.Wait()

	//Wait for a sinal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//
	//slog.Info("shutting down")
	//ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//if err := api.Shutdown(ctxShutdown); err != nil {
	//	slog.Error("Shutdown error: ", err)
	//}

	//if err := conn.Close(); err != nil {
	//	slog.Error("Error closing DB connection: ", err)
	//}

	//ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	//defer stop()
	//
	//<-ctx.Done() // Đợi tín hiệu từ Ctrl+C

	slog.Info("Shutting down server...")
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	slog.Info("Calling Shutdown()...")

	if err := apiInstance.Shutdown(ctxShutdown); err != nil {
		slog.Error("Shutdown error:", err)
	} else {

		slog.Info("Server shutdown completed.")
	}

}

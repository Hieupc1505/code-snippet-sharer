package api

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html/v2"
	"github.com/nedpals/supabase-go"
	"html/template"
	"log"
	"log/slog"
	"os"
	"s-coder-snippet-sharder/api/templateutil"
	"s-coder-snippet-sharder/internal/app"
	"s-coder-snippet-sharder/internal/app/account"
	"s-coder-snippet-sharder/internal/app/snippet"
	"s-coder-snippet-sharder/pkg/config"
	"s-coder-snippet-sharder/pkg/ipcache"
	"s-coder-snippet-sharder/pkg/sup"
	"s-coder-snippet-sharder/pkg/token"
	"sync"
	"time"
)

type Repo struct {
	Snippet snippet.ReadWriter
	Account account.ReadWriter
}

type Services struct {
	SnippetAPI *snippet.Service
	AccountAPI *account.Service
}

type API struct {
	secure bool
	port   int
	domain string

	repo     Repo
	services *Services

	app        *fiber.App
	logger     *slog.Logger
	tasks      *sync.WaitGroup
	template   *template.Template
	sup        *supabase.Client
	tokenMaker token.Maker
	View       *templateutil.View
	IPCache    *ipcache.Cache
}

func New(
	slog *slog.Logger,
	longTasks *sync.WaitGroup,
	services *Services,
	repo Repo,
	opts ...OptFunc,
) *API {
	engine := html.New("./api/web/views", ".tmpl")
	engine.Reload(true)

	templateutil.RegisterEngineFuncs(engine)

	fiberInstance := fiber.New(
		fiber.Config{
			Views:             engine,
			ViewsLayout:       "layouts/main",
			PassLocalsToViews: true,
			ReadTimeout:       5 * time.Second,
		},
	)

	//Tell's the Go server to "serve" the assets (CSS, Images, JS files)
	//todo: handle when static method has been removed
	fiberInstance.Static("/assets", "./api/web/assets")

	tmpl, err := templateutil.ParseFiles()
	if err != nil {
		log.Fatal("Failed to parse template ", "error", err)
	}

	token, err := token.NewPasetoMaker(config.Envs.SymmetricKey) //todo: add key to env file
	if err != nil {
		log.Fatal("Failed to create token maker ", "error", err)
	}

	ipCache := ipcache.New()

	a := &API{
		secure: false,
		port:   8080,
		domain: os.Getenv("DOMAIN"),

		services: services,
		repo:     repo,

		app:        fiberInstance,
		logger:     slog,
		tasks:      longTasks,
		template:   tmpl,
		sup:        sup.CreateClient(),
		tokenMaker: token,
		IPCache:    ipCache,

		View: &templateutil.View{
			NamedEndpoints: templateutil.GetNamedEndpoints(),
			AppName:        app.Name,
			AppTitle:       app.Title,
		},
	}

	for _, opt := range opts {
		opt(a)
	}

	return a

}

func (a *API) Shutdown(ctx context.Context) error {
	return a.app.Shutdown()
}

func (a *API) StartServer(ctx context.Context) error {
	//Define routers
	//a.registerMiddleware()

	a.registerRoutes()

	a.View.NamedEndpoints = templateutil.GetNamedEndpoints()

	//Start server
	a.logger.InfoContext(ctx, "Server listening on ", "port", a.port)
	err := a.app.Listen(fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}

	return nil

}

func (a *API) registerMiddleware() {
	a.app.Use(cors.New(cors.Config{
		AllowOrigins:     fmt.Sprintf("http://localhost:%d", a.port),
		AllowCredentials: true,
	}))
	//a.app.User(middleware.AttachLogger())
	//a.app.Use(middleware.WithAuthenticatedUser(a.repo.Account, a.sup))

	a.app.Use(recover2.New())
	a.app.Use(requestid.New())
	a.app.Use(logger.New())
}

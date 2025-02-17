package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"s-coder-snippet-sharder/api/apierror"
	"s-coder-snippet-sharder/api/handler"
	"s-coder-snippet-sharder/api/handler/account"
	"s-coder-snippet-sharder/api/handler/snippet"
	"s-coder-snippet-sharder/api/templateutil"
)

func (a *API) registerRoutes() {
	h := &handler.Handler{
		App:            a.app,
		View:           a.View,
		NamedEndpoints: templateutil.GetNamedEndpoints(),
		Template:       a.template,
		SB:             a.sup,
		Token:          a.tokenMaker,
		IPCache:        a.IPCache,
		Tasks:          a.tasks,

		SnippetRepo: a.repo.Snippet,
		SnippetAPI:  a.services.SnippetAPI,

		AccountRepo: a.repo.Account,
		AccountAPI:  a.services.AccountAPI,
	}

	//a.broker.Listen(event.EngineerProfileCreatedHandler(h))

	account.Routes(h)
	snippet.Routes(h)

	a.app.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/favicon.ico" {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	})

	//a.app.Get("/chat", websoket.New(ws.HanlderChatConnections(h)))
	a.app.Get("/*", func(c *fiber.Ctx) error {
		return apierror.ErrorResponse(c, fmt.Errorf("Uh oh, this page doesn't exist'"))
	})
}

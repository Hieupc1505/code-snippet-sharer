package templateutil

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
	"maps"
)

const fallbackURL = "/fallback" //Todo: Decide this?

func FlasBackWithSuccess(c *fiber.Ctx, vars fiber.Map) error {
	return flash.WithSuccess(c, vars).RedirectBack(fallbackURL)
}

func FlashBackWithError(c *fiber.Ctx, err error) error {
	return flash.WithError(c, fiber.Map{
		"ErrorMessage": err.Error(),
	}).RedirectBack(fallbackURL)
}

// Render populates template render variables, flash data and default View data into template.
func Render(v *View, c *fiber.Ctx, template string, vars fiber.Map, layouts ...string) error {
	//Lazily init so that we can merge in flash data and View data
	if vars == nil {
		vars = make(fiber.Map)
	}
	//Add in any flash data errors etc
	flashData := flash.Get(c)
	if flashData != nil {
		maps.Copy(vars, flashData)
	}

	//Add default view fields
	vars["View"] = v
	return c.Render(template, vars, layouts...)
}

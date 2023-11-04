package route

import (
	"IoT/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/suhu", handler.SuhuGetAll)
	r.Post("/suhu", handler.AddSuhu)
}

package route

import (
	"IoT/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/suhu", handler.SuhuGetAll)
	r.Get("/suhu/:id", handler.SuhuGetById)
	r.Post("/suhu", handler.AddSuhu)
	r.Put("/suhu/:id", handler.SuhuUpdate)
	r.Delete("/suhu/:id", handler.SuhuDelete)
}

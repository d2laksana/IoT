package handler

import (
	"IoT/database"
	"IoT/model/entity"
	"IoT/model/request"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SuhuGetAll(ctx *fiber.Ctx) error {
	var suhu []entity.Suhu
	result := database.DB.Debug().Find(&suhu)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(suhu)
}

func AddSuhu(ctx *fiber.Ctx) error {
	suhu := new(request.SuhuCreateRequest)
	if err := ctx.BodyParser(suhu); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	newSuhu := entity.Suhu{
		Temp: suhu.Temp,
		Hum:  suhu.Hum,
	}
	result := database.DB.Debug().Create(&newSuhu)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(suhu)
}

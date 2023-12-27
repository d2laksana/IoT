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
	// validate request body not empty
	if suhu.Temp == 0 || suhu.Hum == 0 {
		return ctx.Status(400).SendString("Temp and Hum is required")
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

// get by id
func SuhuGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	// validate if data exist
	var suhu entity.Suhu
	result := database.DB.Debug().Where("id = ?", id).First(&suhu)
	if result.Error != nil {
		return ctx.Status(400).SendString(result.Error.Error())

	}
	return ctx.JSON(suhu)
}

// update
func SuhuUpdate(ctx *fiber.Ctx) error {
	// get id
	id := ctx.Params("id")
	// validate if data exist
	var suhu entity.Suhu
	result := database.DB.Debug().Where("id = ?", id).First(&suhu)
	if result.Error != nil {
		return ctx.Status(400).SendString(result.Error.Error())
	}
	// parse request body
	suhuUpdate := new(request.SuhuUpdateRequest)
	if err := ctx.BodyParser(suhuUpdate); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	// update data change only
	if suhuUpdate.Temp != 0 {
		suhu.Temp = suhuUpdate.Temp
	}
	if suhuUpdate.Hum != 0 {
		suhu.Hum = suhuUpdate.Hum
	}
	// save to database
	result = database.DB.Debug().Save(&suhu)
	if result.Error != nil {
		return ctx.Status(400).SendString(result.Error.Error())
	}
	return ctx.JSON(suhu)

}

// delete
func SuhuDelete(ctx *fiber.Ctx) error {
	// get id
	id := ctx.Params("id")
	// validate if data exist
	var suhu entity.Suhu
	result := database.DB.Debug().Where("id = ?", id).First(&suhu)
	if result.Error != nil {
		return ctx.Status(400).SendString(result.Error.Error())
	}
	// delete data
	result = database.DB.Debug().Delete(&suhu)
	if result.Error != nil {
		return ctx.Status(400).SendString(result.Error.Error())
	}
	return ctx.JSON(suhu)
}

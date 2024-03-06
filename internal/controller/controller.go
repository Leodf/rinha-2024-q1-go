package controller

import (
	"github.com/Leodf/leodf-go/internal/db"
	"github.com/Leodf/leodf-go/internal/dto"
	"github.com/Leodf/leodf-go/internal/model"
	"github.com/Leodf/leodf-go/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func PostTransaction(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	err = utils.IdValidator(id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	var body = &dto.TransactionRequest{}
	err = c.BodyParser(body)
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	err = utils.TransactionReqValidator(body)
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	resp, err := model.SaveTransaction(id, body)
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	return c.JSON(resp)
}

func GetStatment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	err = utils.IdValidator(id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	resp, err := model.GetClientBalance(id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(resp)
}

func ResetDB(c *fiber.Ctx) error {
	_, err := db.PG.Exec(`UPDATE clientes SET saldo = 0`)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	_, err = db.PG.Exec(`TRUNCATE TABLE transacoes`)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}

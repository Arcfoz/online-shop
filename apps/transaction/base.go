package transaction

import (
	infrafiber "online-shop/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB){
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	trxRoute := router.Group("transaction")
	{
		// using middleware
		trxRoute.Use(infrafiber.CheckAuth())


		trxRoute.Post("/checkout", handler.CreateTransaction)
		trxRoute.Get("/user/histories", handler.GetTransactionByUser)
	}
}
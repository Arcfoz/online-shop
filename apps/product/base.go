package product

import (
	"online-shop/apps/auth"
	infrafiber "online-shop/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	productRoute := router.Group("products")
	{
		productRoute.Get("", handler.GetListProduct)
		productRoute.Get("/sku/:sku", handler.GetDetailProduct)

		//auth role only
		productRoute.Post("", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.CreateProduct)
	}
}

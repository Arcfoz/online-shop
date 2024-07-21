package product

import (
	"net/http"
	infrafiber "online-shop/infra/fiber"
	"online-shop/infra/response"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) CreateProduct(ctx *fiber.Ctx) error {
	req := CreateProductRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	if err := h.svc.CreateProduct(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}
	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithMessage("create product successed"),
	).Send(ctx)

}

func (h handler) GetListProduct(ctx *fiber.Ctx) error {
	req := ListProductsRequestPayload{}

	if err := ctx.QueryParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	products, err := h.svc.ListProducts(ctx.UserContext(), req)

	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	ProductsListResponse := NewProductsListResponseFromEntity(products)

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get list products successed"),
		infrafiber.WithPayload(ProductsListResponse),
		infrafiber.WithQuery(req.GenerateDefaultValue()),
	).Send(ctx)

}

func (h handler) GetDetailProduct(ctx *fiber.Ctx) error {

	sku := ctx.Params("sku", "")

	if sku == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	product, err := h.svc.ProductDetail(ctx.UserContext(), sku)

	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	ProductDetailResponse := ProductDetailResponse{
		Id:        product.Id,
		SKU:       product.SKU,
		Name:      product.Name,
		Stock:     product.Stock,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdataeAt: product.UpdataeAt,
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get product detail successed"),
		infrafiber.WithPayload(ProductDetailResponse),
	).Send(ctx)

}

package v1

import (
	"go-apiserver-template/internal/api"
	"go-apiserver-template/internal/ctx"
	"go-apiserver-template/internal/services"
)

// @Summary Get order by id
// @Description Get order by id
// @Tags Order
// @Accept json
// @Produce json
// @Param oid path string true "Order ID"
// @Success 200 {object} api.Response{data=models.Order}
// @Failure 200 {object} api.Response
// @Router /api/v1/orders/{oid} [get]
func GetOrder(c *ctx.Context) {
	oid := c.Param("oid")

	o, err := services.NewOrderService().GetOrder(oid)
	if err != nil {
		api.Error(c, err)
		return
	}

	api.Success(c, o)
}

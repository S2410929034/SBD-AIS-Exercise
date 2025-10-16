package rest

import (
	"encoding/json"
	"net/http"
	"ordersystem/model"
	"ordersystem/repository"
	"time"

	"github.com/go-chi/render"
)

// GetMenu 			godoc
// @tags 			Menu
// @Description 	Returns the menu of all drinks
// @Produce  		json
// @Success 		200 {array} model.Drink
// @Router 			/api/menu [get]
func GetMenu(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo
		// get slice from db
		// render.Status(r, http.StatusOK)
		// render.JSON(w, r, <your-slice>)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, db.GetDrinks())
	}
}

// todo create GetOrders /api/order/all
// GetOrders 		godoc
// @tags 			Order
// @Description 	Returns all orders
// @Produce  		json
// @Success 		200 {array} model.Order
// @Router 			/api/order/all [get]
func GetOrders(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo
		// get slice from db
		// render.Status(r, http.StatusOK)
		// render.JSON(w, r, <your-slice>)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, db.GetOrders())
	}
}

// todo create GetOrdersTotal /api/order/total
// GetOrdersTotalgodoc
// @tags 			Order
// @Description 	Returns all orders
// @Produce  		json
// @Success 		200 {object} map[uint64]uint64
// @Router 			/api/order/totalled [get]
func GetOrdersTotal(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo
		// get slice from db
		// render.Status(r, http.StatusOK)
		// render.JSON(w, r, <your-slice>)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, db.GetTotalledOrders())
	}
}

// PostOrder 		godoc
// @tags 			Order
// @Description 	Adds an order to the db
// @Accept 			json
// @Param 			b body model.OrderDto true "Order"
// @Produce  		json
// @Success 		200
// @Failure     	400
// @Router 			/api/order [post]
func PostOrder(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo
		// declare empty order struct
		// err := json.NewDecoder(r.Body).Decode(&<your-order-struct>)
		// handle error and render Status 400
		// add to db

		dto := model.OrderDto{}
		err := json.NewDecoder(r.Body).Decode(&dto)

		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		order := model.Order{
			DrinkID:   dto.DrinkID,
			Amount:    dto.Amount,
			CreatedAt: time.Now(),
		}
		db.AddOrder(&order)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, "ok")
	}
}

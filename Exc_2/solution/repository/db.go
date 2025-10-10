package repository

import (
	"ordersystem/model"
	"time"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	// Init the drinks slice with some test data
	// drinks := ...

	drinks := []model.Drink{
		{ID: 1, Name: "Coca-Cola", Price: 1.5, Description: "A refreshing soft drink."},
		{ID: 2, Name: "Espresso", Price: 2.0, Description: "Strong and bold coffee."},
		{ID: 3, Name: "Lemonade", Price: 1.2, Description: "Sweet and tangy lemonade."},
	}

	orders := []model.Order{
		{DrinkID: 1, CreatedAt: time.Now().Add(-48 * time.Hour), Amount: 2},
		{DrinkID: 2, CreatedAt: time.Now().Add(-24 * time.Hour), Amount: 1},
		{DrinkID: 1, CreatedAt: time.Now().Add(-12 * time.Hour), Amount: 3},
	}

	// Init orders slice with some test data

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	// calculate total orders
	// key = DrinkID, value = Amount of orders
	// totalledOrders map[uint64]uint64

	orders := db.GetOrders()
	totalledOrders := make(map[uint64]uint64)
	for _, order := range orders {
		totalledOrders[order.DrinkID] += order.Amount
	}

	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	// todo
	// add order to db.orders slice
	db.orders = append(db.orders, *order)
}

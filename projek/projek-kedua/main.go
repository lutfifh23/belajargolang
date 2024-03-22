package main

import (
	"projek-kedua/config"
	crud_order "projek-kedua/crud/order"
	crud_product "projek-kedua/crud/products"
	"projek-kedua/entity"
	"time"
)

func main() {
	config.ConnectDB()

	product := entity.Products{
		KodeProduk: "001",
		NamaProduk: "somay",
		Stok:       "100",
		Harga:      "2000",
	}

	crud_product.InsertProduct(config.DB, &product)

	product := entity.Products{
		Kodeproduk: "002",
		Namaproduk: "bakso",
		Stok:       "100",
		Harga:      `2000`,
	}

	crud_product.InsertProduct(config.DB, &product)

	order := entity.Orders{
		Id_order:      0,
		Tanggalorder:  time.Now(),
		Paymentmethod: "",
		Total:         30000,
	}

	crud_order.InsertOrder(config.DB, &order)

	details := []entity.OrderDetails{
		{
			Orders: order,
			Products: product.entity.Products{
				KodeProduct: "001",
			},
			Qty: 10,
		},
		{
			Orders: order,
			Products: product.entity.Products{
				KodeProduct: "002",
			},
			Qty: 10,
		},
	}
}

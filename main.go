package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"

	"github.com/spf13/viper"
)

type Config struct {
	Port    string `mapstructure:"PORT"`
	DB_CONN string `mapstructure:"DB_CONN"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:    viper.GetString("PORT"),
		DB_CONN: viper.GetString("DB_CONN"),
	}

	// Setup Database
	db, err := database.InitDB(config.DB_CONN)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	// Setup DI
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Setup Routes
	http.HandleFunc("/api/products", productHandler.GetAll)
	http.HandleFunc("/api/product", productHandler.CreateProduct)
	http.HandleFunc("/api/product/", productHandler.HandleProductByID)

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	http.HandleFunc("/list-api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]interface{}{
			"status":  "OK",
			"message": "Kasir API is running",
			"endpoints": []map[string]string{
				{
					"name":        "Get All Products",
					"method":      "GET",
					"url":         "/api/product",
					"description": "Retrieve all products",
				},
				{
					"name":        "Get Product By ID",
					"method":      "GET",
					"url":         "/api/product/{id}",
					"description": "Retrieve product by UUID",
				},
				{
					"name":        "Create Product",
					"method":      "POST",
					"url":         "/api/product",
					"description": "Create a new product",
				},
				{
					"name":        "Update Product",
					"method":      "PUT",
					"url":         "/api/product/{id}",
					"description": "Update product by UUID",
				},
				{
					"name":        "Delete Product",
					"method":      "DELETE",
					"url":         "/api/product/{id}",
					"description": "Delete product by UUID",
				},
			},
		}

		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server running on localhost:" + config.Port)

	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println("Failed to run server", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rigmas/joybox/internal/core/service"
	"github.com/rigmas/joybox/internal/handler/api"
	v1 "github.com/rigmas/joybox/internal/handler/v1"
	"github.com/rigmas/joybox/internal/repository"
	"github.com/rigmas/joybox/pkg/generator"
	"github.com/rigmas/joybox/pkg/handler"
)

func getEnv(key string) string {
	err := godotenv.Load(".env")
	checkError(err)

	return os.Getenv(key)
}
func main() {
	var client *http.Client
	bookApiUrl := getEnv("BOOK_API_URL")
	bookApi, err := api.NewBookApi(client, bookApiUrl)
	checkError(err)
	bookSrv, err := service.NewBookService(bookApi)
	checkError(err)

	ordersGenerator := generator.Order()
	orderRepo, err := repository.NewOrderRepo(ordersGenerator)
	checkError(err)
	orderSrv, err := service.NewOrderService(bookSrv, orderRepo)

	portHost := fmt.Sprintf(":%v", getEnv("PORT"))
	apiHandler := v1.NewApiHandler(portHost, bookSrv, orderSrv)

	registry := handler.NewRegistry()
	registry.Register("API", apiHandler)
	registry.StartAll()
	registry.StopAll()
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

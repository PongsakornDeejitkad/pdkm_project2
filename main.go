package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"order-management/entity"
	adminDelivery "order-management/features/admin/delivery"
	adminRepository "order-management/features/admin/repository"
	adminUsecase "order-management/features/admin/usecase"
	customerDelivery "order-management/features/customer/delivery"
	customerRepository "order-management/features/customer/repository"
	customerUsecase "order-management/features/customer/usecase"
	productDelivery "order-management/features/product/delivery"
	productRepository "order-management/features/product/repository"
	productUsecase "order-management/features/product/usecase"
	"order-management/middleware"
	"order-management/utils"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var runEnv string
var DB *gorm.DB

func serveGracefulShutdown(e *echo.Echo) {
	go func() {
		var port string
		port = os.Getenv("HTTP_PORT")
		if port == "" {
			port = utils.ViperGetString("http.port")
		}

		if err := e.Start(port); err != nil {
			log.Println("shutting down the server", err)

		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	gracefulShutdownTimeout := 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}
}

func migrateDB() {
	DB.AutoMigrate(
		// &OrderItems{},
		// &PaymentDetails{},
		// &CartItems{},
		// &CourierDetails{},
		// &CustomerPayment{},
		// &CustomerAddress{},
		// &Products{},
		// &ProductsCategory{},
		&entity.Admin{},
		&entity.AdminType{},
		// &Sessions{},
		&entity.Customer{},
	// &Orders{},
	)
}

func init() {
	runEnv = os.Getenv("RUN_ENV")
	if runEnv == "" {
		runEnv = "local"
	}

	utils.InitViper(runEnv)

	var err error
	if err = connectDB(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	e := echo.New()

	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	// Unauthenticated route
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
	})

	// Restricted group
	v1 := e.Group("/v1")

	adminGroup := v1.Group("")
	adminGroup.Use(middleware.AdminAuth())

	customerGroup := v1.Group("")
	customerGroup.Use(middleware.CustomerAuth())

	adminDelivery.NewAdminHandler(adminGroup, adminUsecase.NewAdminUsecase(adminRepository.NewAdminRepository(DB)))
	// adminDelivery.NewCustomerHandler(customerGroup, adminUsecase.NewAdminUsecase(adminRepository.NewAdminRepository(DB)))

	productDelivery.NewHandler(customerGroup, productUsecase.NewUsecase(productRepository.NewRepository(DB)))
	customerDelivery.NewHandler(customerGroup, customerUsecase.NewUsecase(customerRepository.NewRepository(DB)))

	serveGracefulShutdown(e)
}

// Feature Customer

func connectDB() error {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		utils.ViperGetString("postgres.host"),
		utils.ViperGetString("postgres.user"),
		utils.ViperGetString("postgres.password"),
		utils.ViperGetString("postgres.dbname"),
		utils.ViperGetString("postgres.port"))

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	// migrateDB()

	return err
}

// Feature Admin
// Feature Admin 2

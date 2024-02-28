package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"order-management/utils"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var runEnv string
var DB *gorm.DB

func init() {
	runEnv = os.Getenv("RUN_ENV")
	if runEnv == "" {
		runEnv = "dev"
	}

	utils.InitViper()

	var err error
	if err = connectDB(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"status": true})
	})

	// v1 := e.Group("/v1")

	serveGracefulShutdown(e)
}

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
	// &Admins{},
	// &AdminType{},
	// &Sessions{},
	// &Customers{},
	// &Orders{},
	)
}

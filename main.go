package main

import (
	"context"
	"database/sql"
	"fmt"
	_healthUsecase "frieda-golang-training-beginner/health/usecase"
	_helloWorldUsecase "frieda-golang-training-beginner/hello-world/usecase"
	"frieda-golang-training-beginner/util"
	"github.com/labstack/gommon/log"

	_healthHttpDirectory "frieda-golang-training-beginner/health/directory/http"
	_helloWorldHttpDirectory "frieda-golang-training-beginner/hello-world/directory/http"
	_paymentCodeHttpDirectory "frieda-golang-training-beginner/payment_code/directory/http"
	"frieda-golang-training-beginner/payment_code/repository"
	_paymentCodeUsecase "frieda-golang-training-beginner/payment_code/usecase"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/robfig/cron"
	"time"
)

func main() {
	e := echo.New()

	helloWorldUsecase := _helloWorldUsecase.HelloWorldUsecase{}
	_helloWorldHttpDirectory.NewHelloWorldHandler(e, helloWorldUsecase)

	healthUsecase := _healthUsecase.HealthUsecase{}
	_healthHttpDirectory.NewHealthHandler(e, healthUsecase)

	db := initDB()
	//paymentCodeRepository := repository.NewPaymentCodeRepository(db)
	paymentCodeRepository := repository.PaymentCodeRepository{Conn: db}
	paymentCodeUsecase := _paymentCodeUsecase.PaymentCodeUsecase{PaymentCodeRepo: paymentCodeRepository, ContextTimeout: time.Duration(100000000)}
	_paymentCodeHttpDirectory.NewPaymentCodeHandler(e, paymentCodeUsecase)

	cr := cron.New()
	_ = cr.AddFunc("*/1 * * * *", func() {
		err := paymentCodeUsecase.Expire(context.Background())
		if err != nil {
			log.Error(err)
		}
	},
	)
	cr.Start()

	log.Fatal(e.Start("localhost:9090"))

}

func initDB() *sql.DB {
	postgresHost := util.MustHaveEnv("POSTGRES_HOST")
	postgresPort := util.MustHaveEnv("POSTGRES_PORT")
	postgresUser := util.MustHaveEnv("POSTGRES_USER")
	postgresPassword := util.MustHaveEnv("POSTGRES_PASSWORD")
	postgresDbname := util.MustHaveEnv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresPassword, postgresDbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

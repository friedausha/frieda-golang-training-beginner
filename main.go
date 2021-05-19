package main

import (
	"database/sql"
	"fmt"
	_healthUsecase "frieda-golang-training-beginner/health/usecase"
	_helloWorldUsecase "frieda-golang-training-beginner/hello-world/usecase"
	"github.com/labstack/gommon/log"

	_healthHttpDirectory "frieda-golang-training-beginner/health/directory/http"
	_helloWorldHttpDirectory "frieda-golang-training-beginner/hello-world/directory/http"
	_paymentCodeHttpDirectory "frieda-golang-training-beginner/payment_code/directory/http"
	"frieda-golang-training-beginner/payment_code/repository"
	_paymentCodeUsecase "frieda-golang-training-beginner/payment_code/usecase"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "frieda"
	password = "namamu"
	dbname   = "golang_training"
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
	paymentCodeUsecase := _paymentCodeUsecase.PaymentCodeUsecase{PaymentCodeRepo: paymentCodeRepository, ContextTimeout: time.Duration(100000000) }
	_paymentCodeHttpDirectory.NewPaymentCodeHandler(e, paymentCodeUsecase)

	log.Fatal(e.Start("localhost:9090"))

}

func initDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

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

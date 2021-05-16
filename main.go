package main

import (
	"database/sql"
	"fmt"
	_healthHttpDirectory "frieda-golang-training-beginner/health/directory/http"
	_healthUseCase "frieda-golang-training-beginner/health/usecase"
	_helloWorldHttpDirectory "frieda-golang-training-beginner/hello-world/directory/http"
	_helloWorldUseCase "frieda-golang-training-beginner/hello-world/usecase"
	"frieda-golang-training-beginner/payment_code/directory/http"
	"frieda-golang-training-beginner/payment_code/repository"
	"frieda-golang-training-beginner/payment_code/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "golang_training"
)

func main() {
	e := echo.New()

	helloWorldUsecase := _helloWorldUseCase.NewHelloWorldUsecase()
	_helloWorldHttpDirectory.NewHelloWorldHandler(e, helloWorldUsecase)

	healthUsecase := _healthUseCase.NewHealthUsecase()
	_healthHttpDirectory.NewHealthHandler(e, healthUsecase)
	log.Fatal(e.Start("localhost:9090"))

	db := initDB()
	paymentCodeRepository := repository.NewPaymentCodeRepository(db)
	paymentCodeUseCase := usecase.NewPaymentCodeUsecase(paymentCodeRepository, time.Duration(10000000))
	http.NewPaymentCodeHandler(e, paymentCodeUseCase)

}

func initDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

package main

import (
	"database/sql"
	"fmt"
	"frieda-golang-training-beginner/aws"
	_healthUsecase "frieda-golang-training-beginner/health/usecase"
	_helloWorldUsecase "frieda-golang-training-beginner/hello-world/usecase"
	"frieda-golang-training-beginner/inquiry/directory/http"
	repository2 "frieda-golang-training-beginner/inquiry/repository"
	"frieda-golang-training-beginner/inquiry/usecase"
	http2 "frieda-golang-training-beginner/payment/directory/http"
	repository3 "frieda-golang-training-beginner/payment/repository"
	usecase2 "frieda-golang-training-beginner/payment/usecase"
	"github.com/labstack/gommon/log"
	"os"

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
	paymentCodeRepository := repository.PaymentCodeRepository{Conn: db}
	paymentCodeUsecase := _paymentCodeUsecase.PaymentCodeUsecase{PaymentCodeRepo: paymentCodeRepository, ContextTimeout: time.Duration(100000000)}
	_paymentCodeHttpDirectory.NewPaymentCodeHandler(e, paymentCodeUsecase)

	inquiryRepository := repository2.InquiryRepository{Conn: db}
	inquiryUsecase := usecase.InquiryUsecase{InquiryRepo: inquiryRepository, ContextTimeout: time.Duration(1000000000)}
	http.NewInquiryHandler(e, inquiryUsecase)

	paymentRepository := repository3.PaymentRepository{Conn: db}

	var QueueUrl = os.Getenv("SQS_QUEUE_URL")
	var AwsID = os.Getenv("AWS_ID")
	var AwsSecret = os.Getenv("AWS_SECRET")
	sess, err := aws.New(aws.Config{
		Address: QueueUrl,
		Region:  "us-east-1",
		ID: AwsID,
		Secret: AwsSecret,
	})
	if err != nil {
		log.Fatal(err)
	}

	sqs := aws.NewSQS(sess, time.Duration(10000000))
	paymentUsecase := usecase2.PaymentUsecase{PaymentRepo: paymentRepository,
		InquiryUsecase: inquiryUsecase, SQS: sqs, ContextTimeout: time.Duration(10000000000)}

	http2.NewPaymentHandler(e, paymentUsecase)

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

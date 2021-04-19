package main

import (
	_helloWorldHttpDirectory "frieda-golang-training-beginner/hello-world/directory/http"
	_helloWorldUseCase "frieda-golang-training-beginner/hello-world/usecase"
	_healthUseCase "frieda-golang-training-beginner/health/usecase"
	_healthHttpDirectory "frieda-golang-training-beginner/health/directory/http"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main(){
	e := echo.New()

	helloWorldUsecase := _helloWorldUseCase.NewHelloWorldUsecase()
	_helloWorldHttpDirectory.NewHelloWorldHandler(e, helloWorldUsecase)


	healthUsecase := _healthUseCase.NewHealthUsecase()
	_healthHttpDirectory.NewHealthHandler(e, healthUsecase)
	log.Fatal(e.Start("localhost:9090"))
}

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sferawann/pinjol/config"
	"github.com/sferawann/pinjol/controller"
	"github.com/sferawann/pinjol/model"
	"github.com/sferawann/pinjol/repository"
	"github.com/sferawann/pinjol/router"
	"github.com/sferawann/pinjol/usecase"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)
	db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Product{},
		&model.PaymentMethod{},
		&model.Transaction{},
		&model.AcceptStatus{},
		&model.Payment{},
	)

	//Init Repository
	userRepo := repository.NewUserRepoImpl(db)
	roleRepo := repository.NewRoleRepoImpl(db)
	proRepo := repository.NewProductRepoImpl(db)
	paymedRepo := repository.NewPaymentMethodRepoImpl(db)
	traRepo := repository.NewTraRepoImpl(db)
	accstatRepo := repository.NewAcceptStatusRepoImpl(db)
	payRepo := repository.NewPaymentRepoImpl(db)

	//Init Usecase
	userUsecase := usecase.NewUserUsecaseImpl(userRepo, roleRepo)
	roleUsecase := usecase.NewRoleUsecaseImpl(roleRepo)
	proUsecase := usecase.NewProductUsecaseImpl(proRepo)
	paymedUsecase := usecase.NewPaymentMethodUsecaseImpl(paymedRepo)
	traUsecase := usecase.NewTraUsecaseImpl(traRepo, userRepo, proRepo)
	accstatUsecase := usecase.NewAcceptStatusUsecaseImpl(accstatRepo, traRepo, proRepo)
	payUsecase := usecase.NewPaymentUsecaseImpl(payRepo, traRepo, proRepo, paymedRepo, userRepo)

	//Init controller
	userCon := controller.NewUserController(userUsecase)
	roleCon := controller.NewRoleController(roleUsecase)
	proCon := controller.NewProductController(proUsecase)
	paymedCon := controller.NewPaymentMethodController(paymedUsecase)
	traCon := controller.NewTraController(traUsecase)
	accstatCon := controller.NewAcceptStatusController(accstatUsecase)
	payCon := controller.NewPaymentController(payUsecase)

	routes := router.NewRouter(userRepo, userCon, roleCon, proCon, paymedCon, traCon, accstatCon, payCon)
	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if server_err := server.ListenAndServe(); err != nil {
		log.Fatal(server_err)
	}
}

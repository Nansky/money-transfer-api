package server

import (
	"context"
	"fmt"
	"log"
	"money-api-transfer/api/config"
	"money-api-transfer/api/db"
	"money-api-transfer/api/handler"
	"money-api-transfer/api/repository"
	"money-api-transfer/api/usecase"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// StartAPI starts API Server
func StartAPI(appConf config.AppConfig) {
	// fmt.Println(appConf)

	// Init DB
	dbCli, err := db.NewAppDB(appConf.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbCli.Close()

	// ctx := context.Background()

	// Init Repository
	tr := repository.NewTransactionRepository(dbCli)
	ar := repository.NewAccountRepository(dbCli)

	repo := repository.Repository{
		AccountRepo:     ar,
		TransactionRepo: tr,
	}

	// Init usecase
	vuc := usecase.NewValidateAccountUsecase(repo)
	tuc := usecase.NewBankTransfertUsecase(repo)
	uc := usecase.Usecases{
		ValidateAccountUc: vuc,
		BankTransferUc:    tuc,
	}

	// Init handler
	accValidationHandler := handler.AccountValidationHandler{}
	accValidationHandler.UseCase = uc
	accValidationHandler.AuthToken = appConf.ProjectSecret

	transferHandler := handler.TransferHandler{}
	transferHandler.UseCase = uc
	transferHandler.AuthToken = appConf.ProjectSecret

	callbackUpdateHandler := handler.TransactionCallbackHandler{}
	callbackUpdateHandler.UseCase = uc
	callbackUpdateHandler.AuthToken = appConf.ProjectSecret

	// setup http server here
	mux := http.NewServeMux()
	mux.Handle("/api/v1/account_validation", accValidationHandler)
	mux.Handle("/api/v1/transfer", transferHandler)
	mux.Handle("/api/v1/transfer/update_status", callbackUpdateHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", appConf.Port),
		Handler: mux,
	}

	go func() {
		log.Printf("Server start at :%d\n", appConf.Port)
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				return
			}
		}
	}()

	// Wait for an interrupt signal to gracefully shut down the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop // Wait for the interrupt signal
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Stop")
	}

	log.Printf("Server Stop Gracefully")
	// log.Fatal(http.ListenAndServe(":"+strconv.Itoa(appConf.Port), mux))
}

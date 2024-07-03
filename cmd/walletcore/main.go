package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com.br/silva4dev/golang-event-driven-arch-project/internal/database"
	"github.com.br/silva4dev/golang-event-driven-arch-project/internal/event"
	"github.com.br/silva4dev/golang-event-driven-arch-project/internal/event/handler"
	"github.com.br/silva4dev/golang-event-driven-arch-project/internal/usecase/create_account"
	"github.com.br/silva4dev/golang-event-driven-arch-project/internal/usecase/create_client"
	"github.com.br/silva4dev/golang-event-driven-arch-project/internal/usecase/create_transaction"
	"github.com.br/silva4dev/golang-event-driven-arch-project/internal/web"
	"github.com.br/silva4dev/golang-event-driven-arch-project/internal/web/webserver"
	"github.com.br/silva4dev/golang-event-driven-arch-project/pkg/events"
	"github.com.br/silva4dev/golang-event-driven-arch-project/pkg/kafka"
	"github.com.br/silva4dev/golang-event-driven-arch-project/pkg/uow"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	transactionCreatedEvent := event.NewTransactionCreated()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent)
	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)

	webserver := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	webserver.Start()
}

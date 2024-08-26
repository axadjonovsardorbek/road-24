package main

import (
	"auth/api"
	"auth/api/handler"
	"auth/config"
	"auth/grpc/clients"
	"auth/service"
	"auth/storage/postgres"
	"log"
	"log/slog"
)

func main() {
	cf := config.Load()

	conn, err := postgres.NewPostgresStorage(cf)

	if err != nil {
		slog.Error("Failed to connect postgres:", err)
	}

	defer conn.Db.Close()

	us := service.NewUsersService(conn)

	services, err := clients.NewGrpcClients(&cf)
	if err != nil {
		log.Fatalf("error while connecting clients. err: %s", err.Error())
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	handler := handler.NewHandler(us, services)

	roter := api.NewApi(handler)
	log.Println("Server is running on port ", cf.AUTH_PORT)
	if err := roter.Run(cf.AUTH_PORT); err != nil {
		slog.Error("Error:", err)
	}
}

package di

import (
	"log"

	"github.com/anjush-bhargavan/golib_admin/config"
	"github.com/anjush-bhargavan/golib_admin/pkg/books"
	"github.com/anjush-bhargavan/golib_admin/pkg/db"
	"github.com/anjush-bhargavan/golib_admin/pkg/handler"
	"github.com/anjush-bhargavan/golib_admin/pkg/repo"
	"github.com/anjush-bhargavan/golib_admin/pkg/server"
	"github.com/anjush-bhargavan/golib_admin/pkg/service"
	"github.com/anjush-bhargavan/golib_admin/pkg/user"
)

func Init() {
	config.LoadConfig()
	db := db.ConnectDB()
	userClient, err := user.ClientDial()
	if err != nil {
		log.Fatalf("something went wrong when dialing user client %v", err)
	}
	bookClient, err := books.ClientDial()
	if err != nil {
		log.Fatalf("something went wrong when dialing book client %v", err)
	}

	adminRepo := repo.NewAdminRepo(db)
	adminService := service.NewAdminService(adminRepo, userClient,bookClient)
	adminHandler := handler.NewAdminHandler(adminService)
	server.NewGrpcServer(adminHandler)
}

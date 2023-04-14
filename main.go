package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_grpcHandler "github.com/xoltawn/simple-file-storage-file-service/delivery/grpc"
	"github.com/xoltawn/simple-file-storage-file-service/delivery/grpc/filepb"
	"github.com/xoltawn/simple-file-storage-file-service/domain"
	"github.com/xoltawn/simple-file-storage-file-service/repository/localstorage"
	_postgres "github.com/xoltawn/simple-file-storage-file-service/repository/postgres"
	"github.com/xoltawn/simple-file-storage-file-service/usecase"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalln(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&domain.File{})
	if err != nil {
		log.Fatalln(err)
	}

	runPort, err := strconv.Atoi(os.Getenv("RUN_PORT"))
	if err != nil {
		log.Fatalln(err)
	}

	listener, err := net.Listen("tcp", fmt.Sprint(":", runPort))
	if err != nil {
		log.Fatal(err)
	}

	linkValidator := usecase.NewLinkValidator()
	bytesToReaderConvertor := usecase.NewBytesToReaderConvertor()
	bytesToLinksConvertor := usecase.NewBytesToLinksConvertor(bytesToReaderConvertor, linkValidator)

	fileStorage := localstorage.NewLocalStorage(os.Getenv("DOWNLOADED_IMAGES_PATH"))
	fileRepository := _postgres.NewFilePostgresRepository(db)
	fileDownloader := usecase.NewFileDownloader()
	fileUsecase := usecase.NewFileUsecase(fileStorage, fileRepository, fileDownloader, os.Getenv("DOWNLOADED_IMAGES_PATH"))

	s := grpc.NewServer()
	filepb.RegisterFileServiceServer(s, _grpcHandler.NewFileGRPCHandler(fileUsecase, bytesToLinksConvertor))

	log.Println("Server is listening on port", runPort)
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

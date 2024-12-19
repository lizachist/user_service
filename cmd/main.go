package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"user_service/internal/handler"
	"user_service/internal/repository"
	"user_service/internal/usecase"
)

func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Подключение к базе данных
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Инициализация репозитория
	userRepo := repository.NewUserRepository(db)

	// Инициализация use case
	userService := usecase.NewUserService(userRepo)

	// Инициализация обработчика
	userHandler := handler.NewUserHandler(userService)

	// Инициализация Echo
	e := echo.New()

	// Регистрация маршрутов
	userHandler.Register(e)

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

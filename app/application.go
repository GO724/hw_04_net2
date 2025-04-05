package application

import (
	"fmt"
	"hw_04_net2/internal/controller"
	"hw_04_net2/internal/service"
	"hw_04_net2/internal/storage"
	"hw_04_net2/internal/tghandlers"
	"log"
	"os"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

const logfile = "log/app.log"

type application struct {
	Log         *log.Logger
	db          *storage.Stor
	controller  *controller.Controller
	Bot         *bot.Bot
	tghandlers  *tghandlers.TgHandlers
	reghandlers map[string]string
}

func New() (*application, error) {

	logger := log.New(os.Stdout, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)

	file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error write log: %v\n", err)
	}
	defer file.Close()
	logger.SetOutput(file)

	fmt.Printf("log in %s\n", logfile)

	logger.Println("Start bootsrap sequence")

	logger.Println("Get environment...")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error load .env: %v\n", err)
	}

	logger.Println("Init telegram api...")
	bot, err := bot.New(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Printf("can't init telegram: %v", err)
		return &application{Log: logger}, fmt.Errorf("can't init telegram: %w", err)
	}

	logger.Println("Init storage...")
	stor, err := storage.New()
	if err != nil {
		logger.Printf("Can't init storage: %v", err)
		return &application{Log: logger}, fmt.Errorf("can't init storage: %w", err)
	}

	logger.Println("Init service...")
	service := service.New(stor)

	logger.Println("Init controller...")
	controller := controller.New(service)

	logger.Println("Init handlers...")
	tghandlers := tghandlers.New(bot, controller)
	registeredHandlers := tghandlers.Bindle()

	logger.Println("Telegram bot ready to dialog")
	fmt.Println("Telegram bot ready to dialog")

	return &application{
		Log:         logger,
		db:          stor,
		controller:  controller,
		Bot:         bot,
		tghandlers:  tghandlers,
		reghandlers: registeredHandlers,
	}, nil
}

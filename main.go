package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type How struct {
	MOTD        []string
	ErrorLog    *log.Logger
	InfoLog     *log.Logger
	Token       string
	CodeWord    string
	Whitelisted interface{}
	RootPath    string
}

func main() {
	cfg := How{}
	rp, _ := os.Getwd()
	err := godotenv.Load(rp + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg.CodeWord = os.Getenv("CODEWORD")
	cfg.Token = os.Getenv("TOKEN")
	cfg.Whitelisted = os.Getenv("WHITELISTED")
	infoLog, errorLog := cfg.startLoggers()
	cfg.ErrorLog = errorLog
	cfg.InfoLog = infoLog
	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	cfg.Menu(dg)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-ch

}
func (c *How) startLoggers() (*log.Logger, *log.Logger) {
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	return infoLog, errorLog
}

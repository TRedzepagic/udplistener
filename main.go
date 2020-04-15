package main

import (
	"log"
	"log/syslog"
	"net"
	"os"
	"time"

	"github.com/TRedzepagic/compositelogger/logs"
	_ "github.com/go-sql-driver/mysql"
)

type msg struct {
	message   string
	timestamp string
}

func main() {
	// Logger creation
	filepath := "serverlog"
	filelogger1 := logs.NewFileLogger(filepath)
	defer filelogger1.Close()

	stdoutLog := logs.NewStdLogger()
	defer stdoutLog.Close()

	systemlogger, _ := logs.NewSysLogger(syslog.LOG_NOTICE, log.LstdFlags)

	databaseLog := logs.NewDBLogger(logs.DatabaseConfiguration())
	defer databaseLog.Close()

	wantDebug := false

	log := logs.NewCustomLogger(wantDebug, filelogger1, stdoutLog, systemlogger, databaseLog)

	arguments := os.Args
	if len(arguments) == 1 {
		log.Warn("No port number specified... Exiting.....")
		return
	}

	PORT := arguments[1]
	resolvedUDPAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:"+PORT)
	if err != nil {
		log.Error(err)
		return
	}

	connection, err := net.ListenUDP("udp4", resolvedUDPAddr)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("Listening...")
	defer connection.Close()
	buffer := make([]byte, 1024)

	for {
		num, addr, err := connection.ReadFromUDP(buffer)
		log.Info(addr.String()+" says: ", string(buffer[0:num-1]))

		StructToSend := msg{message: "This is my reply", timestamp: time.Now().String()}

		data2 := []byte("Time : " + StructToSend.timestamp + " " + "Server Reply : " + StructToSend.message)

		_, err = connection.WriteToUDP(data2, addr)
		if err != nil {
			log.Error(err)
			return

		}

	}
}

package main

import (
	"net"
	"os"
	"time"

	"github.com/TRedzepagic/compositelogger/logs"
	_ "github.com/go-sql-driver/mysql"
)

func timer(log *logs.CompositeLog, connection *net.UDPConn, mapa *map[string]*net.UDPAddr) {
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		datatimer := []byte("Timer Tick! : 5 seconds have elapsed")
		log.Info("5 SECOND TIMER TRIGGERED ! " + string(datatimer))
		for k, v := range *mapa {
			_, err := connection.WriteToUDP(datatimer, v)
			log.Info("Sent timer info to " + k)
			if err != nil {
				log.Error(err)
				return
			}

		}
	}

}

func main() {
	// Logger creation
	filepath := "serverlog"
	filelogger1 := logs.NewFileLogger(filepath)
	defer filelogger1.Close()

	stdoutLog := logs.NewStdLogger()
	defer stdoutLog.Close()

	// OPTIONAL LOGGERS
	// systemlogger, _ := logs.NewSysLogger(syslog.LOG_NOTICE, log.LstdFlags)
	// databaseLog := logs.NewDBLogger(logs.DatabaseConfiguration())
	// defer databaseLog.Close()

	wantDebug := false

	log := logs.NewCustomLogger(wantDebug, filelogger1, stdoutLog)

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
	IPAddresses := make(map[string]*net.UDPAddr)
	go timer(log, connection, &IPAddresses)
	for {
		num, addr, err := connection.ReadFromUDP(buffer)
		IPAddresses[addr.String()] = addr
		log.Info(addr.String()+" says: ", string(buffer[0:num-1]))

		data := []byte("SERVER REPLY : HELLO !")
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			log.Error(err)
			return

		}

	}
}

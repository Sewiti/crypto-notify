package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sewiti/crypto-notify/internal/rules"
	"github.com/Sewiti/crypto-notify/pkg/coinlore"
)

const (
	rulesFile = "data/rules-set-2.json"
	interval  = 30 * time.Second
)

func main() {
	log.Println("started")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go waitSignal(cancel)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

Main:
	for {
		select {
		case <-ctx.Done():
			break Main

		case <-ticker.C:
			if err := exec(ctx, rulesFile); err != nil {
				log.Println(err)
			}
		}
	}
}

func waitSignal(onReceived func()) {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	<-sig
	log.Println("exiting")
	onReceived()
}

func exec(ctx context.Context, rulesFile string) error {
	rul, err := rules.Read(rulesFile)
	if err != nil {
		return err
	}

	rul = filter(rul)
	coins := distinct(rul) // In order to avoid duplicate requests

	coinsMap, err := fetch(ctx, coins)
	if err != nil {
		return err
	}

	any, err := check(&rul, coinsMap)
	if err != nil {
		return err
	}

	if any {
		err = rules.Write(rulesFile, rul)
	}

	return err
}

func check(r *rules.Rules, cm map[int]coinlore.Coin) (anyTrig bool, err error) {
	for i := range *r {
		coin, ok := cm[(*r)[i].CryptoID]
		if !ok {
			return false, fmt.Errorf("coinmap %d: index not found", (*r)[i].CryptoID)
		}

		trig, err := (*r)[i].Check(coin.PriceUSD)
		if err != nil {
			return false, err
		}

		if trig {
			log.Printf("%s id:%s\n", coin.NameID, (*r)[i].String())
			anyTrig = true
		}
	}

	return anyTrig, nil
}

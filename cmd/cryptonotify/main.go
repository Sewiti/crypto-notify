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
	rulesFilePath = "./data/rules-set-2.json"
	interval      = 30 * time.Second
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go setupSignal(cancel)

	ticker := time.NewTicker(interval)

Main:
	for {
		select {
		case <-ctx.Done():
			break Main

		case <-ticker.C:
			tick(ctx)
		}
	}
}

func setupSignal(onReceived func()) {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	<-sig
	log.Println("exiting")
	onReceived()
}

func tick(ctx context.Context) {
	rul, err := rules.Read(rulesFilePath)
	if err != nil {
		log.Println(err)
		return
	}

	rul = filter(rul)
	coins := distinct(rul) // In order to avoid duplicate requests

	coinsMap, err := fetch(ctx, coins)
	if err != nil {
		log.Println(err)
		return
	}

	trig, err := check(&rul, coinsMap)
	if err != nil {
		log.Println(err)
		return
	}

	if trig {
		err = rules.Write(rulesFilePath, rul)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func check(r *rules.Rules, cm map[int]coinlore.Coin) (anyTrig bool, err error) {
	for i, v := range *r {
		coin, ok := cm[v.CryptoID]
		if !ok {
			return false, fmt.Errorf("coinmap %d: index not found", v.CryptoID)
		}

		ruleTr, err := v.Check(coin.PriceUSD)
		if err != nil {
			return false, err
		}

		if ruleTr {
			op, err := formatOp(v.Operator)
			if err != nil {
				// Should never enter here due to rule.Check
				return false, err
			}

			log.Printf("%s (%d) price is %s %.2f\n", coin.Name, coin.ID, op, v.Price)
			(*r)[i].Triggered = true
			anyTrig = true
		}
	}

	return anyTrig, nil
}
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/Sewiti/crypto-notify/internal/rules"
	"github.com/Sewiti/crypto-notify/pkg/coinlore"
)

const rulesFilePath = "./data/rules-set-1.json"
const interval = 3 * time.Second

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go setupSignal(cancel)

	ticker := time.NewTicker(interval)

Main:
	for {
		select {
		case <-ctx.Done():
			break Main

		case <-ticker.C:
			checkRules(ctx)
		}
	}
}

func setupSignal(onReceived func()) {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	<-sig
	onReceived()
}

func checkRules(parent context.Context) {
	r, err := rules.Read(rulesFilePath)
	if err != nil {
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(parent, interval)
	defer cancel()

	trig := false // Any triggered
	wg := &sync.WaitGroup{}
	wg.Add(len(r))

	for i := range r {
		go foreach(ctx, &r[i], wg, &trig)
	}

	if trig {
		rules.Write(rulesFilePath, r)
	}
}

func foreach(ctx context.Context, rule *rules.Rule, wg *sync.WaitGroup, trig *bool) {
	defer wg.Done()

	coin, err := coinlore.GetCoin(ctx, rule.CryptoID)
	if err != nil {
		log.Println(err)
		return
	}

	price, err := strconv.ParseFloat(coin.PriceUSD, 64)
	if err != nil {
		log.Println(err)
		return
	}

	rTrig, err := rule.Check(price)
	if err != nil {
		log.Println(err)
		return
	}

	if rTrig {
		op, err := formatOperator(rule.Operator)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("%s id:%s price is %s %s\n", coin.Name, coin.ID, op, coin.PriceUSD)
		*trig = true
	}
}

func formatOperator(operator string) (string, error) {
	switch strings.ToLower(operator) {
	case "lt":
		return "less than", nil

	case "le":
		return "less than or equals", nil

	case "gt":
		return "greater than", nil

	case "ge":
		return "greater than or equals", nil

	case "eq":
		return "equals", nil

	case "ne":
		return "not equals", nil

	default:
		return "", fmt.Errorf("format %s: invalid operator", operator)
	}
}

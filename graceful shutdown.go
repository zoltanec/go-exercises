package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//server imitation
func server(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.println("Завершаем работу сервера")
			return

		default:
			fmt.println("Сервер работает")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go server(ctx, &wg)

	// Перехватываем мигнал Сигтерм или Сигинт(ктрл+с)
	sigChan := make(chan os.Sygnal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Ожидаем сигнал завершения
	sig := <-sigChan
	fmt.Println("Получен сигнал:", sig)

	// Завершаем сервер
	cancel()
	wg.Wait()

	fmt.Println("Приложение завершено корректно")
}

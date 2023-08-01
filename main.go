package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	userID := 10
	num, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Get: ", num)
	fmt.Println("took: ", time.Since(start))
}

type Rspns struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300) // 시간 제한을 걸고
	defer cancel()                                                // 종료시킨다.
	rspnsCh := make(chan Rspns)                                   // 비동기 수신 채널
	go func() {
		num, err := fetchThirdPartyStuffWhichCanBeSlow()
		rspnsCh <- Rspns{value: num, err: err} // 완료되면 채널에 던져라
	}() // goroutine에 던져버리는 관용문법
	for {
		select {
		case <-ctx.Done(): // 리턴없이 끝남처리(빈 구조체)되었다면
			return 0, fmt.Errorf("fetching data took to long")
		case rspns := <-rspnsCh:
			return rspns.value, rspns.err
		}
	}
}
func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 400)
	return 10, nil
}

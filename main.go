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
func fetchUserData(ctx context.Context, userID int) (int, error) {
	num, err := fetchThirdPartyStuffWhichCanBeSlow()
	if err != nil {
		return 0, err
	}
	return num, nil
}
func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 300)
	return 10, nil
}

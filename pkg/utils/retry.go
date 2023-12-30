package utils

import (
	"time"
	"math"
)

type RetryCallback func() error

func Retry(callback RetryCallback, maxRetries int, delay time.Duration) error {
	ch := make(chan error)
	go func(){
		for i := 0; i < maxRetries; i++ {
			if err := callback();  err == nil{
				 ch <- nil
				 return
			}
			time.Sleep(time.Duration(math.Pow(2, float64(i)) * delay))
		}
		ch <- fmt.Errorf("failed after %d retries",maxRetries)
	}()
	err <- ch
	close(ch)
	return err
}
package boring

import "fmt"
import "time"
import "math/rand"

func Boring01(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}
func Boring10(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

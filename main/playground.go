package main


func main() {
	ch := make(chan int)
	batteryConsumer()
	<-ch
}

const THREADS = 10

func batteryConsumer() {
	for i := 0; i < THREADS; i++ {
		go func() {
			counter := 0
			lowerLimit := 0
			upperLimit := 0

			var increment bool
			for {
				if counter <= lowerLimit {
					increment = true
				} else if counter >= upperLimit {
					increment = false
				}

				if increment {
					counter++
				} else {
					counter--
				}
			}

		}()
	}
}

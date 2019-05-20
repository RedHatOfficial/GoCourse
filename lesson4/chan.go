var channel chan int = make(chan int)
channel <- 5
data := <-channel

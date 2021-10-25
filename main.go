package main

func main() {
	url := "amqp://guest:guest@localhost:5672/"
	queue := "imageID"

	rabbitListen(url, queue, encode, remove)
}

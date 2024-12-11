package main

func main() {
	go StartMqtt()
	go HttpServer()
	select {}
}

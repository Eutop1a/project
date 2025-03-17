package main

//go:generate go run generator.go
func main() {
	genUser()
	genQuestion()
	genKnowledge()
}

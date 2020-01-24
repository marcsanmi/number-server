package server

type NumberServer interface {
	Listen(port string) error
	Run()
	Close()
}

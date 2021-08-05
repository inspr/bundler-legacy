package vm

type Interface interface {
	Run(req Request) chan Response
}

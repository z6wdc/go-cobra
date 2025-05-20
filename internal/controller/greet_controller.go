package controller

import (
	"github.com/z6wdc/go-cobra/internal/presenter"
	"github.com/z6wdc/go-cobra/internal/usecase"
)

func RunGreet(name string) {
	output := usecase.Greet(name)
	presenter.Show(output)
}

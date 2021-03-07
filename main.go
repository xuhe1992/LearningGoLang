package main

import (
	"github.com/xuhe1992/LearningGoLang/structure"
)


func main() {
	a := structure.A{Name: "elyot", Age: 10}
	structure.TestDuck(&a)
}
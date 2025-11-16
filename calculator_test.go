package main

import (
	"fmt"

	"github.com/cucumber/godog"
)

var a, b, result int

func twoNumbers(x, y int) error {
	a, b = x, y
	return nil
}

func iAddThem() error {
	result = a + b
	return nil
}

func theResultShouldBe(expected int) error {
	if result != expected {
		return fmt.Errorf("expected %d, but got %d", expected, result)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
  ctx.Step(`^I have the numbers (\d+) and (\d+)$`, twoNumbers)
  ctx.Step(`^I add them$`, iAddThem)
  ctx.Step(`^the result should be (\d+)$`, theResultShouldBe)
}
package generator

import(
	"testing"
)

func Handler(c chan interface{}) bool{
	c <- "hello"
	return OK
}

func TestRunGenerator(t *testing.T){
	var c chan interface{}

	generator := Generator{
		Handler:Handler,
		Chan:c,
	}

	generator.Run()
	out := <-c
	if out.(string) != "hello"{
		t.Error("expect", "hello", "actual", out)
	}
}

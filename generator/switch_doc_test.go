package generator

import (
	"fmt"
	"log"
)

func ExampleSwitch_Generate() {
	generator := NewSwitch("str")
	generator = generator.AddCase(
		NewCase(`"foo"`, NewRawStatement(`fmt.Printf("str is foo\n")`)),
		NewCase(`"bar"`, NewRawStatement(`fmt.Printf("str is bar\n")`)),
	)
	generator = generator.Default(
		NewDefaultCase(NewRawStatement(`fmt.Printf("here is default\n")`)),
	)

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}

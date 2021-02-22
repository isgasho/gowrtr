package generator

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"
	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateCode(t *testing.T) {
	expected := `// THIS CODE WAS AUTO GENERATED

package mypkg

import (
	"fmt"
)

type MyInterface interface {
	MyFunc(foo string) (string, error)
}

type MyStruct struct {
	Foo string
	Bar int64
}

func (m *MyStruct) MyFunc(foo string) (string, error) {
	{
		str := 		func(bar string) string {
			return bar
		}(foo)

		if str == "" {
			for i := 0; i < 3; i++ {
				fmt.Printf("%d\n", i)
			}
		}

		switch str {
		case "":
			// empty string
		case "foo":
			// foo string
		default:
			// default
		}

		return str, nil
	}
}
`
	myFuncSignature := NewFuncSignature("MyFunc").
		AddParameters(
			NewFuncParameter("foo", "string"),
		).
		AddReturnTypes("string", "error")

	generator := NewRoot(
		NewComment(" THIS CODE WAS AUTO GENERATED"),
		NewNewline(),
		NewPackage("mypkg"),
		NewNewline(),
		NewImport("fmt"),
		NewNewline(),
		NewInterface("MyInterface").
			AddSignatures(myFuncSignature),
		NewNewline(),
		NewStruct("MyStruct").
			AddField("Foo", "string").
			AddField("Bar", "int64"),
		NewNewline(),
	).AddStatements(
		NewFunc(
			NewFuncReceiver("m", "*MyStruct"),
			NewFuncSignature("MyFunc").
				AddParameters(
					NewFuncParameter("foo", "string"),
				).
				AddReturnTypes("string", "error"),
		).AddStatements(
			NewCodeBlock(
				NewRawStatement("str := ").WithNewline(false),
				NewAnonymousFunc(
					false,
					NewAnonymousFuncSignature().
						AddParameters(NewFuncParameter("bar", "string")).
						AddReturnTypes("string"),
					NewReturnStatement("bar"),
				).Invocation(NewFuncInvocation("foo")),
				NewNewline(),
				NewIf(`str == ""`).
					AddStatements(
						NewFor(`i := 0; i < 3; i++`).AddStatements(
							NewRawStatement(`fmt.Printf("%d\n", i)`),
						),
					),
				NewNewline(),
				NewSwitch("str").
					AddCase(
						NewCase(
							`""`,
							NewComment(" empty string"),
						),
						NewCase(
							`"foo"`,
							NewComment(" foo string"),
						),
					).
					Default(
						NewDefaultCase(NewComment(" default")),
					),
				NewNewline(),
				NewReturnStatement("str", "nil"),
			),
		),
	)

	generated, err := generator.EnableSyntaxChecking().Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, generated)

	generated, err = generator.Statements(NewComment("modified")).Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "//modified\n", generated)
}

func TestShouldGenerateCodeWithIndent(t *testing.T) {
	expected := `		// THIS CODE WAS AUTO GENERATED

		package mypkg

		import (
			"fmt"
		)

		type MyInterface interface {
			MyFunc(foo string) (string, error)
		}

		type MyStruct struct {
			Foo string
			Bar int64
		}

		func (m *MyStruct) MyFunc(foo string) (string, error) {
			{
				str := 				func(bar string) string {
					return bar
				}(foo)

				if str == "" {
					for i := 0; i < 3; i++ {
						fmt.Printf("%d\n", i)
					}
				}

				switch str {
				case "":
					// empty string
				case "foo":
					// foo string
				default:
					// default
				}

				return str, nil
			}
		}
`
	myFuncSignature := NewFuncSignature("MyFunc").
		AddParameters(
			NewFuncParameter("foo", "string"),
		).
		AddReturnTypes("string", "error")

	generator := NewRoot(
		NewComment(" THIS CODE WAS AUTO GENERATED"),
		NewNewline(),
		NewPackage("mypkg"),
		NewNewline(),
		NewImport("fmt"),
		NewNewline(),
		NewInterface("MyInterface").
			AddSignatures(myFuncSignature),
		NewNewline(),
		NewStruct("MyStruct").
			AddField("Foo", "string").
			AddField("Bar", "int64"),
		NewNewline(),
	).AddStatements(
		NewFunc(
			NewFuncReceiver("m", "*MyStruct"),
			NewFuncSignature("MyFunc").
				AddParameters(
					NewFuncParameter("foo", "string"),
				).
				AddReturnTypes("string", "error"),
		).AddStatements(
			NewCodeBlock(
				NewRawStatement("str := ").WithNewline(false),
				NewAnonymousFunc(
					false,
					NewAnonymousFuncSignature().
						AddParameters(NewFuncParameter("bar", "string")).
						AddReturnTypes("string"),
					NewReturnStatement("bar"),
				).Invocation(NewFuncInvocation("foo")),
				NewNewline(),
				NewIf(`str == ""`).
					AddStatements(
						NewFor(`i := 0; i < 3; i++`).AddStatements(
							NewRawStatement(`fmt.Printf("%d\n", i)`),
						),
					),
				NewNewline(),
				NewSwitch("str").
					AddCase(
						NewCase(
							`""`,
							NewComment(" empty string"),
						),
						NewCase(
							`"foo"`,
							NewComment(" foo string"),
						),
					).
					Default(
						NewDefaultCase(NewComment(" default")),
					),
				NewNewline(),
				NewReturnStatement("str", "nil"),
			),
		),
	)

	generated, err := generator.Generate(2)

	assert.NoError(t, err)
	assert.Equal(t, expected, generated)
}

func TestShouldGenerateCodeWithGofmt(t *testing.T) {
	expected := `// THIS CODE WAS AUTO GENERATED

package mypkg

import (
	"fmt"
)

type MyInterface interface {
	MyFunc(foo string) (string, error)
}

type MyStruct struct {
	Foo string
	Bar int64
}

func (m *MyStruct) MyFunc(foo string) (string, error) {
	{
		str := func(bar string) string {
			return bar
		}(foo)

		if str == "" {
			for i := 0; i < 3; i++ {
				fmt.Printf("%d\n", i)
			}
		}

		switch str {
		case "":
			// empty string
		case "foo":
			// foo string
		default:
			// default
		}

		return str, nil
	}
}
`
	myFuncSignature := NewFuncSignature("MyFunc").
		AddParameters(
			NewFuncParameter("foo", "string"),
		).
		AddReturnTypes("string", "error")

	generator := NewRoot(
		NewComment(" THIS CODE WAS AUTO GENERATED"),
		NewNewline(),
		NewPackage("mypkg"),
		NewNewline(),
		NewImport("fmt"),
		NewNewline(),
		NewInterface("MyInterface").
			AddSignatures(myFuncSignature),
		NewNewline(),
		NewStruct("MyStruct").
			AddField("Foo", "string").
			AddField("Bar", "int64"),
		NewNewline(),
		NewFunc(
			NewFuncReceiver("m", "*MyStruct"),
			NewFuncSignature("MyFunc").
				AddParameters(
					NewFuncParameter("foo", "string"),
				).
				AddReturnTypes("string", "error"),
		).AddStatements(
			NewCodeBlock(
				NewRawStatement("str := ").WithNewline(false),
				NewAnonymousFunc(
					false,
					NewAnonymousFuncSignature().
						AddParameters(NewFuncParameter("bar", "string")).
						AddReturnTypes("string"),
					NewReturnStatement("bar"),
				).Invocation(NewFuncInvocation("foo")),
				NewNewline(),
				NewIf(`str == ""`).
					AddStatements(
						NewFor(`i := 0; i < 3; i++`).AddStatements(
							NewRawStatement(`fmt.Printf("%d\n", i)`),
						),
					),
				NewNewline(),
				NewSwitch("str").
					AddCase(
						NewCase(
							`""`,
							NewComment(" empty string"),
						),
						NewCase(
							`"foo"`,
							NewComment(" foo string"),
						),
					).
					Default(
						NewDefaultCase(NewComment(" default")),
					),
				NewNewline(),
				NewReturnStatement("str", "nil"),
			),
		),
	).Gofmt("-s")

	generated, err := generator.Generate(0)

	assert.NoError(t, err)
	assert.Equal(t, expected, generated)
}

func TestShouldGenerateCodeWithGoimport(t *testing.T) {
	expected := `// THIS CODE WAS AUTO GENERATED

package mypkg

import "fmt"

type MyInterface interface {
	MyFunc(foo string) (string, error)
}

type MyStruct struct {
	Foo string
	Bar int64
}

func (m *MyStruct) MyFunc(foo string) (string, error) {
	{
		str := func(bar string) string {
			return bar
		}(foo)

		if str == "" {
			for i := 0; i < 3; i++ {
				fmt.Printf("%d\n", i)
			}
		}

		switch str {
		case "":
			// empty string
		case "foo":
			// foo string
		default:
			// default
		}

		return str, nil
	}
}
`
	myFuncSignature := NewFuncSignature("MyFunc").
		AddParameters(
			NewFuncParameter("foo", "string"),
		).
		AddReturnTypes("string", "error")

	generator := NewRoot(
		NewComment(" THIS CODE WAS AUTO GENERATED"),
		NewNewline(),
		NewPackage("mypkg"),
		NewInterface("MyInterface").
			AddSignatures(myFuncSignature),
		NewNewline(),
		NewStruct("MyStruct").
			AddField("Foo", "string").
			AddField("Bar", "int64"),
		NewNewline(),
		NewFunc(
			NewFuncReceiver("m", "*MyStruct"),
			NewFuncSignature("MyFunc").
				AddParameters(
					NewFuncParameter("foo", "string"),
				).
				AddReturnTypes("string", "error"),
		).AddStatements(
			NewCodeBlock(
				NewRawStatement("str := ").WithNewline(false),
				NewAnonymousFunc(
					false,
					NewAnonymousFuncSignature().
						AddParameters(NewFuncParameter("bar", "string")).
						AddReturnTypes("string"),
					NewReturnStatement("bar"),
				).Invocation(NewFuncInvocation("foo")),
				NewNewline(),
				NewIf(`str == ""`).
					AddStatements(
						NewFor(`i := 0; i < 3; i++`).AddStatements(
							NewRawStatement(`fmt.Printf("%d\n", i)`),
						),
					),
				NewNewline(),
				NewSwitch("str").
					AddCase(
						NewCase(
							`""`,
							NewComment(" empty string"),
						),
						NewCase(
							`"foo"`,
							NewComment(" foo string"),
						),
					).
					Default(
						NewDefaultCase(NewComment(" default")),
					),
				NewNewline(),
				NewReturnStatement("str", "nil"),
			),
		),
	).Goimports()

	generated, err := generator.Generate(0)

	assert.NoError(t, err)
	assert.Equal(t, expected, generated)
}

func TestShouldGenerateCodeRaisesError(t *testing.T) {
	generator := NewRoot(
		NewComment(" THIS CODE WAS AUTO GENERATED"),
		NewNewline(),
		NewPackage("mypkg"),
		NewNewline(),
		NewFunc(
			nil,
			nil,
		),
	)

	_, err := generator.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.FuncSignatureIsNilError("").Error(), " ")[0],
	), err.Error())
}

func TestShouldGenerateCodeRaiseErrorWhenCodeFormatterIsExited(t *testing.T) {
	{
		generator := NewRoot(
			NewRawStatement("\timport something"),
		).EnableSyntaxChecking()

		_, err := generator.Generate(0)
		assert.Regexp(t, regexp.MustCompile(`^\[GOWRTR-13\] code formatter raises error: command='gofmt -e'.+`), err.Error())
	}

	{
		generator := NewRoot(
			NewRawStatement("\timport something"),
		).Gofmt()

		_, err := generator.Generate(0)
		assert.Regexp(t, regexp.MustCompile(`^\[GOWRTR-13\] code formatter raises error: command='gofmt'.+`), err.Error())
	}

	{
		generator := NewRoot(
			NewRawStatement("\timport something"),
		).Goimports()

		_, err := generator.Generate(0)
		assert.Regexp(t, regexp.MustCompile(`^\[GOWRTR-13\] code formatter raises error: command='goimports'.+`), err.Error())
	}

	{
		_, err := applyCodeFormatter("", "not-existed-cmd")
		assert.Regexp(t, regexp.MustCompile(`^\[GOWRTR-13\] code formatter raises error: command='not-existed-cmd'.+`), err.Error())
	}
}

// to check STDIN buffer behavior with code formatter
func TestMassiveAmountCode(t *testing.T) {
	generator := NewRoot(
		NewComment(" THIS CODE WAS AUTO GENERATED"),
		NewNewline(),
		NewPackage("mypkg"),
		NewNewline(),
		NewImport("fmt"),
		NewNewline(),
	)

	for i := 0; i <= 1000; i++ {
		generator = generator.AddStatements(
			NewFunc(
				nil,
				NewFuncSignature(fmt.Sprintf("f%d", i)),
				NewRawStatement(
					`fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")`,
				).WithNewline(true),
			),
		)
	}

	generated, err := generator.Gofmt().Generate(0)
	assert.NoError(t, err)
	assert.NotEmpty(t, generated)
}

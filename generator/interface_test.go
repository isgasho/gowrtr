package generator

import (
	"regexp"
	"strings"
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGeneratingInterfaceCodeBeSuccessful(t *testing.T) {
	exp1 := `type myInterface interface {
}
`
	exp2 := `type myInterface interface {
	myFunc()
}
`
	exp3 := `type myInterface interface {
	myFunc1()
	myFunc2(foo string) (string, error)
}
`
	dataset := map[string]*Interface{
		exp1: NewInterface("myInterface"),
		exp2: NewInterface("myInterface").
			AddSignatures(NewFuncSignature("myFunc")),
		exp3: NewInterface(
			"myInterface",
			NewFuncSignature("myFunc1"),
		).AddSignatures(
			NewFuncSignature("myFunc2").
				AddParameters(NewFuncParameter("foo", "string")).
				AddReturnTypes("string", "error"),
		),
	}

	for expected, in := range dataset {
		got, err := in.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	}
}

func TestShouldGeneratingInterfaceCodeWithSetter(t *testing.T) {
	generator := NewInterface(
		"myInterface",
		NewFuncSignature("myFunc1"),
	).AddSignatures(
		NewFuncSignature("myFunc2").
			AddParameters(NewFuncParameter("foo", "string")).
			AddReturnTypes("string", "error"),
	)

	expected := `type myInterface interface {
	myFunc1()
	myFunc2(foo string) (string, error)
}
`
	got, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)

	generator = generator.Signatures(NewFuncSignature("myFunc3"))
	expected = `type myInterface interface {
	myFunc3()
}
`
	got, err = generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestShouldGeneratingInterfaceCodeWithIndentBeSuccessful(t *testing.T) {
	exp1 := `		type myInterface interface {
		}
`
	exp2 := `		type myInterface interface {
			myFunc()
		}
`
	exp3 := `		type myInterface interface {
			myFunc1()
			myFunc2(foo string) (string, error)
		}
`
	dataset := map[string]*Interface{
		exp1: NewInterface("myInterface"),
		exp2: NewInterface("myInterface").
			AddSignatures(NewFuncSignature("myFunc")),
		exp3: NewInterface(
			"myInterface",
			NewFuncSignature("myFunc1"),
		).AddSignatures(
			NewFuncSignature("myFunc2").
				AddParameters(NewFuncParameter("foo", "string")).
				AddReturnTypes("string", "error"),
		),
	}

	for expected, in := range dataset {
		got, err := in.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	}
}

func TestShouldRaiseErrorWhenInterfaceNameIsEmpty(t *testing.T) {
	in := NewInterface("")
	_, err := in.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.InterfaceNameIsEmptyError("").Error(), " ")[0],
	), err.Error())
}

func TestShouldRaiseErrorWhenFuncSignatureRaisesError(t *testing.T) {
	in := NewInterface(
		"myInterface",
		NewFuncSignature(""),
	)
	_, err := in.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.FuncNameIsEmptyError("").Error(), " ")[0],
	), err.Error())
}

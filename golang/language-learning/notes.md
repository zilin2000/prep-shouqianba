> source: 《Go语言学习笔记》  
> https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/install-go




# Hello
> source: https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/hello-world#zai-ci-hui-dao-hello-world



hello.go
```go
package main

import "fmt"

const englishHello = "Hello, "

func Hello(name string) string {
	return englishHello + name
}

func main() {
	fmt.Println(Hello("Bruce"))
}
```

hello_test.go

```go
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Bruce")
	want := "Hello, Bruce"

	if got != want {
		t.Errorf("got '%q want '%q", got, want)
	}
}
```

`const` make sure that we don't write duplicate "Hello, " if we need

# Refactor code

We require that when we enter empty string, we print "Hello, World"

```go
package main

import "fmt"

const englishHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		return englishHello + "World"
	}
	return englishHello + name
}

func main() {
	fmt.Println(Hello("Bruce"))
}


package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectHello := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("this is testing for hello with name", func(t *testing.T) {
		got := Hello("Bruce")
		want := "Hello, Bruce"

		assertCorrectHello(t, got, want)
	})

	t.Run("this is testing for hello with empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectHello(t, got, want)
	})

}
```

`t.Helper()` used for display the line number of error if there is
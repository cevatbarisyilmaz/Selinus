# Selinus

Selinus is a general purpose programming language project that is in very early development.

Some working code examples:

### Hello World
```selinus
println("Hello World!")
```

### Recursive Fibonacci
```selinus
func int fibonacci(int n)
    if n == 0 || n == 1
		return n
	return fibonacci(n - 1) + fibonacci(n - 2)

loop 0 to 10 as i
    println("Fibonacci " + i + ": " + fibonacci(i))
    end
```
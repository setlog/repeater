# repeater ![](https://github.com/setlog/repeater/workflows/Tests/badge.svg)

Repeatedly invoke a function at a given interval until the process receives a termination signal.

Example usage:

```go
var twoWords := [2]string("Hello", "World")

type MyProcessor struct {
    i int
}

func (p *MyProcessor) Process(ctx context.Context) {
    fmt.Println(twoWords[p.i])
    p.i = (p.i + 1) % 2
}

func (p *MyProcessor) CleanUp() {
    if p.i != 0 {
        fmt.Println(twoWords[1])
    }
}

func main() {
    callInterval := time.Second
    makeFirstCallImmediately := true
    repeater.New(&MyProcessor{}).Run(
        context.Background(),
        callInterval,
        makeFirstCallImmediately)
}
```

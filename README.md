# intn: array of n-bit integer

[godoc](https://godoc.org/github.com/wtnb75/intn)

## Usage

```go
import "github.com/wtnb75/intn"
  :
a := intn.NewArrayBit(5)   // array of 5-bit integer
intn.Push(a, 5)
intn.Push(a, 4)
fmt.Println(a)             // [5 4]
intn.Add(a, 0, 2)          // a[0]+=2 -> [7 4]
var sum uint
for v := intn.Each(a) {
  sum += v
}
fmt.Println(a, sum)        // [7 4] 11
```

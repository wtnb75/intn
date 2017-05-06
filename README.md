# intn: array of n-bit integer

[godoc](https://godoc.org/github.com/wtnb75/intn)

## Usage

```go
import "github.com/wtnb75/intn"
  :
a := intn.NewArray(5)   // array of 5-bit integer
a.Append(5)
a.Append(4)
fmt.Println(a)          // [5 4]
a.Add(0, 2)             // a[0]+=2 -> [7 4]
var sum uint
for v := range a.Each() {
  sum += v
}
fmt.Println(a, sum)     // [7 4] 11
```

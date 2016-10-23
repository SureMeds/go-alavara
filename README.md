# go-alavara
By Adi Suresh

Example Usage
```go
import (
  "fmt"
  "github.com/SureMeds/go-alavara"
)
client := alavara.Client{key: "INSERT API KEY KERE"}
zipCode := "10019"
taxRate, err := client.RequestTaxRate(zipCode)
if err != nil {
  fmt.Errorf("%s\n", err.Error())
}
```

# matchers-translate

Translates incompatible Prometheus-like matchers.

## Usage

matchers-translate provides a number of functions to help translate incompatible matchers.

### IsEquivalent

IsEquivalent returns true if the matcher(s) can be parsed with both parsers, and the parsed matcher(s) are equivalent.

```go
func IsEquivalent(s string) (bool, error) 
```

Here is an example of its usage:

```go
package main

import (
	"fmt"

	"github.com/grobinson-grafana/matchers-translate"
)

func main() {
	ok, err := matchers_translate.IsEquivalent("foo=[a-zA-Z]+")
	if err != nil {
		// do something with error
	}
	fmt.Println(ok) // prints false
}
```

### Translate

Translate will translate matchers that can be parsed with the old parser but not the new parser into a matcher that can be parsed with both. If however the matcher can be parsed with both parsers, and the parsed matcher(s) are equivalent, then no translation is required

```go
func Translate(s string) (string, error)
```

Here is an example of its usage:

```go
package main

import (
	"fmt"

	"github.com/grobinson-grafana/matchers-translate"
)

func main() {
	expr, err := matchers_translate.Translate("foo=[a-zA-Z]+")
	if err != nil {
		// do something with error
	}
	fmt.Println(expr) // prints {foo="[a-zA-Z]+"}
}
```

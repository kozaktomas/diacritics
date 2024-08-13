# Remove diacritics

The library can remove diacritics (special characters) from a string.

## Examples:

| Input   | Output  |
|---------|---------|
| "áéíóú" | "aeiou" |
| "çãõ"   | "cao"   |
| "ñ"     | "n"     |

## Usage

```go
package main

import (
	"fmt"
	"github.com/kozaktomas/diacritics"
)

func main() {
	input := "áéíóú"
	output, err := diacritics.Remove(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(output) // Output: "aeiou"
}

```

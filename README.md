# 👋 Welcome

# INDODATE
indodate is a package for golang programming language for date convertion on indonesian format

### How to install
```sh
go get github.com/ekokurniadi/indodate
```

### Then import the indodate package

```go
import "github.com/ekokurniadi/indodate"
```
### Example Project

```go
package main

import (
	"fmt"
	"time"

	"github.com/ekokurniadi/indodate"
)

func main() {
	fmt.Println(indodate.LetterDate(time.Now()))    // output : 23 Oktober 2021
	fmt.Println(indodate.DateWithSlash(time.Now())) // output : 23/10/2021
	fmt.Println(indodate.DateWithMin(time.Now()))   // output : 23-10-2021
}

```

### Have Fun :D

# Go wrapper for JLinkExe

## Installing
```bash
go get github.com/akosmarton/go-jlink-wrapper/jlink
```

## Usage Example

```go
package main

import (
	"log"

	"github.com/akosmarton/go-jlink-wrapper/jlink"
)

func main() {
	var err error

	j := jlink.NewJLink("JLinkExe", "swd", "auto", "MK64FN1M0xxx12 (allow security)")

	err = j.Erase()
	if err != nil {
		log.Fatal(err)
	}

	err = j.LoadBin("file.bin", 0)
	if err != nil {
		log.Fatal(err)
	}

	err = j.Reset()
	if err != nil {
		log.Fatal(err)
	}
}
```

## Supported Commands
- erase
- loadbin
- reset

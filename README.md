# HTMLtoJSON
HTMLtoJSON is a HTML parser, based on net/html package. This package actually just to simplify HTML parsing. If you need more complex HTML processing, please use net/html as its offer more features. The package name is actually is not really fitting for this package purpose, but I use this package for may scraper engines, so I don't really want to bother with changing the package name...

### Installation
HTMLtoJSON requires Golang v1.14 or higher

```sh
$ GO111MODULE=on go get github.com/tamboto2000/htmltojson
```

# Examples

### Example 1
Parse from file

```go
package main

import "github.com/tamboto2000/htmltojson"

func main() {
	// Parse from file
	node, err := htmltojson.ParseFromFile("test.html")
	if err != nil {
		panic(err.Error())
	}

	// Save node
	if err := htmltojson.Save(node); err != nil {
		panic(err.Error())
	}
}
```

### Example 2
Parse from reader
```go
package main

import (
	"os"

	"github.com/tamboto2000/htmltojson"
)

func main() {
	f, err := os.Open("test.html")
	if err != nil {
		panic(err.Error())
	}

	defer f.Close()

	// Parse from io.Reader
	node, err := htmltojson.ParseFromReader(f)
	if err != nil {
		panic(err.Error())
	}

	// Save node
	if err := htmltojson.Save(node); err != nil {
		panic(err.Error())
	}
}
```

### Example 3
Parse from string

```go
package main

import (
	"io/ioutil"
	"os"

	"github.com/tamboto2000/htmltojson"
)

func main() {
	f, err := os.Open("test.html")
	if err != nil {
		panic(err.Error())
	}

	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}

	// Convert to string
	htmlString := string(bytes)

	// Parse from string
	node, err := htmltojson.ParseString(htmlString)
	if err != nil {
		panic(err.Error())
	}

	// Save node
	if err := htmltojson.Save(node); err != nil {
		panic(err.Error())
	}
}
```

### Example 4
Parse from bytes

```go
package main

import (
	"io/ioutil"
	"os"

	"github.com/tamboto2000/htmltojson"
)

func main() {
	f, err := os.Open("test.html")
	if err != nil {
		panic(err.Error())
	}

	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}

	// Parse from bytes
	node, err := htmltojson.ParseBytes(bytes)
	if err != nil {
		panic(err.Error())
	}

	// Save node
	if err := htmltojson.Save(node); err != nil {
		panic(err.Error())
	}
}
```

License
----

MIT

## 🚀 Quick Start

### Installation

```sh
$ go get https://github.com/Tylerholland12/gogin-practice
```

### Instantiate API

Create `main.go`

```go

package main

import (

"github.com/Tylerholland12/gogin-practice/httpd"

)

func main() {

    r := newsfeed.New()
    r.Run()

}

```
Start 

```sh
$ Make dev
```
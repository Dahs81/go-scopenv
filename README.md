## go-scopenv

This is a port of a node.js module written by weisjohn.

## [scopenv](https://github.com/weisjohn/scopenv)

#### Usage

Given these environment variables:

```
HOST="127.0.0.1"
CORE_HOST="154.234.323.2" // This is a random IP
CUSTOMER_CORE_HOST="84.222.43.112"
CORE_TOKEN="jf348943jf8934j8489" // Also fake
```

```
package main

import (
  "fmt"

  "github.com/Dahs81/go-scopenv/scopenv"
  )

  func main() {

    m := []string{"host", "token"}
    test := scopenv.Scopenv(m, "core", "customer")

    fmt.Println(test)
  }

```

##### Output

```
map[host:84.222.43.112 token:jf348943jf8934j8489]
```

If you remove customer:

```
test := scopenv.Scopenv(m, "core")
```

##### Output

```
map[host:154.234.323.2 token:jf348943jf8934j8489]
```

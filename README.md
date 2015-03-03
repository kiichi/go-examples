go-examples
===========

My Go Language Sandbox.


https://golang.org/doc/code.html


!!Setup

```{sh}
export GOPATH=$HOME/work/go
export PATH=$PATH:$GOPATH/bin
```

!!Build

```{sh}
cd ~/work/prj/go/go-examples/user/hello
go install
```

or


```{sh}
go install go-examples/user/hello
```

!!Run

```{sh}
cd ~/work/prj/go/bin/hello
```
or

```{sh}
$GOPATH/bin/hello
```



Testing

1. Create same file name with _test (e.g. mylib_test.go)
2. Create function starting with Upper Case T

```{go}
package jsonexample

import (
    "testing"                                                                                                                                                                                                    
)

func TestMain(m *testing.T) {
	// test code
}

```

More details, see http://golang.org/pkg/testing/

```{sh}
go test
```






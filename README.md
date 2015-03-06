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

func TestMyFunc(m *testing.T) {
	// test code
	// call MyFunc()
	// Assert
}

```

More details, see http://golang.org/pkg/testing/

```{sh}
go test
```

go-vim setup

install neobundle

https://github.com/Shougo/neobundle.vim

then in config file, add

```
NeoBundle 'fatih/vim-go'
```

https://github.com/fatih/vim-go/

after that, in vim install binary

```
:GoInstallBinaries
```

add following lines in .vimrc. First 30 lines are copy from the manual. Last 5 lines are additional settings to highlight method names.


```


" Note: Skip initialization for vim-tiny or vim-small.
 if !1 | finish | endif

 if has('vim_starting')
   if &compatible
     set nocompatible               " Be iMproved
   endif

   " Required:
   set runtimepath+=~/.vim/bundle/neobundle.vim/
 endif

 " Required:
 call neobundle#begin(expand('~/.vim/bundle/'))

 " Let NeoBundle manage NeoBundle
 " Required:
 NeoBundleFetch 'Shougo/neobundle.vim'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
 " My Bundles here:
 " Refer to |:NeoBundle-examples|.
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
NeoBundle 'fatih/vim-go'

 call neobundle#end()

 " Required:
 filetype plugin indent on

 " If there are uninstalled bundles found on startup,
 " this will conveniently prompt you to install them.
 NeoBundleCheck


""""""""""""""""""""""""""""""""""""""""""
" GO
""""""""""""""""""""""""""""""""""""""""""
let g:go_highlight_functions = 1
let g:go_highlight_methods = 1
let g:go_highlight_structs = 1                                                                                                                                                                                   
let g:go_highlight_operators = 1
let g:go_highlight_build_constraints = 1

```







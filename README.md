# cgo-gobject
Small example which shows how to create custom GObject as well as how to connect it to Go.

# The main.go file 

```golang
package main

/*
#cgo LDFLAGS: -L/usr/local/Cellar/glib/2.68.0/lib -L/usr/local/opt/gettext/lib -lgobject-2.0 -lglib-2.0 -lintl
#cgo CFLAGS: -I/usr/local/Cellar/libffi/3.3_3/include -I/usr/local/Cellar/glib/2.68.0/include -I/usr/local/Cellar/glib/2.68.0/include/glib-2.0 -I/usr/local/Cellar/glib/2.68.0/lib/glib-2.0/include -I/usr/local/opt/gettext/include -I/usr/local/Cellar/pcre/8.44/include
#include "test-object.h"
*/
import "C"

import (
	"fmt"
)

type TestObject struct {
	obj *C.TestObject
}

func main() {
	tObj := TestObject{C.get_new_test_object()}
	fmt.Printf("[*] tObj @%p\n", &tObj)
	fmt.Printf("[*] Current value is: %d\n", int(C.test_object_get_value(tObj.obj)))
	fmt.Println("[*] Changing value to 15")
	C.test_object_set_value(tObj.obj, C.guint(15))
	fmt.Printf("[*] Value is now: %d\n", int(C.test_object_get_value(tObj.obj)))
}
```

# Running part

```bash
$ ./cgo-gobject 
[*] tObj @0xc00019c008
[*] Current value is: 0
[*] Changing value to 15
[*] Value is now: 15
```

package api

/*
#cgo LDFLAGS: -L/usr/local/Cellar/glib/2.68.0/lib -L/usr/local/opt/gettext/lib -lgobject-2.0 -lglib-2.0 -lintl
#cgo CFLAGS: -I/usr/local/Cellar/libffi/3.3_3/include -I/usr/local/Cellar/glib/2.68.0/include -I/usr/local/Cellar/glib/2.68.0/include/glib-2.0 -I/usr/local/Cellar/glib/2.68.0/lib/glib-2.0/include -I/usr/local/opt/gettext/include -I/usr/local/Cellar/pcre/8.44/include -I..
#include "test-object.h"

extern void customCallback();

static inline void call_callback() {
	customCallback();
}

static inline void connect(TestObject *obj) {
	g_signal_connect (obj, "yo", G_CALLBACK (call_callback), NULL);
}

*/
import "C"
import "fmt"

type TestObject struct {
	obj *C.TestObject
}

func CreateNewObject() *TestObject {
	return &TestObject{C.get_new_test_object()}
}

func (t *TestObject) GetValue() int {
	return int(C.test_object_get_value(t.obj))
}

func (t *TestObject) SetValue(val int) {
	C.test_object_set_value(t.obj, C.guint(val))
}

func (t *TestObject) Emit() {
	C.test_object_emit_yo(t.obj)
}

func (t *TestObject) ConnectSignal() {
	C.connect(t.obj)
}

func defaultCallback() {
	fmt.Println("I am default callback")
}

var defaultCb = defaultCallback

func (t *TestObject) AddCallback(fn func()) {
	if fn != nil {
		defaultCb = fn
	}
}

//export customCallback
func customCallback() {
	defaultCb()
}

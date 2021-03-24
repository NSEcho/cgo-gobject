#include "test-object.h"

struct _TestObject {
    GObject parent_insance;
    guint value;
};

G_DEFINE_TYPE (TestObject, test_object, G_TYPE_OBJECT)

enum {
	YO,
	LAST_SIGNAL,
};

static guint signals [LAST_SIGNAL];

static void
test_object_class_init (TestObjectClass *klass)
{
	signals[YO] =
		g_signal_new("yo",
			    G_TYPE_FROM_CLASS (klass),
			    G_SIGNAL_RUN_FIRST,
			    0,
			    NULL, NULL, NULL,
			    G_TYPE_NONE, 0);
}

static void
test_object_init (TestObject *self)
{
}

void
test_object_emit_yo (TestObject *self)
{
	g_signal_emit (self, signals[YO], 0);
}

TestObject * get_new_test_object()
{
    return g_object_new (TEST_TYPE_OBJECT, NULL);
}

guint test_object_get_value (TestObject *self)
{
    return self->value;
}

void test_object_set_value (TestObject *self, guint val)
{
    self->value = val;
}

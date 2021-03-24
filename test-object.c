#include "test-object.h"

struct _TestObject {
    GObject parent_insance;
    guint value;
};

G_DEFINE_TYPE (TestObject, test_object, G_TYPE_OBJECT)

static void
test_object_class_init (TestObjectClass *klass)
{
}

static void
test_object_init (TestObject *self)
{
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
#ifndef TEST_OBJECT_H
#define TEST_OBJECT_H

#include <glib-object.h>

G_BEGIN_DECLS

#define TEST_TYPE_OBJECT (test_object_get_type())

G_DECLARE_FINAL_TYPE (TestObject, test_object, TESTO, OBJECT, GObject)

guint test_object_get_value (TestObject *self);
void test_object_set_value (TestObject *self, guint value);
TestObject * get_new_test_object(void);

G_END_DECLS

#endif
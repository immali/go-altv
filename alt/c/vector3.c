#include "vector3.h"

void alt_Vector_float_3_PointLayout_SetX(uintptr_t _instance, float x) {
    ((struct alt_Vector_float_3_PointLayout *)_instance)->x = x;
}

void alt_Vector_float_3_PointLayout_SetY(uintptr_t _instance, float y) {
    ((struct alt_Vector_float_3_PointLayout *)_instance)->y = y;
}

void alt_Vector_float_3_PointLayout_SetZ(uintptr_t _instance, float z) {
    ((struct alt_Vector_float_3_PointLayout *)_instance)->z = z;
}

void alt_Vector_float_3_PointLayout_Set(uintptr_t _instance, float x, float y, float z) {
    struct alt_Vector_float_3_PointLayout* vector = ((struct alt_Vector_float_3_PointLayout *)_instance);
    vector->x = x;
    vector->y = y;
    vector->z = z;
    // ((struct alt_Vector_float_3_PointLayout *)_instance)->z = z;
}
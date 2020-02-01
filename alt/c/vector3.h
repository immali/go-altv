#include <stdint.h>

typedef struct alt_Vector_float_3_PointLayout {
    float x;
    float y;
    float z;
} alt_Vector_float_3_PointLayout;

void alt_Vector_float_3_PointLayout_SetX(uintptr_t _instance, float x);
void alt_Vector_float_3_PointLayout_SetY(uintptr_t _instance, float y);
void alt_Vector_float_3_PointLayout_SetZ(uintptr_t _instance, float z);
void alt_Vector_float_3_PointLayout_Set(uintptr_t _instance, float x, float y, float z);
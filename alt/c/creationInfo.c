#include "creationInfo.h"
#include <stdint.h>

struct alt_String* alt_IResource_CreationInfo_GetType(uintptr_t instance) {
    return &((struct alt_IResource_CreationInfo *)instance)->type;
};
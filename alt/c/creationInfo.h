#include "alt_String.h"

typedef struct alt_IResource_CreationInfo {
    struct alt_String type;
    struct alt_String name;
    struct alt_String main;
    struct alt_IPackage* pkg;
} alt_IResource_CreationInfo;

struct alt_String* alt_IResource_CreationInfo_GetType(uintptr_t instance);

#include <stdio.h>
#include <string.h>
#include "../out/libsample.h"

static int
callback(char *out_char) {
    printf("callback is reached to C!");
    strncpy(out_char, "write buffer in C!", 256);
    return 0;
}

int main(int argc, char** argv) {
    printf("hello from C!\n");

    SetCallback(callback);
    Run();

    return 0;
}



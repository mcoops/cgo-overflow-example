package main

/*
#include <stdlib.h>
#include <string.h>

int test_func(char* str) {
        char buff[5];
        strcpy(buff, str);

        if (strcmp(buff, "abcdef") == 0) {
                return 1;
        } else {
                return 0;
        }
}
*/
import "C"
import (
        "fmt"
        "unsafe"
)

func main() {
        testStr := C.CString("aaaaaaaaaaaaaaaaa")
        r := C.test_func(testStr)
        defer C.free(unsafe.Pointer(testStr))

        fmt.Println(r);
}

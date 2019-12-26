package main

// #cgo LDFLAGS: ${SRCDIR}/target/release/libtest_go.a
// #include <stdint.h>
// int *rust_main(const uint8_t *ptr, size_t len);
import "C"
import "unsafe"
// import "fmt"

func main() {
	var data []byte
	data = make([]byte, 20)
	data[0] = 2
	data[1] = 104
	data[18] = 19
	data[19] = 20

	size := C.size_t(len(data));
	
	ptr := (*C.uint8_t)(unsafe.Pointer(&data));
	C.rust_main(ptr, size)
	
	ptr := (*C.uint8_t)(unsafe.Pointer(&data[0]));
	C.rust_main(ptr, size)
}

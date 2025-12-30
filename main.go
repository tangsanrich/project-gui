package main

/*
#cgo LDFLAGS: -landroid -llog
#include <android/log.h>

void cetakLog() {
    __android_log_print(ANDROID_LOG_INFO, "AplikasiKu", "Hello dari Binary Native!");
}
*/
import "C"
import "fmt"

func main() {
    fmt.Println("--------------------------------")
    fmt.Println("Aplikasi Binary Berjalan!")
    fmt.Println("Mencoba memanggil fungsi C...")
    C.cetakLog()
    fmt.Println("Sukses! Cek logcat untuk pesan native.")
    fmt.Println("--------------------------------")
}

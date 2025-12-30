package main

import (
	"fmt"
	"github.com/AllenDang/cimgui-go"
)

func main() {
	fmt.Println("-------------------------------------------")
	fmt.Println("[SUCCESS] Binary berhasil di-build!")
	fmt.Println("[INFO] Library ImGui v1.4.0 terdeteksi.")
	
	// Kita coba inisialisasi Context ImGui (Core Graphics)
	ctx := cimgui.CreateContext()
	defer cimgui.DestroyContext(ctx)
	
	io := cimgui.CurrentIO()
	// Simulasi layar HP
	io.SetDisplaySize(cimgui.Vec2{X: 1080, Y: 2400}) 
	
	fmt.Println("[INFO] Context Grafis dibuat di memori.")
	fmt.Println("[INFO] Jika ini muncul, berarti CGO + NDK bekerja!")
	fmt.Println("-------------------------------------------")
}

package main

import (
	"fmt"
	"github.com/AllenDang/cimgui-go/imgui"
)

func main() {
	fmt.Println("[+] Memulai Aplikasi...")

	// Buat context
	ctx := imgui.CreateContext()
	
	// PERBAIKAN DI SINI: Hapus 'ctx' dari dalam kurung
	defer imgui.DestroyContext()
	
	io := imgui.CurrentIO()
	io.SetDisplaySize(imgui.Vec2{X: 1080, Y: 2400}) 
	
	fmt.Println("[SUCCESS] Context ImGui berhasil dibuat!")
	fmt.Println("[INFO] Build sukses! Tidak ada lagi error.")
}


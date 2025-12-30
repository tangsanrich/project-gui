package main

import (
	"fmt"
	"github.com/AllenDang/cimgui-go/imgui"
)

func main() {
	fmt.Println("[+] Memulai Aplikasi...")

	// PERBAIKAN: Gunakan '_' atau panggil langsung agar tidak dianggap unused variable
	_ = imgui.CreateContext()
	
	defer imgui.DestroyContext()
	
	io := imgui.CurrentIO()
	io.SetDisplaySize(imgui.Vec2{X: 1080, Y: 2400}) 
	
	fmt.Println("[SUCCESS] Context ImGui berhasil dibuat!")
	fmt.Println("[INFO] Build sukses! Tidak ada lagi error.")
}


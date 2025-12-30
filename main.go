package main

import (
	"fmt"
	
	// SEKARANG INI PASTI BISA (Karena induknya sudah ada):
	"github.com/AllenDang/cimgui-go/imgui"
)

func main() {
	fmt.Println("[+] Memulai Aplikasi...")

	// Perhatikan: Kita panggil pakai 'imgui.' bukan 'cimgui.'
	ctx := imgui.CreateContext()
	defer imgui.DestroyContext(ctx)
	
	io := imgui.CurrentIO()
	io.SetDisplaySize(imgui.Vec2{X: 1080, Y: 2400}) 
	
	fmt.Println("[SUCCESS] Context ImGui berhasil dibuat!")
	fmt.Println("[INFO] Build sukses menggunakan package imgui.")
}


package main

import (
	"fmt"
	"github.com/inkyblackness/imgui-go/v4/imgui"
)

func main() {
	fmt.Println("[InkyBlackness] Memulai inisialisasi...")

	// 1. Create Context
	// Library InkyBlackness butuh context untuk menyimpan state
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	fmt.Println("[Sukses] Context ImGui berhasil dibuat!")

	// 2. Test IO (Input/Output)
	io := imgui.CurrentIO()
	
	// 3. Test Font Atlas (Ini akan memicu kode C++ untuk bekerja)
	fmt.Println("[InkyBlackness] Membangun Font Atlas...")
	io.Fonts().AddFontDefault()
	
	// Fungsi ini akan merender font ke dalam texture data (byte array)
	// Jika langkah ini sukses, berarti CGO compiler bekerja sempurna.
	pixels, width, height, bytesPerPixel := io.Fonts().GetTextureDataAsAlpha8()
	
	fmt.Printf("[Sukses] Font berhasil di-load! Ukuran Texture: %dx%d (%d bytes/pixel)\n", 
		width, height, bytesPerPixel)
	fmt.Printf("[Info] Total data pixels: %d bytes\n", len(pixels))

	fmt.Println("------------------------------------------------")
	fmt.Println("KESIMPULAN: Library InkyBlackness v4 BERHASIL di-build & jalan di Termux!")
	fmt.Println("------------------------------------------------")
}

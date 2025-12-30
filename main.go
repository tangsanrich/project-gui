package main

import (
	"fmt"
	"github.com/rivo/tview"
)

// Variabel status fitur
var (
	wallhackOn = false
	aimbotOn   = false
	statusText = "Status: Idle - Siap digunakan"
)

func main() {
	app := tview.NewApplication()

	// --- Membuat Komponen GUI ---

	// 1. Header Judul
	titleBox := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("[yellow]ULTIMATE TERMUX MOD MENU[white]\n[::b]Versi ELF Binary - No X11 Requitred").
		SetDynamicColors(true)

	// 2. Area Status (Log)
	logView := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText("[red]" + statusText)

	// 3. Formulir Checkbox dan Tombol
	form := tview.NewForm()
	form.SetBorder(true).SetTitle(" Konfigurasi Cheat ").SetTitleAlign(tview.AlignLeft)
	
	// Tambah Checkbox
	form.AddCheckbox("Aktifkan Wallhack ESP", wallhackOn, func(checked bool) {
		wallhackOn = checked
		updateStatus(logView, "Wallhack diubah")
	})
	
	form.AddCheckbox("Aktifkan Aimbot Head", aimbotOn, func(checked bool) {
		aimbotOn = checked
		updateStatus(logView, "Aimbot diubah")
	})

	// Tambah Tombol Aksi
	form.AddButton("INJECT SEKARANG", func() {
		// Di sini nanti logika inject memori
		msg := fmt.Sprintf("[green]Injecting... WH:%v | AIM:%v[white]", wallhackOn, aimbotOn)
		logView.SetText(msg)
	})

	form.AddButton("Keluar (Exit)", func() {
		app.Stop()
	})

	// --- Menyusun Layout ---
	// Menggunakan Flexbox (Atas ke Bawah)
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(titleBox, 3, 1, false).  // Judul tinggi 3 baris
		AddItem(form, 0, 3, true).       // Form mengambil sisa ruang terbanyak, fokus di sini
		AddItem(logView, 3, 1, false)    // Log status tinggi 3 baris

	// --- Jalankan Aplikasi ---
	// EnableMouse(true) agar bisa diklik pakai jari di Termux
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

// Fungsi helper untuk update teks status
func updateStatus(view *tview.TextView, msg string) {
	whStatus := "[red]OFF"
	if wallhackOn { whStatus = "[green]ON" }
	
	aimStatus := "[red]OFF"
	if aimbotOn { aimStatus = "[green]ON" }

	fullMsg := fmt.Sprintf("%s | Status: WH=%s | Aim=%s", msg, whStatus, aimStatus)
	view.SetText(fullMsg)
}

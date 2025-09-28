package main

import (
	"fmt"
	"time"
)

// Struct untuk barang dalam struk belanja
type Item struct {
	Name     string
	Price    float64
	Quantity int
}

// Struct untuk struk belanja
type Receipt struct {
	StoreInfo   string
	Date        time.Time
	Items       []Item
	TotalAmount float64
}

// Method untuk menghitung total harga semua item
func (r *Receipt) CalculateTotal() {
	var total float64
	for _, item := range r.Items {
		total += item.Price * float64(item.Quantity)
	}
	r.TotalAmount = total
}

// Method untuk mencetak struk belanja
func (r Receipt) PrintReceipt() {
	fmt.Println("==========================================")
	fmt.Println(r.StoreInfo)
	fmt.Println("Tanggal:", r.Date.Format("02-01-2006 15:04"))
	fmt.Println("==========================================")
	fmt.Printf("%-15s %-10s %-8s %-10s\n", "Item", "Harga", "Jumlah", "Total")
	fmt.Println("------------------------------------------")

	for _, item := range r.Items {
		itemTotal := item.Price * float64(item.Quantity)
		fmt.Printf("%-15s Rp%-9.2f %-8d Rp%-9.2f\n", item.Name, item.Price, item.Quantity, itemTotal)
	}

	fmt.Println("==========================================")
	fmt.Printf("%-35s Rp%-9.2f\n", "Total Belanja:", r.TotalAmount)
	fmt.Println("==========================================")
	fmt.Println("Terima kasih telah berbelanja!")
}

func main() {
	receipt := Receipt{
		StoreInfo: "Toko Sembako Makmur\nJl. Raya No. 123, Jakarta",
		Date:      time.Now(),
		Items: []Item{
			{Name: "Beras", Price: 12000, Quantity: 5},
			{Name: "Gula", Price: 15000, Quantity: 2},
			{Name: "Minyak", Price: 20000, Quantity: 1},
			{Name: "Telur", Price: 2000, Quantity: 10},
		},
	}

	receipt.CalculateTotal()
	receipt.PrintReceipt()
}

package main

import "time"

type User struct {
	ID   string
	Nama string
}

type Assessment struct {
	IDAssessment string
	IDUser       string
	Tanggal      time.Time
	Jawaban      []int
	SkorTotal    int
}

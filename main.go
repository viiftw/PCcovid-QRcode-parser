package main

import (
	"fmt"
	"os"
	"strings"
)

// User define user's info
// Format: SoGiayTo|HoTen|NgaySinh|LoaiDinhDanh|QID|EXT
// Ví dụ: 001093029317|NGUYEN VAN A|1987-01-18|0|162979388586171100|<1<00<0123232422
type User struct {
	SoGiayTo     string `json:"so_giay_to"`
	HoTen        string `json:"ho_ten"`
	NgaySinh     string `json:"ngay_sinh"`
	LoaiDinhDanh string `json:"loai_dinh_danh"`
	QID          string `json:"qid"`
	EXT          string `json:"ext"`
}

func main() {
	src := "001093029317|NGUYEN VAN A|1987-01-18|0|162979388586171100|<1<00<0123232422"
	fmt.Println("From:  ", src)

	s := strings.Split(src, "|")
	if len(s) != 6 {
		fmt.Println("invalid input")
		os.Exit(1)
	}
	fmt.Println("To:    ", s)

	// initialize a struct having an array or slice field
	u := User{}
	u.SoGiayTo = s[0]
	u.HoTen = s[1]
	u.NgaySinh = s[2]
	u.LoaiDinhDanh = s[3]
	u.QID = s[4]
	u.EXT = s[5]

	// print struct with key
	fmt.Printf("Result: %+v \n", u)
}

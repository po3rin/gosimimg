package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"

	"github.com/po3rin/gosimimg"
)

func main() {
	f1, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	img1, err := jpeg.Decode(f1)
	if err != nil {
		log.Fatal(err)
	}
	img2, err := jpeg.Decode(f2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success to load images")

	src1 := gosimimg.GetImage(img1, 8, 8)
	src2 := gosimimg.GetImage(img2, 8, 8)

	fmt.Println("success to resize & gray images")

	fmt.Println(src1.Bounds())
	fmt.Println(src2.Bounds())

	if !gosimimg.IsSimilar(src1, src2) {
		fmt.Println("not simmilar !!")
		return
	}
	fmt.Println("simmilar !!")
}

# gosimimg

<img src="https://img.shields.io/badge/go-v1.12-blue.svg"/>

-----

## What is

### Is a similar image ? yes !!

<img src="./testdata/similar.png" width="460px">

this package determine if it is a similar image using average hash. The average of the luminance values ​​is calculated, and the 64-bit hash value is calculated as “1” for greater than the average and “0” for less than the average.

## Quick Start as CLI

```
go get -u github.com/po3rin/cmd/gosimimg
gosimimg testdata/sim1_1.jpg testdata/sim1_2.jpg
simmilar !!
```

## As Code

### Simple Usage

```go
func main() {

	// prepare image.Image ...

	// inits config.
	// defaults:
	// s := &Similar{
	// 	Threshold:        10,
	// 	CompressedWidth:  8,
	// 	CompressedHeight: 8,
	// }
	s := gosimimg.NewSimilar()

	// Do Similar image search.
	if s.IsSimilar(img1, img2) {
		fmt.Println("not simmilar !!")
		return
	}
}
```

### With Options

```go
func main() {
	// prepare image.Image ...

	s := gosimimg.NewSimilar(
		gosimimg.SetThreshold(10),
		gosimimg.SetCompressedWidth(16),
		gosimimg.SetCompressedHeight(16),
	)

	// Do Similar image search.
	if s.IsSimilar(img1, img2) {
		fmt.Println("not simmilar !!")
		return
	}
}
```

## reference ( japanese )

[類似画像検索について簡単にまとめてみた](https://qiita.com/hurutoriya/items/88a16d36bafa8d6360e2)

[Go言語を使って類似画像を検索する](https://medium.com/eureka-engineering/go%E8%A8%80%E8%AA%9E%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E9%A1%9E%E4%BC%BC%E7%94%BB%E5%83%8F%E3%82%92%E6%A4%9C%E7%B4%A2%E3%81%99%E3%82%8B-ccb2a0752d04)

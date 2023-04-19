package main

import "fmt"

const ImageWidth = 256;
const ImageHeight = 256;

type point struct {
    x float32;
    y float32;
    z float32;
}

func main() {
	fmt.Println("P3\n", ImageWidth, " ", ImageHeight, "\n255\n")
    for i := 0; i < ImageHeight; i++ {
        for j := 0; j < ImageWidth; j++ {
            r := float32(i) / (ImageHeight-1);
            g := float32(j) / (ImageWidth-1);
            b := 0.25;

            ir := int(255.999 * r);
            ig := int(255.999 * g);
            ib := int(255.999 * b);
            fmt.Print(ir, " ", ig, " ", ib, " ", "\n");
        }
    }
}

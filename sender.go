package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	header = "01111110"
	terminator = "01111110"
)

func calculateHammingBits(dataLength int) int {
	r := 0
	for (1 << r) < (dataLength + r + 1){
		r++
	}
	return r
}

func isPowerOfTwo(n int) bool {
	return n != 0 && (n&(n-1)) == 0
}

func addHammingCode(data string) string {
	dataBits := make([]int, len(data))
	for i, bit := range data {
		if bit == '1'{
			dataBits[i] = 1
		} else {
			dataBits[i] = 0
		}
	}

	r := calculateHammingBits(len(dataBits))

	encodedLength := len(dataBits) + r
	encodedData := make([]int, encodedLength)
	for i := range encodedData {
		encodedData[i] = -1
	}

	j := 0
	for i := 1; i <= encodedLength; i++{
		if !isPowerOfTwo(i){
			encodedData[i-1]=dataBits[j]
			j++
		}
	}

	for i:= 0; i < r; i++{
		parityPosition := (1 << i) - 1
		parityBit := 0

		for j := parityPosition; j < encodedLength; j++{
			if encodedData[j] != -1{
				if ((j+1) & (1 << i)) !=0{
					parityBit ^= encodedData[j]
				}
			}
		}

		encodedData[parityPosition] = parityBit
	}

	var encodedString strings.Builder
	for _, bit := range encodedData{
		encodedString.WriteString(fmt.Sprintf("%d", bit))
	}
	return encodedString.String()

}

func bitStuffing(data string) string {
	var result strings.Builder
	countOnes := 0

	for _, bit:= range data {
		result.WriteRune(bit)

		if bit == '1'{
			countOnes++
			if countOnes == 5{
				result.WriteRune('0')
				countOnes = 0
			}
		} else {
			countOnes = 0
		}
	}

	return result.String()
}

func createFrame(payload string) string {
	encodedPayload := addHammingCode(payload)
	stuffedPayload := bitStuffing(encodedPayload)
	frame := header + stuffedPayload + terminator

	return frame
}

func main(){
	// fmt.Println("Starting the program...")

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Uso: ./sender \"sequência de bits\"\n")
		os.Exit(1)
	}

	payload := os.Args[1]
	// fmt.Println("Payload:", payload)

	for _, bit := range payload {
		if bit != '0' && bit != '1' {
			fmt.Fprintf(os.Stderr, "Erro: O payload deve conter apenas 0s e 1s\n")
			os.Exit(1)
		}
	}

	frame := createFrame(payload)

	fmt.Fprintf(os.Stderr, "Payload original: %s\n", payload)
	fmt.Fprintf(os.Stderr, "Payload com Hamming: %s\n", addHammingCode(payload))
	fmt.Fprintf(os.Stderr, "Frame completo: %s\n", frame)

	fmt.Print(frame)
}
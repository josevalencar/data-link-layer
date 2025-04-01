package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	header = "01111110"
	terminator = "01111110"
)

func isPowerOfTwo(n int) bool {
	return n != 0 && (n&(n-1)) == 0
}

func removeBitStuffing(data string) string {
	var result strings.Builder
	countOnes := 0
	i := 0

	for i < len(data){
		bit := data[i]
		result.WriteByte(bit)

		if bit == '1'{
			countOnes++
			if countOnes == 5 && i+1 < len(data) && data[i+1] == '0'{
				i++
				countOnes = 0
			}
		} else {
			countOnes = 0
		}
		i++
	}

	return result.String()
}

func checkAndCorrectHamming(encodedData string) string {
	encodedBits := make([]int, len(encodedData))
	for i, bit := range encodedData {
		if bit == '1'{
			encodedBits[i] = 1
		} else {
			encodedBits[i] = 0
		}
	}
	encodedLength  := len(encodedBits)

	r := 0
	for (1 << r) < encodedLength {
		r++
	}

	errorPosition := 0
	for i := 0; i < r; i++{
		// parityPosition := (1 << i) - 1
		parityCheck := 0

		for j := 0; j < encodedLength; j++ {
			if ((j+1) & (1<<i)) != 0{
				parityCheck ^= encodedBits[j]
			}
		}

		if parityCheck != 0{
			errorPosition += (1 << i)
		}
	}

	if errorPosition != 0 && errorPosition <= encodedLength {
		fmt.Fprintf(os.Stderr, "Erro detectado na posição %d\n", errorPosition)
		encodedBits[errorPosition-1] ^= 1
	}

	var dataBits []int
	for i := 1; i <= encodedLength; i++{
		if !isPowerOfTwo(i){
			dataBits = append(dataBits, encodedBits[i-1])
		}
	}

	var decodedData strings.Builder
	for _, bit := range dataBits {
		decodedData.WriteString(fmt.Sprintf("%d", bit))
	}
	return decodedData.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frame := scanner.Text()

	startIdx := strings.Index(frame, header)
	if startIdx == -1{
		fmt.Fprintf(os.Stderr, "Erro: cabeçalho não encontrado\n")
		os.Exit(1)
	}
	startIdx += len(header)

	endIdx := strings.LastIndex(frame, terminator)
	if endIdx == -1 || endIdx <= startIdx {
		fmt.Fprintf(os.Stderr, "Erro: Terminador não encontrado após cabeçalho\n")
		os.Exit(1)
	}

	stuffedPayload := frame[startIdx:endIdx]

	fmt.Fprintf(os.Stderr, "Frame recebido: %s\n", frame)
	fmt.Fprintf(os.Stderr, "Payload com bit stuffing: %s\n", stuffedPayload)

	encodedPayload := removeBitStuffing(stuffedPayload)
	fmt.Fprintf(os.Stderr, "Payload com Hamming: %s\n", encodedPayload)

	decodedPayload := checkAndCorrectHamming(encodedPayload)
	fmt.Fprintf(os.Stderr, "Payload decodificado: %s\n", decodedPayload)

	fmt.Print(decodedPayload)
}
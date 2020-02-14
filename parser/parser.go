package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Parse extracts the hidden text from a series of blocks of text where:
// 	- first line is the index
// 	- second line is coordinates
// 	- third line until a blank line is the block of text
func Parse(content []byte) string {
	scanner := bufio.NewScanner(bytes.NewReader(content))
	blocks := [][]string{}

	letterBlock := []string{}
	for i := 0; scanner.Scan(); {
		line := scanner.Text()
		if line != "" {
			letterBlock = append(letterBlock, line)
		} else {
			blocks = append(blocks, letterBlock)
			letterBlock = []string{}
			i++
		}
	}

	mapper := map[int]string{}
	for _, block := range blocks {
		index, err := ParseIndex(block[0])
		if err != nil {
			log.Println("invalid index string", block[0], err)
			continue
		}

		row, col, err := ParseCoordinates(block[1])
		if err != nil {
			log.Println("invalid coordinates", block[1], err)
			continue
		}

		letters, err := ParseBlock(block[2:])
		if err != nil {
			log.Println("invalid block", block[2:], err)
			continue
		}

		mapper[index] = letters[row][col]
	}

	hiddenText := ""
	for i := 0; i < len(mapper); i++ {
		hiddenText += mapper[i]
	}
	return hiddenText
}

// ParseIndex converts the index to an int.
func ParseIndex(line string) (int, error) {
	s := strings.TrimSpace(line)
	return strconv.Atoi(s)
}

// ParseCoordinates converts the given line to row, col.
func ParseCoordinates(line string) (row, col int, err error) {
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		err = fmt.Errorf("invalid coordinates %s", line)
		return
	}

	coordArray := strings.Split(line[1:len(line)-1], ",")
	if len(coordArray) != 2 {
		err = fmt.Errorf("invalid coordinate format, expecting [row, col] but got %s", line)
		return
	}

	if col, err = strconv.Atoi(strings.TrimSpace(coordArray[0])); err != nil {
		return
	}

	if row, err = strconv.Atoi(strings.TrimSpace(coordArray[1])); err != nil {
		return
	}
	return
}

// ParseBlock constracts a 2-D matrix from the block of letters.
func ParseBlock(letters []string) ([][]string, error) {
	matrix := make([][]string, len(letters))
	for i, j := len(letters)-1, 0; i >= 0; i, j = i-1, j+1 {
		line := letters[i]
		matrix[j] = make([]string, len(line))
		for k := 0; k < len(line); k++ {
			matrix[j][k] = string(line[k])
		}
	}
	return matrix, nil
}

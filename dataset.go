package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Piece Define Piece
type Piece struct {
	Points []Point
}

type Obejct struct {
	Hight, Weight int
	PiecesNum     int
	Pieces        []*Piece
}

// GetObjects returns the object from dataset file
func GetObjects(filename string) ([]*Obejct, error) {
	// readfile, Read the file successfully, and return the file content as an array []byte
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// 将读取的二维数组数据存储到lines中
	lines, err := contentStr2Arry(string(bytes))
	if err != nil {
		return nil, err
	}

	var objs []*Obejct
	readLineNum := 2
	// multiple objects
	fmt.Println(lines[0])
	fmt.Println("---------------A---------------------")
	if len(lines[0]) != 1 {
		//the first line has one parameter
		for i := 1; i <= lines[0][0]; i++ {
			objs = append(objs, getObject(lines, lines[0][i], readLineNum, readLineNum+lines[0][i]))
			readLineNum += lines[0][i]
		}
	} else {
		// the first line of *.txt has more than one parameter
		objs = append(objs, getObject(lines, lines[0][0], readLineNum, readLineNum+lines[0][0]))
	}
	return objs, err
}

// getObject
func getObject(lines [][]int, piecesNum, start, end int) (obj *Obejct) {

	obj = &Obejct{}
	// the size of space need to place
	obj.Hight, obj.Weight = lines[1][0], lines[1][1]
	// the number of polygons in the placement interval
	obj.PiecesNum = piecesNum
	for _, line := range lines[start:end] {

		tempPiece := Piece{}
		tempPiece.PointNum = line[0]
		for i := 1; i < len(line); {
			tempPoint := Point{
				X: line[i],
				Y: line[i+1],
			}
			tempPiece.Points = append(tempPiece.Points, tempPoint)
			i += 2
		}
		obj.Pieces = append(obj.Pieces, &tempPiece)
	}
	return
}

// contentStr2Arry txt文本字符串转为二维数组
func contentStr2Arry(conStr string) (lines [][]int, err error) {

	//split used to slice the string with the specified separator and return the sliced string
	lineStrs := strings.Split(strings.TrimSpace(conStr), "\n")

	for _, lineStr := range lineStrs {
		lineStr = strings.TrimSpace(lineStr)
		strs := strings.Split(lineStr, " ")
		line := make([]int, len(strs))
		for i, str := range strs {
			line[i], err = strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
		}
		lines = append(lines, line)
	}
	return
}

func main() {

	//get the information of polygons
	// objs, err := GetObjects("dataset/OpTA001C5.txt")
	objs, err := GetObjects(".jakobsdata.txt")
	if err != nil {
		log.Println(err)
	}
	PrintObjs(objs)
}

// PrintObjs p
func PrintObjs(objs []*Obejct) {
	for _, obj := range objs {
		log.Println("--------------\n--------------\n")
		log.Println("Hight:", obj.Hight, "Weight:", obj.Weight, "PiecesNum:", obj.PiecesNum)
		for _, piece := range obj.Pieces {
			log.Println("\nPointNum:", piece.PointNum)
			for _, point := range piece.Points {
				log.Println("point:", point)
			}
		}
	}
}

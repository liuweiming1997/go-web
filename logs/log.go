package logs

import (
	"bufio"
	"log"
	"os"
	"time"
)

var FileWriter *bufio.Writer

func init() {
	const path = "../logs/log.txt" // fource in here
	_, err := os.Open(path)
	var file *os.File
	if err == nil {
		// append
		file, err = os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
		CheckErr(err)
	} else {
		file, err = os.Create(path)
		CheckErr(err)
	}
	// defer file.Close() can not do it cause you can not write to the file
	file.Sync()
	FileWriter = bufio.NewWriter(file)
}

func LogToFile(str string) {
	// fmt.Println(str)
	_, err := FileWriter.WriteString(time.Now().String()[0:19] + " : " + str)
	if len(str) > 0 && str[len(str)-1] != '\n' || len(str) == 0 {
		FileWriter.WriteString("\n")
	}
	CheckErr(err)
	FileWriter.Flush()
}

func CheckErr(err error) {
	if err != nil {
		LogToFile(err.Error())
		log.Fatal(err.Error())
	}
}

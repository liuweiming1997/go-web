package logs

import (
	"bufio"
	"log"
	"os"
	"time"
)

var FileWriter *bufio.Writer

func init() {
	_, err := os.Open("log.txt")
	var file *os.File
	if err == nil {
		// append
		file, err = os.OpenFile("log.txt", os.O_WRONLY|os.O_APPEND, 0666)
		CheckErr(err)
	} else {
		file, err = os.Create("log.txt")
		CheckErr(err)
	}
	// defer file.Close() can not do it cause you can not write to the file
	file.Sync()
	FileWriter = bufio.NewWriter(file)
}

func LogToFile(str string) {
	_, err := FileWriter.WriteString(time.Now().String()[0:19] + " : " + str)
	if len(str) > 0 && str[len(str)-1] != '\n' {
		FileWriter.WriteString("\n")
	}
	CheckErr(err)
	FileWriter.Flush()
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

package main

import (
	"bytes"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
	"syscall/js"
	"time"
)

//var (
//	document = js.Global().Get("document")
//	numEle   = document.Call("getElementById", "num")
//	ansEle   = document.Call("getElementById", "ans")
//	btnEle   = js.Global().Get("btn")
//)

func fib(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func fibFunc(this js.Value, args []js.Value) interface{} {
	callback := args[len(args)-1]

	go func() {
		time.Sleep(3 * time.Second)
		v := fib(args[0].Int())
		callback.Invoke(v)
	}()

	js.Global().Get("ans").Set("innerHTML", "generating ...")
	return nil
}

func main() {
	done := make(chan int, 0)

	js.Global().Set("fibFunc", js.FuncOf(fibFunc))
	js.Global().Set("readFileFunc", js.FuncOf(readFileFunc))

	<-done
}

func readFileFunc(this js.Value, args []js.Value) interface{} {
	pth := args[0].String()
	callback := args[len(args)-1]

	go func() {
		time.Sleep(3 * time.Second)
		v := readFile(pth)

		callback.Invoke(v)
	}()

	msg := fmt.Sprintf("sleep 1s then read file %s ...", pth)

	js.Global().Get("content").Set("innerHTML", msg)

	return nil
}

func readFile(pth string) (ret string) {
	content, _ := ReadResData(pth)

	excel, err := excelize.OpenReader(bytes.NewReader(content))

	if err != nil {
		log.Printf("fail to read file %s", pth)
		return
	}

	for _, sheet := range excel.GetSheetList() {
		rows, _ := excel.GetRows(sheet)

		var rowArr []string

		for _, row := range rows {
			var colArr []string

			for _, col := range row {
				val := strings.TrimSpace(col)
				colArr = append(colArr, val)
			}

			rowArr = append(rowArr, strings.Join(colArr, ", "))
		}

		ret += strings.Join(rowArr, "<br />")
	}

	return
}

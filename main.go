package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"syscall/js"
)

//var (
//	document = js.Global().Get("document")
//	numEle   = document.Call("getElementById", "num")
//	ansEle   = document.Call("getElementById", "ans")
//	btnEle   = js.Global().Get("btn")
//)

//func fib(i int) int {
//	if i == 0 || i == 1 {
//		return 1
//	}
//	return fib(i-1) + fib(i-2)
//}
//
//func fibFunc(this js.Value, args []js.Value) interface{} {
//	callback := args[len(args)-1]
//
//	go func() {
//		//time.Sleep(3 * time.Second)
//		v := fib(args[0].Int())
//		callback.Invoke(v)
//	}()
//
//	js.Global().Get("ans").Set("innerHTML", "generating ...")
//	return nil
//}

func main() {
	done := make(chan int, 0)

	//js.Global().Set("fibFunc", js.FuncOf(fibFunc))
	//js.Global().Set("readFileFunc", js.FuncOf(readFileFunc))
	js.Global().Set("selectDataFunc", js.FuncOf(selectDataFunc))

	<-done
}

func selectDataFunc(this js.Value, args []js.Value) interface{} {
	//dbName := args[0].String()
	tableName := args[1].String()
	//sql := args[2].String()

	callback := args[len(args)-1]

	go func() {
		//time.Sleep(1 * time.Second)
		//db = createTable(dbName, tableName)
		//v := selectData(db, sql)

		ret, _ := Get("https://baidu.com")

		v := string(ret)
		callback.Invoke(v)
	}()

	msg := fmt.Sprintf("select data from table %s ...", tableName)

	js.Global().Get("content").Set("innerHTML", msg)

	return nil
}

//func createTable(dbName, tableName string) (db *sql.DB) {
//	db, err := sql.Open("sqlite3_js", dbName)
//	if err != nil {
//		log.Fatalf("cannot open test.db: %s", err)
//	}
//
//	schema := fmt.Sprintf("create table %s(id INTEGER PRIMARY KEY, name string)", tableName)
//	_, err = db.Exec(schema)
//	if err != nil {
//		log.Fatalf("cannot create schema: %s", err)
//	}
//
//	_, err = db.Exec(fmt.Sprintf("insert into %s values(1, 'Aaron')", tableName))
//	if err != nil {
//		log.Fatalf("failed to insert: %s", err)
//	}
//
//	_, err = db.Exec(fmt.Sprintf("UPDATE %s SET name='Aaron Chen' WHERE 1=1", tableName))
//	if err != nil {
//		log.Fatalf("failed to update: %s", err)
//	}
//
//	return
//}
//
//func selectData(db *sql.DB, sql string) (ret string) {
//	rows, err := db.Query(sql)
//	if err != nil {
//		log.Fatalf("failed to query: %s", err)
//	}
//
//	var arr []string
//	for rows.Next() {
//		var got string
//		if err := rows.Scan(&got); err != nil {
//			log.Fatalf("failed to scan row: %s", err)
//		}
//		arr = append(arr, got)
//	}
//
//	ret = strings.Join(arr, "<br />")
//
//	return
//}

//
//func readFileFunc(this js.Value, args []js.Value) interface{} {
//	pth := args[0].String()
//	callback := args[len(args)-1]
//
//	go func() {
//		//time.Sleep(1 * time.Second)
//		v := readFile(pth)
//
//		callback.Invoke(v)
//	}()
//
//	msg := fmt.Sprintf("reading file %s ...", pth)
//
//	js.Global().Get("content").Set("innerHTML", msg)
//
//	return nil
//}
//
//func readFile(pth string) (ret string) {
//	content, _ := ReadResData(pth)
//
//	excel, err := excelize.OpenReader(bytes.NewReader(content))
//
//	if err != nil {
//		log.Printf("fail to read file %s", pth)
//		return
//	}
//
//	for _, sheet := range excel.GetSheetList() {
//		rows, _ := excel.GetRows(sheet)
//
//		var rowArr []string
//
//		for _, row := range rows {
//			var colArr []string
//
//			for _, col := range row {
//				val := strings.TrimSpace(col)
//				colArr = append(colArr, val)
//			}
//
//			rowArr = append(rowArr, strings.Join(colArr, ", "))
//		}
//
//		ret += strings.Join(rowArr, "<br />")
//	}
//
//	return
//}

func Get(url string) (ret []byte, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()

	ret, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return
	}

	log.Print(string(ret))

	return
}

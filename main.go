package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

// https://gosamples.dev/read-csv/
func main() {
    // open file
    f, err := os.Open("./dataset/test.csv")
    if err != nil {
        log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(f)

	csvReader.Read()
    for {
        rec, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

		tracedString := strings.ReplaceAll(rec[2],"'",`"`)
		txDataString := strings.ReplaceAll(rec[1],"'",`"`)
	
		txHash := rec[0]
		txTrace := TxTrace{}
		txData := TxData{}
		bt := []byte(tracedString)
		bd := []byte(txDataString)

		// https://surajsharma.net/blog/golang-json-to-struct
		trace_err := json.Unmarshal(bt, &txTrace)
		if trace_err != nil {
			// fmt.Println("Unable to convert tx trace JSON string to a struct; hash", txHash, trace_err)
			continue
		}

		data_err := json.Unmarshal(bd, &txData)
		if data_err != nil {
			// fmt.Println("Unable to convert tx data JSON string to a struct; hash", txHash, data_err)
			continue
		}

		CheckBridge(txHash,txData,txTrace)
    }
}
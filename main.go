package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/gregpr07/bridge-decoder/bridges"
	"github.com/gregpr07/bridge-decoder/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=localhost user=postgres password=postgres dbname=berkeley port=5432 sslmode=disable"

func checkBridges(db *gorm.DB, txHash string,txData utils.TxData,txTrace utils.TxTrace) {
	err := bridges.CheckPolygonPOSDeposit(db, txHash,txData,txTrace)
	if err != nil {
		fmt.Println(err)
	}
}

// https://gosamples.dev/read-csv/
func main() {
    // open file
    f, err := os.Open("./dataset/train.csv")
    if err != nil {
        log.Fatal(err)
    }

	// open database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&bridges.PolygonPOSBridgeTx{})

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
		txTrace := utils.TxTrace{}
		txData := utils.TxData{}
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


		checkBridges(db, txHash,txData,txTrace)
    }
}




import (
	//"PP/controllers"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"log"
	//"net/http"
)

//what it will do

func main() {

	data, errread := os.Open("test.csv")
	writer, errwrite := os.Create("dataMAX.js")

	if errwrite != nil {
		fmt.Println("Unable to create file:", errwrite)
		os.Exit(1)
	}
	if errread != nil {
		log.Fatalf("Error when opening file: %s", errread)
	}
	recordsnodes := make([][]string, 0)
	reader := csv.NewReader(data)
	reader.FieldsPerRecord = 17
	writer.WriteString("var nodes = [\n")
	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Print(e)
			break
		}
		recordsnodes = append(recordsnodes, record)
	}
	for i := 0; i < len(recordsnodes); i++ {
		writer.WriteString(`{id: ` + strconv.Itoa(i+1) + `, label: "` + recordsnodes[i][15] + `" , group: "` + recordsnodes[i][13] + `" },` + "\n")

	}
	writer.WriteString("];\n")
	writer.WriteString("var edges = [\n")
	for i := 0; i < len(recordsnodes)-1; i++ {
		for j := i + 1; j < len(recordsnodes); j++ {
			if recordsnodes[i][15] == recordsnodes[j][15] {

				writer.WriteString("{from: " + strconv.Itoa(i) + ", to: " + strconv.Itoa(j) + "},\n")

			}
		}
	}
	writer.WriteString("];\n")
	data.Close()

}

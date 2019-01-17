// 批量获取面单
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stevzhang01/logistics/label"
	"os"
	"strings"
)

var buf bytes.Buffer
var title = true
var order label.Order
var sku label.Sku

func main() {
	buf.WriteString("序号,订单号,追踪码,面单,请求信息\n")
	input := bufio.NewScanner(os.Stdin)
	i := 0
	for input.Scan() {
		line := input.Text()
		if line == "" {
			continue
		}
		if title {
			title = false
			continue
		}
		params := strings.Split(line, ",")
		err := label.ParseOrder(params, &order)
		if err != nil {
			fmt.Fprintf(os.Stderr, "label.ParseOrder %s: %v\n", line, err)
			os.Exit(1)
		}
		data, err := json.Marshal(order)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Json marshal: %v\n", err)
			os.Exit(1)
		}
		strdata := fmt.Sprintf("%s", data)
		rst, err := label.Request(strdata)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Label request: %v\n", err)
			os.Exit(1)
		}
		item := fmt.Sprintf("%d,%s,%s,%s%s,%s\n", i, params[0], rst.Data.TrackNo, label.LabelPrefix, rst.Data.BillData, rst.Info)
		buf.WriteString(item)
		i++
	}
	fmt.Println(buf.String())
}

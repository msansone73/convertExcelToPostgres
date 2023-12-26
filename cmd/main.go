package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/msansone73/convertExcelToPostgres/internal/model"
)

func main() {
	log.Println("Inicio do programa...")
    f, err := excelize.OpenFile("/Users/msansone/Documents/gitBackup/Go/convertExcelToPostgres/Investimento2023.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
	log.Println("planilha lida....")

    rows := f.GetRows("Transação")
    for _, row := range rows {
		log.Printf("recuperando %s ....", row[1])
		var stock model.Stock
		err:= stock.GetStockByCode(row[1])
		if err != nil {
			continue
		}

		log.Printf("para ticker %s formantando valor %s  ....", row[1], row[5])
		var valor float64
		str := strings.Replace(row[5], ",", ".", -1)
		valor, err = strconv.ParseFloat(str, 32)
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Printf("para ticker %s formantando quantidade %s  ....", row[1], row[7])
		var quantidade float64
		quantidade, err = strconv.ParseFloat(row[7], 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Printf("para ticker %s formantando data %s  ....", row[1], row[4])
		originalDate:= row[4]
		originalDatefloat, err := strconv.ParseFloat(originalDate, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		excelDate := originalDatefloat
		baseDate := time.Date(1899, 12, 31, 0, 0, 0, 0, time.UTC)
		// O Excel considera erroneamente que 1900 é bissexto, então subtrai-se 2 dias se a data for após 28 de fevereiro de 1900
		if excelDate > 60 {
			excelDate -= 2
		}
		// Adiciona o número de dias à data base
		date := baseDate.Add(time.Duration(excelDate*24) * time.Hour)	
		fmt.Println(date.Format("02/01/2006"))

		transacao := model.Transaction{User_id: 1, 
									Stock_id: stock.Id,
									Tipo: "buy",
									Value: float32(valor),
									Quantity: int(quantidade),
									Data_at: date}
		
		fmt.Println(transacao)	
		transacao.Adicionar()					

    }

    fmt.Println("Dados extraídos com sucesso!")
}

package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"proyecto/prueba1/model"
	"time"
)

func GetPurchaseServices(date string, days int64) (model.ConsolidatedPurchase, error)  {
	var rest []byte
	var err error
	var data []model.Purchase
	var consolidated model.ConsolidatedPurchase
	count := 0

	Date, _ := time.Parse("2006-01-02", date)
	finalDate := Date.AddDate(0, 0,int(days)).Format("2006-01-02")

	rest, err = RestPurchase()
	if err != nil {
		return consolidated, err
	}
	data = Parsear(rest)

	for i:= 0; i < len(data); i++ {
		myDate, err := time.Parse("2006-01-02T15:04:05Z07:00", data[i].Date)
		if err != nil {
			panic(err)
		}
		if  myDate.Format("2006-01-02") <= finalDate {
			consolidated.Total = consolidated.Total + data[i].Amount
		}
		if data[i].Tdc == "visa gold" {
			consolidated.PurchaseTDC.Gold =  data[i].Amount
		}
		if data[i].Tdc == "amex" {
			consolidated.PurchaseTDC.Amex =  data[i].Amount
		}
		if data[i].Purchase == false {
			consolidated.NotPurchase =  count + 1
		}
		if data[i].Amount >  consolidated.BuyHigher {
			consolidated.BuyHigher = data[i].Amount
		}
	}
	
	return consolidated, nil

}

func RestPurchase() ([]byte, error) {
	url :="https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/2019-12-01"

	res := &http.Client{}
	req, err := http.NewRequest("GET",url,nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept","application/json")

	resp, errs := res.Do(req)
	if errs != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, er :=  ioutil.ReadAll(resp.Body)
	if er != nil {
		return nil, er
	}

	return response, nil
}

func Parsear(data []byte) []model.Purchase {
	var ResponseDate []model.Purchase

	err := json.Unmarshal(data, &ResponseDate)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	return ResponseDate
}
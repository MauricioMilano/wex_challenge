package fiscaldata

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

type FiscalData struct {
	baseURL string
}

type ExchangeData struct {
	CountryCurrencyDesc string `json:"country_currency_desc"`
	ExchangeRate        string `json:"exchange_rate"`
	RecordDate          string `json:"record_date"`
	Country             string `json:"country"`
}

type FiscalDataResponse struct {
	Data []ExchangeData `json:"data"`
}

func NewFiscalData() FiscalData {
	return FiscalData{
		baseURL: "https://api.fiscaldata.treasury.gov",
	}
}

func (fd *FiscalData) Get(country string, begin_time string, end_time string) ([]ExchangeData, error) {
	resp, err := http.Get(fd.baseURL + "/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?fields=country_currency_desc,exchange_rate,record_date,country&filter=record_date:gte:" + begin_time + ",record_date:lte:" + end_time + ",country:eq:" + country)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data FiscalDataResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data.Data, nil
}

func (fd *FiscalData) GetRate(country string, begin_time string, end_time string) (float32, error) {
	exchanges, err := fd.Get(country, begin_time, end_time)
	if err != nil {
		return 0, err
	}
	if len(exchanges) == 0 {
		return 0, errors.New("exchange doens't exist for this period")
	}
	exchange_choosed := exchanges[0]
	rate, err := strconv.ParseFloat(exchange_choosed.ExchangeRate, 64)
	if err != nil {
		return 0, err
	}
	return float32(rate), nil
}

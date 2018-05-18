package googlefinance

import (
	"context"
	"encoding/csv"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Query query
type Query struct {
	Q  string
	X  string
	I  string
	P  string
	Ts string
	// F  string
}

// Price price
type Price struct {
	Date   time.Time `json:"date"`
	Close  float64   `json:"close"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Open   float64   `json:"open"`
	Volume int64     `json:"volume"`
}

func decodeBody(resp *http.Response, query *Query) (*[]Price, error) {
	defer resp.Body.Close()
	r := csv.NewReader(resp.Body)
	var a, d int64
	var date time.Time
	interval, _ := strconv.ParseInt(query.I, 10, 64)
	prices := []Price{}
	for i := 1; ; i++ {
		row, err := r.Read()
		if err == io.EOF {
			break
		} else if perr, ok := err.(*csv.ParseError); ok && perr.Err == csv.ErrFieldCount {
		} else if err != nil {
			return nil, err
		}

		if i >= 8 {
			if row[0][0:1] == "a" {
				d, _ = strconv.ParseInt(row[0][1:], 10, 64)
				a = d
				date = time.Unix(a, 0)
			} else {

				d, _ = strconv.ParseInt(row[0], 10, 64)
				date = time.Unix(a+(d*interval), 0)
			}
			close, _ := strconv.ParseFloat(row[1], 64)
			high, _ := strconv.ParseFloat(row[2], 64)
			low, _ := strconv.ParseFloat(row[3], 64)
			open, _ := strconv.ParseFloat(row[4], 64)
			volume, _ := strconv.ParseInt(row[5], 10, 64)

			prices = append(prices, Price{
				Date:   date,
				Close:  close,
				High:   high,
				Low:    low,
				Open:   open,
				Volume: volume})
		}
	}
	return &prices, nil
}

// GetPrices get prices
func GetPrices(ctx context.Context, query *Query) (*[]Price, error) {

	u, _ := url.Parse("https://www.google.com/finance/getprices")

	v := url.Values{}

	if query.Q != "" {
		v.Set("q", query.Q)
	}
	if query.X != "" {
		v.Set("x", query.X)
	}
	if query.I != "" {
		v.Set("i", query.I)
	}
	if query.P != "" {
		v.Set("p", query.P)
	}
	// if query.F != "" {
	// 	v.Set("f", query.F)
	// }
	if query.Ts != "" {
		v.Set("ts", query.Ts)
	}

	u.RawQuery = v.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	prices, err := decodeBody(res, query)
	if err != nil {
		return nil, err
	}

	return prices, nil
}

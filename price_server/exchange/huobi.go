package exchange

import (
	"encoding/json"
)

//{"ch":"market.btcusdt.detail.merged","status":"ok","ts":1629098174462,"tick":{"id":271006617773,"version":271006617773,"open":46266.36,"close":47323.59,"low":45477.01,"high":48050.0,"amount":17729.513874233984,"vol":8.266551922141494E8,"count":549488,"bid":[47323.58,1.55321],"ask":[47323.59,0.06133773059609383]}}
type HuobiPriceInfo struct {
	Status string   `json:"status"`
	Tick   TickInfo `json:"tick"`
}

type TickInfo struct {
	Ask []float64 `json:"ask"`
}

func parseHuobiPrice(priceJson string) (float64, error) {
	var huobiPriceInfo HuobiPriceInfo

	err := json.Unmarshal([]byte(priceJson), &huobiPriceInfo)
	if err != nil {
		return 0, err
	}

	return huobiPriceInfo.Tick.Ask[0], nil
}

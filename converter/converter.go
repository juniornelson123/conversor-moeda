package converter

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const Dolar = "USD"
const Euro = "EUR"
const Libra = "GBP"
const Real = "BRL"

//Coin is struct than contain informations about value e kind
type Coin struct {
	Kind         string
	Value        float64
	ValueConvert float64
}

func (c Coin) ConvertCoin(kindFinal string, db *sql.DB) (float64, error) {
	resp, err := http.Get("https://api.fixer.io/latest?base=" + libKind(c.Kind))

	if err != nil {

		return 0.0, fmt.Errorf("Dados invalidos: %s\n", err)

	} else {

		body, errConv := ioutil.ReadAll(resp.Body)

		if errConv != nil {

			return 0.0, fmt.Errorf("Dados invalidos: %s\n", errConv)

		} else {

			r, errJson := toJson(body)
			if err != nil {

				return 0.0, fmt.Errorf("Dados invalidos: %s\n", errJson)

			} else {
				c.ValueConvert = convertCoinFinal(c.Value, r[libKind(kindFinal)].(float64))

				c.insertCoin(db)
				return convertCoinFinal(c.Value, r[libKind(kindFinal)].(float64)), nil

			}
		}
	}
}

func (c Coin) insertCoin(db *sql.DB) (string, error) {
	tx, err := db.Begin()

	if err != nil {
		return "Erro ao salvar conversão\n", fmt.Errorf("Erro ao salvar convsão %s\n", err)
	}

	stmt, err := tx.Prepare("INSERT INTO coins( kind,value,value_convert) values (?,?,?)")
	if err != nil {
		return "Erro ao salvar conversão\n", fmt.Errorf("Erro ao salvar convers %s\n", err)
	}

	_, err = stmt.Exec(c.Kind, c.Value, c.ValueConvert)

	if err != nil {
		return "Erro ao salvar conversão\n", fmt.Errorf("Erro ao salvar converss %s\n", err)
	}

	tx.Commit()

	return "Salvo com sucesso", nil

}

func convertCoinFinal(valueCurrent, valueFinal float64) float64 {
	return valueCurrent * valueFinal
}

func libKind(kindCoin string) string {
	switch {
	case kindCoin == "euro":
		return Euro
	case kindCoin == "libra":
		return Libra
	case kindCoin == "real":
		return Real
	case kindCoin == "dolar":
		return Dolar
	default:
		return Dolar

	}
}

func toJson(body []uint8) (map[string]interface{}, error) {
	r := make(map[string]interface{})

	json.Unmarshal([]byte(body), &r)

	rates, err := json.Marshal(r["rates"])
	if err != nil {
		return r, fmt.Errorf("Dados invalidos: %s\n", err)
	} else {
		json.Unmarshal([]byte(rates), &r)

	}

	return r, nil
}

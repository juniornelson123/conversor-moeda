package converter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"testing"
)

const errorConvertCoin = "Não retornou um float64"

func TestConvertCoinFinal(t *testing.T) {
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", convertCoinFinal(2.2, 2.2)), 64)
	expect := 4.84
	if num != expect {
		t.Errorf("Erro, o valor esperado era %v e o recebido foi %v", expect, num)
	}
}

func TestLibKind(t *testing.T) {
	errorFunc := "O valor esperado era %v e o recebido foi %v"
	tests := []struct {
		actual string
		expect string
	}{
		{"euro", "EUR"},
		{"dolar", "USD"},
		{"libra", "GBP"},
		{"real", "BRL"},
	}

	for _, test := range tests {
		result := libKind(test.actual)
		if result != test.expect {
			t.Errorf(errorFunc, test.expect, result)
		}
	}
}

func TestToJson(t *testing.T) {
	url := "google.com"
	jsonBody := `{"rates":"` + url + `"}`

	byteTest := bytes.NewBufferString(jsonBody)

	actual, _ := ioutil.ReadAll(byteTest)

	result, _ := toJson(actual)

	if result["rates"] != "google.com" {
		t.Errorf("O valor esperado era %v, porém o recebido foi %v", url, result["rates"])
	}

}

func TestConvertCoin(t *testing.T) {

	testes := []struct {
		kind   string
		coin   float64
		expect interface{}
	}{
		{"euro", 13.2, reflect.Float64},
	}
	for _, test := range testes {

		c := Coin{test.kind, test.coin}
		value, _ := c.ConvertCoin("dolar")
		myType := reflect.TypeOf(value)

		if myType.Kind() != test.expect {
			t.Errorf(errorConvertCoin)
		}
	}

}

package converter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
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

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testes := []struct {
		kind   string
		coin   float64
		expect interface{}
	}{
		{"euro", 13.2, reflect.Float64},
	}
	for _, test := range testes {

		c := Coin{test.kind, test.coin, 0.0}
		value, _ := c.ConvertCoin("dolar", db)
		myType := reflect.TypeOf(value)

		if myType.Kind() != test.expect {
			t.Errorf(errorConvertCoin)
		}
	}

}

func TestInsertCoin(t *testing.T) {
	c := Coin{"euro", 12.4, 12.33}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO coins")
	mock.ExpectExec("INSERT INTO coins").WithArgs("euro", 12.4, 12.33).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	resp, err := c.insertCoin(db)

	if err != nil {
		if err.Error() != "Erro ao salvar conversão\n" {
			t.Errorf("Era esperada a mensagem Erro ao salvar conversão, ao inves disso foi recebido %s \n", err.Error())
		}
	} else {
		if resp != "Salvo com sucesso" {
			t.Errorf("Era esperada a mensagem Salvo com sucesso, ao inves disso foi recebido %s \n", resp)
		}

	}

}

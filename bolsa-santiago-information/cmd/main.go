package main

import (
	"encoding/json"
	"fmt"

	"fvillarroelb.cl/bolsa-santiago-information/internal/entity"
)

const URL_BOLSA_UT_MERCADOS = "https://www.bolsadesantiago.com/api/Mercados/getUTMercados"

func main() {
	req := entity.RequestMarketAll{}
	response, errorRepository := repository.getUTMercados(req)
	if errorRepository != nil {
		//FATAL error
	}
	responseToJson := entity.ListaResultArray{}
	json.Unmarshal([]byte(response), &responseToJson)
	fmt.Println(responseToJson)
}

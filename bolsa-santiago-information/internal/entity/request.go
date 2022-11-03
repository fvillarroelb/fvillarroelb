package entity

type RequestMarketAll struct {
	initDate string `json:"fec_pagoini"`
	lastDate string `json:"fec_pagofin"`
	nemo     string `json:"nemo"`
}

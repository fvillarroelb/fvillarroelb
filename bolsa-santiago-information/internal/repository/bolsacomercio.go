package repository

import (
	"io/ioutil"
	"net/http"
	"strings"

	"fvillarroelb.cl/bolsa-santiago-information/internal/entity"
)

func (rq *entity.RequestMercadoAll) getUTMercados() (response string, err string) {

	url := "https://www.bolsadesantiago.com/api/Mercados/getUTMercados"
	method := "POST"

	// Set payload
	payload := strings.NewReader(rq)
	//payload := strings.NewReader(`{"fec_pagoini":"2022-01-01","fec_pagofin":"2022-12-31","nemo":""}`)

	client := &http.Client{}
	req, error := http.NewRequest(method, url, payload)

	if error != nil {
		//log.Default()(err)
		return
	}
	req.Header.Add("Cookie", "BIGipCookie=000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000; __uzma=79e6e54d-0428-459d-87dc-38886db00d91; __uzmb=1663001672; __uzme=5773; BIGipCookie=000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000; f5avraaaaaaaaaaaaaaaa_session_=DKLHEFMBKHADHCDKCHGCDLMOPGNNFKGMGOFOBIJJKCKDOHJMODBJKBCCAOPGEDOPGFBDDBMNJDPMOGGGKMFACGLBPGOHFDPHBPDCHNMBPGEONNPHJPLDBNFCLFIAIGEI; BIGipServerPool-Push_HTML5_corporativa=684715681.20480.0000; __ssds=2; __ssuzjsr2=a9be0cd8e; __uzmaj2=7d3cd1fb-fd7a-4538-9e8a-154eb558831c; __uzmbj2=1663001673; gb-wbchtbt-uid=1663001674795; _ga=GA1.2.1805551646.1663001675; _gid=GA1.2.745080428.1663001675; _csrf=SHnFl6DaiW6u0B6u2zIEAKdd; __gads=ID=4d6cf2779c3a94b1:T=1663001687:S=ALNI_MZzoILu4slPkv0O5jiQuyIfdMtfFQ; __gpi=UID=000009644fd6d262:T=1663001687:RT=1663001687:S=ALNI_MbR93GQ2YGRn3kInX8TUEIU5v_QOA; _gat=1; __uzmcj2=400402251886; __uzmdj2=1663002427; BIGipServerPool-www.bolsadesantiago.com-HTML5_corporativa=735047329.20480.0000; __uzmc=1655310040454; __uzmd=1663002432; BIGipCookie=000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000; f5avraaaaaaaaaaaaaaaa_session_=GOPKMMCPGDCGPEOLNBNAPOGEAODPJENBHFEEDINCCFEGPDOBDCKJAKENNHFLFCJNBFCDJMNPNBKNAHFIENJANGDNLPMIIEBIJOPMHMOJBCPAFONDNKHNFMKCEANKCHGD; BIGipServerPool-Push_HTML5_corporativa=684715681.20480.0000; BIGipServerPool-www.bolsadesantiago.com-HTML5_corporativa=684715681.20480.0000; __uzma=fe31de40-675b-48b6-b761-ff4a7a94c4e2; __uzmb=1663002361; __uzmc=7648910324972; __uzmd=1663037965; __uzme=9505; _csrf=q3bJQH8z97ajc55kn1uInjvW")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36 Edg/105.0.1343.27")
	req.Header.Add("x-csrf-token", "9G8bMVxi-wPn9NrW2JM8AvOoJOdmpVpP0_IQ")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Content-Type", "text/plain")

	res, error2 := client.Do(req)
	if error2 != nil {
		//fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, error3 := ioutil.ReadAll(res.Body)
	if error3 != nil {
		//	fmt.Println(err)
		return
	}
	//log.Default(string(body))
	return string(body), err

}

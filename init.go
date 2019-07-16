package gosoap

import (
	"crypto/tls"
	"net/http"
	"time"
)

var myclient *http.Client

func init() {
	Clientt()
}

func Clientt() {

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		MaxIdleConnsPerHost:300,
		MaxIdleConns: 300 ,
	}
	myclient = &http.Client{Transport: transCfg, Timeout: time.Second * 60}

}

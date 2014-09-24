package main

import (
	"net/http"
	"fmt"
	"strconv"
	"net/http/cookiejar"
	"io/ioutil"
	"strings"
	"crypto/tls"
	"errors"
)
type CasClient struct {
	CasServer string
	AppUrl    string
}

func (this *CasClient)Validate(ticket string) (string,error){
	val_url := this.CasServer + "/validate" + "?service=" + strconv.Quote(this.AppUrl) + "&ticket=" + strconv.Quote(ticket)
	cookieJar, _ := cookiejar.New(nil)
    client := http.Client{
		Jar: cookieJar,
		Transport:&http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp,err := client.Get(val_url)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	res := strings.Split(string(result),"\n")
	if len(res) > 1 && res[0]=="yes"{
		return res[1],nil
	}
	return "", errors.New("认证失败！"+string(result))
}

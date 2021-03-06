package controller

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
	Model "../model"
)

func CrawDataByVN(code string ) []byte  {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	payload := fmt.Sprintf(	`_congbothongke_WAR_coronadvcportlet_ma=%s&_congbothongke_WAR_coronadvcportlet_jsonData=[{"name":"Ha+Noi","ma":"01","soCaNhiem":"0","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":"0"},{"name":"aaaaa","ma":"","soCaNhiem":"20","tuVong":"0","nghiNhiem":"120"},{"name":"bbb","ma":"","soCaNhiem":"20","tuVong":"0","nghiNhiem":"120"},{"ma":"02","soCaNhiem":"0","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":"0"},{"ma":"--Chọn+địa+phương--","soCaNhiem":"","tuVong":"","nghiNhiem":""},{"ma":"VNALL","soCaNhiem":"16","tuVong":"0","nghiNhiem":"61","binhPhuc":"7","cachLy":"1.068"},{"ma":"79","soCaNhiem":"3","tuVong":"0","nghiNhiem":"0","binhPhuc":"3+","cachLy":""},{"ma":"26","soCaNhiem":"11","tuVong":"0","nghiNhiem":"0","binhPhuc":"3","cachLy":""},{"ma":"38","soCaNhiem":"1","tuVong":"0","nghiNhiem":"0","binhPhuc":"1","cachLy":""},{"ma":"56","soCaNhiem":"1","tuVong":"0","nghiNhiem":"0","binhPhuc":"1","cachLy":""},{"ma":"08","soCaNhiem":"0","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":"0"}]`,code)
	body:=strings.NewReader(payload)
	req, err := http.NewRequest("POST", "https://ncov.moh.gov.vn/web/guest/trang-chu?p_p_id=congbothongke_WAR_coronadvcportlet&p_p_lifecycle=2&p_p_state=normal&p_p_mode=view&p_p_resource_id=getItemByMa&p_p_cacheability=cacheLevelPage&p_p_col_id=column-1&p_p_col_count=6",body)
	if err != nil {
		return nil
	}
	req.Header.Set("Authority", "ncov.moh.gov.vn")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://ncov.moh.gov.vn")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.78 Safari/537.36 Edg/80.0.361.45")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Referer", "https://ncov.moh.gov.vn/")
	req.Header.Set("Accept-Language", "en-VN,en-US;q=0.9,en;q=0.8")
	//req.Header.Set("Cookie", "JSESSIONID=E06B9BCBB1497370A7464F89665570A6; COOKIE_SUPPORT=true; GUEST_LANGUAGE_ID=vi_VN; AWSALB=O+r9ELICA7LmRL3nfV1C6VX60EmTZ4cGhs+9CGVBaEvUQ4rOMeAqPIJkRWKFkN+IUdupvY59s1NLnRO9x+ktBa4JsaY4YvG/yrzqi6GFLOcnaC7+KPqQiDwSJxjH; AWSALBCORS=O+r9ELICA7LmRL3nfV1C6VX60EmTZ4cGhs+9CGVBaEvUQ4rOMeAqPIJkRWKFkN+IUdupvY59s1NLnRO9x+ktBa4JsaY4YvG/yrzqi6GFLOcnaC7+KPqQiDwSJxjH")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		return nil
	}
	defer resp.Body.Close()
	data, err:= ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	ioutil.WriteFile("data.json",data,755)
	return  data
}
func CrawDataByCountry(code string) []byte  {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// TODO: This is insecure; use only in dev environments.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	payload := fmt.Sprintf(`_congbothongke_WAR_coronadvcportlet_ma=%s&_congbothongke_WAR_coronadvcportlet_jsonData=[{"name":"01","ma":"","soCaNhiem":"20","tuVong":"0","nghiNhiem":"120"},{"name":"aaaaa","ma":"","soCaNhiem":"20","tuVong":"0","nghiNhiem":"120"},{"name":"bbb","ma":"","soCaNhiem":"20","tuVong":"0","nghiNhiem":"120"},{"ma":"01","soCaNhiem":"","tuVong":"","nghiNhiem":""},{"ma":"02","soCaNhiem":"","tuVong":"","nghiNhiem":""},{"ma":"--Chọn+địa+phương--","soCaNhiem":"6","tuVong":"0","nghiNhiem":"6"},{"ma":"--Chọn+địa+phương--","soCaNhiem":"6","tuVong":"0","nghiNhiem":"99"},{"ma":"AL","soCaNhiem":"","tuVong":"0","nghiNhiem":"","binhPhuc":"","cachLy":""},{"ma":"TW","soCaNhiem":"18","tuVong":"0","nghiNhiem":"0","binhPhuc":"1","cachLy":""},{"ma":"QTALL","soCaNhiem":"69.270","tuVong":"1.669+","nghiNhiem":"","binhPhuc":"9.527","cachLy":""},{"ma":"CN","soCaNhiem":"68.502","tuVong":"1.665","nghiNhiem":"10.978","binhPhuc":"9.419","cachLy":""},{"ma":"JP","soCaNhiem":"408+","tuVong":"1+","nghiNhiem":"0","binhPhuc":"4","cachLy":"0"},{"ma":"TH","soCaNhiem":"34+","tuVong":"0","nghiNhiem":"0","binhPhuc":"14","cachLy":""},{"ma":"SG","soCaNhiem":"72","tuVong":"0","nghiNhiem":"0","binhPhuc":"18+","cachLy":""},{"ma":"KR","soCaNhiem":"28","tuVong":"0","nghiNhiem":"0","binhPhuc":"9","cachLy":""},{"ma":"HK","soCaNhiem":"56","tuVong":"1","nghiNhiem":"0","binhPhuc":"1+","cachLy":""},{"ma":"AU","soCaNhiem":"15","tuVong":"0","nghiNhiem":"0","binhPhuc":"10+","cachLy":""},{"ma":"US","soCaNhiem":"15","tuVong":"0","nghiNhiem":"0","binhPhuc":"3","cachLy":""},{"ma":"CA","soCaNhiem":"7","tuVong":"0","nghiNhiem":"0","binhPhuc":"1","cachLy":"0"},{"ma":"DE","soCaNhiem":"16","tuVong":"0","nghiNhiem":"0","binhPhuc":"3+","cachLy":"0"},{"ma":"PH","soCaNhiem":"3","tuVong":"1","nghiNhiem":"0","binhPhuc":"2","cachLy":""},{"ma":"FI","soCaNhiem":"1","tuVong":"0","nghiNhiem":"0","binhPhuc":"1","cachLy":""},{"ma":"BE","soCaNhiem":"1","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":""},{"ma":"KH","soCaNhiem":"1","tuVong":"0","nghiNhiem":"0","binhPhuc":"1","cachLy":""},{"ma":"IN","soCaNhiem":"3","tuVong":"0","nghiNhiem":"0","binhPhuc":"3+","cachLy":""},{"ma":"NP","soCaNhiem":"1","tuVong":"0","nghiNhiem":"0","binhPhuc":"1+","cachLy":""},{"ma":"FR","soCaNhiem":"12","tuVong":"1+","nghiNhiem":"0","binhPhuc":"4+","cachLy":""},{"ma":"ES","soCaNhiem":"2","tuVong":"0","nghiNhiem":"0","binhPhuc":"2+","cachLy":""},{"ma":"MY","soCaNhiem":"22+","tuVong":"0","nghiNhiem":"0","binhPhuc":"3+","cachLy":"1"},{"ma":"MO","soCaNhiem":"10","tuVong":"0","nghiNhiem":"0","binhPhuc":"1+","cachLy":""},{"ma":"RU","soCaNhiem":"2","tuVong":"0","nghiNhiem":"0","binhPhuc":"2","cachLy":""},{"ma":"SE","soCaNhiem":"1","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":""},{"ma":"VIE","soCaNhiem":"14","tuVong":"0","nghiNhiem":"82","binhPhuc":"3","cachLy":"745"},{"ma":"GB","soCaNhiem":"9","tuVong":"0","nghiNhiem":"0","binhPhuc":"8+","cachLy":"0"},{"ma":"LK","soCaNhiem":"1","tuVong":"0","nghiNhiem":"0","binhPhuc":"1","cachLy":"0"},{"ma":"DJ","soCaNhiem":"5","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":"0"},{"ma":"IT","soCaNhiem":"3+","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":"0"},{"ma":"UY","soCaNhiem":"0","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":"0"},{"ma":"AE","soCaNhiem":"8","tuVong":"0","nghiNhiem":"0","binhPhuc":"1","cachLy":"0"},{"ma":"GE","soCaNhiem":"0","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":"0"},{"ma":"EG","soCaNhiem":"1+","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":"0"}]&_congbothongke_WAR_coronadvcportlet_ma=AQ&_congbothongke_WAR_coronadvcportlet_jsonData=[{"name":"01","ma":"","soCaNhiem":"20","tuVong":"0","nghiNhiem":"120"},{"name":"aaaaa","ma":"","soCaNhiem":"20","tuVong":"0","nghiNhiem":"120"},{"name":"bbb","ma":"","soCaNhiem":"20","tuVong":"0","nghiNhiem":"120"},{"ma":"01","soCaNhiem":"","tuVong":"","nghiNhiem":""},{"ma":"02","soCaNhiem":"","tuVong":"","nghiNhiem":""},{"ma":"--Chọn+địa+phương--","soCaNhiem":"6","tuVong":"0","nghiNhiem":"6"},{"ma":"--Chọn+địa+phương--","soCaNhiem":"6","tuVong":"0","nghiNhiem":"99"},{"ma":"AL","soCaNhiem":"","tuVong":"0","nghiNhiem":"","binhPhuc":"","cachLy":""},{"ma":"TW","soCaNhiem":"18","tuVong":"0","nghiNhiem":"0","binhPhuc":"1","cachLy":""},{"ma":"QTALL","soCaNhiem":"69.270","tuVong":"1.669+","nghiNhiem":"","binhPhuc":"9.527","cachLy":""},{"ma":"CN","soCaNhiem":"68.502","tuVong":"1.665","nghiNhiem":"10.978","binhPhuc":"9.419","cachLy":""},{"ma":"JP","soCaNhiem":"408+","tuVong":"1+","nghiNhiem":"0","binhPhuc":"4","cachLy":"0"},{"ma":"TH","soCaNhiem":"34+","tuVong":"0","nghiNhiem":"0","binhPhuc":"14","cachLy":""},{"ma":"SG","soCaNhiem":"72","tuVong":"0","nghiNhiem":"0","binhPhuc":"18+","cachLy":""},{"ma":"KR","soCaNhiem":"28","tuVong":"0","nghiNhiem":"0","binhPhuc":"9","cachLy":""},{"ma":"HK","soCaNhiem":"56","tuVong":"1","nghiNhiem":"0","binhPhuc":"1+","cachLy":""},{"ma":"AU","soCaNhiem":"15","tuVong":"0","nghiNhiem":"0","binhPhuc":"10+","cachLy":""},{"ma":"US","soCaNhiem":"15","tuVong":"0","nghiNhiem":"0","binhPhuc":"3","cachLy":""},{"ma":"CA","soCaNhiem":"7","tuVong":"0","nghiNhiem":"0","binhPhuc":"1","cachLy":"0"},{"ma":"DE","soCaNhiem":"16","tuVong":"0","nghiNhiem":"0","binhPhuc":"3+","cachLy":"0"},{"ma":"PH","soCaNhiem":"3","tuVong":"1","nghiNhiem":"0","binhPhuc":"2","cachLy":""},{"ma":"FI","soCaNhiem":"1","tuVong":"0","nghiNhiem":"0","binhPhuc":"1","cachLy":""},{"ma":"BE","soCaNhiem":"1","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":""},{"ma":"KH","soCaNhiem":"1","tuVong":"0","nghiNhiem":"0","binhPhuc":"1","cachLy":""},{"ma":"IN","soCaNhiem":"3","tuVong":"0","nghiNhiem":"0","binhPhuc":"3+","cachLy":""},{"ma":"NP","soCaNhiem":"1","tuVong":"0","nghiNhiem":"0","binhPhuc":"1+","cachLy":""},{"ma":"FR","soCaNhiem":"12","tuVong":"1+","nghiNhiem":"0","binhPhuc":"4+","cachLy":""},{"ma":"ES","soCaNhiem":"2","tuVong":"0","nghiNhiem":"0","binhPhuc":"2+","cachLy":""},{"ma":"MY","soCaNhiem":"22+","tuVong":"0","nghiNhiem":"0","binhPhuc":"3+","cachLy":"1"},{"ma":"MO","soCaNhiem":"10","tuVong":"0","nghiNhiem":"0","binhPhuc":"1+","cachLy":""},{"ma":"RU","soCaNhiem":"2","tuVong":"0","nghiNhiem":"0","binhPhuc":"2","cachLy":""},{"ma":"SE","soCaNhiem":"1","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":""},{"ma":"VIE","soCaNhiem":"14","tuVong":"0","nghiNhiem":"82","binhPhuc":"3","cachLy":"745"},{"ma":"GB","soCaNhiem":"9","tuVong":"0","nghiNhiem":"0","binhPhuc":"8+","cachLy":"0"},{"ma":"LK","soCaNhiem":"1","tuVong":"0","nghiNhiem":"0","binhPhuc":"1","cachLy":"0"},{"ma":"DJ","soCaNhiem":"5","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":"0"},{"ma":"IT","soCaNhiem":"3+","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":"0"},{"ma":"UY","soCaNhiem":"0","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":"0"},{"ma":"AE","soCaNhiem":"8","tuVong":"0","nghiNhiem":"0","binhPhuc":"1","cachLy":"0"},{"ma":"GE","soCaNhiem":"0","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":"0"},{"ma":"EG","soCaNhiem":"1+","tuVong":"0","nghiNhiem":"0","binhPhuc":"0","cachLy":"0"}]`,code)
	body:=strings.NewReader(payload)

	req, err := http.NewRequest("POST", "https://ncov.moh.gov.vn/web/guest/trang-chu?p_p_id=congbothongke_WAR_coronadvcportlet&p_p_lifecycle=2&p_p_state=normal&p_p_mode=view&p_p_resource_id=getItemByMa&p_p_cacheability=cacheLevelPage&p_p_col_id=column-1&p_p_col_count=6", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authority", "ncov.moh.gov.vn")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.78 Safari/537.36 Edg/80.0.361.45")
	req.Header.Set("Sec-Fetch-Dest", "iframe")
	req.Header.Set("Accept", "image/webp,image/apng,image/*,*/*;q=0.8")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "no-cors")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Referer", "https://ncov.moh.gov.vn/")
	req.Header.Set("Accept-Language", "en-VN,en-US;q=0.9,en;q=0.8")
	//req.Header.Set("Cookie", "_ga=GA1.3.1845270104.1581762891; _gid=GA1.3.126797458.1581762891; COOKIE_SUPPORT=true; GUEST_LANGUAGE_ID=vi_VN; JSESSIONID=C4379EF56E8382BCA594DCBC90B605ED; _gat=1; AWSALB=Lwifhtoyv6BEsHy3WSxJLG38wv8WQROK080g2wQHrqnQv1RENYPqfqhKi93f0h5sfUPT6b8VLvPW9lkH1RRAEsN5z05lxnOELSLXFVpd+10Kt7GpM6THC0H0bvmT; AWSALBCORS=Lwifhtoyv6BEsHy3WSxJLG38wv8WQROK080g2wQHrqnQv1RENYPqfqhKi93f0h5sfUPT6b8VLvPW9lkH1RRAEsN5z05lxnOELSLXFVpd+10Kt7GpM6THC0H0bvmT; LFR_SESSION_STATE_20159=1581834382544")
	req.Header.Set("Origin", "https://ncov.moh.gov.vn")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Pragma", "no-cache")

	resp, err := client.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	data, err:= ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	ioutil.WriteFile("data.json",data,755)
	return data
}


type CoronaQT struct {
	beego.Controller
}

func (this *CoronaQT) Get()  {
	var (
		Type string
		Code string
	)
	var coronaResp Model.CoronaResponse
	this.Ctx.Input.Bind(&Type,"type")
	this.Ctx.Input.Bind(&Code,"code")
	_ = json.Unmarshal(CrawDataByCountry(Code), &coronaResp)
	var textArr []Model.Text
	textArr= append(textArr,Model.Text{Text:"Tổng số ca nhiễm: "+coronaResp.SoCaNhiem})
	textArr= append(textArr, Model.Text{Text:"Số ca tử vong: "+coronaResp.TuVong})
	textArr= append(textArr, Model.Text{Text:"Số ca cách ly: "+coronaResp.CachLy})
	textArr= append(textArr, Model.Text{Text:"Số ca bình phục: "+coronaResp.BinhPhuc})
	var message Model.Message
	message.Messages = textArr
	this.Data["json"] = message
	this.ServeJSON()

}
type CoronaVN struct {
	beego.Controller
}

func (this *CoronaVN) Get()  {
	var (
		Type string
		Code string
	)
	var coronaResp Model.CoronaResponse
	this.Ctx.Input.Bind(&Type,"type")
	this.Ctx.Input.Bind(&Code,"code")
	_ = json.Unmarshal(CrawDataByVN(Code), &coronaResp)
	var textArr []Model.Text
	textArr= append(textArr,Model.Text{Text:"Tổng số ca nhiễm: "+coronaResp.SoCaNhiem})
	textArr= append(textArr,Model.Text{Text:"Tổng số ca nghi nhiễm: "+coronaResp.NghiNhiem})
	textArr= append(textArr, Model.Text{Text:"Số ca tử vong: "+coronaResp.TuVong})
	textArr= append(textArr, Model.Text{Text:"Số ca cách ly: "+coronaResp.CachLy})
	textArr= append(textArr, Model.Text{Text:"Số ca bình phục: "+coronaResp.BinhPhuc})
	var message Model.Message
	message.Messages = textArr
	this.Data["json"] = message
	this.ServeJSON()

}
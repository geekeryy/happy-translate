// Package baidutranslate @Description  TODO
// @Author  	 jiangyang
// @Created  	 2023/10/16 09:41
package baidutranslate

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func Translate(q string) (*Result, error) {
	appid := os.Getenv("BAIDU_TRANSLATE_APPID")
	appkey := os.Getenv("BAIDU_TRANSLATE_APPKEY")
	from := "auto"
	to := "zh"
	salt := time.Now().Unix()
	sign := Md5(fmt.Sprintf("%s%s%d%s", appid, q, salt, appkey))

	qString := url.QueryEscape(q)
	urlLanguage := fmt.Sprintf("https://fanyi-api.baidu.com/api/trans/vip/language?q=%s&appid=%s&salt=%d&sign=%s", qString, appid, salt, sign)
	resultLanguage, err := Post(urlLanguage)
	if err != nil {
		return nil, err
	}
	if resultLanguage.Data.Src == "zh" {
		to = "en"
	}
	from = resultLanguage.Data.Src

	urlTranslate := fmt.Sprintf("https://api.fanyi.baidu.com/api/trans/vip/translate?q=%s&from=%s&to=%s&appid=%s&salt=%d&sign=%s", qString, from, to, appid, salt, sign)
	resultTranslate, err := Post(urlTranslate)
	if err != nil {
		return nil, err
	}
	return resultTranslate, nil
}

func Post(url string) (*Result, error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("error code:%d", resp.StatusCode))
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result Result
	if err := json.Unmarshal(all, &result); err != nil {
		return nil, errors.Join(err, errors.New(string(all)))
	}
	if result.ErrorCode > 0 {
		return nil, errors.New(result.ErrorMsg)
	}
	return &result, nil
}

type Result struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	From      string `json:"from"`
	To        string `json:"to"`
	Data      struct {
		Src string `json:"src"`
	} `json:"data"`
	TransResult []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

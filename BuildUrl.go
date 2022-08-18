package main

import (
	"net/url"
	"strconv"
)

func BuildUrl(label int64, sum string) string {
	strLabel := strconv.FormatInt(label, 10) //переведём наш label который chatID из цыферек в строку
	u := url.Values{}
	u.Set("receiver", "4100117733655947")
	u.Set("quickpay-form", "shop")
	u.Set("targets", "Tiket")
	u.Set("paymentType", "SB")
	u.Set("sum", sum)
	u.Set("formcomment", "")
	u.Set("short_dest", "")
	u.Set("label", strLabel)
	u.Set("comment", "")
	u.Set("successURL", "")
	u.Set("need_fio", "")
	u.Set("need_email", "")
	u.Set("need_phone", "")
	u.Set("need_address", "")

	/*auth := Info{
		method: "request-payment",
		Url:    "https://yoomoney.ru/api/",
	}*/
	auth := Info{
		method: "quickpay/confirm.xml?",
		Url:    "https://yoomoney.ru/",
	}
	//auth.requestToYooMoney(u)
	auth.Url = auth.Url + auth.method + u.Encode()
	return auth.Url
}

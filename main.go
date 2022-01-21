package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func main() {

	phoneNumbers := flag.String("phoneNumbers", "", "请输入接受手机号码")
	templateCode := flag.String("templateCode", "", "请输入模版号码")
	templateParam := flag.String("templateParam", "", "请输入模版内容")
	signName := flag.String("signName", "文储", "请输入模版内容")

	flag.Parse()
	client, err := dysmsapi.NewClientWithAccessKey("ap-northeast-1", "APPKEY", "秘钥")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.SignName = *signName
	request.PhoneNumbers = *phoneNumbers
	request.TemplateCode = *templateCode
	request.TemplateParam = *templateParam

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(response.String())
}

package main

import (
	"github.com/YunzhanghuOpen/sdk-go/example/authentication"
	"github.com/YunzhanghuOpen/sdk-go/example/dataservice"
	"github.com/YunzhanghuOpen/sdk-go/example/invoice"
	"github.com/YunzhanghuOpen/sdk-go/example/payment"
	"github.com/YunzhanghuOpen/sdk-go/example/tax"
)

func main() {
	payment.Example()
	dataservice.Example()
	tax.Example()
	invoice.Example()
	authentication.Example()
}

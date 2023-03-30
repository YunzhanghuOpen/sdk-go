package main

import (
	"github.com/YunzhanghuOpen/sdk-go/example/apiusersign"
	"github.com/YunzhanghuOpen/sdk-go/example/authentication"
	"github.com/YunzhanghuOpen/sdk-go/example/bizlicxjjh5"
	"github.com/YunzhanghuOpen/sdk-go/example/bizlicxjjh5api"
	"github.com/YunzhanghuOpen/sdk-go/example/dataservice"
	"github.com/YunzhanghuOpen/sdk-go/example/h5usersign"
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
	apiusersign.Example()
	h5usersign.Example()
	bizlicxjjh5.Example()
	bizlicxjjh5api.Example()
}

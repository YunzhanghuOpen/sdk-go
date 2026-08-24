package main

import (
	"github.com/YunzhanghuOpen/sdk-go/example/apiusersign"
	"github.com/YunzhanghuOpen/sdk-go/example/authentication"
	"github.com/YunzhanghuOpen/sdk-go/example/calculatelabor"
	"github.com/YunzhanghuOpen/sdk-go/example/custom"
	"github.com/YunzhanghuOpen/sdk-go/example/customerlink"
	"github.com/YunzhanghuOpen/sdk-go/example/dataservice"
	"github.com/YunzhanghuOpen/sdk-go/example/h5usersign"
	"github.com/YunzhanghuOpen/sdk-go/example/invoice"
	"github.com/YunzhanghuOpen/sdk-go/example/payment"
	"github.com/YunzhanghuOpen/sdk-go/example/realname"
	"github.com/YunzhanghuOpen/sdk-go/example/taxclearrefund"
	"github.com/YunzhanghuOpen/sdk-go/example/usercollect"
	"github.com/YunzhanghuOpen/sdk-go/example/faceauth"
)

func main() {
	payment.Example()
	dataservice.Example()
	invoice.Example()
	authentication.Example()
	apiusersign.Example()
	h5usersign.Example()
	custom.Example()
	usercollect.Example()
	calculatelabor.Example()
	customerlink.Example()
	realname.Example()
	taxclearrefund.Example()
	faceauth.Example()
}

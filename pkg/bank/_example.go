package bank

import (
	"fmt"
	"net/http"
)

const (
	merchantID      = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
	callbackURL     = "http://localhost:8080/verify"
	apiURL          = "https://www.zarinpal.com/pg/services/WebGate/wsdl"
	zarinpalGateURL = "https://www.zarinpal.com/pg/StartPay/"
)

func request(w http.ResponseWriter, r *http.Request) {
	clinet := NewPaymentGatewayImplementationServicePortType("", false, nil)

	// Create a new payment request to Zarinpal
	resp, err := clinet.PaymentRequest(&PaymentRequest{
		MerchantID:  merchantID,
		Amount:      100,
		Description: "توضیحات خرید آزمایشی",
		Email:       "customer@domain.ir",
		Mobile:      "09123456789",
		CallbackURL: callbackURL,
	})

	// Check if response is error free
	if err != nil {
		http.Error(w, fmt.Sprintln("مشکلی در ارتباط رخ داد", err), http.StatusInternalServerError)
		return
	}

	if resp.Status == 100 {
		// redirect user to zarinpal
		http.Redirect(w, r, zarinpalGateURL+resp.Authority, http.StatusFound)
		return
	}

	http.Error(w, fmt.Sprintln("خطایی رخ داد:", resp.Status), http.StatusInternalServerError)
}

func verify(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("Status") != "OK" {
		fmt.Fprintln(w, "تراکنش توسط کاربر خاتمه یافت.")
		return
	}

	clinet := NewPaymentGatewayImplementationServicePortType("", false, nil)

	// Create a new payment request to Zarinpal
	resp, err := clinet.PaymentVerification(&PaymentVerification{
		MerchantID: merchantID,
		Amount:     100,
		Authority:  r.FormValue("Authority"),
	})

	// Check if response is error free
	if err != nil {
		http.Error(w, fmt.Sprintln("مشکلی در ارتباط رخ داد", err), http.StatusInternalServerError)
		return
	}

	if resp.Status == 100 {
		fmt.Fprintln(w, "تراکنش با موفقیت تایید شد.", resp.RefID)
	} else {
		fmt.Fprintln(w, "تراکنش تایید نشد.", resp.Status)
	}
}

func main() {
	// New request to redirect user to Zarinpal
	http.HandleFunc("/request", request)

	// Verify redirected user response from Zarinpal
	http.HandleFunc("/verify", verify)

	fmt.Println("Please visit: http://localhost:8080/request")

	http.ListenAndServe(":8080", nil)
}

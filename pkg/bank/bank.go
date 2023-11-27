package bank

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	// "log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// PaymentRequest stores a payment request data
type PaymentRequest struct {
	XMLName xml.Name `xml:"http://zarinpal.com/ PaymentRequest"`

	MerchantID string `xml:"MerchantID,omitempty"`

	Amount int32 `xml:"Amount,omitempty"`

	Description string `xml:"Description,omitempty"`

	Email string `xml:"Email,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	CallbackURL string `xml:"CallbackURL,omitempty"`
}

// PaymentRequestResponse stores a payment request response data
type PaymentRequestResponse struct {
	XMLName xml.Name `xml:"http://zarinpal.com/ PaymentRequestResponse"`

	Status int32 `xml:"Status,omitempty"`

	Authority string `xml:"Authority,omitempty"`
}

// PaymentRequestWithExtra stores a payment request with additional data
type PaymentRequestWithExtra struct {
	XMLName xml.Name `xml:"http://zarinpal.com/ PaymentRequestWithExtra"`

	MerchantID string `xml:"MerchantID,omitempty"`

	Amount int32 `xml:"Amount,omitempty"`

	Description string `xml:"Description,omitempty"`

	AdditionalData string `xml:"AdditionalData,omitempty"`

	Email string `xml:"Email,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	CallbackURL string `xml:"CallbackURL,omitempty"`
}

// PaymentRequestWithExtraResponse stores a payment request response with additional data
type PaymentRequestWithExtraResponse struct {
	XMLName xml.Name `xml:"http://zarinpal.com/ PaymentRequestWithExtraResponse"`

	Status int32 `xml:"Status,omitempty"`

	Authority string `xml:"Authority,omitempty"`
}

// PaymentVerification stores a payment verification request data
type PaymentVerification struct {
	XMLName xml.Name `xml:"http://zarinpal.com/ PaymentVerification"`

	MerchantID string `xml:"MerchantID,omitempty"`

	Authority string `xml:"Authority,omitempty"`

	Amount int32 `xml:"Amount,omitempty"`
}

// PaymentVerificationResponse stores a payment verification response data
type PaymentVerificationResponse struct {
	XMLName xml.Name `xml:"http://zarinpal.com/ PaymentVerificationResponse"`

	Status int32 `xml:"Status,omitempty"`

	RefID int64 `xml:"RefID,omitempty"`
}

// PaymentVerificationWithExtra stores a payment verification request with extra data
type PaymentVerificationWithExtra struct {
	XMLName xml.Name `xml:"http://zarinpal.com/ PaymentVerificationWithExtra"`

	MerchantID string `xml:"MerchantID,omitempty"`

	Authority string `xml:"Authority,omitempty"`

	Amount int32 `xml:"Amount,omitempty"`
}

// PaymentVerificationWithExtraResponse stores a payment verification response with extra data
type PaymentVerificationWithExtraResponse struct {
	XMLName xml.Name `xml:"http://zarinpal.com/ PaymentVerificationWithExtraResponse"`

	Status int32 `xml:"Status,omitempty"`

	RefID int64 `xml:"RefID,omitempty"`

	ExtraDetail string `xml:"ExtraDetail,omitempty"`
}

// GetUnverifiedTransactions is used to get unverified transactions
type GetUnverifiedTransactions struct {
	XMLName xml.Name `xml:"http://zarinpal.com/ GetUnverifiedTransactions"`

	MerchantID string `xml:"MerchantID,omitempty"`
}

// GetUnverifiedTransactionsResponse is used to retrieve unverfied transaction response
type GetUnverifiedTransactionsResponse struct {
	XMLName xml.Name `xml:"http://zarinpal.com/ GetUnverifiedTransactionsResponse"`

	Status int32 `xml:"Status,omitempty"`

	Authorities string `xml:"Authorities,omitempty"`
}

// RefreshAuthority is used to refresh the authority
type RefreshAuthority struct {
	XMLName xml.Name `xml:"http://zarinpal.com/ RefreshAuthority"`

	MerchantID string `xml:"MerchantID,omitempty"`

	Authority string `xml:"Authority,omitempty"`

	ExpireIn int32 `xml:"ExpireIn,omitempty"`
}

// RefreshAuthorityResponse stores refresh authority response
type RefreshAuthorityResponse struct {
	XMLName xml.Name `xml:"http://zarinpal.com/ RefreshAuthorityResponse"`

	Status int32 `xml:"Status,omitempty"`
}

// PaymentGatewayImplementationServicePortType stores a soapClient related to this implementation
type PaymentGatewayImplementationServicePortType struct {
	client *SOAPClient
}

// NewPaymentGatewayImplementationServicePortType will return a new payment gateway with a soapClient
func NewPaymentGatewayImplementationServicePortType(url string, tls bool, auth *BasicAuth) *PaymentGatewayImplementationServicePortType {
	if url == "" {
		url = "https://www.zarinpal.com/pg/services/WebGate/service"
	}
	client := NewSOAPClient(url, tls, auth)

	return &PaymentGatewayImplementationServicePortType{
		client: client,
	}
}

// PaymentRequest will generate a new payment request and returns the response and error, if any
func (service *PaymentGatewayImplementationServicePortType) PaymentRequest(request *PaymentRequest) (*PaymentRequestResponse, error) {
	response := new(PaymentRequestResponse)
	err := service.client.Call("#PaymentRequest", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// PaymentRequestWithExtra will generate a new payment with extra data and returns the response and error, if any
func (service *PaymentGatewayImplementationServicePortType) PaymentRequestWithExtra(request *PaymentRequestWithExtra) (*PaymentRequestWithExtraResponse, error) {
	response := new(PaymentRequestWithExtraResponse)
	err := service.client.Call("#PaymentRequestWithExtra", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// PaymentVerification will generate a payment verification request and returns the response and error, if any
func (service *PaymentGatewayImplementationServicePortType) PaymentVerification(request *PaymentVerification) (*PaymentVerificationResponse, error) {
	response := new(PaymentVerificationResponse)
	err := service.client.Call("#PaymentVerification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// PaymentVerificationWithExtra will generate a payment verification request with extra data and returns the response and error, if any
func (service *PaymentGatewayImplementationServicePortType) PaymentVerificationWithExtra(request *PaymentVerificationWithExtra) (*PaymentVerificationWithExtraResponse, error) {
	response := new(PaymentVerificationWithExtraResponse)
	err := service.client.Call("#PaymentVerificationWithExtra", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetUnverifiedTransactions returns the unverified transactions
func (service *PaymentGatewayImplementationServicePortType) GetUnverifiedTransactions(request *GetUnverifiedTransactions) (*GetUnverifiedTransactionsResponse, error) {
	response := new(GetUnverifiedTransactionsResponse)
	err := service.client.Call("#GetUnverifiedTransactions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// RefreshAuthority will refresh the authority
func (service *PaymentGatewayImplementationServicePortType) RefreshAuthority(request *RefreshAuthority) (*RefreshAuthorityResponse, error) {
	response := new(RefreshAuthorityResponse)
	err := service.client.Call("#RefreshAuthority", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

// SOAPEnvelope entity
type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

// SOAPHeader entity
type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

// SOAPBody entity
type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

// SOAPFault entity
type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

// BasicAuth entity
type BasicAuth struct {
	Login    string
	Password string
}

// SOAPClient SOAPClient
type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
}

// UnmarshalXML unmarshals a given XML
func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

// Error returns the error of the SOAPFault
func (f *SOAPFault) Error() string {
	return f.String
}

// NewSOAPClient creates a new SOAPClient
func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

// Call calls the webservice method
func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{
		//Header:        SoapHeader{},
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	// log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		// log.Println("empty response")
		return nil
	}

	// log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}

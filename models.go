package bpaygo

type (

	// login request and response
	BpayLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	BpayLoginResponse struct {
		BpayResponse
		TokenType    string `json:"tokenType"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    int    `json:"expiresIn"`
		AccessToken  string `json:"accessToken"`
		UserId       int    `json:"userId"`
		RoleId       int    `json:"roleId"`
		JTI          string `json:"jti"`
		Username     string `json:"username"`
	}

	// Customer request and response
	BpayCustomerRegisterRequest struct {
		UserID string `json:"userId"`
		Email  string `json:"email"`
	}
	BpayCustomerRegisterResponse struct {
		BpayResponse
		Data string `json:"data"` // Хэрэглэгчийн BPAY рүү нэвтрэхэд таних код
	}

	BpayCustomerLoginRequest struct {
		UserID   string `json:"userId"`
		BpayCOde string `json:"bpayCode"`
	}
	BpayCustomerLoginResponse struct {
		BpayResponse
	}
	BpayCustomerCheckRequest struct {
		UserID string `json:"userId"`
	}
	BpayCustomerCheckResponse struct {
		BpayResponse
		Data string `json:"data"` // Хэрэглэгчийн BPAY рүү нэвтрэхэд таних код
	}

	//Constants request and response
	BpayConstantResponse struct {
		BpayResponse
		Data []BpayConstantData `json:"data"`
	}
	BpayConstantData struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	// Group request and response
	BpayGroupCreateRequest struct {
		Name string `json:"name"`
	}
	BpayGroupCreateResponse struct {
		BpayResponse
	}
	BpayGroupEditRequest struct {
		Name string `json:"name"`
	}
	BpayGroupEditResponse struct {
		BpayResponse
	}

	BpayGroupListRequest struct {
		PageNo  int               `json:"pageNo"`
		PerPage int               `json:"perPage"`
		Sort    string            `json:"sort"`
		FIlter  []BpayGroupFilter `json:"filter"`
	}
	BpayGroupFilter struct {
		FieldName string `json:"fieldName"`
		Operation string `json:"operation"`
		Value     string `json:"value"`
		FieldType string `json:"fieldType"`
	}

	BpayGroupListResponse struct {
		BpayResponse
		Data []BpayGroupData `json:"data"`
	}
	BpayGroupData struct {
		ID         int64  `json:"id"`
		Name       string `json:"name"`
		CustomerID int64  `json:"customerId"`
	}
	BpayGroupAddBillsRequest struct {
		BillIds []int64 `json:"billIds"`
	}
	BpayGroupAddBillsResponse struct {
		BpayResponse
	}

	BpayGroupBillsResponse struct {
		BpayResponse
		Data []BpayBillData `json:"data"`
	}

	// Find request and response
	BpayFindResponse struct {
		BpayResponse
		Data []BpayFindData `json:"data"`
	}
	BpayFindData struct {
		ID          int64          `json:"id"`
		Name        string         `json:"name"`
		Code        string         `json:"code"`
		TotalAmount int            `json:"totalAmount"`
		ProviderID  int64          `json:"providerId"`
		BIlls       []BpayBillData `json:"bills"`
	}
	BpayBillData struct {
		ID          int64  `json:"id"`
		BillID      string `json:"billId"`
		Code        string `json:"code"`        // Хэрэглэгчийн CID код
		BillAmount  int    `json:"billAmount"`  // Төлөх дүн
		LossAmount  int    `json:"lossAmount"`  // Алдангийн дүн
		TotalAmount int    `json:"totalAmount"` // Нэхэмжилсэн дүн
		PaidAmount  int    `json:"paidAmount"`  // Төлбөл зохих дүн
		Year        int    `json:"year"`
		Month       int    `json:"month"`
		Name        string `json:"name"`
		OrgTypeID   int64  `json:"orgTypeId"`
		OrgName     string `json:"orgName"`
		ProviderID  int64  `json:"providerId"`
		CustomerID  int64  `json:"customerId"`
		StatusID    int64  `json:"statusId"`
	}

	// Invoice request and response
	BpayInvoiceCreateRequest struct {
		BillIds []int `json:"billIds"`
	}
	BpayInvoiceResponse struct {
		BpayResponse
		ID          int64          `json:"id"`
		TotalAmount int            `json:"totalAmount"`
		CustomerID  int            `json:"customerId"`
		StatusID    int            `json:"statusId"`
		BIlls       []BpayBillData `json:"bills"`
	}

	BpayInvoiceTransactionCreateRequest struct {
		InvoiceID int    `json:"invoiceId"`
		IsOrg     bool   `json:"isOrg"`
		VatInfo   string `json:"vatInfo"` // Company Register
	}
	BpayInvoiceTransactionCreateResponse struct {
		BpayResponse
		InvoiceID    string        `json:"invoiceId"`
		QrText       string        `json:"qr_text"`
		QrImage      string        `json:"qr_image"`
		QpayShrotUrl string        `json:"qPay_shortUrl"`
		Urls         []BpayUrlData `json:"urls"`
	}
	BpayUrlData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Logo        string `json:"logo"`
		Link        string `json:"link"`
	}
	BpayResponse struct {
		ResponseCode bool   `json:"responseCode"`
		ResponseMsg  string `json:"responseMsg"`
	}
)

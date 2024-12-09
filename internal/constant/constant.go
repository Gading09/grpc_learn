package constant

const (
	ACCEPT             = "Accept"
	XAPPNAME           = "X-Appname"
	XAPIKEY            = "X-Api-Key"
	CLIENTSIGNATURE    = "Client-Signature"
	AUTHORIZATION      = "Authorization"
	XCTO               = "X-Content-Type-Options"
	XCTO_VALUE         = "nosniff"
	HSTS               = "Strict-Transport-Security"
	HSTS_VALUE         = "max-age=31536000; includeSubDomains"
	CC                 = "Cache-Control"
	CC_VALUE           = "no-store"
	ACAO               = "Access-Control-Allow-Origin"
	ACAO_VALUE         = "*"
	ACAM               = "Access-Control-Allow-Methods"
	ACAM_VALUE         = "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS"
	ACAH               = "Access-Control-Allow-Headers"
	ACAH_VALUE         = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Origin, Cookie, X-Appname, X-Api-Key, Signature, Grpc-Metadata-Signature, Timestamp, Grpc-Metadata-Timestamp, Grpc-Metadata-Client, Grpc-Metadata-Secret, Grpc-Metadata-Device, Client-Signature"
	ACAC               = "Access-Control-Allow-Credentials"
	ACAC_VALUE         = "false"
	GRPC_METHOD        = "Grpc-Metadata-Method"
	GRPC_HEADER_CLIENT = "client"
	GRPC_HEADER_SECRET = "secret"

	// GENERATE DOCX TO PDF
	DIRDOC            = "standardDocuments"
	OUTDIR            = "--outdir"
	LOWRITER          = "lowriter"
	INVISIBLE         = "--invisible" // This command is optional, it will help to disable the splash screen of LibreOffice.
	CONVERT_TO        = "--convert-to"
	PDF_WRITER_EXPORT = "pdf:writer_pdf_Export"
	PK                = "standardDocuments/PK SME"
	APLIKASI          = "standardDocuments/APLIKASI SME"

	// MAX 10 MB REQUEST AND RESPONSE GRPC
	MAX_SIZE_GRPC = 1024 * 1024 * 50

	// URI USER MAGIC
	URI = "https://mf-sme-flamingo.ddb.dev.bri.co.id/?token="
	//URI = "https://mf-sme-flamingo.ddb.dev.bri.co.id"

	// HTTP
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_JSON = "application/json"

	IdentifierId = "identifierId"

	XREQUESTUSER_BRISPOT = "1101"

	APP_NAME      = "DELIMA"
	TELLER_ID     = "891"
	SUPERVISOR_ID = "892"
	CHANNEL_ID    = "DLMA"

	SERVICE_ID_SAVING_TRANSFER            = "00015"
	SERVICE_ID_LOAN_TRANSFER              = "000KB"
	SERVICE_ID_MULTI_CIF_LOAN             = "000BI"
	SERVICE_ID_INQUIRY_LOAN               = "00004"
	SERVICE_ID_INQUIRY_CIF_BY_ACCOUNT     = "0008L"
	SERVICE_ID_HOLD_SAVING                = "0002S"
	SERVICE_ID_CASA_INFORMATION           = "0008O"
	SERVICE_ID_PAYMENT_7Y                 = "000M5"
	SERVICE_ID_INQUIRY_PN_BY_LOAN_ACCOUNT = "000LG"

	INVOICE_TYPE_AP = "AP"
	INVOICE_TYPE_AR = "AR"

	PAYMENT_METHOD_DF = "DF"
	PAYMENT_METHOD_VF = "VF"

	TRANSCODE_MBASE_TRANSACTION = "4106"

	INDEX_OF_INTEREST_BILLING_DATE = 15

	CURRENCY           = "IDR"
	TRX_TYPE_FEE_BASED = "F"

	ACCOUNT_TYPE_L = "L"
)

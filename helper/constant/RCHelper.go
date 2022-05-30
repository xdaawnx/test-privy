package constant

const Success = "200"
const SuccessCreate = "201"
const Paymentprocess = "001"
const Decryptfailed = "100"
const Invalidparentformat = "101"
const Invalidchildformat = "102"
const Dataexpired = "103"
const Encryptfailed = "104"
const Invalidformat = "108"
const Invalidipaddress = "109"
const Invalidaccesstoken = "116"
const Incorrectcredential = "117"
const Paymentfailed = "121"
const InvalidAmountRequest = "122"
const Numberonly = "124"
const Responsecannotbeparsed = "333"
const Badrequest = "400"
const Authenticationfailed = "401"
const Nodata = "404"
const Notallowed = "405"
const Validationfailed = "422"
const Validateparamfailed = "423"
const Internalerror = "500"

const Requestfailed = "600"
const InvalidVaType = "667"
const LogoutFailed = "999"

const InvalidRequestList = "110"
const InvalidRequestCreate = "111"
const InvalidRequestUpdate = "112"
const InvalidRequestDetail = "113"
const FailedCreateData = "221"
const EmptyRequest = "222"
const NoDataToUpdate = "223"
const NoDataInvoice = "224"
const FailedUpdateData = "225"
const CurrentDate = "123"

const IncorrectUsernamePassword = "888"

// Error struct
type Error struct {
	Attribute *string `json:"attribute"`
	Message   string  `json:"message"`
	error
}

// RC struct for get value from db
type RC struct {
	Code        string `gorm:"column:code"`
	Message     string `gorm:"column:message"`
	MessageID   string `gorm:"column:messageID"`
	Description string `gorm:"column:description"`
}

// TableName is for declaring value table name
func (RC) TableName() string {
	return "ref_error_codes"
}

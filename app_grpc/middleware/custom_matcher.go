package middleware

import (
	"fmt"
	"grpc/internal/constant"
	"regexp"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func CustomMatcher(key string) (string, bool) {
	switch key {
	case constant.ACCEPT:
		return key, true
	case constant.XAPPNAME:
		return key, true
	case constant.XAPIKEY:
		return key, true
	case constant.AUTHORIZATION:
		return key, true
	case constant.CLIENTSIGNATURE:
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func ValidateRequestBody(bodyString string) (string, error) {
	// validate if thre is some special character
	positiveinput := regexp.MustCompile(`^[a-zA-Z0-9_.,:/()\t\n\v\f\r]*$`).MatchString
	inputan := positiveinput(bodyString)
	if !strings.Contains(bodyString, "http://") && !strings.Contains(bodyString, "https://") && !strings.Contains(bodyString, "://") && !strings.Contains(bodyString, "./") {
		if strings.Contains(bodyString, "data:image/jpeg;base64") || strings.Contains(bodyString, "data:image/pjpeg;base64") || strings.Contains(bodyString, "data:image/png;base64") || strings.Contains(bodyString, "data:application/pdf;base64") {
			return bodyString, nil
		} else if !inputan {
			output := Replaceinputhelp(bodyString)
			return output, nil
		} else {
			return bodyString, nil
		}
	} else {
		return "", fmt.Errorf("body contain unknown type")
	}
}

func Replaceinputhelp(input string) (output string) {
	request := "{\"body\":" + input + "}"
	// replace := ""
	// iteration := -1
	// symbol := map[int]string{
	//	1:  "%",
	//	2:  "*",
	//	3:  "#",
	//	4:  "!",
	//	5:  "^",
	//	6:  "<",
	//	7:  ">",
	//	8:  "?",
	//	9:  "&",
	//	10: "$",
	//	11: ";",
	// }
	//
	// replacedstring1 := strings.Replace(input, symbol[2], replace, iteration)
	// replacedstring2 := strings.Replace(replacedstring1, symbol[3], replace, iteration)
	// replacedstring3 := strings.Replace(replacedstring2, symbol[4], replace, iteration)
	// replacedstring4 := strings.Replace(replacedstring3, symbol[5], replace, iteration)
	// replacedstring5 := strings.Replace(replacedstring4, symbol[6], replace, iteration)
	// replacedstring6 := strings.Replace(replacedstring5, symbol[7], replace, iteration)
	// replacedstring7 := strings.Replace(replacedstring6, symbol[8], replace, iteration)
	// replacedstring8 := strings.Replace(replacedstring7, symbol[10], replace, iteration)
	// replacedstring9 := strings.Replace(replacedstring8, symbol[11], replace, iteration)
	// output = replacedstring9
	output = request
	return output
}

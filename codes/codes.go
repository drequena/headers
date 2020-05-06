package codes

// Source https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
var validCodes = map[int]bool{
	100: true,
	101: true,
	102: true,
	103: true,
	200: true,
	202: true,
	203: true,
	204: true,
	205: true,
	206: true,
	207: true,
	208: true,
	226: true,
	300: true,
	301: true,
	302: true,
	303: true,
	304: true,
	305: true,
	306: true,
	307: true,
	308: true,
	400: true,
	401: true,
	402: true,
	403: true,
	404: true,
	405: true,
	406: true,
	407: true,
	408: true,
	409: true,
	410: true,
	411: true,
	412: true,
	413: true,
	414: true,
	415: true,
	416: true,
	417: true,
	421: true,
	422: true,
	423: true,
	424: true,
	425: true,
	426: true,
	427: true,
	428: true,
	429: true,
	430: true,
	431: true,
	451: true,
	500: true,
	501: true,
	502: true,
	503: true,
	504: true,
	505: true,
	506: true,
	507: true,
	508: true,
	509: true,
	510: true,
	511: true}

// CheckHTTPCODE check if a int is a valid HTTP STATUS CODE
func CheckHTTPCODE(code int) bool {

	if validCodes[code] {
		return true
	}

	return false
}

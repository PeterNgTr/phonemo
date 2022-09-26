package phonemo

import (
	"math/rand"
	"strconv"
	"time"
)

var supportedCountryList []string = []string{"DE", "SA", "AE", "KW", "BH", "QA", "OM", "VN"}
var countryPhoneCodeMapping = map[string]string{"DE": "+49", "SA": "+966", "AE": "+971", "KW": "+965", "BH": "+973", "QA": "+974", "OM": "+968", "VN": "+84"}
var phoneNumberLengthMapping = map[string]int{"DE": 7, "SA": 7, "AE": 7, "KW": 6, "BH": 5, "QA": 6, "OM": 6, "VN": 7}
var phonePrefixMapping = map[string][]string{
	"DE": {"151", "160", "170", "171", "175", "152", "162", "172", "173", "174", "155", "157", "159", "163", "176", "177", "178", "179"},
	"SA": {"50", "55"},
	"AE": {"50", "55"},
	"KW": {"55", "54", "57"},
	"BH": {"322", "383", "384", "388", "340", "341", "377", "343", "344", "345"},
	"QA": {"33", "44", "55", "66", "77", "31"},
	"OM": {"22", "23", "24", "25", "26", "77", "71", "72", "78", "79", "90", "91", "92", "93", "94", "97", "98", "99"},
	"VN": {"91", "92", "93", "94", "96", "97", "98", "99", "32", "33", "34", "35", "36", "37", "38", "52", "56", "58", "59", "70", "76", "77", "78", "79", "81", "82", "83", "84", "85", "86", "87", "88", "89"},
}

func contains(s [](string), e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func supportedCountry(country string) string {
	if contains(supportedCountryList, country) == true {
		return country
	}
	return "SA"
}

func getRandom(list []string) string {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(list)
	return list[n]
}

func getRandomInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(max-min+1) + min)
}

func generatePhoneNumber(phoneNumberLength int) string {
	phoneNumber := ""
	for i := 0; i < phoneNumberLength; i++ {
		phoneNumber += strconv.FormatInt(int64(getRandomInt(0, 9)), 10)
	}

	return phoneNumber
}

type PhoneNumberGenerator struct {
	Country         string
	WithCountryCode bool
}

func (png *PhoneNumberGenerator) PhoneNumberGenerator() string {
	country := supportedCountry(png.Country)
	phoneNumber := getRandom(phonePrefixMapping[country]) + generatePhoneNumber(phoneNumberLengthMapping[country])
	if png.WithCountryCode == true {
		phoneNumber = countryPhoneCodeMapping[country] + phoneNumber
	}

	return phoneNumber
}

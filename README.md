# phonemo

Creates random mobile phone numbers with a given country id.
Right now, supported countries are Germany(DE), Vietnam(VN), Saudi(SA), UAE(AE), Kuwait(KW), Bahrain(BH), Qatar(QA), Oman(OM). More on the way...

# Installation

```
go get github.com/peterngtr/phonemo
```

# How to use

````
import "github.com/peterngtr/phonemo"

func main() {
  phoneOne := phonemo.PhoneData{Country: "SA", WithCountryCode: false}
	phoneTwo := phonemo.PhoneData{Country: "SA", WithCountryCode: true}
	fmt.Println(phoneOne.PhoneNumberGenerator())
	fmt.Println(phoneTwo.PhoneNumberGenerator())
}


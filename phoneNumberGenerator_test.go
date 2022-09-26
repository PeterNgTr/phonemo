package phonemo

import (
	"strconv"
	"testing"
)

func Test_contains(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Including in array",
			args: args{[]string{"DE", "SA", "AE", "KW", "BH", "QA", "OM", "VN"}, "DE"},
			want: true,
		},
		{
			name: "Not including in array",
			args: args{[]string{"DE", "SA", "AE", "KW", "BH", "QA", "OM", "VN"}, "AB"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_supportedCountry(t *testing.T) {
	type args struct {
		country string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Supporting country",
			args: args{"DE"},
			want: "DE",
		},
		{
			name: "Supporting country",
			args: args{"AB"},
			want: "SA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supportedCountry(tt.args.country); got != tt.want {
				t.Errorf("supportedCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRandom(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Return random item in array",
			args: args{[]string{"DE"}},
			want: "DE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRandom(tt.args.list); got != tt.want {
				t.Errorf("getRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRandomInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Return random int",
			args: args{0, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRandomInt(tt.args.min, tt.args.max); contains([]string{strconv.Itoa(tt.args.min), strconv.Itoa(tt.args.max)}, strconv.Itoa(got)) != true {
				t.Errorf("getRandomInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneNumberGenerator_PhoneNumberGenerator(t *testing.T) {
	type fields struct {
		Country         string
		WithCountryCode bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "Return random number KSA",
			fields: fields{"SA", false},
			want:   9,
		},
		{
			name:   "Return random number KSA with country code",
			fields: fields{"SA", true},
			want:   13,
		},
		{
			name:   "Return random number DE",
			fields: fields{"DE", false},
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			png := &PhoneData{
				Country:         tt.fields.Country,
				WithCountryCode: tt.fields.WithCountryCode,
			}
			if got := png.PhoneNumberGenerator(); len(got) != tt.want {
				t.Errorf("PhoneData.PhoneNumberGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}

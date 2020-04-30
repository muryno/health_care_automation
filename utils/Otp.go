package utils



import (
	"encoding/base32"
	"flag"
	"fmt"
	"strings"

	"github.com/hgfischer/go-otp"
)

const (
	DefaultLength             = 5   // Default length of the generated tokens
	DefaultPeriod             = 4  // Default time period for TOTP tokens, in seconds

)

var (
	secret   = flag.String("secret", "lifetrusty-ng", "Secret key")
	isBase32 = flag.Bool("base32", true, "If true, the secret is interpreted as a Base32 string")
	length   = flag.Uint("length", DefaultLength, "OTP length")
	period   = flag.Uint("period", DefaultPeriod, "Period in seconds")

	)

func GetOtp() string {
	flag.Parse()

	key := *secret
	if !*isBase32 {
		key = base32.StdEncoding.EncodeToString([]byte(*secret))
	}

	key = strings.ToUpper(key)
	if !isGoogleAuthenticatorCompatible(key) {
		fmt.Println("WARN: Google Authenticator requires 16 chars base32 secret, without padding")
	}


	totp := &otp.TOTP{
		Secret:         key,
		Length:         uint8(*length),
		Period:         uint8(*period),
		IsBase32Secret: true,
	}

	otp := totp.Get()
	fmt.Println("TOTP:", otp)

return otp

}

func isGoogleAuthenticatorCompatible(base32Secret string) bool {
	cleaned := strings.Replace(base32Secret, "=", "", -1)
	cleaned = strings.Replace(cleaned, " ", "", -1)
	return len(cleaned) == 16
}

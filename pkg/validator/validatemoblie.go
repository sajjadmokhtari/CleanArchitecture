package validator

import "regexp"

// فقط چک می‌کند: 11 رقم باشد و با 09 شروع شود
var iranMobileRegex = regexp.MustCompile(`^09\d{9}$`)

func IsValidIranianMobile(phone string) bool {
    return iranMobileRegex.MatchString(phone)
}

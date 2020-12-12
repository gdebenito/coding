package validation

// ValidatePassport ...
// ...
func ValidatePassport(passport map[string]string) bool {

	return checkBirthday(passport) &&
		checkYear(passport) &&
		checkExpirationYear(passport) &&
		checkHeight(passport) &&
		checkHairColor(passport) &&
		checkEyeColor(passport) &&
		checkPassportID(passport)
}

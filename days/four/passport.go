package four

import (
	"regexp"
	"strconv"
)

func validatePassportFields(fields map[string]bool) bool {
	requiredFields := [7]string{"byr", "ecl", "eyr", "hcl", "hgt", "iyr", "pid"}
	valid := true
	for _, key := range requiredFields {
		_, exists := fields[key]
		if !exists {
			valid = false
			break
		}
	}
	return valid
}

func validatePassport(fields map[string]string) bool {
	requiredFields := map[string]func(string) bool{
		"byr": func(v string) bool { return validateYear(v, 1920, 2002) },
		"ecl": validateEyeColour,
		"eyr": func(v string) bool { return validateYear(v, 2020, 2030) },
		"hcl": validateHairColour,
		"hgt": validateHeight,
		"iyr": func(v string) bool { return validateYear(v, 2010, 2020) },
		"pid": validatePassportNumber,
	}
	valid := true
	for key, validator := range requiredFields {
		_, exists := fields[key]
		if !exists {
			valid = false
		} else {
			valid = validator(fields[key])
		}
		if !valid {
			break
		}
	}
	return valid
}

func validateYear(year string, min, max int) bool {
	yearI, _ := strconv.Atoi(year)
	return yearI <= max && yearI >= min
}

func validateHeight(h string) bool {
	re := regexp.MustCompile("([0-9]+)(cm|in)")
	if !re.MatchString(h) {
		return false
	}
	submatches := re.FindStringSubmatch(h)
	height, _ := strconv.Atoi(submatches[1])
	if submatches[2] == "cm" {
		return height >= 150 && height <= 193
	}
	return height >= 59 && height <= 76
}

func validateHairColour(c string) bool {
	re := regexp.MustCompile("^#[0-9a-f]{6}$")
	return re.MatchString(c)
}

func validateEyeColour(c string) bool {
	re := regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$")
	return re.MatchString(c)
}

func validatePassportNumber(n string) bool {
	re := regexp.MustCompile("^[0-9]{9}$")
	return re.MatchString(n)
}

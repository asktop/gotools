package acast

// ToStrIfaceMap casts an interface to a map[string]interface{} type.
func ToStrIfaceMap(i interface{}) map[string]interface{} {
	v, _ := ToStrIfaceMapE(i)
	return v
}

// ToStrBoolMap casts an interface to a map[string]bool type.
func ToStrBoolMap(i interface{}) map[string]bool {
	v, _ := ToStrBoolMapE(i)
	return v
}

// ToStrIntMap casts an interface to a map[string]int type.
func ToStrIntMap(i interface{}) map[string]int {
	v, _ := ToStrIntMapE(i)
	return v
}

// ToStrInt64Map casts an interface to a map[string]int64 type.
func ToStrInt64Map(i interface{}) map[string]int64 {
	v, _ := ToStrInt64MapE(i)
	return v
}

// ToStrStrMap casts an interface to a map[string]string type.
func ToStrStrMap(i interface{}) map[string]string {
	v, _ := ToStrStrMapE(i)
	return v
}

// ToStrStrsMap casts an interface to a map[string][]string type.
func ToStrStrsMap(i interface{}) map[string][]string {
	v, _ := ToStrStrsMapE(i)
	return v
}

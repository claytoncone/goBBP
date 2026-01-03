package convert

// --- Byte to String Conversions ---

const hexChars = "0123456789ABCDEF"
const hexCharsLower = "0123456789abcdef"

// HexBytesToString converts raw hex bytes (0-15) to uppercase hex string
func HexBytesToString(hex []byte) string {
	if len(hex) == 0 {
		return ""
	}
	result := make([]byte, len(hex))
	for i, b := range hex {
		result[i] = hexChars[b&0x0F]
	}
	return string(result)
}

// HexBytesToLowerString converts raw hex bytes (0-15) to lowercase hex string
func HexBytesToLowerString(hex []byte) string {
	if len(hex) == 0 {
		return ""
	}
	result := make([]byte, len(hex))
	for i, b := range hex {
		result[i] = hexCharsLower[b&0x0F]
	}
	return string(result)
}

// HexBytesToAlpha converts raw hex bytes (0-15) to ASCII hex bytes ('0'-'9', 'A'-'F')
func HexBytesToAlpha(hex []byte) []byte {
	if len(hex) == 0 {
		return nil
	}
	result := make([]byte, len(hex))
	for i, b := range hex {
		result[i] = hexChars[b&0x0F]
	}
	return result
}

// DecBytesToString converts decimal bytes (0-9) to string
func DecBytesToString(dec []byte) string {
	if len(dec) == 0 {
		return ""
	}
	result := make([]byte, len(dec))
	for i, b := range dec {
		result[i] = '0' + b
	}
	return string(result)
}

// BytesEqual compares two byte slices for equality
func BytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

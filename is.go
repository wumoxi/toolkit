package toolkit

// IsString used to determine whether a value is a string type
func IsString(v interface{}) bool {
	switch v.(type) {
	case string:
		return true
	}
	return false
}

// IsBool Used to determine whether a value is a bool type
func IsBool(v interface{}) bool {
	switch v.(type) {
	case bool:
		return true
	}
	return false
}

// IsFloat32 Used to determine whether a value is a float32 type
func IsFloat32(v interface{}) bool {
	switch v.(type) {
	case float32:
		return true
	}
	return false
}

// IsFloat64 Used to determine whether a value is a float64 type
func IsFloat64(v interface{}) bool {
	switch v.(type) {
	case float64:
		return true
	}
	return false
}

// IsInt Used to determine whether a value is a int type
func IsInt(v interface{}) bool {
	switch v.(type) {
	case int:
		return true
	}
	return false
}

// IsUint Used to determine whether a value is a uint type
func IsUint(v interface{}) bool {
	switch v.(type) {
	case uint:
		return true
	}
	return false
}

// IsInt8 Used to determine whether a value is a int8 type
func IsInt8(v interface{}) bool {
	switch v.(type) {
	case int8:
		return true
	}
	return false
}

// IsInt8 Used to determine whether a value is a uint8 type
func IsUint8(v interface{}) bool {
	switch v.(type) {
	case uint8:
		return true
	}
	return false
}

// IsInt16 Used to determine whether a value is a int16 type
func IsInt16(v interface{}) bool {
	switch v.(type) {
	case int16:
		return true
	}
	return false
}

// IsUint16 Used to determine whether a value is a uint16 type
func IsUint16(v interface{}) bool {
	switch v.(type) {
	case uint16:
		return true
	}
	return false
}

// IsInt32 Used to determine whether a value is a int32 type
func IsInt32(v interface{}) bool {
	switch v.(type) {
	case int32:
		return true
	}
	return false
}

// IsUint32 Used to determine whether a value is a uint32 type
func IsUint32(v interface{}) bool {
	switch v.(type) {
	case uint32:
		return true
	}
	return false
}

// IsInt64 Used to determine whether a value is a int64 type
func IsInt64(v interface{}) bool {
	switch v.(type) {
	case int64:
		return true
	}
	return false
}

// IsUint64 Used to determine whether a value is a uint64 type
func IsUint64(v interface{}) bool {
	switch v.(type) {
	case uint64:
		return true
	}
	return false
}

// IsUintptr Used to determine whether a value is a uintptr type
func IsUintptr(v interface{}) bool {
	switch v.(type) {
	case uintptr:
		return true
	}
	return false
}

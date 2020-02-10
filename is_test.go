package toolkit

import "testing"

func TestIsBool(t *testing.T) {
	if actualTrue := IsBool(true); !actualTrue {
		t.Errorf("Type error: %t\n", actualTrue)
	}

	if actualFalse := IsBool(123); actualFalse {
		t.Errorf("Type error: %t\n", actualFalse)
	}
}

func TestIsInt(t *testing.T) {
	if actualTrue := IsInt(12); !actualTrue {
		t.Errorf("Type error: %t\n", actualTrue)
	}

	if actualFalse := IsInt(true); actualFalse {
		t.Errorf("Type error: %t\n", actualFalse)
	}
}

func TestIsInt8(t *testing.T) {
	var expectedInt8 int8 = 12
	if actualTrue := IsInt8(expectedInt8); !actualTrue {
		t.Errorf("Type error: %t\n", actualTrue)
	}

	if actualFalse := IsInt8(true); actualFalse {
		t.Errorf("Type error: %t\n", actualFalse)
	}
}

func TestIsInt16(t *testing.T) {
	var expectedInt16 int16 = 12
	if actualTrue := IsInt16(expectedInt16); !actualTrue {
		t.Errorf("Type error: %t\n", actualTrue)
	}

	if actualFalse := IsInt16(true); actualFalse {
		t.Errorf("Type error: %t\n", actualFalse)
	}
}

func TestIsInt32(t *testing.T) {
	var expectedInt32 int32 = 12
	if actualTrue := IsInt32(expectedInt32); !actualTrue {
		t.Errorf("Type error: %t\n", actualTrue)
	}

	if actualFalse := IsInt32(true); actualFalse {
		t.Errorf("Type error: %t\n", actualFalse)
	}
}

func TestIsInt64(t *testing.T) {
	var expectedInt64 int64 = 12
	if actualTrue := IsInt64(expectedInt64); !actualTrue {
		t.Errorf("Type error: %t\n", actualTrue)
	}

	if actualFalse := IsInt64(true); actualFalse {
		t.Errorf("Type error: %t\n", actualFalse)
	}
}

func TestIsUint(t *testing.T) {
	var expectedUint uint = 12
	if actualTrue := IsUint(expectedUint); !actualTrue {
		t.Errorf("Type error: %t\n", actualTrue)
	}

	if actualFalse := IsUint(12); actualFalse {
		t.Errorf("Type error: %t\n", actualFalse)
	}
}

func TestIsUint8(t *testing.T) {
	var expectedUint8 uint8 = 12
	if actualTrue := IsUint8(expectedUint8); !actualTrue {
		t.Errorf("Type error: %t\n", actualTrue)
	}

	if actualFalse := IsUint8(12); actualFalse {
		t.Errorf("Type error: %t\n", actualFalse)
	}
}

func TestIsUint16(t *testing.T) {
	var expectedUint16 uint16 = 12
	if actualTrue := IsUint16(expectedUint16); !actualTrue {
		t.Errorf("Type error: %t\n", actualTrue)
	}

	if actualFalse := IsUint16(12); actualFalse {
		t.Errorf("Type error: %t\n", actualFalse)
	}
}

func TestIsUint32(t *testing.T) {
	var expectedUint32 uint32 = 12
	if actualTrue := IsUint32(expectedUint32); !actualTrue {
		t.Errorf("Type error: %t\n", actualTrue)
	}

	if actualFalse := IsUint32(12); actualFalse {
		t.Errorf("Type error: %t\n", actualFalse)
	}
}

func TestIsUint64(t *testing.T) {
	var expectedUint64 uint64 = 12
	if actualTrue := IsUint64(expectedUint64); !actualTrue {
		t.Errorf("Type error: %t\n", actualTrue)
	}

	if actualFalse := IsUint64(12); actualFalse {
		t.Errorf("Type error: %t\n", actualFalse)
	}
}

func TestIsFloat32(t *testing.T) {
	var expectedFloat32 float32 = 3.1415926
	if actualTrue := IsFloat32(expectedFloat32); !actualTrue {
		t.Errorf("Type error: %t\n", actualTrue)
	}

	if actualFalse := IsFloat32(12); actualFalse {
		t.Errorf("Type error: %t\n", actualFalse)
	}
}

func TestIsFloat64(t *testing.T) {
	var expectedFloat64 float64 = 3.1415926
	if actualTrue := IsFloat64(expectedFloat64); !actualTrue {
		t.Errorf("Type error: %t\n", actualTrue)
	}

	if actualFalse := IsFloat64(12); actualFalse {
		t.Errorf("Type error: %t\n", actualFalse)
	}
}

func TestIsUintptr(t *testing.T) {
	var expectedUintptr uintptr = 7788556699
	if actualTrue := IsUintptr(expectedUintptr); !actualTrue {
		t.Errorf("Type error: %t\n", actualTrue)
	}

	if actualFalse := IsUintptr("0x4567888"); actualFalse {
		t.Errorf("Type error: %t\n", actualFalse)
	}
}

package convert

func PtrBool(b bool) *bool {
	return &b
}

func PtrInt32(i int32) *int32 {
	return &i
}

func PtrUint32(i uint32) *uint32 {
	return &i
}

func PtrInt64(i int64) *int64 {
	return &i
}

func PtrUint64(i uint64) *uint64 {
	return &i
}

func PtrString(s string) *string {
	return &s
}

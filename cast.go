package cast

func ToBool(i interface{}) bool {
	v, _ := ToBoolE(i)
	return v
}

package utils

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func ExtractStrArrFromInterface(dataValue interface{}) []string {
	interfaceArrs, ok := dataValue.([]interface{})
	if !ok {
		return nil
	}
	strs := make([]string, 0, len(interfaceArrs))
	for _, arr := range interfaceArrs {
		s, ok1 := arr.(string)
		if !ok1 {
			continue
		}
		strs = append(strs, s)
	}
	return strs
}

func ExtractIntArrFromInterface(dataValue interface{}) []int {
	interfaceArrs, ok := dataValue.([]interface{})
	if !ok {
		return nil
	}
	ints := make([]int, 0, len(interfaceArrs))
	for _, arr := range interfaceArrs {
		s, ok1 := arr.(int)
		if !ok1 {
			continue
		}
		ints = append(ints, s)
	}
	return ints
}

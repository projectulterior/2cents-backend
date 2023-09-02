package utils

// func Union[T comparable](s []T, b []T) []T {
// 	m := make(map[T]bool)

// 	for _, item := range s {
// 		m[item] = true
// 	}

// 	for _, item := range b {
// 		if _, ok := m[item]; !ok {
// 			s = append(s, item)
// 		}
// 	}
// 	return s
// }

func Union(s []string, b []string) []string {
	m := make(map[string]bool)

	for _, item := range s {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			s = append(s, item)
		}
	}
	return s
}

func Intersect(s []string, b []string) []string {
	m := make(map[string]bool)

	for _, item := range s {
		m[item] = true
	}

	var toRet []string

	for _, item := range b {
		if _, ok := m[item]; ok {
			toRet = append(toRet, item)
		}
	}

	return toRet
}

func Differences(s []string, b []string) ([]string, []string) {
	m := make(map[string]bool)

	for _, item := range Intersect(s, b) {
		m[item] = true
	}

	var sDisjoint []string
	var bDisjoint []string

	for _, item := range b {
		if _, ok := m[item]; !ok {
			bDisjoint = append(bDisjoint, item)
		}
	}

	for _, item := range s {
		if _, ok := m[item]; !ok {
			sDisjoint = append(sDisjoint, item)
		}
	}

	return sDisjoint, bDisjoint
}	
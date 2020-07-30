package commons

func Compare(a, b []string) []string {
	for i := len(a) - 1; i >= 0; i-- {
		for _, v := range b {
			if a[i] == v {
				a = append(a[:i], a[i+1:]...)
				break
			}
		}
	}
	return a
}

func Deduplicate(a []string) []string {

	l := []string{}
	m := make(map[string]bool)

	for _, e := range a {
		if _, v := m[e]; !v {
			m[e] = true
			l = append(l, e)
		}
	}

	return l

}

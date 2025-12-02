package time_based_key_value_store

type pair struct {
	value     string
	timestamp int
}

type TimeMap struct {
	m map[string][]pair
}

func Constructor() TimeMap {
	m := make(map[string][]pair)
	return TimeMap{m}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.m[key]; !ok {
		this.m[key] = make([]pair, 0)
	}

	this.m[key] = append(this.m[key], pair{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.m[key]; !ok {
		return ""
	}

	vals := this.m[key]

	l := 0
	r := len(vals) - 1
	i := 0

	for l <= r {
		mid := (l + r) / 2
		if timestamp > vals[mid].timestamp {
			l = mid + 1
			i = mid
		} else if timestamp < vals[mid].timestamp {
			r = mid - 1
		} else {
			i = mid
			break
		}
	}

	if vals[i].timestamp > timestamp {
		return ""
	} else {
		return vals[i].value
	}
}

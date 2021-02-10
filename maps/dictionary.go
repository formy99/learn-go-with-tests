package maps

type Dictionary map[string]string

func Search(d Dictionary, key string) (value string) {
	return d[key]
}

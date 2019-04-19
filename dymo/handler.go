package function

import (
	"strings"
)

func reverse(domain string) string {

	if !strings.Contains(domain, ".") {
		return domain
	}

	withoutScheme := strings.TrimPrefix(strings.TrimPrefix(domain, "http://"), "https://")
	parts := strings.Split(withoutScheme, "/")
	elems := strings.Split(parts[0], ".")

	for i := len(elems)/2 - 1; i >= 0; i-- {
		opp := len(elems) - 1 - i
		elems[i], elems[opp] = elems[opp], elems[i]
	}

	return strings.Join(elems, ".")

}

// Handle a serverless request
func Handle(req []byte) string {

	return reverse(string(req))

}

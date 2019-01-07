package function

import (
	"fmt"
	"strings"
)

// Handle a serverless request
func Handle(req []byte) string {

	elems := strings.Split(string(req), ".")

	for i := len(elems)/2 - 1; i >= 0; i-- {
		opp := len(elems) - 1 - i
		elems[i], elems[opp] = elems[opp], elems[i]
	}
	return fmt.Sprintf("%s", strings.Join(elems, "."))
}

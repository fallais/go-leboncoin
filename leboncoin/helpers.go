package leboncoin

// contains returns true if slice `s` contains element `e`. False otherwise.
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

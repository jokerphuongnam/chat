package utils

func ArrayExists[T comparable](a T, list []T) bool {
	for _, b := range list {
        if a == b {
            return true
        }
    }
    return false
}
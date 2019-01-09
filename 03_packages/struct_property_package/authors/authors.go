package authors

import "fmt"

// AuthorInfo struct
type AuthorInfo struct {
	Name, Email string // Public
	age         int    // Private
}

// GetAuthor Func
func (a AuthorInfo) GetAuthor() string {
	return fmt.Sprintf("*** Author Info ***\n Name: %s\n Email: %s\n", a.Name, a.Email)
}

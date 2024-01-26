package masker_test

import (
	"github/Igo87/links/masker"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMasksLInks(t *testing.T) {
	/* tabular testing */
	testCases := []struct {
		name     string
		s        string
		url      string
		expected string
	}{
		{"success", "https://www.google.com", "https://", "https://**************"},
		{"successToo", "hello world", "https://", "hello world"},
		// {"fail", "https://www.google.com", "https://", "https://www.google.com"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, masker.GetMasksLInks(tc.s, tc.url))

		})
	}

}

func FuzzGetMasksLInks(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string, url string) {
		masker.GetMasksLInks(s, url)
	})
}

package stringutil

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/mail"
	"strings"
)

// HashMailboxName accepts a mailbox name and hashes it.  filestore uses this as
// the directory to house the mailbox
func HashMailboxName(mailbox string) string {
	h := sha1.New()
	if _, err := io.WriteString(h, mailbox); err != nil {
		// This shouldn't ever happen
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// StringAddressList converts a list of addresses to a list of strings
func StringAddressList(addrs []*mail.Address) []string {
	s := make([]string, len(addrs))
	for i, a := range addrs {
		if a != nil {
			s[i] = a.String()
		}
	}
	return s
}

// SliceContains returns true if s is present in slice.
func SliceContains(slice []string, s string) bool {
	for _, v := range slice {
		if s == v {
			return true
		}
	}
	return false
}

// SliceToLower lowercases the contents of slice of strings.
func SliceToLower(slice []string) {
	for i, s := range slice {
		slice[i] = strings.ToLower(s)
	}
}

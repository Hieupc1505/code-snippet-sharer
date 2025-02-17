package snippet

import (
	"errors"
	"strings"
)

type Snippet string

func (sn Snippet) String() string { return string(sn) }

// NewSnippet tạo một Snippet mới với validation
func NewSnippet(sn string) (Snippet, error) {
	// 1. Kiểm tra nếu snippet trống
	if strings.TrimSpace(sn) == "" {
		return "", errors.New("snippet không được để trống")
	}

	// 2. Kiểm tra độ dài tối đa (giả sử 10,000 ký tự)
	const maxLength = 10000
	if len(sn) > maxLength {
		return "", errors.New("snippet quá dài, tối đa 10,000 ký tự")
	}

	// 3. Kiểm tra ký tự nguy hiểm (đơn giản)
	blockedWords := []string{"<script>", "DROP TABLE", "DELETE FROM"}
	for _, word := range blockedWords {
		if strings.Contains(strings.ToLower(sn), strings.ToLower(word)) {
			return "", errors.New("snippet chứa nội dung không hợp lệ")
		}
	}

	return Snippet(sn), nil
}

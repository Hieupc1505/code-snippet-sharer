package templateutil

// NamedEndpoints lưu trữ danh sách các endpoint có tên
type NamedEndpoints struct {
	endpoints map[string]string
}

// NewNamedEndpoints khởi tạo NamedEndpoints
func NewNamedEndpoints() NamedEndpoints {
	return NamedEndpoints{
		endpoints: make(map[string]string),
	}
}

// Add thêm một endpoint với tên và URL
func (ne NamedEndpoints) Add(name, path string) {
	ne.endpoints[name] = path
}

// Get lấy đường dẫn của endpoint theo tên
func (ne NamedEndpoints) Get(name string) string {
	if path, exists := ne.endpoints[name]; exists {
		return path
	}
	return "/" // Trả về trang chủ nếu không tìm thấy
}

// GetNamedEndpoints khởi tạo danh sách các endpoint có tên
func GetNamedEndpoints() NamedEndpoints {
	ne := NewNamedEndpoints()
	ne.Add("home", "/")
	ne.Add("login", "/login")
	ne.Add("dashboard", "/dashboard")
	ne.Add("profile", "/profile")
	return ne
}

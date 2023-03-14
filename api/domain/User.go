package domain

// User 用户，目前有且只有一个用户，用户名/密码：admin/P@$$w0rd!@#3166
type User struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

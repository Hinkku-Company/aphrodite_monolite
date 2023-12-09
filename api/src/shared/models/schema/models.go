package schema

type Access struct {
	Token        string `json:"token"`
	TokenRefresh string `json:"tokenRefresh"`
}

type AccessToken struct {
	Token        string `json:"token"`
	TokenRefresh string `json:"tokenRefresh"`
}

type Credentials struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Response struct {
	Status *bool `json:"status,omitempty"`
}

package piholeapi

const (
	AuthEndpoint	= "/auth"
)

type LoginRequest struct {
	Password string `json:"password"`
}

type LoginSession struct {
        Valid bool `json:"valid"`
        Totp bool `json:"totp"`
        Sid string `json:"sid"`
        Csrf string `json:"csrf"`
        Validity int `json:"validity"`
        Message string `json:"message"` 
}

type LoginResponse struct {
        Session LoginSession `json:"session"`
        Took float64 `json:"took"`
}


func (c *Client) Auth(password string) error {
	var LoginResp LoginResponse

	err := c.PostRequest(
		AuthEndpoint,
		LoginRequest{
			Password: password,
		},
		&LoginResp,
	)
	if err != nil {
		return err
	}

	c.Sid = LoginResp.Session.Sid

	return nil
}

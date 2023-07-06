package youtube

import "github.com/circle2jt/go-castv2/primitives"

//GetScreenIDRequest for getting a screen ID for an existing youtube application.
type GetScreenIDRequest struct {
	primitives.PayloadHeaders
	ScreenID int `json:"screen_ids"`
}

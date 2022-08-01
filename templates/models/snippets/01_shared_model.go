package snippets

import ()

type Shared struct {
	ID       int32  `json:"id"`
	Uid      string `json:"uid"`
	Created  int    `json:"created"`
	Updated  int    `json:"updated"`
	Archived *int   `json:"archived"`
}

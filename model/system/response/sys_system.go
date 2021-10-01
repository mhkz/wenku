package response

import "wenku/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}

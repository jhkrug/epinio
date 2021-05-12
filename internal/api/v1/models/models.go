// Package models contains structures to encapsulate the requested and
// response data used by the communication between cli and api server.
package models

type ServiceResponse struct {
	Name      string   `json:"name"`
	BoundApps []string `json:"boundapps"`
}

type ServiceResponseList []ServiceResponse

type CatalogCreateRequest struct {
	Name             string            `json:"name"`
	Class            string            `json:"class"`
	Plan             string            `json:"plan"`
	Data             map[string]string `json:"data"`
	WaitForProvision bool              `json:"waitforprovision"`
}

type CustomCreateRequest struct {
	Name string            `json:"name"`
	Data map[string]string `json:"data"`
}

type DeleteRequest struct {
	Unbind bool `json:"unbind"`
}

type DeleteResponse struct {
	BoundApps []string `json:"boundapps"`
}

type BindRequest struct {
	Names []string `json:"names"`
}

type UpdateAppRequest struct {
	Instances string `json:"instances"`
}

// TODO: CreateOrgRequest

// UploadRequest is a multipart form

type UploadResponse struct {
	Image ImageRef `json:"image,omitempty"`
	Git   *GitRef  `json:"git,omitempty"`
	Route string   `json:"route,omitempty"`
}

type StageRequest struct {
	App   AppRef   `json:"app,omitempty"`
	Image ImageRef `json:"image,omitempty"`
	Git   *GitRef  `json:"git,omitempty"`
}

type StageResponse struct {
	Stage StageRef `json:"stage,omitempty"`
}

package permission

type permission struct {
	Action string `json:"action"`
	Object string `json:"object"`
}

type roles struct {
	SuperAdmin []permission `json:"super_admin"`
	Admin      []permission `json:"admin"`
	Reporter   []permission `json:"reporter"`
	User       []permission `json:"user"`
}

type authorizationData struct {
	Roles roles `json:"roles"`
}

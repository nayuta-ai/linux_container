package main

type Config struct {
	OciVersion string  `json:"ociVersion"`
	Process    Process `json:"process"`
	Root       Root    `json:"root"`
	HostName   string  `json:"hostname"`
	Mounts     []Mount `json:"mounts"`
	Linux      Linux   `json:"linux"`
}

type Process struct {
	Terminal        bool         `json:"terminal"`
	User            User         `json:"user"`
	Args            []string     `json:"args"`
	Env             []string     `json:"env"`
	Cwd             string       `json:"cwd"`
	Capabilities    Capabilities `json:"capabilities"`
	Rlimits         []Rlimit     `json:"rlimits"`
	NoNewPrivileges bool         `json:"noNewPrivileges"`
}

type User struct {
	Uid int `json:"uid"`
	Gid int `json:"gid"`
}

type Capabilities struct {
	Bounding  []string `json:"bounding"`
	Effective []string `json:"effective"`
	Permitted []string `json:"permitted"`
	Ambient   []string `json:"ambient"`
}

type Rlimit struct {
	Type string `json:"type"`
	Hard int    `json:"hard"`
	Soft int    `json:"soft"`
}

type Root struct {
	Path     string `json:"path"`
	ReadOnly bool   `json:"readonly"`
}

type Mount struct {
	Destination string   `json:"destination"`
	Type        string   `json:"type"`
	Source      string   `json:"source"`
	Options     []string `json:"options"`
}

type Linux struct {
	Resources     Resources   `json:"resources"`
	Namespaces    []Namespace `json:"namespaces"`
	MaskedPaths   []string    `json:"maskedPaths"`
	ReadonlyPaths []string    `json:"readonlypaths"`
}

type Resources struct {
	Devices []Device `json:"devices"`
}

type Device struct {
	Allow  bool   `json:"allow"`
	Access string `json:"access"`
}

type Namespace struct {
	Type string `json:"type"`
}

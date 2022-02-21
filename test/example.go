package test

type configuration struct { //nolint:deadcode,unused
	needHelp     bool   //nolint:structcheck
	url          string `yaml:"url"`            //nolint:structcheck
	maxDepth     uint64 `yaml:"max_depth"`      //nolint:structcheck
	timeout      int    `yaml:"timeout"`        //nolint:structcheck
	depthIncStep int    `yaml:"depth_inc_step"` //nolint:structcheck
	output       string `yaml:"output"`         //nolint:structcheck
	jsonLog      bool   `yaml:"json_log"`       //nolint:structcheck
	jsonLogNext  bool   `yaml:"json_log"`       //nolint:structcheck
	withPanic    bool   `yaml:"with_panic"`     //nolint:structcheck
}

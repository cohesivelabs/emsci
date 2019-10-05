package parser

type Command struct {
	Image      string
	WorkingDir string
	Entrypoint string
	Command    string
	Args       []string
}

type ContainerCommand struct {
	Command string
	Args    []string
}

type Dep struct {
	Commands []Command
	Image    string
}

type CIConfig struct {
	Test       []ContainerCommand
	Env        []string
	Deps       []Dep
	PreStart   []Command
	SharedDeps []Dep
}

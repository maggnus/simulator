package api

type Command string

const (
	StartCommand   Command = "start"
	StopCommand    Command = "stop"
	LeftCommand    Command = "left"
	RightCommand   Command = "right"
	ForwardCommand Command = "forward"
	BackCommand    Command = "back"
)

func (c *Command) Get() []Command {
	return []Command{
		StartCommand,
		StopCommand,
		LeftCommand,
		RightCommand,
		ForwardCommand,
		BackCommand,
	}
}

func (c *Command) GetAll() []string {
	var arr []string

	for _, cmd := range c.Get() {
		arr = append(arr, string(cmd))
	}

	return arr
}

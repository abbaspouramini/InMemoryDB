package commandline_handler

type CommandFactory interface {
	RecognizeCommand(input string)
}

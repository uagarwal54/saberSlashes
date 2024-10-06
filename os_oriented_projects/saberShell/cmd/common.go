package cmd

type (
	// Command struct will contain all the functions and data for working with the commands being executed by the user
	Command struct {
		Name           string
		Args           []string
		IsBackground   string
		InputRedirect  string
		OutputRedirect string
	}
)

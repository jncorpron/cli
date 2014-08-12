package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shell struct {
	App    *App
	Prompt func() string
}

func NewShell(app *App) *Shell {
	return &Shell{App: app}
}

func (s *Shell) Run() {
	bufStdin := bufio.NewReader(os.Stdin)
	for {
		if s.Prompt != nil {
			prompt := s.Prompt()
			fmt.Print(prompt)
		} else {
			fmt.Printf("%s>", s.App.Name)
		}

		// read a line from input
		line, err := bufStdin.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading stdin: ", err)
			return
		}

		// trim the newline at the end
		line = line[:len(line)-1]

		// split the line into arguments
		args := splitArgs(line)

		// append the app name onto the current arguments
		args = append([]string{s.App.Name}, args...)

		// run the app with these arguments
		s.App.Run(args)
	}
}

func splitArgs(line string) []string {
	return strings.Split(line, " ")
}

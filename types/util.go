package types

import (
	"fmt"
	"os"
	"strings"
)

func FakeHandlerCommand(command string, args ...string) *Handler {
	cs := []string{os.Args[0], "-test.run=TestHelperHandlerProcess", "--", command}
	cs = append(cs, args...)
	cmdStr := strings.Join(cs, " ")
	trimmedCmd := strings.Trim(cmdStr, " ")
	env := "GO_WANT_HELPER_HANDLER_PROCESS=1"

	handler := &Handler{
		Command: fmt.Sprintf("%s %s", env, trimmedCmd),
	}

	return handler
}

func FakeMutatorCommand(command string, args ...string) *Mutator {
	cs := []string{os.Args[0], "-test.run=TestHelperMutatorProcess", "--", command}
	cs = append(cs, args...)
	cmdStr := strings.Join(cs, " ")
	trimmedCmd := strings.Trim(cmdStr, " ")
	env := "GO_WANT_HELPER_MUTATOR_PROCESS=1"

	mutator := &Mutator{
		Command: fmt.Sprintf("%s %s", env, trimmedCmd),
	}

	return mutator
}

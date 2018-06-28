package selenium

import (
	"os/exec"
	"strconv"
)

type server struct {
	cmd *exec.Cmd
}

func Server(path string, port int, arguments ...string) (*server, error) {
	s := &server{exec.Command("java", "-jar", path, "-port", strconv.Itoa(port))}
	s.cmd.Args = append(s.cmd.Args, arguments...)
	return s, nil
}

func (server *server) SetServerExitListener(listener chan error) {

	go func() { listener <- server.cmd.Wait() }()

}

func (s *server) Start() error {

	stdout, err := s.cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := s.cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := s.cmd.Start(); err != nil {
		return err
	}

	go func() { Log(stdout) }()
	go func() { Log(stderr) }()

	return nil

}

package exec

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/fatih/color"
)

var (
	nilWrite  io.Writer
	nilReader io.Reader
)

type Printf = func(format string, a ...interface{})
type Option struct {
	PWD    string
	Envs   []string
	Input  io.Reader
	Output io.Writer
	Errput io.Writer
}

func ToCommand(cmd string) *exec.Cmd {
	cmds := strings.Fields(cmd)
	c := exec.Command(cmds[0], cmds[1:]...)
	return c
}

func MustSHRun(cmdStr string, envs ...string) {
	cmd := exec.Command("sh", "-c", cmdStr)
	RunCmdWithOutput(cmd, envs...)
}

func Run(cmd string, envs ...string) (string, error) {
	c := ToCommand(cmd)
	color.Green(c.String())
	env := os.Environ()
	if len(envs) > 0 {
		env = append(env, envs...)
	}
	c.Env = env
	res, err := c.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s %w", string(res), err)
	}
	return string(res), nil
}

func RunWithOutput(cmd string, envs ...string) error {
	c := ToCommand(cmd)
	return RunCmdWithOutput(c, envs...)
}

func RunCmdWithOutput(cmd *exec.Cmd, envs ...string) error {
	env := os.Environ()
	if len(envs) > 0 {
		env = append(env, envs...)
	}
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	io.WriteString(cmd.Stdout, cmd.String())
	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func runCmdWithOutput(c *exec.Cmd) error {
	color.Green(c.String())
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}

	if err = c.Start(); err != nil {
		return err
	}

	wait := &sync.WaitGroup{}
	wait.Add(3)
	flag := false

	go readPipi(wait, stdout, color.Green)
	go func() {
		flag = readPipi(wait, stderr, color.Red)
	}()

	go func() {
		c.Wait()
		wait.Done()
	}()
	wait.Wait()
	if flag {
		panic("")
	}
	return nil
}
func MustRun(cmd string, envs ...string) {
	res, err := Run(cmd, envs...)
	CheckErr(err)
	color.Green(string(res))
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readPipi(wait *sync.WaitGroup, reader io.Reader, p Printf) bool {
	defer func() {
		wait.Done()
	}()
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	flag := false
	for scanner.Scan() {
		m := scanner.Text()
		p(m)
		flag = true
	}
	return flag
}

func readPipi2(wait *sync.WaitGroup, reader io.Reader, p Printf) bool {
	defer func() {
		wait.Done()
	}()
	flag := false
	for {
		tmp := make([]byte, 1024)
		_, err := reader.Read(tmp)
		p(string(tmp))
		if err != nil {
			color.Red("read std err fail, %s", err.Error())
			break
		}
		flag = true
	}
	return flag
}

package containerd

import "golang.org/x/net/context"

type ContainerInfo struct {
	ID      string
	Runtime string
}

type Container interface {
	// Information of the container
	Info() ContainerInfo
	// Start the container's user defined process
	Start(context.Context) error
	// State returns the container's state
	State(context.Context) (State, error)
	// Pause pauses the container process
	Pause(context.Context) error
	// Resume unpauses the container process
	Resume(context.Context) error
	// Kill signals a container
	Kill(context.Context, uint32, bool) error
	// Exec adds a process into the container
	Exec(context.Context, ExecOpts) (Process, error)
	// Pty resizes the processes pty/console
	Pty(context.Context, uint32, ConsoleSize) error
	// CloseStdin closes the processes stdin
	CloseStdin(context.Context, uint32) error
}

type LinuxContainer interface {
	Container
}

type ExecOpts struct {
	Spec []byte
	IO   IO
}

type Process interface {
	// State returns the process state
	State(context.Context) (State, error)
	// Kill signals a container
	Kill(context.Context, uint32, bool) error
}

type ConsoleSize struct {
	Width  uint32
	Height uint32
}

type Status int

const (
	CreatedStatus Status = iota + 1
	RunningStatus
	StoppedStatus
	DeletedStatus
	PausedStatus
)

type State interface {
	// Status is the current status of the container
	Status() Status
	// Pid is the main process id for the container
	Pid() uint32
}

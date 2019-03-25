package info

const data = "test-content"

type Info struct {
	Metadata string
}

func DefaultInfo() *Info {
	return &Info{Metadata: data}
}

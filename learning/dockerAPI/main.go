package main

import (
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {

	ctx := context.Background()

	//cli, err := client.NewEnvClient()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}

	imageName := "hello-world"

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		fmt.Println("Error pulling image")
		panic(err)
	}
	io.Copy(os.Stdout, out)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
	}, nil, nil, "")

	if err != nil {
		panic(err)
	}

	err1 := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(resp.ID)

}

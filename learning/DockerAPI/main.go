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
	defer out.Close()
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

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("List of existing containers is...")
	for _, container := range containers {
		fmt.Printf("Container ID=%s Container Status=%s Container Label=%s Container Name=%s\n", container.ID, container.Status, container.Labels, container.Names)
	}

	// List images

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}

	ilr, err := cli.ImageLoad(ctx, "/infoblox/data/config", true)
	if err != nil {
		panic(err)
	}

	if images == true {
		fmt.Println("Image load successful")
	}

}

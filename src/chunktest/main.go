package main

import "kserver/world"

func main()  {
	world := new(world.World)
	world.Init()
	world.Destroy()
}

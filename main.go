// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// An app that draws a green triangle on a red background.
package main

import (
	"time"

	"github.com/tbruyelle/fsm"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/app/debug"
	"golang.org/x/mobile/event"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/gl"
	"golang.org/x/mobile/sprite"
	"golang.org/x/mobile/sprite/clock"
	"golang.org/x/mobile/sprite/glsprite"
	_ "image/png"
)

var (
	startTime = time.Now()
	lastClock = clock.Time(-1)

	eng              = glsprite.Engine()
	scene            *fsm.Object
	tiles            map[rune]sprite.SubTex
	player           *fsm.Object
	screenW, screenH float32
	hexSprite        sprite.SubTex
)

func main() {
	app.Run(app.Callbacks{
		//Start: start,
		Stop:  stop,
		Draw:  draw,
		Touch: touch,
	})
}

func start() {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	scene = &fsm.Object{Width: 1, Height: 1}
	scene.Register(nil, eng)

	// Screen dimensions
	screenW, screenH = float32(geom.Width), float32(geom.Height)

	// background
	bg := &fsm.Object{Width: screenW, Height: screenH}
	bg.Register(scene, eng)
	bg.Sprite = fsm.SubTex(fsm.MustLoadTexture(eng, "bg0.png"), 0, 0, 1920, 1080)

	// the hex sprite
	hexSprite = fsm.SubTex(fsm.MustLoadTexture(eng, "hex-point-topped.png"), 0, 0, 183, 200)

	NewBoard()
}

func stop() {
}

func touch(t event.Touch) {
	if t.Type == event.TouchEnd {
		// Handle click
	}
}

func draw() {
	// Keep until golang.org/x/mogile/x11.go handle Start callback
	if scene == nil {
		start()
	}
	now := clock.Time(time.Since(startTime) * 60 / time.Second)
	if now == lastClock {
		// TODO: figure out how to limit draw callbacks to 60Hz instead of
		// burning the CPU as fast as possible.
		// TODO: (relatedly??) sync to vblank?
		return
	}
	lastClock = now

	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// test collisions

	eng.Render(scene.Node, now)
	debug.DrawFPS()
}

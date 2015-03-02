// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// An app that draws a green triangle on a red background.
package main

import (
	"encoding/binary"
	"time"

	"github.com/tbruyelle/fsm"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/app/debug"
	"golang.org/x/mobile/event"
	"golang.org/x/mobile/f32"
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

var triangleData = f32.Bytes(binary.LittleEndian,
	0.0, 0.4, 0.0, // top left
	0.0, 0.0, 0.0, // bottom left
	0.4, 0.0, 0.0, // bottom right
)

const (
	coordsPerVertex = 3
	vertexCount     = 3
)

const vertexShader = `#version 100
uniform vec2 offset;

attribute vec4 position;
void main() {
	// offset comes in with x/y values between 0 and 1.
	// position bounds are -1 to 1.
	vec4 offset4 = vec4(2.0*offset.x-1.0, 1.0-2.0*offset.y, 0, 0);
	gl_Position = position + offset4;
}`

const fragmentShader = `#version 100
precision mediump float;
uniform vec4 color;
void main() {
	gl_FragColor = color;
}`

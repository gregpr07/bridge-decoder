package tracer

import (
	"github.com/gregpr07/bridge-decoder/tracer/assets"
)

type TracerType int8

const (
	JSTracer TracerType = iota
	NativeTracer
)

const NativeTracerContent = "@raw_tracer"

var JSTracerContent string

func init() {
	JSTracerContent = string(assets.MustAsset("call_tracer.js"))
}

type TracerConfig struct {
	ty TracerType
}

func JSTracerConfig() TracerConfig {
	return TracerConfig{JSTracer}
}

func NativeTracerConfig() TracerConfig {
	return TracerConfig{NativeTracer}
}

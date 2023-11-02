module gowas/example/uilabels

go 1.21.1

replace github.com/rocco-gossmann/GoWas => ../../

replace github.com/rocco-gossmann/GoWas/examples/ui-labels => ./

require github.com/rocco-gossmann/GoWas v0.0.0-00010101000000-000000000000

require (
	github.com/rocco-gossmann/go_throwable v0.0.0-20230909155809-0ca53d5db119 // indirect
	github.com/rocco-gossmann/go_wasmcanvas v1.1.1 // indirect
)

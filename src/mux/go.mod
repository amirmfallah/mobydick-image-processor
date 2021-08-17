module image-processor-dokkaan.ir/mux

go 1.16

require image-processor-dokkaan.ir/muxlib v0.0.0

replace (
	image-processor-dokkaan.ir/lib => ../lib
	image-processor-dokkaan.ir/muxlib => ./muxlib
)

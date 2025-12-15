module main

go 1.22.3

replace todos => ./todos

require todos v0.0.0-00010101000000-000000000000

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/xeonx/timeago v1.0.0-rc5 // indirect
)

module verify

go 1.24.0

toolchain go1.24.11

require github.com/openai/openai-go v1.12.0

require backend/common v1.0.0

replace backend/common => ../common

require (
	github.com/tidwall/gjson v1.18.0 // indirect
	github.com/tidwall/match v1.2.0 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
)

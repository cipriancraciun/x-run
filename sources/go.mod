module github.com/cipriancraciun/z-run

go 1.17

require (
	github.com/Pallinder/go-randomdata v1.2.0 // indirect
	github.com/colinmarc/cdb v0.0.0-20190223170904-60f317823f70
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eiannone/keyboard v0.0.0-20200508000154-caf4b762e807
	github.com/junegunn/fzf v0.0.0-20210815070326-3c804bcfec7c
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.13
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mattn/go-shellwords v1.0.12 // indirect
	github.com/peterh/liner v1.2.1
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/vbauerster/mpb/v5 v5.4.0
	go.starlark.net v0.0.0-20210602144842-1cdb82c9e17a
	golang.org/x/crypto v0.0.0-20210813211128-0a44fdfbc16e // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210816183151-1e6c022a8912
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

require (
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d // indirect
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/gdamore/tcell v1.4.0 // indirect
	github.com/saracen/walker v0.1.2 // indirect
)

replace (
	github.com/colinmarc/cdb => github.com/cipriancraciun/go-cdb-lib v0.0.0-20190809203657-d959ce9cc674
	github.com/junegunn/fzf => github.com/cipriancraciun/fzf v0.0.0-20210213184424-8610a98b3b9b
)

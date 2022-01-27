module carefreex

go 1.16

replace (
	github.com/carefreex-io/config v0.0.0-20220124114250-915f601e9d3e => /Users/weijiexu/Work/framwork/config
	github.com/carefreex-io/logger v0.0.0-20220126142718-64cf9688a43e => /Users/weijiexu/Work/framwork/logger
	github.com/carefreex-io/rpcxserver v0.0.0-20220124131535-4239a9958640 => /Users/weijiexu/Work/framwork/rpcxserver
	github.com/carefreex-io/rpcxclient v0.0.0-20220125015040-a1c90cf18fa0 => /Users/weijiexu/Work/framwork/rpcxclient
)

require (
	github.com/carefreex-io/config v0.0.0-20220124114250-915f601e9d3e
	github.com/carefreex-io/dbdao v0.0.0-20220127091322-0875875b30cb
	github.com/carefreex-io/example v0.0.0-20220127101006-d3dfc81952ab
	github.com/carefreex-io/logger v0.0.0-20220126142718-64cf9688a43e
	github.com/carefreex-io/rpcxclient v0.0.0-20220125015040-a1c90cf18fa0
	github.com/carefreex-io/rpcxserver v0.0.0-20220124131535-4239a9958640
	gorm.io/gorm v1.22.5
)

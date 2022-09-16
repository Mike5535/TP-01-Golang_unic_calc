module mod

go 1.18

require (
	github.com/mike5535/uniq v0.0.0
)

replace github.com/mike5535/uniq => ../uniq

replace github.com/mike5535/uniq/cut_string => ../uniq/cut_string

replace github.com/mike5535/uniq/uniq_types => ../uniq/uniq_types

replace github.com/mike5535/uniq/uniq_read => ../uniq/uniq_read

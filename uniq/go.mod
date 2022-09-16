module github.com/mike5535/uniq

go 1.18

require (
    github.com/mike5535/uniq/uniq_types v0.0.0
    github.com/mike5535/uniq/cut_string v0.0.0
    github.com/mike5535/uniq/uniq_read  v0.0.0
)

replace github.com/mike5535/uniq/uniq_types => ./uniq_types

replace github.com/mike5535/uniq/cut_string => ./cut_string

replace github.com/mike5535/uniq/uniq_read => ./uniq_read

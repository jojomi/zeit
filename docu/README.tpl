# zeit

`zeit` does contain a few time abstractions that might be useful in other projects.

[![Go Reference](https://pkg.go.dev/badge/github.com/jojomi/zeit.svg)](https://pkg.go.dev/github.com/jojomi/zeit)

## Key structs

{{ $zeitDir := exec "ff -f zeit" -}}
{{ $godocBaseLink := "https://pkg.go.dev/github.com/jojomi/zeit" -}}
{{ $structs := makeStringList "Date" "Time" "Duration" "Month" "TimeRange" -}}
{{ range $structs.All -}}
* **[ {{- . -}} ]( {{- $godocBaseLink -}} # {{- . -}} )**: {{ exec (printf "gocumentation %s zeit %s" $zeitDir .) }} (full source: [{{- . | toSnakeCase -}}.go]( {{- . | toSnakeCase -}}.go))
{{ end }}

## Tests

{{ $command := "go test -count=1 -v ./..." -}}

``` shell
{{ $command -}}
{{- newline -}}
{{ execWd "go test -count=1 ./..." ".." }}
```

<details>
  <summary>All test results</summary>

  ```
  {{ execWd $command ".." | trim }}
  ```
</details>

### Test Coverage

```
{{ $command := "go test -cover -count=1 ./..." -}}
{{ $command }}
{{ execWd $command ".." | trim }}
```

## Latest changes

See the (https://github.com/jojomi/zeit/commits/master)[commits on master].

## Why the name?

"Zeit" is German for "time".

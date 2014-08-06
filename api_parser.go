package api_parser

import (
  "errors"
  "os/exec"
  "strings"
)

const (
  Json = "json"
  Yaml = "yaml"
)

func Parse(s, s_type string) ([]byte, error) {
  path, err := exec.LookPath("snowcrash")
  if err != nil {
    return nil, errors.New("Couldn't find snowcrash. Please install it first https://github.com/apiaryio/snowcrash")
  }

  cmd := exec.Command(path, "--format", s_type)
  a := strings.NewReader(s)
  cmd.Stdin = a
  return cmd.Output()
}
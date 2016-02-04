package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type config struct {
	repo *string
	version *string
	status *string
	description *string
	provider *string
	box *string
	checksumType *string
	checksum *string
}

func (c *config) addflags(s *flag.FlagSet) {
	repo := s.String("repo", "", "path to metadata.json")
	c.repo = repo

	version := s.String("version", "", "version to add (SemVer)")
	c.version = version

	status := s.String("status", "", "status of the version")
	c.status = status

	description := s.String("description", "", "description of the version (markdown)")
	s.StringVar(description, "desc", "", "Shortcut for --description")
	c.description = description

	provider := s.String("provider", "", "provider the box supports")
	c.provider = provider

	box := s.String("box", "", "url of the box")
	c.box = box

	checksumType := s.String("checksum-type", "", "type of the checksum")
	c.checksumType = checksumType

	checksum := s.String("checksum", "", "checksum of the box")
	c.checksum = checksum
}

func usage(s *flag.FlagSet) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <command> [flags]\n", filepath.Base(os.Args[0]))
		s.PrintDefaults()
	}
}

func parseFlags(s *flag.FlagSet, args []string) *config {
	c := &config{}
	c.addflags(s)
	s.Usage = usage(s)
	s.Parse(args)
	return c
}

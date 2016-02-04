package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"encoding/json"

	"github.com/golang/glog"
)

func main() {
	flagSet := flag.CommandLine
	c := parseFlags(flagSet, os.Args[1:])

	defer glog.Flush()
	glog.V(1).Info("Executing: ", strings.Join(os.Args, " "))

	// non-flag args
	args := flagSet.Args()
	if len(args) == 0 {
		fmt.Fprint(os.Stderr, "Error: No command specified - Expected \"add\"\n")
		flagSet.Usage()
		os.Exit(2)
	}

	if len(args) > 1 {
		fmt.Fprint(os.Stderr, "Error: Too many arguments specified: %v\n", args)
		flagSet.Usage()
		os.Exit(2)
	}

//	fmt.Fprintf(os.Stderr, "repo: %s\n", *c.repo)
//	fmt.Fprintf(os.Stderr, "version: %s\n", *c.version)
//	fmt.Fprintf(os.Stderr, "description: %s\n", *c.description)
//	fmt.Fprintf(os.Stderr, "checksum: %s\n", *c.checksum)

	cmdArg := args[0]

	var err error
	switch cmdArg {
	case "add":
		err = add(c)
	default:
		fmt.Fprint(os.Stderr, "Error: Unknown command - Expected \"add\"\n")
		flagSet.Usage()
		os.Exit(2)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func add(c *config) error {

	// check required fields
	missingFields := []string{}
	if c.repo == nil || *c.repo == "" {
		missingFields = append(missingFields, "repo")
	}
	if c.version == nil || *c.version == "" {
		missingFields = append(missingFields, "version")
	}
	if c.status == nil || *c.status == "" {
		missingFields = append(missingFields, "status")
	}
	if c.description == nil || *c.description == "" {
		missingFields = append(missingFields, "description")
	}
	if c.provider == nil || *c.provider == "" {
		missingFields = append(missingFields, "provider")
	}
	if c.box == nil || *c.box == "" {
		missingFields = append(missingFields, "box")
	}
	if c.checksumType == nil || *c.checksumType == "" {
		missingFields = append(missingFields, "checksum-type")
	}
	if c.checksum == nil || *c.checksum == "" {
		missingFields = append(missingFields, "checksum")
	}
	if len(missingFields) > 0 {
		return fmt.Errorf("Missing config fields: %s", missingFields)
	}

	// read repo file
	file, err := ioutil.ReadFile(*c.repo)
	if err != nil {
		return fmt.Errorf("Reading repo file: %v", err)
	}

	// parse repo json
	var repo Metadata
	err = json.Unmarshal(file, &repo)
	if err != nil {
		return fmt.Errorf("Unmarshalling repo json: %v", err)
	}

	// create new version record
	version := Version{
		Version: *c.version,
		Status: *c.status,
		//DescriptionHTML: c.description,
		DescriptionMarkdown: *c.description,
		Providers: []Provider{
			{
				Name: *c.provider,
				URL: *c.box,
				ChecksumType: *c.checksumType,
				Checksum: *c.checksum,
			},
		},
	}

	// prepend version record into version list
	repo.Versions = append([]Version{version}, repo.Versions...)

	// format repo json
	repoBytes, err := json.MarshalIndent(repo, "", "  ")
	if err != nil {
		return fmt.Errorf("Marshalling repo json: %v", err)
	}

	// write repo file
	err = ioutil.WriteFile(*c.repo, repoBytes, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Writing repo file: %v", err)
	}

	return nil
}
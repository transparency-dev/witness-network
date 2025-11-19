// Copyright 2025 The Witness Network authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// validate is a tool to check a log list file for correctness.
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"iter"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/mod/sumdb/note"
)

func init() {
	flag.Var(&logLists, "loglist", "Path or glob to one or more log list files. Can be provided multiple times to specify individual files or multiple globs for convenience.")
}

var (
	logLists multiStringFlag
)

func main() {
	flag.Parse()

	fail := false

	for _, n := range mustResolveFiles() {
		log.Printf("Examining %q...", n)
		p, err := os.ReadFile(n)
		if err != nil {
			log.Printf("  ❌ Failed to read file contents: %v", err)
			fail = true
			continue
		}
		if err := validateLogList(p); err != nil {
			log.Printf("  ❌ Failed to parse contents of file: %v", err)
			fail = true
			continue
		}
		log.Printf("  ✅ File appears valid")
	}
	if fail {
		log.Printf("❌ One or more files had severe errors.")
		os.Exit(-1)
	}
	log.Print("✅ All inspected files appear valid")
}

// These iotas represent a simple state machine for parsing the log configs.
const (
	logStateUnknown = "unknown"
	logStateVkey    = "vkey"
	logStateOrigin  = "origin"
	logStateQPD     = "qpd"
	logStateContact = "contact"

	logHeaderV0 = "logs/v0"
)

// validateLogList checks that the bytes provided are a valid loglist, returning an error otherwise.
func validateLogList(b []byte) error {
	expectHeader := true
	expect := logStateVkey
	s := bufio.NewScanner(bytes.NewReader(b))

	for line, num := range lineIter(s) {
		if expectHeader {
			if line != logHeaderV0 {
				return fmt.Errorf("line %d: expected %s header, but found %q", num, logHeaderV0, line)
			}
			expectHeader = false
			expect = logStateVkey
			continue
		}

		bits := strings.Split(line, " ")
		// Empty lines should never be returned by the iterator
		if len(bits) == 0 {
			panic("encountered empty line, this shouldn't happen!")
		}

		switch kw, params := bits[0], bits[1:]; kw {
		case logStateVkey:
			if expect != logStateVkey {
				return fmt.Errorf("line %d: found vkey keyword, but expected %q", num, expect)
			}
			if err := expectParts(logStateVkey, len(params), 1); err != nil {
				return fmt.Errorf("line %d: %v", num, err)
			}
			if _, err := note.NewVerifier(params[0]); err != nil {
				return fmt.Errorf("line %d: invalid vkey - %v", num, err)
			}
			expect = logStateOrigin
			continue
		case logStateOrigin:
			if expect != logStateOrigin {
				return fmt.Errorf("line %d: found origin keyword, but expcted %q", num, expect)
			}
			if len(params) == 0 {
				return fmt.Errorf("line: %d: expected origin parameter, but none found", num)
			}
			expect = logStateQPD
			continue
		case logStateQPD:
			// Origin is optional
			if expect == logStateOrigin {
				expect = logStateQPD
			}
			if expect != logStateQPD {
				return fmt.Errorf("line %d: found qpd keyword, but expected %q", num, expect)
			}
			if err := expectParts(logStateQPD, len(params), 1); err != nil {
				return fmt.Errorf("line %d: %v", num, err)
			}
			if v, err := strconv.ParseInt(params[0], 10, 64); err != nil {
				return fmt.Errorf("line %d: invalid QPD value - %v", num, err)
			} else if v <= 0 {
				return fmt.Errorf("line %d: invalid QPD value <= 0", num)
			}
			expect = logStateContact
			continue
		case logStateContact:
			if expect != logStateContact {
				return fmt.Errorf("line %d: found contact keyword, but expected %q", num, expect)
			}
			if len(params) == 0 {
				return fmt.Errorf("line %d: no contact details provided", num)
			}
			expect = logStateVkey
			continue
		}
	}
	if err := s.Err(); err != nil {
		return err
	}
	if expect != logStateVkey {
		return fmt.Errorf("incomplete log config at end of file")
	}
	return nil
}

// expectParts returns a descriptive error if got != want.
// Used above to return helpful errors if an incorrect number of tokens are present on a log config line.
func expectParts(k string, got, want int) error {
	if got < want {
		return fmt.Errorf("missing parameter to %s keyword", k)
	} else if got > want {
		return fmt.Errorf("too many parameters to %s keyword", k)
	}
	return nil
}

// lineIter returns an iterator over filtered lines from the provided bytes.
//
// The filtering removes whitespace from all lines, and will skip the line entirely if it is either
// empty, or contains only a comment.
func lineIter(s *bufio.Scanner) iter.Seq2[string, int] {
	lineNum := 0
	return func(yield func(string, int) bool) {
		for s.Scan() {
			line := s.Text()
			lineNum++

			line = strings.TrimSpace(line)
			// Annoyingly, it seems we can't return nil from stripSplit when the whole line is a comment or the scanner
			// simply gives up, so handle empty lines here instead.
			if line == "" {
				continue
			}
			if line[0] == '#' {
				continue
			}
			if !yield(line, lineNum) {
				return
			}
		}
	}
}

// mustResolveFiles returns a list of filepaths matching the provided loglist glob flags, or an error if no paths match.
func mustResolveFiles() []string {
	n := []string{}
	for _, g := range logLists {
		m, err := filepath.Glob(g)
		if err != nil {
			log.Printf("Failed to glob pattern %q: %v", g, err)
			os.Exit(1)
		}
		n = append(n, m...)
	}
	if len(n) == 0 {
		log.Print("No log lists found.")
		os.Exit(1)
	}
	return n
}

// multiStringFlag allows a flag to be specified multiple times on the command
// line, and stores all of these values.
type multiStringFlag []string

func (ms *multiStringFlag) String() string {
	return strings.Join(*ms, ",")
}

func (ms *multiStringFlag) Set(w string) error {
	*ms = append(*ms, w)
	return nil
}

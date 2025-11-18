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

	f_note "github.com/transparency-dev/formats/note"
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
	logStateVkey    = iota
	logStateOrigin  = iota
	logStateQPD     = iota
	logStateContact = iota

	logHeaderV0 = "logs/v0"
)

// validateLogList checks that the bytes provided are a valid loglist, returning an error otherwise.
func validateLogList(b []byte) error {
	expectHeader := true
	logState := logStateVkey
	s := bufio.NewScanner(bytes.NewReader(b))
	s.Split(stripSplit)

	for line, num := range lineIter(s) {
		if expectHeader {
			if line != logHeaderV0 {
				return fmt.Errorf("line %d: expected %s header, but found %q", num, logHeaderV0, line)
			}
			expectHeader = false
			continue
		}

		bits := strings.Split(line, " ")
		// Empty lines should never be returned by the iterator
		if len(bits) == 0 {
			panic("encountered empty line, this shouldn't happen!")
		}
		switch logState {
		case logStateVkey:
			if bits[0] != "vkey" {
				return fmt.Errorf("line %d: expected vkey keyword, but found %q", num, bits[0])
			}
			if err := expectParts("vkey", len(bits), 2); err != nil {
				return fmt.Errorf("line %d: %v", num, err)
			}
			if _, err := f_note.NewVerifier(bits[1]); err != nil {
				return fmt.Errorf("line %d: invalid vkey - %v", num, err)
			}
			logState = logStateOrigin
			continue
		case logStateOrigin:
			if bits[0] == "qpd" {
				goto originNotPresent
			}
			if bits[0] != "origin" {
				return fmt.Errorf("line %d: expected origin or qpd keyword, but found %q", num, bits[0])
			}
			if len(bits) == 1 {
				return fmt.Errorf("line: %d: expected origin parameter, but none found", num)
			}
			logState = logStateQPD
			continue
		originNotPresent:
			fallthrough
		case logStateQPD:
			if bits[0] != "qpd" {
				return fmt.Errorf("line %d: expected qpd keyword, but found %q", num, bits[0])
			}
			if err := expectParts("qpd", len(bits), 2); err != nil {
				return fmt.Errorf("line %d: %v", num, err)
			}
			if v, err := strconv.ParseInt(bits[1], 10, 64); err != nil {
				return fmt.Errorf("line %d: invalid QPD value - %v", num, err)
			} else if v <= 0 {
				return fmt.Errorf("line %d: invalid QPD value <= 0", num)
			}
			logState = logStateContact
			continue
		case logStateContact:
			if bits[0] != "contact" {
				return fmt.Errorf("line %d: expected contact keyword, but found %q", num, bits[0])
			}
			if len(bits) == 1 {
				return fmt.Errorf("line %d: no contact details provided", num)
			}
			logState = logStateVkey
			continue
		}
	}
	if err := s.Err(); err != nil {
		return err
	}
	if logState != logStateVkey {
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

			// Annoyingly, it seems we can't return nil from stripSplit when the whole line is a comment or the scanner
			// simply gives up, so handle empty lines here instead.
			if line == "" {
				continue
			}
			if !yield(line, lineNum) {
				return
			}
		}
	}
}

// stripSplit is a bufio.SplitFunc which strips out leading/trailing whitespace and comments.
func stripSplit(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, filterLine(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), filterLine(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// filterLine drops a terminal \r from the data and removes surrounding whitespace and comments.
func filterLine(line []byte) []byte {
	if len(line) > 0 && line[len(line)-1] == '\r' {
		line = line[0 : len(line)-1]
	}

	line = bytes.TrimSpace(line)
	if len(line) > 0 && line[0] == '#' {
		return []byte{}
	}
	return line
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

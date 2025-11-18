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

package main

import (
	"testing"
)

func TestValidateLogList(t *testing.T) {
	for _, test := range []struct {
		name    string
		logList string
		wantErr bool
	}{
		{
			name:    "valid empty",
			logList: "logs/v0",
		}, {
			name: "valid with logs and comments",
			logList: `
			# These are some comments
			####### more comments
			# 
			# There is a URL below which contains a # character too.

			logs/v0

			vkey log.staging.ct.example.com+3af057ed+AcOM/FdR90fZeCLT4OGd4F+RA38KwLzJ1vdJvX+3LMJW
			origin some legacy origin line
			qpd 86400
			contact me@example.com

			vkey log.example.com/v1/tree/4e89cc51651f0d95f3c6127c15e1a42e3ddf7046c5b17b752689c402e773bb4d+d15ef221+AehD64OcCnZ3q4cJrhnAHgjSjoZq3gKPDUGOgkAqokJG
			qpd 8640
			contact logs (at) example.se | https://some.url.example.se/with/a#fragment
			`,
		}, {
			name: "empty ok, ignores comments",
			logList: `
				# Comment
						   # Indented comment
				logs/v0`,
		}, {
			name:    "missing header",
			logList: "something else",
			wantErr: true,
		}, {
			name: "vkey missing",
			logList: `
			logs/v0
			origin example.com/somewhere
			qpd 86400
			contact me@example.com
			`,
			wantErr: true,
		}, {
			name: "vkey no parameter",
			logList: `
			logs/v0
			vkey
			origin example.com/somewhere
			qpd 86400
			contact me@example.com
			`,
			wantErr: true,
		}, {
			name: "vkey too many parameters",
			logList: `
			logs/v0
			vkey log.staging.ct.example.com+3af057ed+AcOM/FdR90fZeCLT4OGd4F+RA38KwLzJ1vdJvX+3LMJW and some bananas too
			origin example.com/somewhere
			qpd 86400
			contact me@example.com
			`,
			wantErr: true,
		}, {
			name: "vkey is not a vkey",
			logList: `
			logs/v0
			vkey bananas
			origin example.com/somewhere
			qpd 86400
			contact me@example.com
			`,
			wantErr: true,
		}, {
			name: "origin no parameter",
			logList: `
			logs/v0
			vkey log.staging.ct.example.com+3af057ed+AcOM/FdR90fZeCLT4OGd4F+RA38KwLzJ1vdJvX+3LMJW
			origin
			qpd 86400
			contact me@example.com
			`,
			wantErr: true,
		}, {
			name: "qpd no parameter",
			logList: `
			logs/v0
			vkey log.staging.ct.example.com+3af057ed+AcOM/FdR90fZeCLT4OGd4F+RA38KwLzJ1vdJvX+3LMJW
			qpd
			contact me@example.com
			`,
			wantErr: true,
		}, {
			name: "qpd too many parameters",
			logList: `
			logs/v0
			vkey log.staging.ct.example.com+3af057ed+AcOM/FdR90fZeCLT4OGd4F+RA38KwLzJ1vdJvX+3LMJW
			qpd 86400 432534
			contact me@example.com
			`,
			wantErr: true,
		}, {
			name: "qpd not a number",
			logList: `
			logs/v0
			vkey log.staging.ct.example.com+3af057ed+AcOM/FdR90fZeCLT4OGd4F+RA38KwLzJ1vdJvX+3LMJW
			qpd pineapple
			contact me@example.com
			`,
			wantErr: true,
		}, {
			name: "qpd not positive",
			logList: `
			logs/v0
			vkey log.staging.ct.example.com+3af057ed+AcOM/FdR90fZeCLT4OGd4F+RA38KwLzJ1vdJvX+3LMJW
			qpd 0
			contact me@example.com
			`,
			wantErr: true,
		}, {
			name: "contact no parameter",
			logList: `
			logs/v0
			vkey log.staging.ct.example.com+3af057ed+AcOM/FdR90fZeCLT4OGd4F+RA38KwLzJ1vdJvX+3LMJW
			qpd 3
			contact
			`,
			wantErr: true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			err := validateLogList([]byte(test.logList))
			if gotErr := err != nil; gotErr != test.wantErr {
				t.Fatalf("Got %v, want error %t", err, test.wantErr)
			}
		})
	}

}

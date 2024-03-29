/*
Copyright 2019 - The TXTDirect Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package minitxtd

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		txtRecord string
		expected  record
		err       error
	}{
		{
			"v=txtv0;to=https://example.com/;code=302",
			record{
				Version: "txtv0",
				To:      "https://example.com/",
				Code:    302,
				Type:    "host",
			},
			nil,
		},
		{
			"v=txtv0;to=https://example.com/",
			record{
				Version: "txtv0",
				To:      "https://example.com/",
				Code:    302,
				Type:    "host",
			},
			nil,
		},
		{
			"v=txtv0;to=https://example.com/;code=302",
			record{
				Version: "txtv0",
				To:      "https://example.com/",
				Code:    302,
				Type:    "host",
			},
			nil,
		},
		{
			"v=txtv0;to=https://example.com/;code=302;vcs=hg;type=gometa",
			record{
				Version: "txtv0",
				To:      "https://example.com/",
				Code:    302,
				Vcs:     "hg",
				Type:    "gometa",
			},
			nil,
		},
		{
			"v=txtv0;to=https://example.com/;code=302;type=gometa;vcs=git",
			record{
				Version: "txtv0",
				To:      "https://example.com/",
				Code:    302,
				Vcs:     "git",
				Type:    "gometa",
			},
			nil,
		},
		{
			"v=txtv0;to=https://example.com/;code=test",
			record{},
			fmt.Errorf("could not parse status code"),
		},
		{
			"v=txtv1;to=https://example.com/;code=test",
			record{},
			fmt.Errorf("unhandled version 'txtv1'"),
		},
		{
			"v=txtv0;https://example.com/",
			record{},
			fmt.Errorf("arbitrary data not allowed"),
		},
		{
			"v=txtv0;to=https://example.com/caddy;type=path;code=302",
			record{
				Version: "txtv0",
				To:      "https://example.com/caddy",
				Type:    "path",
				Code:    302,
			},
			nil,
		},
		{
			"v=txtv0;to=https://example.com/;key=value",
			record{
				Version: "txtv0",
				To:      "https://example.com/",
				Code:    302,
				Type:    "host",
			},
			nil,
		},
		{
			"v=txtv0;to={?url}",
			record{
				Version: "txtv0",
				To:      "https://example.com/testing",
				Code:    302,
				Type:    "host",
			},
			nil,
		},
		{
			"v=txtv0;to={?url};from={method}",
			record{
				Version: "txtv0",
				To:      "https://example.com/testing",
				Code:    302,
				Type:    "host",
				From:    "GET",
			},
			nil,
		},
	}

	for i, test := range tests {
		r := record{}
		c := Config{
			Enable: []string{test.expected.Type},
		}
		req, _ := http.NewRequest("GET", "http://example.com?url=https://example.com/testing", nil)
		err := r.Parse(test.txtRecord, req, c)

		if err != nil {
			if test.err == nil || !strings.HasPrefix(err.Error(), test.err.Error()) {
				t.Errorf("Test %d: Unexpected error: %s", i, err)
			}
			continue
		}
		if err == nil && test.err != nil {
			t.Errorf("Test %d: Expected error, got nil", i)
			continue
		}

		if got, want := r.Version, test.expected.Version; got != want {
			t.Errorf("Test %d: Expected Version to be '%s', got '%s'", i, want, got)
		}
		if got, want := r.To, test.expected.To; got != want {
			t.Errorf("Test %d: Expected To to be '%s', got '%s'", i, want, got)
		}
		if got, want := r.Code, test.expected.Code; got != want {
			t.Errorf("Test %d: Expected Code to be '%d', got '%d'", i, want, got)
		}
		if got, want := r.Type, test.expected.Type; got != want {
			t.Errorf("Test %d: Expected Type to be '%s', got '%s'", i, want, got)
		}
		if got, want := r.Vcs, test.expected.Vcs; got != want {
			t.Errorf("Test %d: Expected Vcs to be '%s', got '%s'", i, want, got)
		}
	}
}

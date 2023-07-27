package main

import (
	"strconv"
	"testing"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		input       []string
		want        string
		omitMmsig   bool
		paramsTypes string
	}{
		{
			input:     []string{`AQEByM06Awa3RMCs_e0J1uBvFQAB`, `AQEByM06Awa3RMCs/e0J1uBvFQAB`},
			omitMmsig: true,
			want:      `{"version":1,"protocol_id":1,"follow_id":"c8cd3a03-06b7-44c0-acfd-ed09d6e06f15","action":1}`,
		},
		{
			input: []string{
				"AQEBecIa9NuuSuqFjx6vl4I8swADAQAIrowoFSlDh7MN7WVBRYfkAAAAAAJPkCcG",
			},
			paramsTypes: `["uuid","string","decimal"]`,
			want:        `{"version":1,"protocol_id":1,"follow_id":"79c21af4-dbae-4aea-858f-1eaf97823cb3","action":3,"mmsig":{"version":1,"members":[],"threshold":0},"params":["uuid:08ae8c28-1529-4387-b30d-ed65414587e4","string:","decimal:99.2478183"]}`,
		},
		{
			input:       []string{"AQQBs7TAnCQhQey41L3t-71YsQABAAAAAQE=", "AQQBs7TAnCQhQey41L3t+71YsQABAAAAAQE="},
			paramsTypes: `["int32","uint8"]`,
			omitMmsig:   true,
			want:        `{"version":1,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":1,"params":["int32:1","uint8:1"]}`,
		},
		{
			input:       []string{"AQQBs7TAnCQhQey41L3t-71YsQACAAAAAQExAQ==", "AQQBs7TAnCQhQey41L3t+71YsQACAAAAAQExAQ=="},
			paramsTypes: `["int32","string","uint8"]`,
			omitMmsig:   true,
			want:        `{"version":1,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":2,"params":["int32:1","string:1","uint8:1"]}`,
		},

		// with checksum
		{
			input:       []string{"AgQBs7TAnCQhQey41L3t-71YsQABAAAAAQE9SmiZ", "AgQBs7TAnCQhQey41L3t+71YsQABAAAAAQE9SmiZ"},
			paramsTypes: `["int32","uint8"]`,
			omitMmsig:   true,
			want:        `{"version":2,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":1,"params":["int32:1","uint8:1"]}`,
		},
		{
			input:       []string{"AgQBs7TAnCQhQey41L3t-71YsQACAAAAAQExAV5Gv0A=", "AgQBs7TAnCQhQey41L3t+71YsQACAAAAAQExAV5Gv0A="},
			paramsTypes: `["int32","string","uint8"]`,
			omitMmsig:   true,
			want:        `{"version":2,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":2,"params":["int32:1","string:1","uint8:1"]}`,
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			for _, str := range c.input {
				got, err := Decode(str, c.omitMmsig, c.paramsTypes)
				if err != nil {
					t.Error(err)
				}
				if got != c.want {
					t.Errorf("Decode(%q) == %q, want %q", c.input, got, c.want)
				}
			}
		})
	}
}

func TestEncode(t *testing.T) {
	cases := []struct {
		input, want string
		b64Method   string
	}{
		{
			input:     `{"version":1,"protocol_id":1,"follow_id":"c8cd3a03-06b7-44c0-acfd-ed09d6e06f15","action":1}`,
			b64Method: "url",
			want:      "AQEByM06Awa3RMCs_e0J1uBvFQAB",
		},
		{
			input:     `{"version":1,"protocol_id":1,"follow_id":"79c21af4-dbae-4aea-858f-1eaf97823cb3","action":3,"mmsig":{"version":1,"members":[],"threshold":0},"params":["uuid:08ae8c28-1529-4387-b30d-ed65414587e4","string:","decimal:99.2478183"]}`,
			b64Method: "url",
			want:      "AQEBecIa9NuuSuqFjx6vl4I8swADAQAIrowoFSlDh7MN7WVBRYfkAAAAAAJPkCcG",
		},
		{
			input:     `{"version":1,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":1,"params":["int32:1","uint8:1"]}`,
			b64Method: "url",
			want:      "AQQBs7TAnCQhQey41L3t-71YsQABAAAAAQE=",
		},
		{
			input:     `{"version":1,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":2,"params":["int32:1","string:1","uint8:1"]}`,
			b64Method: "url",
			want:      "AQQBs7TAnCQhQey41L3t-71YsQACAAAAAQExAQ==",
		},

		// version 2, with checksum
		{
			input:     `{"version":2,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":1,"params":["int32:1","uint8:1"]}`,
			b64Method: "url",
			want:      "AgQBs7TAnCQhQey41L3t-71YsQABAAAAAQE9SmiZ",
		},
		{
			input:     `{"version":2,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":2,"params":["int32:1","string:1","uint8:1"]}`,
			b64Method: "url",
			want:      "AgQBs7TAnCQhQey41L3t-71YsQACAAAAAQExAV5Gv0A=",
		},

		// base64 standard
		{
			input:     `{"version":1,"protocol_id":1,"follow_id":"c8cd3a03-06b7-44c0-acfd-ed09d6e06f15","action":1}`,
			b64Method: "std",
			want:      "AQEByM06Awa3RMCs/e0J1uBvFQAB",
		},
		{
			input:     `{"version":1,"protocol_id":1,"follow_id":"79c21af4-dbae-4aea-858f-1eaf97823cb3","action":3,"mmsig":{"version":1,"members":[],"threshold":0},"params":["uuid:08ae8c28-1529-4387-b30d-ed65414587e4","string:","decimal:99.2478183"]}`,
			b64Method: "std",
			want:      "AQEBecIa9NuuSuqFjx6vl4I8swADAQAIrowoFSlDh7MN7WVBRYfkAAAAAAJPkCcG",
		},
		{
			input:     `{"version":1,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":1,"params":["int32:1","uint8:1"]}`,
			b64Method: "std",
			want:      "AQQBs7TAnCQhQey41L3t+71YsQABAAAAAQE=",
		},
		{
			input:     `{"version":1,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":2,"params":["int32:1","string:1","uint8:1"]}`,
			b64Method: "std",
			want:      "AQQBs7TAnCQhQey41L3t+71YsQACAAAAAQExAQ==",
		},

		// version 2, with checksum
		{
			input:     `{"version":2,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":1,"params":["int32:1","uint8:1"]}`,
			b64Method: "std",
			want:      "AgQBs7TAnCQhQey41L3t+71YsQABAAAAAQE9SmiZ",
		},
		{
			input:     `{"version":2,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":2,"params":["int32:1","string:1","uint8:1"]}`,
			b64Method: "std",
			want:      "AgQBs7TAnCQhQey41L3t+71YsQACAAAAAQExAV5Gv0A=",
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			got, err := Encode(c.input, c.b64Method)
			if err != nil {
				t.Error(err)
			}
			if got != c.want {
				t.Errorf("Encode(%q) == %q, want %q", c.input, got, c.want)
			}
		})
	}
}

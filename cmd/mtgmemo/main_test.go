package main

import (
	"testing"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		input, want string
		omitMmsig   bool
		paramsTypes string
	}{
		{
			input:     `AQEByM06Awa3RMCs/e0J1uBvFQAB`,
			omitMmsig: true,
			want:      `{"version":1,"protocol_id":1,"follow_id":"c8cd3a03-06b7-44c0-acfd-ed09d6e06f15","action":1}`,
		},
		{
			input:       "AQEBecIa9NuuSuqFjx6vl4I8swADAQAIrowoFSlDh7MN7WVBRYfkAAAAAAJPkCcG",
			paramsTypes: `["uuid","string","decimal"]`,
			want:        `{"version":1,"protocol_id":1,"follow_id":"79c21af4-dbae-4aea-858f-1eaf97823cb3","action":3,"mmsig":{"version":1,"members":[],"threshold":0},"params":["uuid:08ae8c28-1529-4387-b30d-ed65414587e4","string:","decimal:99.2478183"]}`,
		},
		{
			input:       "AQQBs7TAnCQhQey41L3t+71YsQABAAAAAQE=",
			paramsTypes: `["int32","uint8"]`,
			omitMmsig:   true,
			want:        `{"version":1,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":1,"params":["int32:1","uint8:1"]}`,
		},
		{
			input:       "AQQBs7TAnCQhQey41L3t+71YsQACAAAAAQExAQ==",
			paramsTypes: `["int32","string","uint8"]`,
			omitMmsig:   true,
			want:        `{"version":1,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":2,"params":["int32:1","string:1","uint8:1"]}`,
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			got, err := Decode(c.input, c.omitMmsig, c.paramsTypes)
			if err != nil {
				t.Error(err)
			}
			if got != c.want {
				t.Errorf("Decode(%q) == %q, want %q", c.input, got, c.want)
			}
		})
	}

}

func TestEncode(t *testing.T) {
	cases := []struct {
		input, want string
	}{
		{
			input: `{"version":1,"protocol_id":1,"follow_id":"c8cd3a03-06b7-44c0-acfd-ed09d6e06f15","action":1}`,
			want:  "AQEByM06Awa3RMCs/e0J1uBvFQAB",
		},
		{
			input: `{"version":1,"protocol_id":1,"follow_id":"79c21af4-dbae-4aea-858f-1eaf97823cb3","action":3,"mmsig":{"version":1,"members":[],"threshold":0},"params":["uuid:08ae8c28-1529-4387-b30d-ed65414587e4","string:","decimal:99.2478183"]}`,
			want:  "AQEBecIa9NuuSuqFjx6vl4I8swADAQAIrowoFSlDh7MN7WVBRYfkAAAAAAJPkCcG",
		},
		{
			input: `{"version":1,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":1,"params":["int32:1","uint8:1"]}`,
			want:  "AQQBs7TAnCQhQey41L3t+71YsQABAAAAAQE=",
		},
		{
			input: `{"version":1,"protocol_id":4,"follow_id":"b3b4c09c-2421-41ec-b8d4-bdedfbbd58b1","action":2,"params":["int32:1","string:1","uint8:1"]}`,
			want:  "AQQBs7TAnCQhQey41L3t+71YsQACAAAAAQExAQ==",
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			got, err := Encode(c.input)
			if err != nil {
				t.Error(err)
			}
			if got != c.want {
				t.Errorf("Encode(%q) == %q, want %q", c.input, got, c.want)
			}
		})
	}
}

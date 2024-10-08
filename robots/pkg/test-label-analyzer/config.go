/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2023 Red Hat, Inc.
 */

package test_label_analyzer

import (
	"regexp"
	"strings"
)

// A LabelCategory defines a category of tests that share a common label either in their test name or as a Ginkgo label
type LabelCategory struct {
	Name            string  `json:"name"`
	TestNameLabelRE *Regexp `json:"test_name_label_re"`
	GinkgoLabelRE   *Regexp `json:"ginkgo_label_re"`
}

func (c *LabelCategory) String() string {
	return c.Name
}

// Config defines the configuration file structure that is required to map tests to categories.
type Config struct {
	Categories []LabelCategory `json:"categories"`
}

func (c *Config) String() string {
	var s []string
	for _, cat := range c.Categories {
		s = append(s, cat.String())
	}
	return strings.Join(s, ", ")
}

// Regexp adds unmarshalling from json for regexp.Regexp
type Regexp struct {
	*regexp.Regexp
}

func NewRegexp(value string) *Regexp {
	return &Regexp{Regexp: regexp.MustCompile(value)}
}

// UnmarshalText unmarshals json into a regexp.Regexp
func (r *Regexp) UnmarshalText(b []byte) error {
	regex, err := regexp.Compile(string(b))
	if err != nil {
		return err
	}

	r.Regexp = regex

	return nil
}

// MarshalText marshals regexp.Regexp as string
func (r *Regexp) MarshalText() ([]byte, error) {
	if r.Regexp != nil {
		return []byte(r.Regexp.String()), nil
	}

	return nil, nil
}
func NewQuarantineDefaultConfig() *Config {
	return &Config{
		Categories: []LabelCategory{
			{
				Name:            "Quarantine",
				TestNameLabelRE: NewRegexp("\\[QUARANTINE\\]"),
				GinkgoLabelRE:   NewRegexp("Quarantine"),
			},
		},
	}
}

func NewTestNameDefaultConfig(partialTestName string) *Config {
	return &Config{
		Categories: []LabelCategory{
			{
				Name:            "PartialTestNameCategory",
				TestNameLabelRE: NewRegexp(partialTestName),
			},
		},
	}
}

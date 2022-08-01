package utils

import (
	"api-builder/utils/name"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// TODO: lifted from old generator. Verify its not shit later.
func FileVersion(path string) (version int, err error) {

	var versions []int

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return version, err
	}

	if len(files) == 1 {
		return 1, nil
	}

	for i := 0; i < len(files); i++ {

		if files[i].Name() == ".gitkeep" {
			continue
		}

		if (i % 2) == 0 {
			version, err := strconv.Atoi(strings.Trim(strings.Split(strings.Split(files[i].Name(), "_")[0], ".")[0], "0"))
			if err != nil {
				return version, err
			}

			versions = append(versions, version)
		}

	}

	sort.Ints(versions)

	return len(versions) + 1, nil
}

// For you viewing pleasure:
// https://stackoverflow.com/questions/64158699/how-do-i-remove-duplicates-from-a-struct-slice
func DeduplicateNames(names []name.Name) (deduplicated []name.Name) {

	m := map[name.Name]struct{}{}

	for _, name := range names {
		if _, ok := m[name]; !ok {
			deduplicated = append(deduplicated, name)
			m[name] = struct{}{}
		}
	}

	return deduplicated
}

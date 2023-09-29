package main

import "strings"

func SplitRepositoryID(name string) (string, string) {
	split := strings.Split(name, "/")
	return strings.ToLower(split[0]), strings.ToLower(split[1])
}

func ParseRepository(name string) string {
	if !strings.Contains(name, "/") {
		return "noahonfyre/" + strings.ToLower(name)
	} else {
		return strings.ToLower(name)
	}
}

func ParsePackage(name string) string {
	if !strings.Contains(name, ".") {
		return "noahonfyre." + strings.ToLower(name)
	} else {
		return strings.ToLower(name)
	}
}

func ToPackageName(name string) string {
	return strings.ReplaceAll(strings.ToLower(name), "/", ".")
}

func ToRepositoryName(name string) string {
	return strings.ReplaceAll(strings.ToLower(name), ".", "/")
}

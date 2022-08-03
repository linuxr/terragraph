package models

type Node struct {
	Id    string
	Name  string
	Group string
}

type Edge struct {
	SourceId string
	TargetId string
}

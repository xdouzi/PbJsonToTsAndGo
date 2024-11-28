package main

type PbToMain struct {
}

func (m PbToMain) CheckToTime() bool {

	return false
}

func (m PbToMain) Hang() {

}

func (m PbToMain) DeleteFiles(output string, s string) {

}

func NewPbToMain() *PbToMain {
	t := &PbToMain{}
	return t
}

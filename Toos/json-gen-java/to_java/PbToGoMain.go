package to_java

type PbToJavaMain struct {
}

func NewPbToJavaMain() *PbToJavaMain {
	t := &PbToJavaMain{}
	return t
}

func (t *PbToJavaMain) CheckToTime() bool {
	return false
}

func (t *PbToJavaMain) Hang() {

}

func (t *PbToJavaMain) DeleteFiles(output string, s string) {

}

/**
 * Just a little trick to use to see what's going on.
 */
package debug

import (
	"fmt"
	"path"
	"runtime"
)

func Trace() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	s := fmt.Sprintf("%s:%d %s", path.Base(file), line, f.Name())
	out(">>", s)
	return s
}
func Un(s string) {
	out("<<", s)
}
func Here() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	s := fmt.Sprintf("%s:%d %s", path.Base(file), line, f.Name())
	out("HERE", s)
}
func out(p string, s string) {
	fmt.Println("%s %s\n", p, s)
}

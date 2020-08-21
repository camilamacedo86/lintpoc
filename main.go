package main

import (
	"k8s.io/kubernetes/pkg/kubelet/kubeletconfig/util/log"
	"os"
	"strings"
)

const NEVER_USED = "should fail"

func main() {
	if err := testErrorFormat(); err != nil {
		// printf: Errorf call has arguments but no formatting directives (govet)
		log.Errorf("error", err)

		// Solution
		log.Errorf("error %v:", err)

	}
}

func testErrorFormat() error {
	return os.Chdir("test")
}

// The following will fail with (golint)

type AppService struct {
	Name string
}

// Should fail with GoSimple
func (s *AppService) validateContent() bool {
	if strings.Count(s.Name, " ") > 4 {
		return false
	}
	return true
}

// Should fail with GoSimple
func (s *AppService) validateContent2() bool {
	return strings.Count(s.Name, " ") > 4
}

// The following wil fail with (golint) (deadcode)
// Should
type AppServiceCommentStyle struct {
	Name string
}

// The following will fail with (structcheck)
// NotOPtime
type NotOPtime struct {
	Name string
	value bool
	Number int
}

// The following will fail with (maligned)
// NotOptimesed is an test to check lint issues
type NotOptimesed struct {
	Check bool
	Mapped []string
	Value bool
	Number int
	Name string
}

func (f *NotOptimesed) testUseNotOptimesed() error {
	return nil
}


func(f *NotOptimesed)  pleaseLetCheckYourCode(value string) string {
	if f.Name == "test" && f.Value == true && len(value) > 0 && value == "i will write something big here just to show it failing" {
		return "worked"
	}
	return "fail"
}
package starter

import (
	"os"
	"strconv"
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	s := "1451606400" // 1/1/2016
	b := "testing"
	h := "a10600b129253b1aaaa860778bef2043ee40c715"

	os.Setenv("GIT_TIME", s)
	os.Setenv("GIT_BRANCH", b)
	os.Setenv("GIT_HASH", h)
	c := NewConfig()
	if c.GitTime != s {
		t.Fatalf("Did not get expected time back, got %v\n", s)
	}
	if c.GitBranch != b {
		t.Fatalf("Did not get expected branch back, got %v\n", b)
	}
	if c.GitHash != h {
		t.Fatalf("Did not get expected hash back, got %v\n", h)
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		t.Fatalf("err from parseint %v\n", err)
	}

	gt := time.Unix(i, 0)
	if gt.Equal(c.BuildTime) != true {
		t.Fatalf("Did not get expected buildtime got %v\n", c.BuildTime)
	}
}

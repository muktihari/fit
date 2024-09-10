package opener

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	cd                = filepath.Dir(filename)
	testdata          = filepath.Join(cd, "..", "..", "..", "testdata")
	fromGarminForums  = filepath.Join(testdata, "from_garmin_forums")
	fromOfficialSDK   = filepath.Join(testdata, "from_official_sdk")
)

func TestOpen(t *testing.T) {
	numCPU = 2 // Override NumCPU

	paths := []string{
		filepath.Join(fromGarminForums, "triathlon_summary_first.fit"),
		filepath.Join(fromGarminForums, "triathlon_summary_last.fit"),
		filepath.Join(fromOfficialSDK, "activity_developerdata.fit"),
		filepath.Join(fromOfficialSDK, "Activity.fit"),
	}

	fits, err := Open(context.Background(), paths)
	if err != nil {
		t.Fatalf("expected error nil, got: %v", err)
	}
	if len(paths) != len(fits) {
		t.Fatalf("expected len(fits) is %d, got: %d", len(paths), len(fits))
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

//go:build snapshot

package cloudsearchdomain

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	"io"
	"io/fs"
	"os"
	"testing"
)

var quiet bool

func init() {
	flag.BoolVar(&quiet, "quiet", false, "suppress full snapshot diffs")
}

const ssprefix = "snapshot"

type snapshotOK struct{}

func (snapshotOK) Error() string { return "error: success" }

func createp(path string) (*os.File, error) {
	if err := os.Mkdir(ssprefix, 0700); err != nil && !errors.Is(err, fs.ErrExist) {
		return nil, err
	}
	return os.Create(path)
}

func sspath(op string) string {
	return fmt.Sprintf("%s/api_op_%s.go.snap", ssprefix, op)
}

func updateSnapshot(stack *middleware.Stack, operation string) error {
	f, err := createp(sspath(operation))
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write([]byte(stack.String())); err != nil {
		return err
	}
	return snapshotOK{}
}

func testSnapshot(stack *middleware.Stack, operation string) error {
	f, err := os.Open(sspath(operation))
	if errors.Is(err, fs.ErrNotExist) {
		return snapshotOK{}
	}
	if err != nil {
		return err
	}
	defer f.Close()
	expected, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	if actual := stack.String(); actual != string(expected) {
		if quiet {
			return fmt.Errorf("snapshot mismatch: %s", operation)
		}
		return fmt.Errorf("%s != %s", expected, actual)
	}
	return snapshotOK{}
}
func TestCheckSnapshot_Search(t *testing.T) {
	svc := New(Options{})
	_, err := svc.Search(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "Search")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_Suggest(t *testing.T) {
	svc := New(Options{})
	_, err := svc.Suggest(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "Suggest")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UploadDocuments(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UploadDocuments(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UploadDocuments")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}
func TestUpdateSnapshot_Search(t *testing.T) {
	svc := New(Options{})
	_, err := svc.Search(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "Search")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_Suggest(t *testing.T) {
	svc := New(Options{})
	_, err := svc.Suggest(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "Suggest")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UploadDocuments(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UploadDocuments(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UploadDocuments")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

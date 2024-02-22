// Code generated by smithy-go-codegen DO NOT EDIT.

//go:build snapshot

package b2bi

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
func TestCheckSnapshot_CreateCapability(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateCapability(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "CreateCapability")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_CreatePartnership(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreatePartnership(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "CreatePartnership")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_CreateProfile(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateProfile(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "CreateProfile")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_CreateTransformer(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateTransformer(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "CreateTransformer")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeleteCapability(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteCapability(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeleteCapability")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeletePartnership(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeletePartnership(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeletePartnership")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeleteProfile(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteProfile(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeleteProfile")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeleteTransformer(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteTransformer(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeleteTransformer")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_GetCapability(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetCapability(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "GetCapability")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_GetPartnership(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetPartnership(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "GetPartnership")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_GetProfile(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetProfile(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "GetProfile")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_GetTransformer(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetTransformer(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "GetTransformer")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_GetTransformerJob(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetTransformerJob(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "GetTransformerJob")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListCapabilities(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListCapabilities(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListCapabilities")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListPartnerships(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListPartnerships(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListPartnerships")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListProfiles(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListProfiles(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListProfiles")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListTagsForResource(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListTagsForResource(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListTagsForResource")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListTransformers(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListTransformers(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListTransformers")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_StartTransformerJob(t *testing.T) {
	svc := New(Options{})
	_, err := svc.StartTransformerJob(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "StartTransformerJob")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_TagResource(t *testing.T) {
	svc := New(Options{})
	_, err := svc.TagResource(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "TagResource")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_TestMapping(t *testing.T) {
	svc := New(Options{})
	_, err := svc.TestMapping(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "TestMapping")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_TestParsing(t *testing.T) {
	svc := New(Options{})
	_, err := svc.TestParsing(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "TestParsing")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UntagResource(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UntagResource(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UntagResource")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UpdateCapability(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateCapability(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UpdateCapability")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UpdatePartnership(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdatePartnership(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UpdatePartnership")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UpdateProfile(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateProfile(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UpdateProfile")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UpdateTransformer(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateTransformer(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UpdateTransformer")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}
func TestUpdateSnapshot_CreateCapability(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateCapability(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "CreateCapability")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_CreatePartnership(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreatePartnership(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "CreatePartnership")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_CreateProfile(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateProfile(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "CreateProfile")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_CreateTransformer(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateTransformer(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "CreateTransformer")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeleteCapability(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteCapability(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeleteCapability")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeletePartnership(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeletePartnership(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeletePartnership")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeleteProfile(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteProfile(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeleteProfile")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeleteTransformer(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteTransformer(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeleteTransformer")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_GetCapability(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetCapability(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "GetCapability")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_GetPartnership(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetPartnership(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "GetPartnership")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_GetProfile(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetProfile(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "GetProfile")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_GetTransformer(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetTransformer(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "GetTransformer")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_GetTransformerJob(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetTransformerJob(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "GetTransformerJob")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListCapabilities(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListCapabilities(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListCapabilities")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListPartnerships(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListPartnerships(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListPartnerships")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListProfiles(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListProfiles(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListProfiles")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListTagsForResource(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListTagsForResource(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListTagsForResource")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListTransformers(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListTransformers(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListTransformers")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_StartTransformerJob(t *testing.T) {
	svc := New(Options{})
	_, err := svc.StartTransformerJob(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "StartTransformerJob")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_TagResource(t *testing.T) {
	svc := New(Options{})
	_, err := svc.TagResource(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "TagResource")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_TestMapping(t *testing.T) {
	svc := New(Options{})
	_, err := svc.TestMapping(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "TestMapping")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_TestParsing(t *testing.T) {
	svc := New(Options{})
	_, err := svc.TestParsing(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "TestParsing")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UntagResource(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UntagResource(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UntagResource")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UpdateCapability(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateCapability(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UpdateCapability")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UpdatePartnership(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdatePartnership(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UpdatePartnership")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UpdateProfile(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateProfile(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UpdateProfile")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UpdateTransformer(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateTransformer(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UpdateTransformer")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

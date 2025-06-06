package testrunner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunnerSanitizeAndParseArgs(t *testing.T) {
	t.Run("Passthrough", func(t *testing.T) {
		r := newRunner("gotestsum")
		args, err := r.sanitizeAndParseArgs([]string{
			"--junitfile=test.xml",
			"-foo",
			"bar",
			"--retries=3",
			"--",
			"-coverprofile=test.cover.out",
			"baz",
		})
		require.NoError(t, err)
		require.Equal(t, []string{
			"--junitfile=test.xml",
			"-foo",
			"bar",
			// retries has been stripped
			"--",
			"-coverprofile=test.cover.out",
			"baz",
		}, args)
		require.Equal(t, "test.xml", r.junitOutputPath)
		require.Equal(t, 3, r.retries)
		require.Equal(t, "test.cover.out", r.coverProfilePath)
	})

	t.Run("RetriesMissing", func(t *testing.T) {
		r := newRunner("gotestsum")
		_, err := r.sanitizeAndParseArgs([]string{
			"--junitfile=test.xml",
			"-foo",
			"bar",
			// missing:
			// "--retries=0",
			"--",
			"-coverprofile=test.cover.out",
			"baz",
		})
		require.ErrorContains(t, err, `missing required argument "--retries="`)
	})
	t.Run("RetriesInvalid1", func(t *testing.T) {
		r := newRunner("gotestsum")
		_, err := r.sanitizeAndParseArgs([]string{
			"--junitfile=test.xml",
			"-foo",
			"bar",
			"--retries=0", // invalid!
			"--",
			"-coverprofile=test.cover.out",
			"baz",
		})
		require.ErrorContains(t, err, `invalid argument "--retries=": must be greater than zero`)
	})
	t.Run("RetriesInvalid2", func(t *testing.T) {
		r := newRunner("gotestsum")
		_, err := r.sanitizeAndParseArgs([]string{
			"--junitfile=test.xml",
			"-foo",
			"bar",
			"--retries=invalid", // invalid!
			"--",
			"-coverprofile=test.cover.out",
			"baz",
		})
		require.ErrorContains(t, err, `invalid argument "--retries=": strconv.Atoi: parsing "invalid"`)
	})

	t.Run("JunitfileMissing", func(t *testing.T) {
		r := newRunner("gotestsum")
		_, err := r.sanitizeAndParseArgs([]string{
			// missing:
			// "--junitfile=test.xml"
			"-foo",
			"bar",
			"--retries=3",
			"--",
			"-coverprofile=test.cover.out",
			"baz",
		})
		require.ErrorContains(t, err, `missing required argument "--junitfile="`)
	})

	t.Run("CoverprofileMissing", func(t *testing.T) {
		r := newRunner("gotestsum")
		_, err := r.sanitizeAndParseArgs([]string{
			"--junitfile=test.xml",
			"-foo",
			"bar",
			"--retries=3",
			"--",
			// missing:
			// "-coverprofile=test.cover.out",
			"baz",
		})
		require.ErrorContains(t, err, `missing required argument "-coverprofile="`)
	})
}

func TestStripRunFromArgs(t *testing.T) {
	t.Run("OneArg", func(t *testing.T) {
		args := stripRunFromArgs([]string{"-foo", "bar", "-run=A"})
		require.Equal(t, []string{"-foo", "bar"}, args)
	})

	t.Run("TwoArgs", func(t *testing.T) {
		args := stripRunFromArgs([]string{"-foo", "bar", "-run", "A"})
		require.Equal(t, []string{"-foo", "bar"}, args)
	})
}

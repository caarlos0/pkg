package acceptance

import "testing"

func TestSimpleRPM(t *testing.T) {
	t.Run("amd64", func(t *testing.T) {
		t.Parallel()
		accept(t, "simple_rpm", "simple.yaml", "rpm", "rpm.dockerfile")
	})
	t.Run("i386", func(t *testing.T) {
		t.Parallel()
		accept(t, "simple_rpm_386", "simple.386.yaml", "rpm", "rpm.386.dockerfile")
	})
}

func TestComplexRPM(t *testing.T) {
	t.Run("amd64", func(t *testing.T) {
		t.Parallel()
		accept(t, "complex_rpm", "complex.yaml", "rpm", "rpm.complex.dockerfile")
	})
	t.Run("i386", func(t *testing.T) {
		t.Parallel()
		accept(t, "complex_rpm_386", "complex.386.yaml", "rpm", "rpm.386.complex.dockerfile")
	})
}

func TestComplexOverridesRPM(t *testing.T) {
	t.Run("amd64", func(t *testing.T) {
		t.Parallel()
		accept(t, "overrides_rpm", "overrides.yaml", "rpm", "rpm.overrides.dockerfile")
	})
}

func TestMinRPM(t *testing.T) {
	t.Run("amd64", func(t *testing.T) {
		t.Parallel()
		accept(t, "min_rpm", "min.yaml", "rpm", "rpm.min.dockerfile")
	})
}

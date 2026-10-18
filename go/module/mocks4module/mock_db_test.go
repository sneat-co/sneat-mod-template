package mocks4module

import "testing"

func TestNewTestContext(t *testing.T) {
	ctx, db := NewTestContext(t)
	if ctx == nil {
		t.Fatal("expected non-nil context")
	}
	if db == nil {
		t.Fatal("expected non-nil db")
	}
}

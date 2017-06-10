package noface_test

import (
  "github.com/picatz/noface"
  "testing"
)

func TestFirstIface(t *testing.T) {

  _, err := noface.FirstIface()
  if err != nil {
    t.Errorf("FirstIface should not fail with error: %v", err)
  }

}

package majorityvote_test

import (
    "testing"
    "majorityvote"
)

func TestMajorityVote(t *testing.T) {
    t1 := []string{"a", "b", "c", "a", "a", "b", "d", "a", "a"} // 'a' is majority vote
    if (majorityvote.MajorityVote(t1) != "a") {
        t.Error("Expected a result of 'a'.")
    }

    t2 := []string{"a", "b", "c", "a", "a", "b", "d"}   // no majority vote
    if(majorityvote.MajorityVote(t2) != "") {
        t.Error("Expected a result of '' (No majority).")
    }
}

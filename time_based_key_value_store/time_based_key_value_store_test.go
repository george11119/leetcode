package time_based_key_value_store

import "testing"

func TestTimeMap(t *testing.T) {
	got := ""
	want := ""

	tm := Constructor()

	tm.Set("alice", "happy", 1)

	got = tm.Get("alice", 1)
	want = "happy"
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = tm.Get("alice", 2)
	want = "happy"
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	tm.Set("alice", "sad", 3)

	got = tm.Get("alice", 3)
	want = "sad"
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}
func TestTimeMap2(t *testing.T) {
	got := ""
	want := ""

	tm := Constructor()

	tm.Set("alice", "a", 1)
	tm.Set("alice", "b", 3)
	tm.Set("alice", "c", 5)
	tm.Set("alice", "d", 7)
	tm.Set("alice", "e", 9)
	tm.Set("alice", "f", 11)
	tm.Set("alice", "g", 13)
	tm.Set("alice", "h", 15)

	got = tm.Get("alice", 12)
	want = "f"
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}

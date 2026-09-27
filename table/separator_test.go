package table

import "testing"

func TestSeparator_TopUsesRoundedCorners(t *testing.T) {
	table := &Table{}
	result := table.Separator([]int{3, 3}, 0)
	if result != "┏━━━┳━━━┓\n" {
		t.Errorf("unexpected top separator: %q", result)
	}
}

func TestSeparator_MidUsesCross(t *testing.T) {
	table := &Table{}
	result := table.Separator([]int{3, 3}, 1)
	if result != "┣━━━╋━━━┫\n" {
		t.Errorf("unexpected mid separator: %q", result)
	}
}

func TestSeparator_BottomUsesRoundedCorners(t *testing.T) {
	table := &Table{}
	result := table.Separator([]int{3, 3}, 2)
	if result != "┗━━━┻━━━┛\n" {
		t.Errorf("unexpected bottom separator: %q", result)
	}
}

package solution

import "testing"

func TestShouldReturnProperValueForTestInputPart1(t *testing.T) {
	res := handlePart1("1x1x10")
	expected := 43
	if res != expected {
		t.Fatal("Returned: ", res, " should be: ", expected)
	}

	res = handlePart1("2x3x4")
	expected = 58
	if res != expected {
		t.Fatal("Returned: ", res, " should be: ", expected)
	}

	res = handlePart1(`
4x23x21
22x29x19
29x2x30
`)

	expected = 7068
	if res != expected {
		t.Fatal("Returned: ", res, " should be: ", expected)
	}
}

func TestShouldReturnProperValueForProperInputPart1(t *testing.T) {
	res := part1()
	exepcted := 1598415
	if res != exepcted {
		t.Fatal("Returned: ", res, " should be: ", exepcted)
	}
}

func TestGetSmallest(t *testing.T) {
	res := getSmallest(1, 2, 3)
	if res != 1 {
		t.Fatal("returned: ", res)
	}
	res = getSmallest(2, 2, 3)
	if res != 2 {
		t.Fatal("returned: ", res)
	}

	res = getSmallest(2, 2, 2)
	if res != 2 {
		t.Fatal("returned: ", res)
	}

	res = getSmallest(2, 2, 1)
	if res != 1 {
		t.Fatal("returned: ", res)
	}
	res = getSmallest(2, 1, 2)
	if res != 1 {
		t.Fatal("returned: ", res)
	}

	res = getSmallest(3, 2, 1)
	if res != 1 {
		t.Fatal("returned: ", res)
	}
	res = getSmallest(29, 2, 30)
	if res != 2 {
		t.Fatal("returned: ", res)
	}
	res = getSmallest(58, 870, 60)
	if res != 58 {
		t.Fatal("returned: ", res)
	}
}

func TestShouldReturnProperValueForTestPart1(t *testing.T) {
	res := handlePart2("2x3x4")
	expected := 34
	if res != expected {
		t.Fatal("Returned: ", res, " should be: ", expected)
	}
}

func TestShouldReturnProperValueForProperInputPart2(t *testing.T) {
	res := part2()
	exepcted := 3812909
	if res != exepcted {
		t.Fatal("Returned: ", res, " should be: ", exepcted)
	}
}

package srel

import "testing"

func TestParse(t *testing.T) {
	testQuery := `STARTOFLINE ^ THEN www. MAYBE "walla walla"`

	testResults := []string {
		"STARTOFLINE",
		"^",
		"THEN",
		"www.",
		"MAYBE",
		"walla walla",
	}

	parseResults := parse(testQuery)

	for i, e := range parseResults {
		if i >= len(testResults) {
			t.Error("Too many items of the stack:", e)
			return
		}

		if e != testResults[i] {
			t.Error("Result mismatch: Expected:", testResults[i], "Acc:", e)
		}
	}
}

func TestURL(t *testing.T) {

	stack := parse(`STARTOFLINE THEN http MAYBE s THEN :// MAYBE www. ALLBUT " " EOL`)

	b, err := run(stack)
	if err != nil {
		t.Fatal("run call returned an error,", err.Error())
	}

	t.Log("builder:", b)

	urlRX, err := b.compile()
	if err != nil {
		t.Fatal("compile call returned and error,", err.Error())
	}

	t.Log(urlRX)

	ptests := []string {
		"https://www.google.com",
		"http://www.google.com",
		"https://google.com",
		"http://google.com",
	}

	ftests := []string{
		"www.google.com",
		"https:// .com",
	}

	t.Log("Checking Passing strings")
	for _, e := range ptests {
		if urlRX.Match([]byte(e)) == false {
			t.Error(e, "Failed")
		}
	}
	t.Log("------------------------")

	t.Log("Checking Failing string")
	for _, e := range ftests {
		if urlRX.Match([]byte(e)) {
			t.Error(e, "Failed")
		}
	}
	t.Log("------------------------")
}
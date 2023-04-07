package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	t.Log("Given the need to test isPrime.")
	{
		for testID, e := range primeTests {
			result, msg := isPrime(e.testNum)
			if e.expected && !result {
				t.Errorf("\t%sTest %d %s: expected true but got false", failed, testID, e.name)
			} else {
				t.Logf("\t%sTest %d %s: Should return true", succeed, testID, e.name)

			}

			if !e.expected && result {
				t.Errorf("\t%sTest %d %s: expected false but got true", failed, testID, e.name)
			} else {
				t.Logf("\t%sTest %d %s: Should return false", succeed, testID, e.name)

			}

			if e.msg != msg {
				t.Errorf("\t%sTest %d %s: expected %s but got %s", failed, testID, e.name, e.msg, msg)
			} else {
				t.Logf("\t%sTest %d %s: Should return %s", succeed, testID, e.name, msg)
			}
		}
	}

}

func TestPrompt(t *testing.T) {
	t.Log("Testing prompt function")
	{
		old := os.Stdout
		r, w, _ := os.Pipe() //used for temporary replacing instead of os.stdout
		os.Stdout = w

		prompt()

		w.Close()
		os.Stdout = old

		var buf bytes.Buffer
		io.Copy(&buf, r)

		if output := buf.String(); output != "-> " {
			t.Errorf("%s should return '-> ' instead of %s", failed, output)
		} else {
			t.Logf("%s promt should return '-> ' ", succeed)
		}

	}

}

func TestIntro(t *testing.T) {
	expectedAsw := "Is it Prime?\n------------\nEnter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n-> "
	t.Log("Testing intro function")
	{
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		intro()

		w.Close()
		os.Stdout = old

		var buf bytes.Buffer
		io.Copy(&buf, r)

		if output := buf.String(); output != expectedAsw {
			t.Errorf("%s Test 0: should return %s instead of %s", failed, expectedAsw, output)
		} else {
			t.Logf("%s Test 0 : should return :\n%s ", succeed, expectedAsw)
		}
	}
}

func TestCheckNumbers(t *testing.T) {
	//Table Unit test
	testTable := []struct {
		name            string
		input           string
		expectedMessage string
		expectedBool    bool
	}{
		{"Quit", "q", "", true},
		{"Not number", "s", "Please enter a whole number!", false},
		{"Prime", "7", "7 is a prime number!", false},
	}

	t.Log("Given the need to test checkNumbers function.")
	{
		for testID, test := range testTable {
			r := strings.NewReader(test.input)

			res, done := checkNumbers(bufio.NewScanner(r))

			if res != test.expectedMessage {
				t.Errorf("\t%s Test %d %s: expected '%s' but got '%s'", failed, testID, test.name, test.expectedMessage, res)
			} else {
				t.Logf("\t%s Test %d %s: should return '%s'", succeed, testID, test.name, test.expectedMessage)
			}

			if done != test.expectedBool {
				t.Errorf("\t%s Test %d %s: expected '%t' but got '%t'", failed, testID, test.name, test.expectedBool, done)
			} else {
				t.Logf("\t%s Test %d %s: should return '%t'", succeed, testID, test.name, test.expectedBool)
			}

		}

	}
}

func TestReadUserInput(t *testing.T) {

	doneChan := make(chan bool)
	input := "1\n2\n3\n7\nh\nq\n"
	expectedOutput := "1 is not prime, by definition!\n-> 2 is a prime number!\n-> 3 is a prime number!\n-> 7 is a prime number!\n-> Please enter a whole number!\n-> "

	t.Log("Given the need to test readUserInput function.")
	{
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		go readUserInput(strings.NewReader(input), doneChan)

		<-doneChan

		w.Close()
		os.Stdout = old

		var buf bytes.Buffer
		io.Copy(&buf, r)
		if output := buf.String(); output != expectedOutput {
			t.Errorf("%s Test 0 : Output does not match expected output. Got %q, expected %q", failed, output, expectedOutput)
		} else {
			t.Logf("%s Test 0 : Output does match expected output. Got %q", succeed, expectedOutput)
		}

	}
}

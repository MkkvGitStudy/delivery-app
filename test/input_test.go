package test

import (
	"bytes"
	"delivery-app/input"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

const ENTER_FLOAT_PROMPT = "Enter a float number:"
const ENTER_INT_PROMPT = "Enter an int value:"

func TestGetFloatInput(t *testing.T) {
	floatIp := input.FloatInput{}
	tests := []struct {
		name        string
		input       string
		expected    float64
		expectError bool
	}{
		{"valid - positive float val", "3.1415\n", 3.1415, false},
		{"valid - negative float val", "-273.15\n", -273.15, false},
		{"valid - float val 0", "0\n", 0.0, false},
		{"invalid - string val", "abc\n42.42\n", 42.42, true}, //appending float number to it, as the prompt will not close without a valid float value
		{"invalid - number and special character val", "42,42\n42.42\n", 42.42, true},
		{"invalid - number as string val val", "one hundred\n100\n", 100, true},
	}

	for _, test := range tests {
		// Simulate user input
		r, w, _ := os.Pipe()
		os.Stdin = r

		go func() {
			w.WriteString(test.input)
			w.Close()
		}()

		// Capture standard output
		oldStdout := os.Stdout
		rOut, wOut, _ := os.Pipe()
		os.Stdout = wOut

		var buf bytes.Buffer
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			io.Copy(&buf, rOut)
		}()

		// Call the function and check the result
		result := floatIp.GetFloatInput(ENTER_FLOAT_PROMPT)
		wOut.Close()
		wg.Wait() // Wait for the output capture to complete
		rOut.Close()
		os.Stdout = oldStdout

		if result != test.expected {
			t.Errorf("For input %s, expected %f but got %f", test.input, test.expected, result)
		}

		// Check if an error message was expected
		if test.expectError {
			output := buf.String()
			if !strings.Contains(output, "Invalid input. Please enter a valid float number.") {
				t.Errorf("%s expected invalid input warning for input %s, but got %s", test.name, test.input, output)
			}
		}
	}

	// Reset os.Stdin after tests
	os.Stdin = os.NewFile(0, "/dev/tty")
}

func TestGetIntInput(t *testing.T) {
	intIp := input.IntInput{}
	tests := []struct {
		name        string
		input       string
		expected    int
		expectError bool
	}{
		{"valid - positive int val", "3\n", 3, false},
		{"valid - negative int val", "-5\n", -5, false},
		{"valid - int val 0", "0\n", 0.0, false},
		{"invalid - string val", "abc\n42\n", 42, true}, //appending float number to it, as the prompt will not close without a valid float value
		{"invalid - number and special character val", "42,42\n42\n", 42, true},
		{"invalid - number as string val val", "one hundred\n100\n", 100, true},
	}

	for _, test := range tests {
		// Simulate user input
		r, w, _ := os.Pipe()
		os.Stdin = r

		go func() {
			w.WriteString(test.input)
			w.Close()
		}()

		// Capture standard output
		oldStdout := os.Stdout
		rOut, wOut, _ := os.Pipe()
		os.Stdout = wOut

		var buf bytes.Buffer
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			io.Copy(&buf, rOut)
		}()

		// Call the function and check the result
		result := intIp.GetIntInput(ENTER_INT_PROMPT)
		wOut.Close()
		wg.Wait() // Wait for the output capture to complete
		rOut.Close()
		os.Stdout = oldStdout

		if result != test.expected {
			t.Errorf("For input %s, expected %d but got %d", test.input, test.expected, result)
		}

		// Check if an error message was expected
		if test.expectError {
			output := buf.String()
			if !strings.Contains(output, "Invalid input. Please enter a valid integer.") {
				t.Errorf("%s expected invalid input warning for input %s, but got %s", test.name, test.input, output)
			}
		}
	}

	// Reset os.Stdin after tests
	os.Stdin = os.NewFile(0, "/dev/tty")
}

func TestGetPositiveFloatValue(t *testing.T) {
	tests := []struct {
		name          string
		inputs        []float64
		prompt        string
		zeroInclusive bool
		expected      float64
		expectedMsgs  []string
	}{
		{"valid float number", []float64{42.42}, ENTER_FLOAT_PROMPT, false, 42.42, []string{}},
		{"negative number with zero inclusive", []float64{-1, 3.14159}, ENTER_FLOAT_PROMPT, true, 3.14159, []string{"Please enter a positive value"}},
		{"negative number with zero exclusive", []float64{-1, 3.14159}, ENTER_FLOAT_PROMPT, false, 3.14159, []string{"Please enter a positive non zero value"}},
		{"zero with zero exclusive", []float64{0, 3.14159}, ENTER_FLOAT_PROMPT, false, 3.14159, []string{"Please enter a positive non zero value"}},
		{"zero with zero inclusive", []float64{0}, ENTER_FLOAT_PROMPT, true, 0, []string{}},
	}

	for _, test := range tests {
		mockGetter := &input.MockFloatInputGetter{Responses: test.inputs}

		// Capture standard output
		oldStdout := os.Stdout
		rOut, wOut, _ := os.Pipe()
		os.Stdout = wOut

		var buf bytes.Buffer
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			io.Copy(&buf, rOut)
		}()

		// Call the function and check the result
		result := input.GetPositiveFloatValue(test.prompt, test.zeroInclusive, mockGetter)
		wOut.Close()
		wg.Wait() // Wait for the output capture to complete
		os.Stdout = oldStdout

		if result != test.expected {
			t.Errorf("%s - for inputs %v, expected %f but got %f", test.name, test.inputs, test.expected, result)
		}

		// Check if the correct prompts were printed
		output := buf.String()
		for _, expectedMsg := range test.expectedMsgs {
			if !strings.Contains(output, expectedMsg) {
				t.Errorf("%s - expected output to contain message %s", test.name, expectedMsg)
			}
		}
	}
}

func TestGetPositiveIntValue(t *testing.T) {
	tests := []struct {
		name         string
		inputs       []int
		prompt       string
		expected     int
		expectedMsgs []string
	}{
		{"valid positive number", []int{42}, ENTER_INT_PROMPT, 42, []string{}},
		{"negative number ", []int{-1, 3}, ENTER_INT_PROMPT, 3, []string{"Please enter a positive value"}},
		{"zero", []int{0, 3}, ENTER_INT_PROMPT, 3, []string{}},
	}

	for _, test := range tests {
		mockGetter := &input.MockIntInputGetter{Responses: test.inputs}

		// Capture standard output
		oldStdout := os.Stdout
		rOut, wOut, _ := os.Pipe()
		os.Stdout = wOut

		var buf bytes.Buffer
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			io.Copy(&buf, rOut)
		}()

		// Call the function and check the result
		result := input.GetPositiveIntValue(test.prompt, mockGetter)
		wOut.Close()
		wg.Wait() // Wait for the output capture to complete
		os.Stdout = oldStdout

		if result != test.expected {
			t.Errorf("%s - for inputs %v, expected %d but got %d", test.name, test.inputs, test.expected, result)
		}

		// Check if the correct prompts were printed
		output := buf.String()
		for _, expectedMsg := range test.expectedMsgs {
			if !strings.Contains(output, expectedMsg) {
				t.Errorf("%s - expected output to contain message %s", test.name, expectedMsg)
			}
		}
	}
}

func TestGetStringInput(t *testing.T) {
	tests := []struct {
		input    string
		prompt   string
		expected string
	}{
		{"hello\n", "Enter text: ", "hello"},
		{" world \n", "Type something: ", "world"},
		{"  go  \n", "Input: ", "go"},
	}

	for _, test := range tests {
		// Simulate user input
		r, w, _ := os.Pipe()
		oldStdin := os.Stdin
		os.Stdin = r

		go func() {
			w.WriteString(test.input)
			w.Close()
		}()

		// Capture standard output
		oldStdout := os.Stdout
		rOut, wOut, _ := os.Pipe()
		os.Stdout = wOut

		var buf bytes.Buffer
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			io.Copy(&buf, rOut)
		}()

		// Call the function and check the result
		result := input.GetStringInput(test.prompt)
		wOut.Close()
		wg.Wait() // Wait for the output capture to complete
		os.Stdout = oldStdout
		os.Stdin = oldStdin

		if result != test.expected {
			t.Errorf("For input %s, expected %s but got %s", test.input, test.expected, result)
		}

		// Check if the correct prompt was printed
		output := buf.String()
		if !strings.Contains(output, test.prompt) {
			t.Errorf("Expected prompt %s but got %s", test.prompt, output)
		}
	}
}

func TestGetCouponCodeInput(t *testing.T) {
	const testPrompt = "Enter coupon code: "
	tests := []struct {
		name          string
		inputs        []string
		prompt        string
		expected      string
		expectedError bool
	}{
		{"valid code - caps alphanumeric", []string{"CODE123\n"}, testPrompt, "CODE123", false},
		{"valid code - number only", []string{"123\n"}, testPrompt, "123", false},
		{"valid code - alphabets only", []string{"ABcd\n"}, testPrompt, "ABcd", false},
		{"valid code - small case alphanumeric", []string{"abc123\n"}, testPrompt, "abc123", false},
		{"invalid code - with space", []string{"abc 123\n", "abc123\n"}, testPrompt, "abc123", true},
		{"invalid code - special characters", []string{"@@@###\n", "CODE123\n"}, testPrompt, "CODE123", true},
	}

	for _, test := range tests {
		// Simulate user input
		r, w, _ := os.Pipe()
		oldStdin := os.Stdin
		os.Stdin = r

		go func() {
			for _, input := range test.inputs {
				w.WriteString(input)
			}
			w.Close()
		}()

		// Capture standard output
		oldStdout := os.Stdout
		rOut, wOut, _ := os.Pipe()
		os.Stdout = wOut

		var buf bytes.Buffer
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			io.Copy(&buf, rOut)
		}()

		// Call the function and check the result
		result := input.GetCouponCodeInput(test.prompt)
		wOut.Close()
		wg.Wait() // Wait for the output capture to complete
		os.Stdout = oldStdout
		os.Stdin = oldStdin

		if result != test.expected {
			t.Errorf("For inputs %v, expected %s but got %s", test.inputs, test.expected, result)
		}

		// Check if the correct prompts and error messages were printed
		output := buf.String()
		if !strings.Contains(output, test.prompt) {
			t.Errorf("%s - expected prompt %s but got %s", test.name, test.prompt, output)
		}

		if test.expectedError {
			if !strings.Contains(output, "Invalid input. Please enter a valid alphanumeric string without spaces.") {
				t.Errorf("Expected error message for invalid input, but got %s", output)
			}
		}
	}
}

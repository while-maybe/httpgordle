package gordle

import "testing"

// Benchmark the string concatenation with only one value in feedback
func BenchmarkStringConcat1(b *testing.B) {
	fb := Feedback{absentCharacter}
	for b.Loop() {
		_ = fb.StringConcat()
	}
}

// Benchmark the string concatenation with two hints in feedback
func BenchmarkStringConcat2(b *testing.B) {
	fb := Feedback{correctPosition, absentCharacter}
	for b.Loop() {
		_ = fb.StringConcat()
	}
}

// Benchmark the string concatenation with three hints in feedback
func BenchmarkStringConcat3(b *testing.B) {
	fb := Feedback{absentCharacter, absentCharacter, wrongPosition}
	for b.Loop() {
		_ = fb.StringConcat()
	}
}

// Benchmark the string concatenation with four hints in feedback
func BenchmarkStringConcat4(b *testing.B) {
	fb := Feedback{absentCharacter, correctPosition, correctPosition, wrongPosition}
	for b.Loop() {
		_ = fb.StringConcat()
	}
}

// Benchmark the string concatenation with five hints in feedback
func BenchmarkStringConcat5(b *testing.B) {
	fb := Feedback{absentCharacter, correctPosition, correctPosition, absentCharacter, wrongPosition}
	for b.Loop() {
		_ = fb.StringConcat()
	}
}

// Benchmark the string Builder with only one value in feedback
func BenchmarkStringBuilder1(b *testing.B) {
	fb := Feedback{absentCharacter}
	for b.Loop() {
		_ = fb.String()
	}
}

// Benchmark the string Builder with two hints in feedback
func BenchmarkStringBuilder2(b *testing.B) {
	fb := Feedback{correctPosition, absentCharacter}
	for b.Loop() {
		_ = fb.String()
	}
}

// Benchmark the string Builder with three hints in feedback
func BenchmarkStringBuilder3(b *testing.B) {
	fb := Feedback{absentCharacter, absentCharacter, wrongPosition}
	for b.Loop() {
		_ = fb.String()
	}
}

// Benchmark the string Builder with four hints in feedback
func BenchmarkStringBuilder4(b *testing.B) {
	fb := Feedback{absentCharacter, correctPosition, correctPosition, wrongPosition}
	for b.Loop() {
		_ = fb.String()
	}
}

// Benchmark the string Builder with five hints in feedback
func BenchmarkStringBuilder5(b *testing.B) {
	fb := Feedback{absentCharacter, correctPosition, correctPosition, absentCharacter, wrongPosition}
	for b.Loop() {
		_ = fb.String()
	}
}

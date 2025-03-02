package clip

import (
	"slices"
	"testing"
)

var arguments = []string{
	"--doubleValue", "value1",
	"-singleValue", "value2",
	"--doubleNoValue", "-singleNoValue",
	"value3", "value4",
}

func TestArgumentParser(t *testing.T) {

	set := parseArguments(arguments)
	expectedTypes := []argumentKind{
		argumentKindLong, argumentKindValue,
		argumentKindShort, argumentKindValue,
		argumentKindLong, argumentKindShort,
		argumentKindValue, argumentKindValue,
	}
	expectedValues := []string{
		"doubleValue", "value1",
		"singleValue", "value2",
		"doubleNoValue", "singleNoValue",
		"value3", "value4",
	}

	for i, v := range expectedTypes {
		a := set.next()
		if a.kind != v {
			t.Errorf("expected kind %v but got %v at index %d", v, set.next().kind, i)
		}
		if a.value != expectedValues[i] {
			t.Errorf("expected value %v but got %v at index %d", expectedValues[i], set.next().kind, i)
		}
	}
}

func TestClipParser(t *testing.T) {
	app := New()
	app.String("doubleValue")
	app.String("singleV").Shorthand("singleValue")
	app.Boolean("doubleNoValue")
	app.Boolean("singleNoV").Shorthand("singleNoValue")
	app.KVFlag("kv")
	app.StringSlice("slice")
	app.Options("opts", "a", "b", "c")
	app.Int("int")
	app.Int8("int8")
	app.Int16("int16")
	app.Int32("int32")
	app.Int64("int64")
	app.Uint("uint")
	app.Uint8("uint8")
	app.Uint16("uint16")
	app.Uint32("uint32")
	app.Uint64("uint64")

	cli := app.parse(append([]string{
		"--kv", "a=b",
		"--kv", "c=d",
		"--slice", "a",
		"--slice", "b",
		"--slice", "c",
		"--opts", "c",
		"--int", "1",
		"--int8", "2",
		"--int16", "3",
		"--int32", "4",
		"--int64", "5",
		"--uint", "6",
		"--uint8", "7",
		"--uint16", "8",
		"--uint32", "9",
		"--uint64", "10",
	}, arguments...))
	v1 := cli.String("doubleValue")
	v2 := cli.String("singleV")
	v3 := cli.Boolean("doubleNoValue")
	v4 := cli.Boolean("singleNoV")
	v5 := cli.KV("kv")
	v6 := cli.StringSlice("slice")
	v7 := cli.String("opts")

	if v1 != "value1" {
		t.Errorf("expected value %v but got %v", "value1", v1)
	}
	if v2 != "value2" {
		t.Errorf("expected value %v but got %v", "value2", v2)
	}
	if v3 != true {
		t.Errorf("expected value %v but got %v", true, v3)
	}
	if v4 != true {
		t.Errorf("expected value %v but got %v", true, v4)
	}

	kvValues := map[string]string{"a": "b", "c": "d"}
	for k, v := range kvValues {
		if v5[k] != v {
			t.Errorf("expected value %v but got %v", v, v5[k])
		}
	}

	if !slices.Equal([]string{"a", "b", "c"}, v6) {
		t.Errorf("expected value %v but got %v", []string{"a", "b", "c"}, v6)
	}

	if v7 != "c" {
		t.Errorf("expected value %v but got %v", "c", v7)
	}

	v8 := cli.Int("int")
	v9 := cli.Int8("int8")
	v10 := cli.Int16("int16")
	v11 := cli.Int32("int32")
	v12 := cli.Int64("int64")
	v13 := cli.Uint("uint")
	v14 := cli.Uint8("uint8")
	v15 := cli.Uint16("uint16")
	v16 := cli.Uint32("uint32")
	v17 := cli.Uint64("uint64")

	if v8 != 1 {
		t.Errorf("expected value %v but got %v", 1, v8)
	}
	if v9 != 2 {
		t.Errorf("expected value %v but got %v", 2, v9)
	}
	if v10 != 3 {
		t.Errorf("expected value %v but got %v", 3, v10)
	}
	if v11 != 4 {
		t.Errorf("expected value %v but got %v", 4, v11)
	}
	if v12 != 5 {
		t.Errorf("expected value %v but got %v", 5, v12)
	}
	if v13 != 6 {
		t.Errorf("expected value %v but got %v", 6, v13)
	}
	if v14 != 7 {
		t.Errorf("expected value %v but got %v", 7, v14)
	}
	if v15 != 8 {
		t.Errorf("expected value %v but got %v", 8, v15)
	}
	if v16 != 9 {
		t.Errorf("expected value %v but got %v", 9, v16)
	}
	if v17 != 10 {
		t.Errorf("expected value %v but got %v", 10, v17)
	}

	if cli.NArgs() != 2 {
		t.Errorf("expected value %v but got %v", 2, cli.NArgs())
	}

	if cli.Arg(0) != "value3" {
		t.Errorf("expected value %v but got %v", "value3", cli.Arg(0))
	}
	if cli.Arg(1) != "value4" {
		t.Errorf("expected value %v but got %v", "value4", cli.Arg(1))
	}
}

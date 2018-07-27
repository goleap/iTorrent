package common

import "testing"

func TestEncodeInt(t *testing.T) {
	if EncodeInt(0) != "i0e" {
		t.Fatalf("Expect i0e, got %s\n", EncodeInt(0))
	}
	if EncodeInt(42) != "i42e" {
		t.Fatalf("Expect i42e, got %s\n", EncodeInt(42))
	}
	if EncodeInt(-42) != "i-42e" {
		t.Fatalf("Expect i-42e, got %s\n", EncodeInt(-42))
	}
}

func TestEncodeString(t *testing.T) {
	if EncodeString("spam") != "4:spam" {
		t.Fatalf("Expect 4:spam, got %s\n", EncodeString("spam"))
	}
}

func TestEncodeList(t *testing.T) {
	if EncodeList([]interface{}{"spam", 42}) != "l4:spami42ee" {
		t.Fatalf("Expect l4:spami42ee, got %s\n", EncodeList([]interface{}{"spam", 42}))
	}
}

func TestEncodeMap(t *testing.T) {
	data := map[string]interface{}{"bar": "spam", "foo": 42}
	target := "d3:bar4:spam3:fooi42ee"
	actual := EncodeMap(data)
	if actual != target {
		t.Fatalf("Expect %s, got %s\n", target, actual)
	}
}

func TestDecodeInt(t *testing.T) {
	num,eat,err := decodeInt("i-42e")
	if err != nil {
		t.Fatalf("Expect -42, got error: %v\n",err)
	} else if num != -42 {
		t.Fatalf("Expect -42, got %d\n",num)
	}
	if eat != 5 {
		t.Fatalf("Expect advance 5 character, got %d\n",eat)
	}
}

func TestDecodeString(t *testing.T) {
	str,eat,err := decodeString("4:spam")
	if err != nil {
		t.Fatalf("Expect 'spam', got error: %v\n",err)
	} else if str != "spam" {
		t.Fatalf("Expect 'spam', got '%s'\n",str)
	}
	if eat != 6 {
		t.Fatalf("Expect advance 6 character, got %d\n",eat)
	}
}

func TestDecodeList(t *testing.T) {
	data := "l4:spami42ee"
	list,eat,err := decodeList(data)
	if err != nil {
		t.Fatalf("Expect ['spam',42], got error: %v\n",err)
	}
	if eat != len(data) {
		t.Fatalf("Expect advance %d character, got %d\n",len(data),eat)
	}
	if len(list) == 2 {
		if str,ok := list[0].(string); !ok || str != "spam"{
			t.Fatalf("First item is expected to be string 'spam', got %v[%T]\n",list[0],list[0])
		}
		if num,ok := list[1].(int); !ok || num != 42 {
			t.Fatalf("Second item is expected to be int 42, got %d[%T]\n",list[1],list[1])
		}
	} else {
		t.Fatalf("Expect ['spam',42], got %v\n",list)
	}
}

func TestDecodeMap(t *testing.T) {
	target := map[string]interface{}{"bar": "spam", "foo": 42}
	data := "d3:bar4:spam3:fooi42ee"
	mp,eat,err := decodeMap(data)
	if err != nil {
		t.Fatalf("Expect {\"bar\": \"spam\", \"foo\": 42}, got error: %v\n",err)
	} else if eat != len(data) {
		t.Fatalf("Expect advance %d character, got %d\n",len(data),eat)
	}

	for k,v := range mp {
		if v2,ok := target[k]; !ok || v != v2 {
			t.Fatalf("Expect get %v, actual get %v\n",v2,v)
		}
	}
}

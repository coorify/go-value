# go-value

get value from struct with path

## usage

```go
	type object0 struct {
		A string
		B int
	}

	type object1 struct {
		C string
	}

	type object2 struct {
		D string
	}

	type object3 struct {
		object1

		o0 object0
		O0 object0

		O2 []object2
		O3 map[string]string
		O4 map[int]string
	}

	src := object3{
		object1: object1{
			C: "cccc",
		},
		o0: object0{
			A: "aaaa",
			B: -9999,
		},
		O0: object0{
			A: "AAAA",
			B: 999,
		},
		O2: []object2{
			{
				D: "D1",
			},
			{
				D: "D2",
			},
		},
		O3: map[string]string{
			"1": "O3 1",
		},
		O4: map[int]string{
			2: "O4 1",
		},
	}


  value.Get(src,"C") // cccc
  value.Get(src,"O0.A") // aaaa
  value.Get(src,"O2.0.D") // D1
  value.Get(src,"O3.1") // 03 1
```

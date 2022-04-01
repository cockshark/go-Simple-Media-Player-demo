package parser

var seeds = [][]byte{
	nil,
	[]byte("123"),
	[]byte("(12)"),
}

type ParseError interface {

}

//func FuzzParserSomething(f *testing.F)  {
//	for _, seed := range seeds {
//		f.Add(seed)
//	}
//	f.Fuzz(func(t *testing.T, input []byte) {
//		err := ParseSomething(input)
//		if err == nil {
//			return
//		}
//		if parseErr := (*ParseError)(nil);!errors.As(err, parseErr) {
//			t.Fatal(err)
//		}
//	})
//}

func ParseSomething(input []byte) interface{} {
	return nil
}

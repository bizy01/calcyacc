package parse

import (
   "testing"
)

func TestParse(t *testing.T) {
   tests := []struct {
       code string
       value Result
   }{
       {"1", 1},
       {"1+1", 2},
       {"1+2*3", 7},
       {"1+3/1", 4},
       {"1-1", 0},
       {"(1+2)*3", 9},
       {"(1+2)/3", 1},
       {"1+(2*3)", 7},
       {"3*2+1", 7},
       {"3*(2+1)", 9},
       {"1*2*3", 6},
       {"1+2+3", 6},
       {"1/2", 0.5},
       {"2/3", 0.6666666666666666},
   }

   for _, test := range tests {
       	value, err := Parse(test.code)

       	if err != nil {
       		t.Errorf("err %v", err)
       	} else if value != test.value {
           t.Errorf("err Actual: %.16f Expect: %f", value, test.value)
       	} else {
           t.Logf(" sucess %s = %f", test.code, test.value)
       	}
   }
}
package clevercalc

import (
	"reflect"
	"testing"

	"github.com/LexusEgorov/api-calculator/internal/models"
)

func Test_removeSpaces(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "regular case",
			input: "1 + 2 + 3 + 4",
			want:  "1+2+3+4",
		},
		{
			name:  "multy spaces",
			input: "1 + 2     -         4",
			want:  "1+2-4",
		},
		{
			name:  "without spaces",
			input: "1+2",
			want:  "1+2",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeSpaces(tt.input); got != tt.want {
				t.Errorf("removeSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addNum(t *testing.T) {
	testDestination := make([]string, 0)

	type args struct {
		destination *[]string
		num         string
	}
	tests := []struct {
		name      string
		input     string
		want      string
		args      args
		expectAdd bool
	}{
		{
			name:  "regular case",
			input: "10",
			args: args{
				destination: &testDestination,
			},
			want:      "10",
			expectAdd: true,
		},
		{
			name:  "with spaces",
			input: " 10",
			args: args{
				destination: &testDestination,
			},
			want:      "10",
			expectAdd: true,
		},
		{
			name:  "with spaces",
			input: " 10",
			args: args{
				destination: &testDestination,
			},
			want:      "10",
			expectAdd: true,
		},
		{
			name:  "only spaces",
			input: "  ",
			args: args{
				destination: &testDestination,
			},
			expectAdd: false,
		},
		{
			name:  "empty string",
			input: "",
			args: args{
				destination: &testDestination,
			},
			expectAdd: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.num = tt.input
			sizeBefore := len(*tt.args.destination)

			addNum(tt.args.destination, &tt.args.num)

			sizeAfter := len(*tt.args.destination)

			if !tt.expectAdd {
				if sizeAfter != sizeBefore {
					t.Errorf("parser.addNum(\"%s\"): destination slice was changed, but expected not", tt.input)
				}
				return
			}

			if sizeAfter-sizeBefore != 1 {
				t.Errorf("parser.addNum(\"%s\"): destination slice len %d, but expected %d", tt.input, len(*tt.args.destination), sizeBefore)
			}

			if dest := *tt.args.destination; dest[sizeAfter-1] != tt.want {
				t.Errorf("parser.addNum(\"%s\"): add want '%s', got '%s'", tt.input, tt.want, dest[sizeAfter-1])
			}

			if tt.args.num != "" {
				t.Errorf("source value didn't clean")
			}
		})
	}
}

func Test_parser_getActions(t *testing.T) {
	testStack := Stack{}
	testParser := newParser()

	type args struct {
		from Stack
		rank int
	}
	tests := []struct {
		name string
		p    parser
		args args
		init []string
		want []string
	}{
		{
			name: "lower rank",
			p:    *testParser,
			args: args{
				from: testStack,
				rank: 1,
			},
			init: []string{
				models.OperationSum,
				models.OperationSub,
				models.OperationMult,
				models.OperationDiv,
			},
			want: []string{
				models.OperationDiv,
				models.OperationMult,
				models.OperationSub,
				models.OperationSum,
			},
		},
		{
			name: "middle rank",
			p:    *testParser,
			args: args{
				from: testStack,
				rank: 2,
			},
			init: []string{
				models.OperationSum,
				models.OperationSub,
				models.OperationMult,
				models.OperationDiv,
			},
			want: []string{
				models.OperationDiv,
				models.OperationMult,
			},
		},
		{
			name: "middle rank (low upper)",
			p:    *testParser,
			args: args{
				from: testStack,
				rank: 2,
			},
			init: []string{
				models.OperationMult,
				models.OperationDiv,
				models.OperationSum,
				models.OperationSub,
			},
			want: []string{},
		},
		{
			name: "high rank",
			p:    *testParser,
			args: args{
				from: testStack,
				rank: 3,
			},
			init: []string{
				models.OperationSum,
				models.OperationSub,
				models.OperationMult,
				models.OperationDiv,
				models.OperationPow,
			},
			want: []string{
				models.OperationPow,
			},
		},
		{
			name: "high rank (without high)",
			p:    *testParser,
			args: args{
				from: testStack,
				rank: 3,
			},
			init: []string{
				models.OperationSum,
				models.OperationSub,
				models.OperationMult,
				models.OperationDiv,
			},
			want: []string{},
		},
		{
			name: "empty stack",
			p:    *testParser,
			args: args{
				from: testStack,
				rank: 1,
			},
			init: []string{},
			want: []string{},
		},
		{
			name: "brakes blocker",
			p:    *testParser,
			args: args{
				from: testStack,
				rank: 1,
			},
			init: []string{
				models.OperationSum,
				models.OperationSum,
				models.OpeningBrake,
				models.OperationSum,
				models.OperationSum,
			},
			want: []string{
				models.OperationSum,
				models.OperationSum,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.from.data = tt.init
			tt.args.from.size = len(tt.init)

			if got := tt.p.getActions(&tt.args.from, tt.args.rank); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parser.getActions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_prepareSimplePart(t *testing.T) {
	testParser := newParser()

	type args struct {
		part string
	}
	tests := []struct {
		name string
		p    parser
		args args
		want string
	}{
		{
			name: "empty string",
			p:    *testParser,
			args: args{
				part: "",
			},
			want: "",
		},
		{
			name: "starts from positive number",
			p:    *testParser,
			args: args{
				part: "10-4",
			},
			want: "10-4",
		},
		{
			name: "only sub symb",
			p:    *testParser,
			args: args{
				part: "-",
			},
			want: "0-",
		},
		{
			name: "negative brakes",
			p:    *testParser,
			args: args{
				part: "-(123*5)",
			},
			want: "(0-(123*5))",
		},
		{
			name: "starts from negative number",
			p:    *testParser,
			args: args{
				part: "-17*2)",
			},
			want: "(0-17)*2)",
		},
		{
			name: "negative number",
			p:    *testParser,
			args: args{
				part: "-4",
			},
			want: "(0-4)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.prepareSimplePart(tt.args.part); got != tt.want {
				t.Errorf("parser.prepareSimplePart(\"%v\") = %v, want %v", tt.args.part, got, tt.want)
			}
		})
	}
}

func Test_parser_prepare(t *testing.T) {
	testParser := newParser()

	type args struct {
		input string
	}
	tests := []struct {
		name string
		p    parser
		args args
		want string
	}{
		{
			name: "without negative numbers",
			p:    *testParser,
			args: args{
				input: "2+2",
			},
			want: "2+2",
		},
		{
			name: "negative numbers (not at start)",
			p:    *testParser,
			args: args{
				input: "2-2",
			},
			want: "2-2",
		},
		{
			name: "start from negative number",
			p:    *testParser,
			args: args{
				input: "-2+2",
			},
			want: "(0-2)+2",
		},
		{
			name: "action in brakes starts from negative number",
			p:    *testParser,
			args: args{
				input: "4*(-2+2)",
			},
			want: "4*((0-2)+2)",
		},
		{
			name: "after operator going negative operator",
			p:    *testParser,
			args: args{
				input: "4*-(2+2)",
			},
			want: "4*(0-(2+2))",
		},
		{
			name: "empty string",
			p:    *testParser,
			args: args{
				input: "",
			},
			want: "",
		},
		{
			name: "spaces",
			p:    *testParser,
			args: args{
				input: "2+            2",
			},
			want: "2+2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.prepare(tt.args.input); got != tt.want {
				t.Errorf("parser.prepare(\"%v\") = %v, want %v", tt.args.input, got, tt.want)
			}
		})
	}
}

func Test_parser_parse(t *testing.T) {
	testParser := newParser()

	type args struct {
		input string
	}
	tests := []struct {
		name    string
		p       parser
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "simple",
			p:    *testParser,
			args: args{
				input: "2+2",
			},
			want: []string{
				"2",
				"2",
				"+",
			},
			wantErr: false,
		},
		{
			name: "many actions",
			p:    *testParser,
			args: args{
				input: "8*((4+3/3)/8)+15",
			},
			want: []string{
				"8",
				"4",
				"3",
				"3",
				"/",
				"+",
				"8",
				"/",
				"*",
				"15",
				"+",
			},
			wantErr: false,
		},
		{
			name: "starts from negative",
			p:    *testParser,
			args: args{
				input: "-2+2",
			},
			want: []string{
				"0",
				"2",
				"-",
				"2",
				"+",
			},
			wantErr: false,
		},
		{
			name: "brakes problem #1",
			p:    *testParser,
			args: args{
				input: "4*(2+2",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "brakes problem #2",
			p:    *testParser,
			args: args{
				input: "4*(2+2))",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.parse(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parser.parse(\"%v\") error = %v, wantErr %v", tt.args.input, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parser.parse(\"%v\") = %v, want %v", tt.args.input, got, tt.want)
			}
		})
	}
}

package mappers

import (
	"container/list"
	"flag"
	"fmt"
	"os"

	"miller/clitypes"
	"miller/lib"
	"miller/mapping"
	"miller/types"
)

// ----------------------------------------------------------------
var TailSetup = mapping.MapperSetup{
	Verb:         "tail",
	ParseCLIFunc: mapperTailParseCLI,
	IgnoresInput: false,
}

func mapperTailParseCLI(
	pargi *int,
	argc int,
	args []string,
	errorHandling flag.ErrorHandling, // ContinueOnError or ExitOnError
	_ *clitypes.TReaderOptions,
	__ *clitypes.TWriterOptions,
) mapping.IRecordMapper {

	// Get the verb name from the current spot in the mlr command line
	argi := *pargi
	verb := args[argi]
	argi++

	// Parse local flags
	flagSet := flag.NewFlagSet(verb, errorHandling)

	//Usage: mlr tail [options]
	//-n {count}    Tail count to print; default 10
	//-g {a,b,c}    Optional group-by-field names for tail counts

	pTailCount := flagSet.Uint64(
		"n",
		10,
		`Tail count to print`,
	)

	pGroupByFieldNames := flagSet.String(
		"g",
		"",
		"Optional group-by-field names for tail counts, e.g. a,b,c",
	)

	flagSet.Usage = func() {
		ostream := os.Stderr
		if errorHandling == flag.ContinueOnError { // help intentionally requested
			ostream = os.Stdout
		}
		mapperTailUsage(ostream, args[0], verb, flagSet)
	}
	flagSet.Parse(args[argi:])
	if errorHandling == flag.ContinueOnError { // help intentioally requested
		return nil
	}

	// Find out how many flags were consumed by this verb and advance for the
	// next verb
	argi = len(args) - len(flagSet.Args())

	mapper, _ := NewMapperTail(
		*pTailCount,
		*pGroupByFieldNames,
	)

	*pargi = argi
	return mapper
}

func mapperTailUsage(
	o *os.File,
	argv0 string,
	verb string,
	flagSet *flag.FlagSet,
) {
	fmt.Fprintf(o, "Usage: %s %s [options]\n", argv0, verb)
	fmt.Fprint(o,
		`Passes through the first n records, optionally by category.  Without -g, ceases
consuming more input (i.e. is fast) when n records have been read.
`)
	// flagSet.PrintDefaults() doesn't let us control stdout vs stderr
	flagSet.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(o, " -%v (default %v) %v\n", f.Name, f.Value, f.Usage) // f.Name, f.Value
	})
}

// ----------------------------------------------------------------
type MapperTail struct {
	// input
	tailCount            uint64
	groupByFieldNameList []string

	// state
	recordListsByGroup map[string]*list.List
}

func NewMapperTail(
	tailCount uint64,
	groupByFieldNames string,
) (*MapperTail, error) {

	groupByFieldNameList := lib.SplitString(groupByFieldNames, ",")

	this := &MapperTail{
		tailCount:            tailCount,
		groupByFieldNameList: groupByFieldNameList,

		recordListsByGroup: make(map[string]*list.List),
	}

	return this, nil
}

// ----------------------------------------------------------------
func (this *MapperTail) Map(
	inrecAndContext *types.RecordAndContext,
	outputChannel chan<- *types.RecordAndContext,
) {
	inrec := inrecAndContext.Record
	if inrec != nil { // not end of record stream

		groupByKey, ok := inrec.GetSelectedValuesJoined(this.groupByFieldNameList)
		if !ok {
			return
		}

		recordListForGroup, present := this.recordListsByGroup[groupByKey]
		if !present { // first time
			recordListForGroup = list.New()
			this.recordListsByGroup[groupByKey] = recordListForGroup
		}

		recordListForGroup.PushBack(inrecAndContext)
		for uint64(recordListForGroup.Len()) > this.tailCount {
			recordListForGroup.Remove(recordListForGroup.Front())
		}

	} else {
		for _, recordListForGroup := range this.recordListsByGroup {
			for entry := recordListForGroup.Front(); entry != nil; entry = entry.Next() {
				outputChannel <- entry.Value.(*types.RecordAndContext)
			}
		}
		outputChannel <- inrecAndContext // end-of-stream marker
	}
}

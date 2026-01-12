package report

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func DisplayResults(results []map[string]string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "PACKAGE\tVERSION\tID\tSUMMARY")
	
	for _, res := range results {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", res["pkg"], res["ver"], res["id"], res["summary"])
	}
	w.Flush()
}
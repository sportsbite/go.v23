// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/peterh/liner"
	"v.io/syncbase/v23/syncbase/nosql/internal/query"
	"v.io/syncbase/v23/syncbase/nosql/internal/query/demo/db"
	"v.io/syncbase/v23/syncbase/nosql/internal/query/demo/reader"
	"v.io/syncbase/v23/syncbase/nosql/internal/query/query_db"
)

func dumpDB(d query_db.Database) {
	for _, table := range db.GetTableNames() {
		fmt.Printf("table: %s\n", table)
		queryExec(d, fmt.Sprintf("select k, v from %s", table))
	}
}

// Split an error message into an offset and the remaining (i.e., rhs of offset) message.
// The convention for syncql is "<module><optional-rpc>[offset}<remaining-message>".
func splitError(err error) (int64, string) {
	errMsg := err.Error()
	idx1 := strings.Index(errMsg, "[")
	idx2 := strings.Index(errMsg, "]")
	if idx1 == -1 || idx2 == -1 {
		return 0, errMsg
	}
	offsetString := errMsg[idx1+1 : idx2]
	offset, err := strconv.ParseInt(offsetString, 10, 64)
	if err != nil {
		return 0, errMsg
	}
	return offset, errMsg[idx2+1:]
}

func queryExec(d query_db.Database, q string) {
	if columnNames, rs, err := query.Exec(d, q); err != nil {
		off, msg := splitError(err)
		fmt.Fprintf(os.Stderr, "%s\n", q)
		fmt.Fprintf(os.Stderr, "%s^\n", strings.Repeat(" ", int(off)))
		fmt.Fprintf(os.Stderr, "Error: %s\n", msg)
	} else {
		sep := ""
		for _, cName := range columnNames {
			fmt.Printf("%s%s", sep, cName)
			sep = " | "
		}
		fmt.Printf("\n")
		for rs.Advance() {
			sep = ""
			for _, column := range rs.Result() {
				fmt.Printf("%s%v", sep, column)
				sep = " | "
			}
			fmt.Printf("\n")
		}
		if err := rs.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		}
	}
}

func main() {
	// TODO(kash): Detect if this is not a terminal and use
	// reader.NewNonInteractive.
	input := reader.NewInteractive()
	defer input.Close()

	d := db.GetDatabase()
	for true {
		if q, err := input.GetQuery(); err != nil {
			if err == io.EOF {
				// ctrl-d
				fmt.Println()
				break
			} else {
				// ctrl-c
				break
			}
		} else {
			if q == "" || strings.ToLower(q) == "exit" || strings.ToLower(q) == "quit" {
				break
			}
			line.AppendHistory(q)
			if strings.ToLower(q) == "dump" {
				dumpDB(d)
			} else {
				queryExec(d, q)
			}
		}
	}
}

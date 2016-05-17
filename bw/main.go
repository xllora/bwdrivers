// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Binary btbolt injects the btbolt driver into the bw command line tool.
package main

import (
	"flag"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/google/badwolf/storage"
	"github.com/google/badwolf/storage/memory"
	"github.com/google/badwolf/tools/vcli/bw/common"
	"github.com/google/badwolf/triple/literal"
	"github.com/xllora/bwdrivers/bwbolt"
)

var (
	// drivers contains the registed drivers available for this command line tool.
	registeredDrivers map[string]common.StoreGenerator

	// Available flags.
<<<<<<< HEAD
	driverName = flag.String("driver", "VOLATILE", "The storage driver to use {VOLATILE|BWBOLT}.")

	bqlChannelSize       = flag.Int("bql_channel_size", 0, "Internal channel size to use on BQL queries.")
	bulkTripleOpSize     = flag.Int("bulk_triple_op_size", 1000, "Number of triples to use in bulk load operations.")
	bulkTriplBuildersize = flag.Int("bulk_triple_builder_size_in_bytes", 1000, "Maximum size of literals when parsing a triple.")

=======
	driverName           = flag.String("driver", "VOLATILE", "The storage driver to use {VOLATILE|BWBOLT}.")
	bqlChannelSize       = flag.Int("bql_channel_size", 0, "Internal channel size to use on BQL queries.")
	bulkTripleOpSize     = flag.Int("bulk_triple_op_size", 1000, "Number of triples to use in bulk load operations.")
	bulkTriplBuildersize = flag.Int("bulk_triple_builder_size_in_bytes", 1000, "Maximum size of literals when parsing a triple.")
>>>>>>> 46f45036c6745a0ef584dfeb30c1c293dadbf4fa
	// Add your driver flags below.

	// BwBolt driver.
	boltDBPath   = flag.String("bolt_db_path", "", "The path to the Bolt database to use.")
	boldTimeout  = flag.Duration("bold_timeout", 3*time.Second, "The duration of the timeout while opening the Bolt database.")
	boltReadOnly = flag.Bool("bolt_db_read_only", false, "Use te Bolt DB only in read only mode.")

	// Driver specific variables.
	db *bolt.DB
)

// Registers the available drivers.
func registerDrivers() {
	registeredDrivers = map[string]common.StoreGenerator{
		// Memory only storage driver.
		"VOLATILE": func() (storage.Store, error) {
			return memory.NewStore(), nil
		},
		"BWBOLT": func() (storage.Store, error) {
			s, bdb, err := bwbolt.New(*boltDBPath, literal.DefaultBuilder(), *boldTimeout, *boltReadOnly)
			db = bdb
			return s, err
		},
	}
}

func main() {
	flag.Parse()
	registerDrivers()
	ret := common.Run(*driverName, registeredDrivers, *bqlChannelSize, *bulkTripleOpSize, *bulkTriplBuildersize)
	// Clean up.
	if db != nil {
		db.Close()
	}
	os.Exit(ret)
}

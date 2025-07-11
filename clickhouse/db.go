package main

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDb(dsn string) {
	db, err := gorm.Open(clickhouse.New(clickhouse.Config{
		DSN:                          dsn,
		DisableDatetimePrecision:     true,     // disable datetime64 precision, not supported before clickhouse 20.4
		DontSupportRenameColumn:      true,     // rename column not supported before clickhouse 20.4
		DontSupportEmptyDefaultValue: false,    // do not consider empty strings as valid default values
		SkipInitializeWithVersion:    false,    // smart configure based on used version
		DefaultGranularity:           3,        // 1 granule = 8192 rows
		DefaultCompression:           "LZ4",    // default compression algorithm. LZ4 is lossless
		DefaultIndexType:             "minmax", // index stores extremes of the expression
		DefaultTableEngineOpts:       "ENGINE=MergeTree() ORDER BY tuple()",
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}

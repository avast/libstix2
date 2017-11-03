// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/vocabs"
	"log"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// CreateAllVocabTables - This method will create all of the tables needed to store
// STIX content in the database.
func (ds *Sqlite3DatastoreType) CreateAllVocabTables() {
	ds.createTable(DB_TABLE_VOCAB_ATTACK_MOTIVATIONS, ds.vocabProperties())
	ds.createTable(DB_TABLE_VOCAB_ATTACK_RESOURCE_LEVEL, ds.vocabProperties())
	ds.createTable(DB_TABLE_VOCAB_IDENTITY_CLASS, ds.vocabProperties())
	ds.createTable(DB_TABLE_VOCAB_INDICATOR_LABEL, ds.vocabProperties())
	ds.createTable(DB_TABLE_VOCAB_INDUSTRY_SECTOR, ds.vocabProperties())
	ds.createTable(DB_TABLE_VOCAB_MALWARE_LABEL, ds.vocabProperties())
	ds.createTable(DB_TABLE_VOCAB_REPORT_LABEL, ds.vocabProperties())
	ds.createTable(DB_TABLE_VOCAB_THREAT_ACTOR_LABEL, ds.vocabProperties())
	ds.createTable(DB_TABLE_VOCAB_THREAT_ACTOR_ROLE, ds.vocabProperties())
	ds.createTable(DB_TABLE_VOCAB_THREAT_ACTOR_SOPHISTICATION, ds.vocabProperties())
	ds.createTable(DB_TABLE_VOCAB_TOOL_LABEL, ds.vocabProperties())
}

// PopulateAllVocabTables - This method will insert all of the vocabulary data
// into the right database tables.
func (ds *Sqlite3DatastoreType) PopulateAllVocabTables() {
	ds.insertVocabData(DB_TABLE_VOCAB_ATTACK_MOTIVATIONS, vocabs.AttackMotivation)
	ds.insertVocabData(DB_TABLE_VOCAB_ATTACK_RESOURCE_LEVEL, vocabs.AttackResourceLevel)
	ds.insertVocabData(DB_TABLE_VOCAB_IDENTITY_CLASS, vocabs.IdentityClass)
	ds.insertVocabData(DB_TABLE_VOCAB_INDICATOR_LABEL, vocabs.IndicatorLabel)
	ds.insertVocabData(DB_TABLE_VOCAB_INDUSTRY_SECTOR, vocabs.IndustrySector)
	ds.insertVocabData(DB_TABLE_VOCAB_MALWARE_LABEL, vocabs.MalwareLabel)
	ds.insertVocabData(DB_TABLE_VOCAB_REPORT_LABEL, vocabs.ReportLabel)
	ds.insertVocabData(DB_TABLE_VOCAB_THREAT_ACTOR_LABEL, vocabs.ThreatActorLabel)
	ds.insertVocabData(DB_TABLE_VOCAB_THREAT_ACTOR_ROLE, vocabs.ThreatActorRole)
	ds.insertVocabData(DB_TABLE_VOCAB_THREAT_ACTOR_SOPHISTICATION, vocabs.ThreatActorSophistication)
	ds.insertVocabData(DB_TABLE_VOCAB_TOOL_LABEL, vocabs.ToolLabel)
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

// vocabProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) vocabProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"value" text NOT NULL
	`
}

// InsertVocabData - This method will add a vocabulary item to its table
func (ds *Sqlite3DatastoreType) insertVocabData(name string, data []string) {
	var stmt = `INSERT INTO "` + name + `" (value) values (?)`

	var err error
	for _, value := range data {
		_, err = ds.DB.Exec(stmt, value)
	}

	if err != nil {
		log.Println("ERROR: The vocabulary item could not be inserted in to the", name, "table")
	}
}

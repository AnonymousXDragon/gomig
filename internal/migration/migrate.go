package migration

import (
	"database/sql"
	"fmt"
	"gomig/internal/parser"
	"path/filepath"
	"sort"
)

type Migrater struct {
	db         *sql.DB
	migrations []*parser.Migration
}

func NewMigrater(db *sql.DB, dir string) (*Migrater, error) {
	migrations, err := loadMigrationsFrom(dir)

	if err != nil {
		return nil, err
	}

	return &Migrater{
		db,
		migrations,
	}, nil
}

func loadMigrationsFrom(dir string) ([]*parser.Migration, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.sql"))

	if err != nil {
		return nil, err
	}

	var migrations []*parser.Migration

	for _, file := range files {
		migrateObj, err := parser.ParseMigFile(file)
		if err != nil {
			return nil, err
		}

		migrations = append(migrations, migrateObj)
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil

}

func (m *Migrater) Up() error {

	for _, migFile := range m.migrations {
		if err := m.InsertData(migFile.Up); err != nil {
			return err
		}

		fmt.Printf("successfully migrated: %d-%s \n", migFile.Version, migFile.Name)
	}

	return nil
}

func (m *Migrater) Down() error {

	for _, migFile := range m.migrations {
		if err := m.InsertData(migFile.Down); err != nil {
			return err
		}

		fmt.Printf("successfully reverted: %d-%s \n", migFile.Version, migFile.Name)
	}
	return nil
}

func (m *Migrater) InsertData(sql string) error {
	_, err := m.db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

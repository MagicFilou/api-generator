package resources

import (
	"fmt"
	"io/ioutil"
	"strings"

	cfg "api-builder/configs"
	"api-builder/utils/translator"
)

func (r Resource) ToSQLCreateTable() string {

	b := strings.Builder{}

	b.WriteString("CREATE OR REPLACE FUNCTION updated_timestamp()\n")
	b.WriteString("RETURNS TRIGGER AS $updated_timestamp$\n")
	b.WriteString("BEGIN\n")
	b.WriteString("  NEW.updated = date_part('epoch'::text, now());\n")
	b.WriteString("  RETURN NEW;\n")
	b.WriteString("END;\n")
	b.WriteString("$updated_timestamp$ language 'plpgsql';\n\n")

	b.WriteString("CREATE TABLE IF NOT EXISTS\n")
	b.WriteString(r.PluralUnderscored + " (\n")

	for index, field := range DefaultFields {
		b.WriteString(field.Name)
		b.WriteString(" " + translator.ToSQL(field.DataType))
		if field.Name == "id" {
			b.WriteString(" PRIMARY KEY")
		}
		if field.Name == "created" || field.Name == "updated" {
			b.WriteString(" DEFAULT date_part('epoch'::text, now())")
		}
		if index != len(DefaultFields)-1 || len(r.Storage.Fields) != 0 {
			b.WriteString(",\n")
		} else {
			b.WriteString("\n")
		}
	}

	// extract(epoch from now())

	for index, field := range r.Storage.Fields {
		b.WriteString(field.Name)
		b.WriteString(" " + translator.ToSQL(field.DataType))
		for _, constraint := range field.Constraints {
			if constraint.Type == "sql" {
				b.WriteString(" " + strings.ToUpper(constraint.Value))
			}
		}
		if field.Default != "" {
			b.WriteString(" DEFAULT '" + field.Default + "'")
		}
		if field.Relation.Resource != "" {
			b.WriteString(",\nFOREIGN KEY (" + field.Name + ") REFERENCES " + field.Relation.Resource + "(" + field.Relation.Field + ") ON DELETE CASCADE")
		}

		if index != len(r.Storage.Fields)-1 {
			b.WriteString(",\n")
		}
	}

	b.WriteString("\n);\n")

	if r.Storage.Config.DefaultFields {
		b.WriteString("\nCREATE TRIGGER " + cfg.UPDATE_TIMESTAMP_TRIGGER + " BEFORE UPDATE ON " + r.PluralUnderscored + " FOR EACH ROW EXECUTE FUNCTION " + cfg.UPDATE_TIMESTAMP_TRIGGER + "();")
	}

	return b.String()
}

func (r Resource) ToSQLDropTable() string {

	b := strings.Builder{}

	if r.Storage.Config.DefaultFields {
		b.WriteString("DROP TRIGGER IF EXISTS " + cfg.UPDATE_TIMESTAMP_TRIGGER + " ON " + r.PluralUnderscored + ";\n")
	}
	b.WriteString("DROP TABLE IF EXISTS " + r.PluralUnderscored + ";")

	return b.String()
}

func (r Resource) WriteSQLFiles(path string, migrationVersion int) error {

	// fmt.Println("write up")
	err := ioutil.WriteFile(path+"/"+fmt.Sprintf("%06d", migrationVersion)+"_"+r.Name.PluralUnderscored+".up.sql", []byte(r.ToSQLCreateTable()), 0644)
	if err != nil {
		return err
	}

	// fmt.Println("write down")
	err = ioutil.WriteFile(path+"/"+fmt.Sprintf("%06d", migrationVersion)+"_"+r.Name.PluralUnderscored+".down.sql", []byte(r.ToSQLDropTable()), 0644)
	if err != nil {
		return err
	}

	fmt.Println("done")
	return nil
}

//State disabled
// func (r ResourceWithState) HandleFieldChanges(rootPath string, migrationVersion int) error {

// 	newFields, removedFields, changed := r.calculateChanges()

// 	if changed {
// 		for _, field := range newFields {

// 			err := field.WriteAddMigration(rootPath, r.PluralUnderscored, migrationVersion)
// 			if err != nil {
// 				return err
// 			}
// 			migrationVersion++
// 		}

// 		for _, field := range removedFields {
// 			err := field.WriteDropMigration(rootPath, r.PluralUnderscored, migrationVersion)
// 			if err != nil {
// 				return err
// 			}
// 			migrationVersion++
// 		}
// 	}

// 	return nil
// }

func (f Field) WriteAddMigration(path, tablename string, migrationVersion int) error {

	err := ioutil.WriteFile(path+"/"+fmt.Sprintf("%06d", migrationVersion)+"_"+tablename+"-add-"+f.Name+".down.sql", []byte(f.ToSQLDropColumn(tablename)), 0644)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path+"/"+fmt.Sprintf("%06d", migrationVersion)+"_"+tablename+"-add-"+f.Name+".up.sql", []byte(f.toSQLAddColumn(tablename)), 0644)
}

func (f Field) WriteDropMigration(path, tablename string, migrationVersion int) error {

	err := ioutil.WriteFile(path+"/"+fmt.Sprintf("%06d", migrationVersion)+"_"+tablename+"-drop-"+f.Name+".down.sql", []byte(f.toSQLAddColumn(tablename)), 0644)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path+"/"+fmt.Sprintf("%06d", migrationVersion)+"_"+tablename+"-drop-"+f.Name+".up.sql", []byte(f.ToSQLDropColumn(tablename)), 0644)
}

func (f Field) toSQLAddColumn(tablename string) string {

	b := strings.Builder{}

	b.WriteString("ALTER TABLE " + tablename)
	b.WriteString(" ADD COLUMN " + f.Name)
	b.WriteString(" " + translator.ToSQL(f.DataType) + ";\n")

	return b.String()
}

func (f Field) ToSQLDropColumn(tablename string) string {

	b := strings.Builder{}

	b.WriteString("ALTER TABLE " + tablename)
	b.WriteString(" DROP COLUMN " + f.Name + ";")

	return b.String()
}

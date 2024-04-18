package main

import (
	"fmt"
	"generationsTables/schema"
	"reflect"
	"strings"
)

func generateSQLFromStruct(schema interface{}, tableName string) string {
	var sql strings.Builder

	sql.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", tableName))

	t := reflect.TypeOf(schema)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name
		fieldType := field.Tag.Get("size") // Obtém o tipo do campo da tag 'size'
		fieldSize := field.Tag.Get("maxlen") // Obtém o tamanho máximo do campo da tag 'maxlen'
		nullable := true                   // Define como true por padrão

		// Verifica se o campo é obrigatório ('fixed' ou 'llvar')
		if fieldType == "fixed" || fieldType == "llvar" || fieldType == "lllvar" {
			nullable = false
		}

		sql.WriteString(fmt.Sprintf("    %s %s", fieldName, getSQLType(fieldType, fieldSize)))

		// Adiciona a cláusula NOT NULL se o campo não for nulo
		if !nullable {
			sql.WriteString(" NOT NULL")
		}

		sql.WriteString(",\n")
	}

	// Remove a vírgula extra no final da última linha
	sqlStr := strings.TrimSuffix(sql.String(), ",\n")
	sql.Reset()
	sql.WriteString(sqlStr)

	sql.WriteString("\n)")

	return sql.String()
}


// Função auxiliar para obter o tipo SQL correspondente ao tipo de campo especificado
func getSQLType(fieldType, size string) string {
	switch fieldType {
	case "fixed":
		return fmt.Sprintf("VARCHAR(%s)", size)
	case "llvar":
		return "VARCHAR(255)" // Ou use o tamanho máximo desejado
	case "lllvar":
		return "TEXT"
	default:
		return "VARCHAR(255)"
	}
}

func main() {
	schema := schema.Person{}

	sql := generateSQLFromStruct(schema, "Person")
	fmt.Println(sql)
}

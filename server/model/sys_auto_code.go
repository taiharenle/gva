package model

// 初始版本自动化代码工具
type AutoCodeStruct struct {
	StructName         string  `json:"structName"`
	PackageName        string  `json:"packageName"`
	Abbreviation       string  `json:"abbreviation"`
	Description        string  `json:"description"`
	AutoCreateApiToSql bool    `json:"autoCreateApiToSql"`
	Fields             []Field `json:"fields"`
}

type Field struct {
	FieldName       string `json:"fieldName"`
	FieldDesc       string `json:"fieldDesc"`
	FieldType       string `json:"fieldType"`
	FieldJson       string `json:"fieldJson"`
	Comment         string `json:"comment"`
	ColumnName      string `json:"columnName"`
	FieldSearchType string `json:"fieldSearchType"`
}

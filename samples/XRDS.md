
### xjetpostgresqls.database.wk
| Name | Claim |
|------|-------|
| xjetpostgresqls.database.wk | JetPostgreSQL |
#### Specs
##### Version: v1alpha1
| Field | Path | Type | Description | Required |
|------|-------|------|-------|-------|
| spec |  | object | Required spec field for your API | false |
| parameters | spec | object |  | true |
| dbName | spec.parameters | string | name of the new DB inside the DB instance - string | true |
| instanceSize | spec.parameters | string | instance size - string | true |
| storageGB | spec.parameters | integer | size of the Database in GB - integer | true |


### xjetpostgresqls.database.wk
| Name | Claim |
|------|-------|
| xjetpostgresqls.database.wk | JetPostgreSQL |
#### Specs
##### Version: v1alpha1
| Field | Path | Type | Description | Default |  Required |
|------|-------|------|-------|-------|-------|
| spec |  | object | Required spec field for your API | n/a | false |
| parameters | spec | object |  | n/a | true |
| instanceSize | spec.parameters | string | instance size - string | "large" | true |
| storageGB | spec.parameters | integer | size of the Database in GB - integer | 2 | true |
| dbName | spec.parameters | string | name of the new DB inside the DB instance - string | n/a | true |

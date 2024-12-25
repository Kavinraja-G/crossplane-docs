
### jetpostgresql.gcp.database.wk
| XR Kind | XR Version |
|---------|-------------|
| XJetPostgreSQL | database.wk/v1alpha1 |
#### XRD
| Name | Claim |
|------|-------|
| xjetpostgresqls.database.wk | JetPostgreSQL |
#### XRD Spec
##### Version: v1alpha1
| Field | Path | Type | Description | Default | Required |
|------|-------|------|-------|-------|-------|
| spec |  | object | Required spec field for your API | n/a | false |
| parameters | spec | object |  | n/a | true |
| dbName | spec.parameters | string | name of the new DB inside the DB instance - string | n/a | true |
| instanceSize | spec.parameters | string | instance size - string | "large" | true |
| storageGB | spec.parameters | integer | size of the Database in GB - integer | 2 | true |
#### Resources
| Name | Kind | API Version |
|------|------|-------------|
| cloudsqlinstance | DatabaseInstance | sql.gcp.jet.crossplane.io/v1alpha2 |
| cloudsqldb | Database | sql.gcp.jet.crossplane.io/v1alpha2 |
| cloudsqldbuser | User | sql.gcp.jet.crossplane.io/v1alpha2 |

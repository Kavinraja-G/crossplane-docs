
### xjetpostgresqls.database.wk
| Name | Claim |
|------|-------|
| xjetpostgresqls.database.wk | JetPostgreSQL |
#### Specs
##### Version: v1alpha1
##### Inputs
| Field | Path | Type | Description | Default |  Required |
|------|-------|------|-------|-------|-------|
| spec |  | object | Required spec field for your API | n/a | false |
| parameters | spec | object |  | n/a | true |
| dbName | spec.parameters | string | name of the new DB inside the DB instance - string | n/a | true |
| instanceSize | spec.parameters | string | instance size - string | "large" | true |
| storageGB | spec.parameters | integer | size of the Database in GB - integer | 2 | true |
##### Outputs
| Field | Path | Type | Description | Default |  Required |
|------|-------|------|-------|-------|-------|
| status |  | object |  | n/a | false |
| dbConnectionURL | status | string | String that contains the URL of the Database connection | n/a | false |
| status | status | string | Status of the Database itself | n/a | false |
| upgradeEligibility | status | boolean | Flag to check eligibility for DB upgrade | n/a | false |

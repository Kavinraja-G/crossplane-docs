
### table-1.dynamodb.awsblueprints.io
| XR Kind | XR APIVersion |
|---------|-------------|
| XDynamoDBTable | awsblueprints.io/v1alpha1 |
#### XRD
| Name | Claim |
|------|-------|
|  |  |
#### Resources
| Name | Kind | API Version |
|------|------|-------------|
| costTable | Table | dynamodb.aws.upbound.io/v1beta1 |
| genericTable | Table | dynamodb.aws.upbound.io/v1beta1 |
### table-2.dynamodb.awsblueprints.io
| XR Kind | XR APIVersion |
|---------|-------------|
| SampleDBTable | awsblueprints.io/v1alpha1 |
#### XRD
| Name | Claim |
|------|-------|
|  |  |
#### Resources
| Name | Kind | API Version |
|------|------|-------------|
| userTable | Table | dynamodb.aws.upbound.io/v1beta1 |
| adminTable | Table | dynamodb.aws.upbound.io/v1beta1 |
### jetpostgresql.gcp.database.wk
| XR Kind | XR APIVersion |
|---------|-------------|
| XJetPostgreSQL | database.wk/v1alpha1 |
#### XRD
| Name | Claim |
|------|-------|
| xjetpostgresqls.database.wk | JetPostgreSQL |
#### Resources
| Name | Kind | API Version |
|------|------|-------------|
| cloudsqlinstance | DatabaseInstance | sql.gcp.jet.crossplane.io/v1alpha2 |
| cloudsqldb | Database | sql.gcp.jet.crossplane.io/v1alpha2 |
| cloudsqldbuser | User | sql.gcp.jet.crossplane.io/v1alpha2 |

// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AnalysisErrorType string

// Enum values for AnalysisErrorType
const (
	AnalysisErrorTypeAccessDenied                 AnalysisErrorType = "ACCESS_DENIED"
	AnalysisErrorTypeSourceNotFound               AnalysisErrorType = "SOURCE_NOT_FOUND"
	AnalysisErrorTypeDataSetNotFound              AnalysisErrorType = "DATA_SET_NOT_FOUND"
	AnalysisErrorTypeInternalFailure              AnalysisErrorType = "INTERNAL_FAILURE"
	AnalysisErrorTypeParameterValueIncompatible   AnalysisErrorType = "PARAMETER_VALUE_INCOMPATIBLE"
	AnalysisErrorTypeParameterTypeInvalid         AnalysisErrorType = "PARAMETER_TYPE_INVALID"
	AnalysisErrorTypeParameterNotFound            AnalysisErrorType = "PARAMETER_NOT_FOUND"
	AnalysisErrorTypeColumnTypeMismatch           AnalysisErrorType = "COLUMN_TYPE_MISMATCH"
	AnalysisErrorTypeColumnGeographicRoleMismatch AnalysisErrorType = "COLUMN_GEOGRAPHIC_ROLE_MISMATCH"
	AnalysisErrorTypeColumnReplacementMissing     AnalysisErrorType = "COLUMN_REPLACEMENT_MISSING"
)

// Values returns all known values for AnalysisErrorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AnalysisErrorType) Values() []AnalysisErrorType {
	return []AnalysisErrorType{
		"ACCESS_DENIED",
		"SOURCE_NOT_FOUND",
		"DATA_SET_NOT_FOUND",
		"INTERNAL_FAILURE",
		"PARAMETER_VALUE_INCOMPATIBLE",
		"PARAMETER_TYPE_INVALID",
		"PARAMETER_NOT_FOUND",
		"COLUMN_TYPE_MISMATCH",
		"COLUMN_GEOGRAPHIC_ROLE_MISMATCH",
		"COLUMN_REPLACEMENT_MISSING",
	}
}

type AnalysisFilterAttribute string

// Enum values for AnalysisFilterAttribute
const (
	AnalysisFilterAttributeQuicksightUser AnalysisFilterAttribute = "QUICKSIGHT_USER"
)

// Values returns all known values for AnalysisFilterAttribute. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AnalysisFilterAttribute) Values() []AnalysisFilterAttribute {
	return []AnalysisFilterAttribute{
		"QUICKSIGHT_USER",
	}
}

type AssignmentStatus string

// Enum values for AssignmentStatus
const (
	AssignmentStatusEnabled  AssignmentStatus = "ENABLED"
	AssignmentStatusDraft    AssignmentStatus = "DRAFT"
	AssignmentStatusDisabled AssignmentStatus = "DISABLED"
)

// Values returns all known values for AssignmentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssignmentStatus) Values() []AssignmentStatus {
	return []AssignmentStatus{
		"ENABLED",
		"DRAFT",
		"DISABLED",
	}
}

type ColumnDataType string

// Enum values for ColumnDataType
const (
	ColumnDataTypeString   ColumnDataType = "STRING"
	ColumnDataTypeInteger  ColumnDataType = "INTEGER"
	ColumnDataTypeDecimal  ColumnDataType = "DECIMAL"
	ColumnDataTypeDatetime ColumnDataType = "DATETIME"
)

// Values returns all known values for ColumnDataType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ColumnDataType) Values() []ColumnDataType {
	return []ColumnDataType{
		"STRING",
		"INTEGER",
		"DECIMAL",
		"DATETIME",
	}
}

type ColumnTagName string

// Enum values for ColumnTagName
const (
	ColumnTagNameColumnGeographicRole ColumnTagName = "COLUMN_GEOGRAPHIC_ROLE"
	ColumnTagNameColumnDescription    ColumnTagName = "COLUMN_DESCRIPTION"
)

// Values returns all known values for ColumnTagName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ColumnTagName) Values() []ColumnTagName {
	return []ColumnTagName{
		"COLUMN_GEOGRAPHIC_ROLE",
		"COLUMN_DESCRIPTION",
	}
}

type DashboardBehavior string

// Enum values for DashboardBehavior
const (
	DashboardBehaviorEnabled  DashboardBehavior = "ENABLED"
	DashboardBehaviorDisabled DashboardBehavior = "DISABLED"
)

// Values returns all known values for DashboardBehavior. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DashboardBehavior) Values() []DashboardBehavior {
	return []DashboardBehavior{
		"ENABLED",
		"DISABLED",
	}
}

type DashboardErrorType string

// Enum values for DashboardErrorType
const (
	DashboardErrorTypeAccessDenied                 DashboardErrorType = "ACCESS_DENIED"
	DashboardErrorTypeSourceNotFound               DashboardErrorType = "SOURCE_NOT_FOUND"
	DashboardErrorTypeDataSetNotFound              DashboardErrorType = "DATA_SET_NOT_FOUND"
	DashboardErrorTypeInternalFailure              DashboardErrorType = "INTERNAL_FAILURE"
	DashboardErrorTypeParameterValueIncompatible   DashboardErrorType = "PARAMETER_VALUE_INCOMPATIBLE"
	DashboardErrorTypeParameterTypeInvalid         DashboardErrorType = "PARAMETER_TYPE_INVALID"
	DashboardErrorTypeParameterNotFound            DashboardErrorType = "PARAMETER_NOT_FOUND"
	DashboardErrorTypeColumnTypeMismatch           DashboardErrorType = "COLUMN_TYPE_MISMATCH"
	DashboardErrorTypeColumnGeographicRoleMismatch DashboardErrorType = "COLUMN_GEOGRAPHIC_ROLE_MISMATCH"
	DashboardErrorTypeColumnReplacementMissing     DashboardErrorType = "COLUMN_REPLACEMENT_MISSING"
)

// Values returns all known values for DashboardErrorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DashboardErrorType) Values() []DashboardErrorType {
	return []DashboardErrorType{
		"ACCESS_DENIED",
		"SOURCE_NOT_FOUND",
		"DATA_SET_NOT_FOUND",
		"INTERNAL_FAILURE",
		"PARAMETER_VALUE_INCOMPATIBLE",
		"PARAMETER_TYPE_INVALID",
		"PARAMETER_NOT_FOUND",
		"COLUMN_TYPE_MISMATCH",
		"COLUMN_GEOGRAPHIC_ROLE_MISMATCH",
		"COLUMN_REPLACEMENT_MISSING",
	}
}

type DashboardFilterAttribute string

// Enum values for DashboardFilterAttribute
const (
	DashboardFilterAttributeQuicksightUser DashboardFilterAttribute = "QUICKSIGHT_USER"
)

// Values returns all known values for DashboardFilterAttribute. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DashboardFilterAttribute) Values() []DashboardFilterAttribute {
	return []DashboardFilterAttribute{
		"QUICKSIGHT_USER",
	}
}

type DashboardUIState string

// Enum values for DashboardUIState
const (
	DashboardUIStateExpanded  DashboardUIState = "EXPANDED"
	DashboardUIStateCollapsed DashboardUIState = "COLLAPSED"
)

// Values returns all known values for DashboardUIState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DashboardUIState) Values() []DashboardUIState {
	return []DashboardUIState{
		"EXPANDED",
		"COLLAPSED",
	}
}

type DataSetImportMode string

// Enum values for DataSetImportMode
const (
	DataSetImportModeSpice       DataSetImportMode = "SPICE"
	DataSetImportModeDirectQuery DataSetImportMode = "DIRECT_QUERY"
)

// Values returns all known values for DataSetImportMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataSetImportMode) Values() []DataSetImportMode {
	return []DataSetImportMode{
		"SPICE",
		"DIRECT_QUERY",
	}
}

type DataSourceErrorInfoType string

// Enum values for DataSourceErrorInfoType
const (
	DataSourceErrorInfoTypeAccessDenied              DataSourceErrorInfoType = "ACCESS_DENIED"
	DataSourceErrorInfoTypeCopySourceNotFound        DataSourceErrorInfoType = "COPY_SOURCE_NOT_FOUND"
	DataSourceErrorInfoTypeTimeout                   DataSourceErrorInfoType = "TIMEOUT"
	DataSourceErrorInfoTypeEngineVersionNotSupported DataSourceErrorInfoType = "ENGINE_VERSION_NOT_SUPPORTED"
	DataSourceErrorInfoTypeUnknownHost               DataSourceErrorInfoType = "UNKNOWN_HOST"
	DataSourceErrorInfoTypeGenericSqlFailure         DataSourceErrorInfoType = "GENERIC_SQL_FAILURE"
	DataSourceErrorInfoTypeConflict                  DataSourceErrorInfoType = "CONFLICT"
	DataSourceErrorInfoTypeUnknown                   DataSourceErrorInfoType = "UNKNOWN"
)

// Values returns all known values for DataSourceErrorInfoType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceErrorInfoType) Values() []DataSourceErrorInfoType {
	return []DataSourceErrorInfoType{
		"ACCESS_DENIED",
		"COPY_SOURCE_NOT_FOUND",
		"TIMEOUT",
		"ENGINE_VERSION_NOT_SUPPORTED",
		"UNKNOWN_HOST",
		"GENERIC_SQL_FAILURE",
		"CONFLICT",
		"UNKNOWN",
	}
}

type DataSourceType string

// Enum values for DataSourceType
const (
	DataSourceTypeAdobeAnalytics      DataSourceType = "ADOBE_ANALYTICS"
	DataSourceTypeAmazonElasticsearch DataSourceType = "AMAZON_ELASTICSEARCH"
	DataSourceTypeAthena              DataSourceType = "ATHENA"
	DataSourceTypeAurora              DataSourceType = "AURORA"
	DataSourceTypeAuroraPostgresql    DataSourceType = "AURORA_POSTGRESQL"
	DataSourceTypeAwsIotAnalytics     DataSourceType = "AWS_IOT_ANALYTICS"
	DataSourceTypeGithub              DataSourceType = "GITHUB"
	DataSourceTypeJira                DataSourceType = "JIRA"
	DataSourceTypeMariadb             DataSourceType = "MARIADB"
	DataSourceTypeMysql               DataSourceType = "MYSQL"
	DataSourceTypeOracle              DataSourceType = "ORACLE"
	DataSourceTypePostgresql          DataSourceType = "POSTGRESQL"
	DataSourceTypePresto              DataSourceType = "PRESTO"
	DataSourceTypeRedshift            DataSourceType = "REDSHIFT"
	DataSourceTypeS3                  DataSourceType = "S3"
	DataSourceTypeSalesforce          DataSourceType = "SALESFORCE"
	DataSourceTypeServicenow          DataSourceType = "SERVICENOW"
	DataSourceTypeSnowflake           DataSourceType = "SNOWFLAKE"
	DataSourceTypeSpark               DataSourceType = "SPARK"
	DataSourceTypeSqlserver           DataSourceType = "SQLSERVER"
	DataSourceTypeTeradata            DataSourceType = "TERADATA"
	DataSourceTypeTwitter             DataSourceType = "TWITTER"
	DataSourceTypeTimestream          DataSourceType = "TIMESTREAM"
	DataSourceTypeAmazonOpensearch    DataSourceType = "AMAZON_OPENSEARCH"
)

// Values returns all known values for DataSourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceType) Values() []DataSourceType {
	return []DataSourceType{
		"ADOBE_ANALYTICS",
		"AMAZON_ELASTICSEARCH",
		"ATHENA",
		"AURORA",
		"AURORA_POSTGRESQL",
		"AWS_IOT_ANALYTICS",
		"GITHUB",
		"JIRA",
		"MARIADB",
		"MYSQL",
		"ORACLE",
		"POSTGRESQL",
		"PRESTO",
		"REDSHIFT",
		"S3",
		"SALESFORCE",
		"SERVICENOW",
		"SNOWFLAKE",
		"SPARK",
		"SQLSERVER",
		"TERADATA",
		"TWITTER",
		"TIMESTREAM",
		"AMAZON_OPENSEARCH",
	}
}

type Edition string

// Enum values for Edition
const (
	EditionStandard   Edition = "STANDARD"
	EditionEnterprise Edition = "ENTERPRISE"
)

// Values returns all known values for Edition. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Edition) Values() []Edition {
	return []Edition{
		"STANDARD",
		"ENTERPRISE",
	}
}

type EmbeddingIdentityType string

// Enum values for EmbeddingIdentityType
const (
	EmbeddingIdentityTypeIam        EmbeddingIdentityType = "IAM"
	EmbeddingIdentityTypeQuicksight EmbeddingIdentityType = "QUICKSIGHT"
	EmbeddingIdentityTypeAnonymous  EmbeddingIdentityType = "ANONYMOUS"
)

// Values returns all known values for EmbeddingIdentityType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EmbeddingIdentityType) Values() []EmbeddingIdentityType {
	return []EmbeddingIdentityType{
		"IAM",
		"QUICKSIGHT",
		"ANONYMOUS",
	}
}

type ExceptionResourceType string

// Enum values for ExceptionResourceType
const (
	ExceptionResourceTypeUser                ExceptionResourceType = "USER"
	ExceptionResourceTypeGroup               ExceptionResourceType = "GROUP"
	ExceptionResourceTypeNamespace           ExceptionResourceType = "NAMESPACE"
	ExceptionResourceTypeAccountSettings     ExceptionResourceType = "ACCOUNT_SETTINGS"
	ExceptionResourceTypeIampolicyAssignment ExceptionResourceType = "IAMPOLICY_ASSIGNMENT"
	ExceptionResourceTypeDataSource          ExceptionResourceType = "DATA_SOURCE"
	ExceptionResourceTypeDataSet             ExceptionResourceType = "DATA_SET"
	ExceptionResourceTypeVpcConnection       ExceptionResourceType = "VPC_CONNECTION"
	ExceptionResourceTypeIngestion           ExceptionResourceType = "INGESTION"
)

// Values returns all known values for ExceptionResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExceptionResourceType) Values() []ExceptionResourceType {
	return []ExceptionResourceType{
		"USER",
		"GROUP",
		"NAMESPACE",
		"ACCOUNT_SETTINGS",
		"IAMPOLICY_ASSIGNMENT",
		"DATA_SOURCE",
		"DATA_SET",
		"VPC_CONNECTION",
		"INGESTION",
	}
}

type FileFormat string

// Enum values for FileFormat
const (
	FileFormatCsv  FileFormat = "CSV"
	FileFormatTsv  FileFormat = "TSV"
	FileFormatClf  FileFormat = "CLF"
	FileFormatElf  FileFormat = "ELF"
	FileFormatXlsx FileFormat = "XLSX"
	FileFormatJson FileFormat = "JSON"
)

// Values returns all known values for FileFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FileFormat) Values() []FileFormat {
	return []FileFormat{
		"CSV",
		"TSV",
		"CLF",
		"ELF",
		"XLSX",
		"JSON",
	}
}

type FilterOperator string

// Enum values for FilterOperator
const (
	FilterOperatorStringEquals FilterOperator = "StringEquals"
)

// Values returns all known values for FilterOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FilterOperator) Values() []FilterOperator {
	return []FilterOperator{
		"StringEquals",
	}
}

type FolderFilterAttribute string

// Enum values for FolderFilterAttribute
const (
	FolderFilterAttributeParentFolderArn FolderFilterAttribute = "PARENT_FOLDER_ARN"
)

// Values returns all known values for FolderFilterAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FolderFilterAttribute) Values() []FolderFilterAttribute {
	return []FolderFilterAttribute{
		"PARENT_FOLDER_ARN",
	}
}

type FolderType string

// Enum values for FolderType
const (
	FolderTypeShared FolderType = "SHARED"
)

// Values returns all known values for FolderType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FolderType) Values() []FolderType {
	return []FolderType{
		"SHARED",
	}
}

type GeoSpatialCountryCode string

// Enum values for GeoSpatialCountryCode
const (
	GeoSpatialCountryCodeUs GeoSpatialCountryCode = "US"
)

// Values returns all known values for GeoSpatialCountryCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GeoSpatialCountryCode) Values() []GeoSpatialCountryCode {
	return []GeoSpatialCountryCode{
		"US",
	}
}

type GeoSpatialDataRole string

// Enum values for GeoSpatialDataRole
const (
	GeoSpatialDataRoleCountry   GeoSpatialDataRole = "COUNTRY"
	GeoSpatialDataRoleState     GeoSpatialDataRole = "STATE"
	GeoSpatialDataRoleCounty    GeoSpatialDataRole = "COUNTY"
	GeoSpatialDataRoleCity      GeoSpatialDataRole = "CITY"
	GeoSpatialDataRolePostcode  GeoSpatialDataRole = "POSTCODE"
	GeoSpatialDataRoleLongitude GeoSpatialDataRole = "LONGITUDE"
	GeoSpatialDataRoleLatitude  GeoSpatialDataRole = "LATITUDE"
)

// Values returns all known values for GeoSpatialDataRole. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GeoSpatialDataRole) Values() []GeoSpatialDataRole {
	return []GeoSpatialDataRole{
		"COUNTRY",
		"STATE",
		"COUNTY",
		"CITY",
		"POSTCODE",
		"LONGITUDE",
		"LATITUDE",
	}
}

type IdentityStore string

// Enum values for IdentityStore
const (
	IdentityStoreQuicksight IdentityStore = "QUICKSIGHT"
)

// Values returns all known values for IdentityStore. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IdentityStore) Values() []IdentityStore {
	return []IdentityStore{
		"QUICKSIGHT",
	}
}

type IdentityType string

// Enum values for IdentityType
const (
	IdentityTypeIam        IdentityType = "IAM"
	IdentityTypeQuicksight IdentityType = "QUICKSIGHT"
)

// Values returns all known values for IdentityType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (IdentityType) Values() []IdentityType {
	return []IdentityType{
		"IAM",
		"QUICKSIGHT",
	}
}

type IngestionErrorType string

// Enum values for IngestionErrorType
const (
	IngestionErrorTypeFailureToAssumeRole             IngestionErrorType = "FAILURE_TO_ASSUME_ROLE"
	IngestionErrorTypeIngestionSuperseded             IngestionErrorType = "INGESTION_SUPERSEDED"
	IngestionErrorTypeIngestionCanceled               IngestionErrorType = "INGESTION_CANCELED"
	IngestionErrorTypeDataSetDeleted                  IngestionErrorType = "DATA_SET_DELETED"
	IngestionErrorTypeDataSetNotSpice                 IngestionErrorType = "DATA_SET_NOT_SPICE"
	IngestionErrorTypeS3UploadedFileDeleted           IngestionErrorType = "S3_UPLOADED_FILE_DELETED"
	IngestionErrorTypeS3ManifestError                 IngestionErrorType = "S3_MANIFEST_ERROR"
	IngestionErrorTypeDataToleranceException          IngestionErrorType = "DATA_TOLERANCE_EXCEPTION"
	IngestionErrorTypeSpiceTableNotFound              IngestionErrorType = "SPICE_TABLE_NOT_FOUND"
	IngestionErrorTypeDataSetSizeLimitExceeded        IngestionErrorType = "DATA_SET_SIZE_LIMIT_EXCEEDED"
	IngestionErrorTypeRowSizeLimitExceeded            IngestionErrorType = "ROW_SIZE_LIMIT_EXCEEDED"
	IngestionErrorTypeAccountCapacityLimitExceeded    IngestionErrorType = "ACCOUNT_CAPACITY_LIMIT_EXCEEDED"
	IngestionErrorTypeCustomerError                   IngestionErrorType = "CUSTOMER_ERROR"
	IngestionErrorTypeDataSourceNotFound              IngestionErrorType = "DATA_SOURCE_NOT_FOUND"
	IngestionErrorTypeIamRoleNotAvailable             IngestionErrorType = "IAM_ROLE_NOT_AVAILABLE"
	IngestionErrorTypeConnectionFailure               IngestionErrorType = "CONNECTION_FAILURE"
	IngestionErrorTypeSqlTableNotFound                IngestionErrorType = "SQL_TABLE_NOT_FOUND"
	IngestionErrorTypePermissionDenied                IngestionErrorType = "PERMISSION_DENIED"
	IngestionErrorTypeSslCertificateValidationFailure IngestionErrorType = "SSL_CERTIFICATE_VALIDATION_FAILURE"
	IngestionErrorTypeOauthTokenFailure               IngestionErrorType = "OAUTH_TOKEN_FAILURE"
	IngestionErrorTypeSourceApiLimitExceededFailure   IngestionErrorType = "SOURCE_API_LIMIT_EXCEEDED_FAILURE"
	IngestionErrorTypePasswordAuthenticationFailure   IngestionErrorType = "PASSWORD_AUTHENTICATION_FAILURE"
	IngestionErrorTypeSqlSchemaMismatchError          IngestionErrorType = "SQL_SCHEMA_MISMATCH_ERROR"
	IngestionErrorTypeInvalidDateFormat               IngestionErrorType = "INVALID_DATE_FORMAT"
	IngestionErrorTypeInvalidDataprepSyntax           IngestionErrorType = "INVALID_DATAPREP_SYNTAX"
	IngestionErrorTypeSourceResourceLimitExceeded     IngestionErrorType = "SOURCE_RESOURCE_LIMIT_EXCEEDED"
	IngestionErrorTypeSqlInvalidParameterValue        IngestionErrorType = "SQL_INVALID_PARAMETER_VALUE"
	IngestionErrorTypeQueryTimeout                    IngestionErrorType = "QUERY_TIMEOUT"
	IngestionErrorTypeSqlNumericOverflow              IngestionErrorType = "SQL_NUMERIC_OVERFLOW"
	IngestionErrorTypeUnresolvableHost                IngestionErrorType = "UNRESOLVABLE_HOST"
	IngestionErrorTypeUnroutableHost                  IngestionErrorType = "UNROUTABLE_HOST"
	IngestionErrorTypeSqlException                    IngestionErrorType = "SQL_EXCEPTION"
	IngestionErrorTypeS3FileInaccessible              IngestionErrorType = "S3_FILE_INACCESSIBLE"
	IngestionErrorTypeIotFileNotFound                 IngestionErrorType = "IOT_FILE_NOT_FOUND"
	IngestionErrorTypeIotDataSetFileEmpty             IngestionErrorType = "IOT_DATA_SET_FILE_EMPTY"
	IngestionErrorTypeInvalidDataSourceConfig         IngestionErrorType = "INVALID_DATA_SOURCE_CONFIG"
	IngestionErrorTypeDataSourceAuthFailed            IngestionErrorType = "DATA_SOURCE_AUTH_FAILED"
	IngestionErrorTypeDataSourceConnectionFailed      IngestionErrorType = "DATA_SOURCE_CONNECTION_FAILED"
	IngestionErrorTypeFailureToProcessJsonFile        IngestionErrorType = "FAILURE_TO_PROCESS_JSON_FILE"
	IngestionErrorTypeInternalServiceError            IngestionErrorType = "INTERNAL_SERVICE_ERROR"
	IngestionErrorTypeRefreshSuppressedByEdit         IngestionErrorType = "REFRESH_SUPPRESSED_BY_EDIT"
	IngestionErrorTypePermissionNotFound              IngestionErrorType = "PERMISSION_NOT_FOUND"
	IngestionErrorTypeElasticsearchCursorNotEnabled   IngestionErrorType = "ELASTICSEARCH_CURSOR_NOT_ENABLED"
	IngestionErrorTypeCursorNotEnabled                IngestionErrorType = "CURSOR_NOT_ENABLED"
)

// Values returns all known values for IngestionErrorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IngestionErrorType) Values() []IngestionErrorType {
	return []IngestionErrorType{
		"FAILURE_TO_ASSUME_ROLE",
		"INGESTION_SUPERSEDED",
		"INGESTION_CANCELED",
		"DATA_SET_DELETED",
		"DATA_SET_NOT_SPICE",
		"S3_UPLOADED_FILE_DELETED",
		"S3_MANIFEST_ERROR",
		"DATA_TOLERANCE_EXCEPTION",
		"SPICE_TABLE_NOT_FOUND",
		"DATA_SET_SIZE_LIMIT_EXCEEDED",
		"ROW_SIZE_LIMIT_EXCEEDED",
		"ACCOUNT_CAPACITY_LIMIT_EXCEEDED",
		"CUSTOMER_ERROR",
		"DATA_SOURCE_NOT_FOUND",
		"IAM_ROLE_NOT_AVAILABLE",
		"CONNECTION_FAILURE",
		"SQL_TABLE_NOT_FOUND",
		"PERMISSION_DENIED",
		"SSL_CERTIFICATE_VALIDATION_FAILURE",
		"OAUTH_TOKEN_FAILURE",
		"SOURCE_API_LIMIT_EXCEEDED_FAILURE",
		"PASSWORD_AUTHENTICATION_FAILURE",
		"SQL_SCHEMA_MISMATCH_ERROR",
		"INVALID_DATE_FORMAT",
		"INVALID_DATAPREP_SYNTAX",
		"SOURCE_RESOURCE_LIMIT_EXCEEDED",
		"SQL_INVALID_PARAMETER_VALUE",
		"QUERY_TIMEOUT",
		"SQL_NUMERIC_OVERFLOW",
		"UNRESOLVABLE_HOST",
		"UNROUTABLE_HOST",
		"SQL_EXCEPTION",
		"S3_FILE_INACCESSIBLE",
		"IOT_FILE_NOT_FOUND",
		"IOT_DATA_SET_FILE_EMPTY",
		"INVALID_DATA_SOURCE_CONFIG",
		"DATA_SOURCE_AUTH_FAILED",
		"DATA_SOURCE_CONNECTION_FAILED",
		"FAILURE_TO_PROCESS_JSON_FILE",
		"INTERNAL_SERVICE_ERROR",
		"REFRESH_SUPPRESSED_BY_EDIT",
		"PERMISSION_NOT_FOUND",
		"ELASTICSEARCH_CURSOR_NOT_ENABLED",
		"CURSOR_NOT_ENABLED",
	}
}

type IngestionRequestSource string

// Enum values for IngestionRequestSource
const (
	IngestionRequestSourceManual    IngestionRequestSource = "MANUAL"
	IngestionRequestSourceScheduled IngestionRequestSource = "SCHEDULED"
)

// Values returns all known values for IngestionRequestSource. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IngestionRequestSource) Values() []IngestionRequestSource {
	return []IngestionRequestSource{
		"MANUAL",
		"SCHEDULED",
	}
}

type IngestionRequestType string

// Enum values for IngestionRequestType
const (
	IngestionRequestTypeInitialIngestion   IngestionRequestType = "INITIAL_INGESTION"
	IngestionRequestTypeEdit               IngestionRequestType = "EDIT"
	IngestionRequestTypeIncrementalRefresh IngestionRequestType = "INCREMENTAL_REFRESH"
	IngestionRequestTypeFullRefresh        IngestionRequestType = "FULL_REFRESH"
)

// Values returns all known values for IngestionRequestType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IngestionRequestType) Values() []IngestionRequestType {
	return []IngestionRequestType{
		"INITIAL_INGESTION",
		"EDIT",
		"INCREMENTAL_REFRESH",
		"FULL_REFRESH",
	}
}

type IngestionStatus string

// Enum values for IngestionStatus
const (
	IngestionStatusInitialized IngestionStatus = "INITIALIZED"
	IngestionStatusQueued      IngestionStatus = "QUEUED"
	IngestionStatusRunning     IngestionStatus = "RUNNING"
	IngestionStatusFailed      IngestionStatus = "FAILED"
	IngestionStatusCompleted   IngestionStatus = "COMPLETED"
	IngestionStatusCancelled   IngestionStatus = "CANCELLED"
)

// Values returns all known values for IngestionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IngestionStatus) Values() []IngestionStatus {
	return []IngestionStatus{
		"INITIALIZED",
		"QUEUED",
		"RUNNING",
		"FAILED",
		"COMPLETED",
		"CANCELLED",
	}
}

type IngestionType string

// Enum values for IngestionType
const (
	IngestionTypeIncrementalRefresh IngestionType = "INCREMENTAL_REFRESH"
	IngestionTypeFullRefresh        IngestionType = "FULL_REFRESH"
)

// Values returns all known values for IngestionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IngestionType) Values() []IngestionType {
	return []IngestionType{
		"INCREMENTAL_REFRESH",
		"FULL_REFRESH",
	}
}

type InputColumnDataType string

// Enum values for InputColumnDataType
const (
	InputColumnDataTypeString   InputColumnDataType = "STRING"
	InputColumnDataTypeInteger  InputColumnDataType = "INTEGER"
	InputColumnDataTypeDecimal  InputColumnDataType = "DECIMAL"
	InputColumnDataTypeDatetime InputColumnDataType = "DATETIME"
	InputColumnDataTypeBit      InputColumnDataType = "BIT"
	InputColumnDataTypeBoolean  InputColumnDataType = "BOOLEAN"
	InputColumnDataTypeJson     InputColumnDataType = "JSON"
)

// Values returns all known values for InputColumnDataType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InputColumnDataType) Values() []InputColumnDataType {
	return []InputColumnDataType{
		"STRING",
		"INTEGER",
		"DECIMAL",
		"DATETIME",
		"BIT",
		"BOOLEAN",
		"JSON",
	}
}

type JoinType string

// Enum values for JoinType
const (
	JoinTypeInner JoinType = "INNER"
	JoinTypeOuter JoinType = "OUTER"
	JoinTypeLeft  JoinType = "LEFT"
	JoinTypeRight JoinType = "RIGHT"
)

// Values returns all known values for JoinType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JoinType) Values() []JoinType {
	return []JoinType{
		"INNER",
		"OUTER",
		"LEFT",
		"RIGHT",
	}
}

type MemberType string

// Enum values for MemberType
const (
	MemberTypeDashboard MemberType = "DASHBOARD"
	MemberTypeAnalysis  MemberType = "ANALYSIS"
	MemberTypeDataset   MemberType = "DATASET"
)

// Values returns all known values for MemberType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MemberType) Values() []MemberType {
	return []MemberType{
		"DASHBOARD",
		"ANALYSIS",
		"DATASET",
	}
}

type NamespaceErrorType string

// Enum values for NamespaceErrorType
const (
	NamespaceErrorTypePermissionDenied     NamespaceErrorType = "PERMISSION_DENIED"
	NamespaceErrorTypeInternalServiceError NamespaceErrorType = "INTERNAL_SERVICE_ERROR"
)

// Values returns all known values for NamespaceErrorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NamespaceErrorType) Values() []NamespaceErrorType {
	return []NamespaceErrorType{
		"PERMISSION_DENIED",
		"INTERNAL_SERVICE_ERROR",
	}
}

type NamespaceStatus string

// Enum values for NamespaceStatus
const (
	NamespaceStatusCreated             NamespaceStatus = "CREATED"
	NamespaceStatusCreating            NamespaceStatus = "CREATING"
	NamespaceStatusDeleting            NamespaceStatus = "DELETING"
	NamespaceStatusRetryableFailure    NamespaceStatus = "RETRYABLE_FAILURE"
	NamespaceStatusNonRetryableFailure NamespaceStatus = "NON_RETRYABLE_FAILURE"
)

// Values returns all known values for NamespaceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NamespaceStatus) Values() []NamespaceStatus {
	return []NamespaceStatus{
		"CREATED",
		"CREATING",
		"DELETING",
		"RETRYABLE_FAILURE",
		"NON_RETRYABLE_FAILURE",
	}
}

type ResourceStatus string

// Enum values for ResourceStatus
const (
	ResourceStatusCreationInProgress ResourceStatus = "CREATION_IN_PROGRESS"
	ResourceStatusCreationSuccessful ResourceStatus = "CREATION_SUCCESSFUL"
	ResourceStatusCreationFailed     ResourceStatus = "CREATION_FAILED"
	ResourceStatusUpdateInProgress   ResourceStatus = "UPDATE_IN_PROGRESS"
	ResourceStatusUpdateSuccessful   ResourceStatus = "UPDATE_SUCCESSFUL"
	ResourceStatusUpdateFailed       ResourceStatus = "UPDATE_FAILED"
	ResourceStatusDeleted            ResourceStatus = "DELETED"
)

// Values returns all known values for ResourceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceStatus) Values() []ResourceStatus {
	return []ResourceStatus{
		"CREATION_IN_PROGRESS",
		"CREATION_SUCCESSFUL",
		"CREATION_FAILED",
		"UPDATE_IN_PROGRESS",
		"UPDATE_SUCCESSFUL",
		"UPDATE_FAILED",
		"DELETED",
	}
}

type RowLevelPermissionFormatVersion string

// Enum values for RowLevelPermissionFormatVersion
const (
	RowLevelPermissionFormatVersionVersion1 RowLevelPermissionFormatVersion = "VERSION_1"
	RowLevelPermissionFormatVersionVersion2 RowLevelPermissionFormatVersion = "VERSION_2"
)

// Values returns all known values for RowLevelPermissionFormatVersion. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (RowLevelPermissionFormatVersion) Values() []RowLevelPermissionFormatVersion {
	return []RowLevelPermissionFormatVersion{
		"VERSION_1",
		"VERSION_2",
	}
}

type RowLevelPermissionPolicy string

// Enum values for RowLevelPermissionPolicy
const (
	RowLevelPermissionPolicyGrantAccess RowLevelPermissionPolicy = "GRANT_ACCESS"
	RowLevelPermissionPolicyDenyAccess  RowLevelPermissionPolicy = "DENY_ACCESS"
)

// Values returns all known values for RowLevelPermissionPolicy. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RowLevelPermissionPolicy) Values() []RowLevelPermissionPolicy {
	return []RowLevelPermissionPolicy{
		"GRANT_ACCESS",
		"DENY_ACCESS",
	}
}

type Status string

// Enum values for Status
const (
	StatusEnabled  Status = "ENABLED"
	StatusDisabled Status = "DISABLED"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"ENABLED",
		"DISABLED",
	}
}

type TemplateErrorType string

// Enum values for TemplateErrorType
const (
	TemplateErrorTypeSourceNotFound  TemplateErrorType = "SOURCE_NOT_FOUND"
	TemplateErrorTypeDataSetNotFound TemplateErrorType = "DATA_SET_NOT_FOUND"
	TemplateErrorTypeInternalFailure TemplateErrorType = "INTERNAL_FAILURE"
	TemplateErrorTypeAccessDenied    TemplateErrorType = "ACCESS_DENIED"
)

// Values returns all known values for TemplateErrorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TemplateErrorType) Values() []TemplateErrorType {
	return []TemplateErrorType{
		"SOURCE_NOT_FOUND",
		"DATA_SET_NOT_FOUND",
		"INTERNAL_FAILURE",
		"ACCESS_DENIED",
	}
}

type TextQualifier string

// Enum values for TextQualifier
const (
	TextQualifierDoubleQuote TextQualifier = "DOUBLE_QUOTE"
	TextQualifierSingleQuote TextQualifier = "SINGLE_QUOTE"
)

// Values returns all known values for TextQualifier. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TextQualifier) Values() []TextQualifier {
	return []TextQualifier{
		"DOUBLE_QUOTE",
		"SINGLE_QUOTE",
	}
}

type ThemeErrorType string

// Enum values for ThemeErrorType
const (
	ThemeErrorTypeInternalFailure ThemeErrorType = "INTERNAL_FAILURE"
)

// Values returns all known values for ThemeErrorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ThemeErrorType) Values() []ThemeErrorType {
	return []ThemeErrorType{
		"INTERNAL_FAILURE",
	}
}

type ThemeType string

// Enum values for ThemeType
const (
	ThemeTypeQuicksight ThemeType = "QUICKSIGHT"
	ThemeTypeCustom     ThemeType = "CUSTOM"
	ThemeTypeAll        ThemeType = "ALL"
)

// Values returns all known values for ThemeType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ThemeType) Values() []ThemeType {
	return []ThemeType{
		"QUICKSIGHT",
		"CUSTOM",
		"ALL",
	}
}

type UserRole string

// Enum values for UserRole
const (
	UserRoleAdmin            UserRole = "ADMIN"
	UserRoleAuthor           UserRole = "AUTHOR"
	UserRoleReader           UserRole = "READER"
	UserRoleRestrictedAuthor UserRole = "RESTRICTED_AUTHOR"
	UserRoleRestrictedReader UserRole = "RESTRICTED_READER"
)

// Values returns all known values for UserRole. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (UserRole) Values() []UserRole {
	return []UserRole{
		"ADMIN",
		"AUTHOR",
		"READER",
		"RESTRICTED_AUTHOR",
		"RESTRICTED_READER",
	}
}

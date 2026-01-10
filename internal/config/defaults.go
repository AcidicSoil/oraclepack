package config

const (
	EnvOutputVerify          = "ORACLEPACK_OUTPUT_VERIFY"
	EnvOutputRetries         = "ORACLEPACK_OUTPUT_RETRIES"
	EnvOutputRequireHeadings = "ORACLEPACK_OUTPUT_REQUIRE_HEADINGS"
	EnvOutputChunkMode       = "ORACLEPACK_OUTPUT_CHUNK_MODE"
)

const (
	DefaultOutputVerify          = false
	DefaultOutputRetries         = 0
	DefaultOutputRequireHeadings = false
	DefaultOutputChunkMode       = "auto"
)

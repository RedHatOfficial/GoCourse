type StorageConfiguration struct {
        Driver        string `mapstructure:"db_driver"       toml:"db_driver"`
        PGUsername    string `mapstructure:"pg_username"     toml:"pg_username"`
        PGPassword    string `mapstructure:"pg_password"     toml:"pg_password"`
        PGHost        string `mapstructure:"pg_host"         toml:"pg_host"`
        PGPort        int    `mapstructure:"pg_port"         toml:"pg_port"`
        PGDBName      string `mapstructure:"pg_db_name"      toml:"pg_db_name"`
        PGParams      string `mapstructure:"pg_params"       toml:"pg_params"`
        LogSQLQueries bool   `mapstructure:"log_sql_queries" toml:"log_sql_queries"`
}

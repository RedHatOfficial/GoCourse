type LoggingConfiguration struct {
        // Debug enables pretty colored logging
        Debug bool `mapstructure:"debug" toml:"debug"`
        LogLevel string `mapstructure:"log_level" toml:"log_level"`
}

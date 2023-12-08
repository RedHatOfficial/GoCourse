type ConfigStruct struct {
        Logging       LoggingConfiguration       `mapstructure:"logging" toml:"logging"`
        Storage       StorageConfiguration       `mapstructure:"storage" toml:"storage"`
        Kafka         KafkaConfiguration         `mapstructure:"kafka_broker" toml:"kafka_broker"`
        ServiceLog    ServiceLogConfiguration    `mapstructure:"service_log" toml:"service_log"`
        Dependencies  DependenciesConfiguration  `mapstructure:"dependencies" toml:"dependencies"`
        Notifications NotificationsConfiguration `mapstructure:"notifications" toml:"notifications"`
        Metrics       MetricsConfiguration       `mapstructure:"metrics" toml:"metrics"`
        Cleaner       CleanerConfiguration       `mapstructure:"cleaner" toml:"cleaner"`
        Processing    ProcessingConfiguration    `mapstructure:"processing" toml:"processing"`
}

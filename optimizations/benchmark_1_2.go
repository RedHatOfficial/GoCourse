// loadConfiguration function loads configuration prepared to be used by
// benchmarks
func loadConfiguration() (conf.ConfigStruct, error) {
        os.Clearenv()

        err := os.Setenv(configFileEnvVariableName, defaultConfigFileName)
        if err != nil {
                return conf.ConfigStruct{}, err
        }

        config, err := conf.LoadConfiguration(configFileEnvVariableName, defaultConfigFileName)
        if err != nil {
                return conf.ConfigStruct{}, err
        }

        return config, nil
}

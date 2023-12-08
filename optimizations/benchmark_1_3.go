func mustLoadBenchmarkConfiguration(b *testing.B) conf.ConfigStruct {
        configuration, err := loadConfiguration()
        if err != nil {
                b.Fatal(err)
        }
        return configuration
}

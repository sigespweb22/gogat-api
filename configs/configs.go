func GetConfig() *Config {
	APP := serviceConfig {
		APP_DEBUG: viper.GetBool("APP_DEBUG"),
		APP_PORT: viper.GetInt("APP_PORT"),
		USERS_ENDPOINT: viper.GetString("USERS_ENDPOINT"),
        TODO_ENDPOINT: viper.GetInt("TODO_ENDPOINT"),
        AUTH_ENDPOINT: viper.GetString("AUTH_ENDPOINT"),
	}

	return &Config {
		Service: APP
	}
}
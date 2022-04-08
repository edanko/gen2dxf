package cmd

import (
	"fmt"

	"github.com/muesli/coral"
	"github.com/spf13/viper"
)

func initParams(cmd *coral.Command, params any) error {
	v := viper.New()

	if err := v.BindPFlags(cmd.Flags()); err != nil {
		return fmt.Errorf("couldn't process command flags, %w", err)
	}

	/* configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("couldn't find config directory, %v", err)
	}
	v.AddConfigPath(configDir) */
	v.AddConfigPath(".")
	v.SetConfigName("gen2dxf")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("couldn't read config file, %w", err)
		}
	}

	if err := v.Unmarshal(params); err != nil {
		return fmt.Errorf("couldn't parse config, %w", err)
	}

	return nil
}

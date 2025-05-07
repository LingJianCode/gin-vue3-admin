package config

type Server struct {
	Zap        Zap        `mapstructure:"zap" json:"zap" yaml:"zap"`
	PostgreSQL PostgreSQL `mapstructure:"postgresql" json:"postgresql" yaml:"postgresql"`
}

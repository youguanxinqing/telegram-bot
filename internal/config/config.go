package config

type basic struct {
	BotToken    string `toml:"bot_token"`
	ChatId      int64  `toml:"chat_id"`
	UrlToken    string `toml:"url_token"`
	DbDSN       string `toml:"db_dsn"`
	StarDictDSN string `toml:"star_dict_dsn"`
}

var Basic basic

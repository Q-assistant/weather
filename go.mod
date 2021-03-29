module "weather"

go 1.16

require (
	github.com/q-assistant/sdk v1.0.0
)

replace github.com/q-assistant/sdk => ../../sdk
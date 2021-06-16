package golangmodule_v0

import logger "github.com/ElrondNetwork/elrond-go-logger"

func WriteSomethingV1(str string) {
	log := logger.GetOrCreate("core/indexer")

	log.Warn("Something", "str", str)
}

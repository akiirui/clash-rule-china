package main

import (
	"flag"
	"os"

	P "github.com/metacubex/mihomo/constant/provider"
	RP "github.com/metacubex/mihomo/rules/provider"
)

var (
	inputFile = flag.String("i","china.yaml", "input file path")
	outputFile = flag.String("o", "china.mrs", "output file path")
)

func ConvertMetaRuleSet(buf []byte, b string, f string, outputPath string) error {
	behavior, err := P.ParseBehavior(b)
	if err != nil {
		return err
	}
	format, err := P.ParseRuleFormat(f)
	if err != nil {
		return err
	}
	targetFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	err = RP.ConvertToMrs(buf, behavior, format, targetFile)
	closeErr := targetFile.Close()
	if err != nil {
		return err
	}
	return closeErr
}

func main() {
	flag.Parse()

	data, err := os.ReadFile(*inputFile);
	if err != nil {
		panic(err)
	}

	ConvertMetaRuleSet(data, "domain", "yaml", *outputFile)
}

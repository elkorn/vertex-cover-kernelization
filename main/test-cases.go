package main

var testCases map[string]func(string) = map[string]func(string){
	"MeasureBnb":                                      MeasureBnb,
	"MeasureBnbPreprocessing":                         MeasureBnbPreprocessing,
	"MeasureNaive":                                    MeasureNaive,
	"MeasureNaivePreprocessing":                       MeasureNaivePreprocessing,
	"MeasureVCCrownReduction":                         MeasureVCCrownReduction,
	"MeasureVCNetworkFlow":                            MeasureVCNetworkFlow,
	"MeasureVCCrownReductionPreprocessing":            MeasureVCCrownReductionPreprocessing,
	"MeasureVCNetworkFlowPreprocessing":               MeasureVCNetworkFlowPreprocessing,
	"MeasureKernelizationCrownReduction":              MeasureKernelizationCrownReduction,
	"MeasureKernelizationNetworkFlow":                 MeasureKernelizationNetworkFlow,
	"MeasureKernelizationCrownReductionPreprocessing": MeasureKernelizationCrownReductionPreprocessing,
	"MeasureKernelizationNetworkFlowPreprocessing":    MeasureKernelizationNetworkFlowPreprocessing,
	"MeasurePreprocessing":                            MeasurePreprocessing,
	"WriteGraphSizes":                                 WriteGraphSizes,
}

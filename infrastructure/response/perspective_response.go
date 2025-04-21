package response

type SummaryScore struct {
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

type ToxicityScore struct {
	SummaryScore SummaryScore `json:"summaryScore"`
}

type AttributeScores struct {
	Toxicity ToxicityScore `json:"TOXICITY"`
}

type PerspectiveResponse struct {
	AttributeScores AttributeScores `json:"attributeScores"`
	Languages       []string        `json:"languages"`
}

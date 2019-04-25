package err

const (
	Unknown = "unknown"
)

type M map[string]interface{}

type E struct {
	Code    string `json:"code"`
	Meta    M      `json:"meta,omitempty"`
	Reasons []E    `json:"reasons,omitempty"`
}

func New(code string, meta M, reasons ...E) E {
	return E{
		Code:    code,
		Meta:    meta,
		Reasons: reasons,
	}
}

func (e E) Error() string {
	return e.Code
}

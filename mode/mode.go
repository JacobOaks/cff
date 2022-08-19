package mode

const (
	// Unknown Generation Mode.
	Unknown GenerationMode = iota

	// Base generates CFF code without modification.
	Base

	// SourceMap generates CFF code with line directives to remap
	// generated code locations to source.
	SourceMap
)

// GenerationMode represents the mode of CFF generated code.
type GenerationMode uint8

// SourceMap returns whether SourceMap mode is enabled.
func (m GenerationMode) SourceMap() bool {
	return m == SourceMap
}

func (m GenerationMode) String() string {
	switch m {
	case SourceMap:
		return "source-map"
	case Base:
		return "base"
	default:
		return "unknown"
	}
}

// UnmarshalText unmarshals a GenerationMode.
func (m *GenerationMode) UnmarshalText(text []byte) {
	switch s := string(text); s {
	case "source-map":
		*m = SourceMap
	case "base":
		*m = Base
	default:
		*m = Unknown
	}
}
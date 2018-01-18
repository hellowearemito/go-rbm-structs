package send

// FileUploadRequest represents the file upload request structure.
type FileUploadRequest struct {
	FileURL string
}

// FileUploadResponse represents the file upload response structure.
type FileUploadResponse struct {
	Name string
}

// BaseMessage represents the base message structure.
type BaseMessage struct {
	Suggestions *Suggestion
}

// AddSuggestion assigns the given suggestion to the base message.
func (bm *BaseMessage) AddSuggestion(s Suggestion) {
	bm.Suggestions = &s
}

// TextMessage represents the text message structure.
type TextMessage struct {
	BaseMessage
	Text string
}

// FileMessage represents the file message structure.
type FileMessage struct {
	BaseMessage
	FileName string
}

// RichCardMessage represents the rich card message structure.
type RichCardMessage struct {
	CarouselCard   *CarouselCard
	StandaloneCard *StandaloneCard
}

// Width represents the width enum.
type Width string

const (
	// SmallWidth represents the small enum value.
	SmallWidth = Width("SMALL")
	// MediumWidth represents the medium enum value.
	MediumWidth = Width("MEDIUM")
)

// CarouselCard represents the carousel card structure.
type CarouselCard struct {
	CardWidth    Width
	CardContents []CardContent
}

// StandaloneCard represents the standalone card structure.
type StandaloneCard struct {
	CardOrientation         Orientation
	ThumbnailImageAlignment Alignment
	CardContent             CardContent
}

// Orientation represents the orientation enum.
type Orientation string

const (
	// Horizontal represents the horizontal enum value.
	Horizontal = Orientation("HORIZONTAL")
	// Vertical represents the vertical enum value.
	Vertical = Orientation("VERTICAL")
)

// Alignment represents the alignment enum.
type Alignment string

const (
	// Left represents the alignment value.
	Left = Alignment("LEFT")
	// Right represents the alignment value.
	Right = Alignment("RIGHT")
)

// CardContent represents the card content structure.
type CardContent struct {
	Title       string
	Description string
	Media       Media
	Suggestions Suggestion
}

// Height represents the height enum.
type Height string

const (
	// ShortHeight represents the short height enum value.
	ShortHeight = Height("SHORT")
	// MediumHeight represents the medium heigth enum value.
	MediumHeight = Height("MEDIUM")
	// TallHeight represents the tall height enum value.
	TallHeight = Height("TALL")
)

// Media represents the media sturcture.
type Media struct {
	FileName string
	Height   Height
}

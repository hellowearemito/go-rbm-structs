package send

// FileUploadRequest represents the file upload request structure.
type FileUploadRequest struct {
	FileURL string `json:"fileUrl"`
}

// FileUploadResponse represents the file upload response structure.
type FileUploadResponse struct {
	Name string `json:"name"`
}

// BaseMessage represents the base message structure.
type BaseMessage struct {
	Text        string           `json:"text,omitempty"`
	RichCard    *RichCardMessage `json:"richCard,omitempty"`
	Suggestions []Suggestion     `json:"suggestions,omitempty"`
}

// SuggestionMessage represents the suggestion message structure.
type SuggestionMessage struct {
	BaseMessage
}

// TextMessage represents the text message structure.
type TextMessage struct {
	BaseMessage
	Text string `json:"text"`
}

// FileMessage represents the file message structure.
type FileMessage struct {
	BaseMessage
	FileName string `json:"fileName"`
}

type RichCardMessage struct {
	BaseMessage
	Card interface{} `json:"richCard"`
}

// CarouselCardMessage is represents the carousel card message structure. The CarouselCardMessage is a kind of rich card.
type CarouselCardMessage struct {
	CarouselCard CarouselCard `json:"carouselCard"`
}

// StandaloneCardMessage is represents the standalone card message structure. The StandaloneCardMessage is a kind of rich card.
type StandaloneCardMessage struct {
	StandaloneCard StandaloneCard `json:"standaloneCard"`
}

const (
	// SmallWidth represents the small enum value.
	SmallWidth = "SMALL"
	// MediumWidth represents the medium enum value.
	MediumWidth = "MEDIUM"
)

// CarouselCard represents the carousel card structure.
type CarouselCard struct {
	CardWidth    string        `json:"cardWidth"`
	CardContents []CardContent `json:"cardContents"`
}

// StandaloneCard represents the standalone card structure.
type StandaloneCard struct {
	CardOrientation         string      `json:"cardOrientation"`
	ThumbnailImageAlignment string      `json:"thumbnailImageAlignment"`
	CardContent             CardContent `json:"cardContent"`
}

const (
	// Horizontal represents the horizontal enum value.
	Horizontal = "HORIZONTAL"
	// Vertical represents the vertical enum value.
	Vertical = "VERTICAL"
)

const (
	// Left represents the alignment value.
	Left = "LEFT"
	// Right represents the alignment value.
	Right = "RIGHT"
)

// CardContent represents the card content structure.
type CardContent struct {
	Title       *string      `json:"title"`
	Description *string      `json:"description"`
	Media       Media        `json:"media"`
	Suggestions []Suggestion `json:"suggestions"`
}

const (
	// ShortHeight represents the short height enum value.
	ShortHeight = "SHORT"
	// MediumHeight represents the medium heigth enum value.
	MediumHeight = "MEDIUM"
	// TallHeight represents the tall height enum value.
	TallHeight = "TALL"
)

// Media represents the media sturcture.
type Media struct {
	FileName string `json:"fileName"`
	Height   string `json:"height"`
}

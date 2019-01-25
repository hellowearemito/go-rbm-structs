package send

// FileUploadRequest represents the file upload request structure.
type FileUploadRequest struct {
	FileURL string `json:"fileUrl,omitempty"`
}

// FileUploadResponse represents the file upload response structure.
type FileUploadResponse struct {
	Name string `json:"name,omitempty"`
}

// BaseMessage represents the base message structure.
type BaseMessage struct {
	Suggestions []Suggestion `json:"suggestions,omitempty"`
}

// SuggestionMessage represents the suggestion message structure.
type SuggestionMessage struct {
	Text     string           `json:"text,omitempty"`
	FileName string           `json:"fileName,omitempty"`
	RichCard *RichCardMessage `json:"richCard,omitempty"`
	BaseMessage
}

// TextMessage represents the text message structure.
type TextMessage struct {
	BaseMessage
	Text string `json:"text,omitempty"`
}

// FileMessage represents the file message structure.
type FileMessage struct {
	BaseMessage
	FileName string `json:"fileName,omitempty"`
}

type RichCardMessage struct {
	BaseMessage
	Card interface{} `json:"richCard,omitempty"`
}

// CarouselCardMessage is represents the carousel card message structure. The CarouselCardMessage is a kind of rich card.
type CarouselCardMessage struct {
	CarouselCard CarouselCard `json:"carouselCard,omitempty"`
}

// StandaloneCardMessage is represents the standalone card message structure. The StandaloneCardMessage is a kind of rich card.
type StandaloneCardMessage struct {
	StandaloneCard StandaloneCard `json:"standaloneCard,omitempty"`
}

const (
	// SmallWidth represents the small enum value.
	SmallWidth = "SMALL"
	// MediumWidth represents the medium enum value.
	MediumWidth = "MEDIUM"
)

// CarouselCard represents the carousel card structure.
type CarouselCard struct {
	CardWidth    string        `json:"cardWidth,omitempty"`
	CardContents []CardContent `json:"cardContents,omitempty"`
}

// StandaloneCard represents the standalone card structure.
type StandaloneCard struct {
	CardOrientation         string      `json:"cardOrientation,omitempty"`
	ThumbnailImageAlignment string      `json:"thumbnailImageAlignment,omitempty"`
	CardContent             CardContent `json:"cardContent,omitempty"`
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
	Title       *string      `json:"title,omitempty"`
	Description *string      `json:"description,omitempty"`
	Media       *Media       `json:"media,omitempty"`
	Suggestions []Suggestion `json:"suggestions,omitempty"`
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
	FileName    string       `json:"fileName,omitempty"`
	Height      string       `json:"height"`
	ContentInfo *ContentInfo `json:"contentInfo,omitempty"`
}

// ContentInfo containing the content information.
type ContentInfo struct {
	FileURL      string `json:"fileUrl"`
	ThumbnailURL string `json:"thumbnailUrl,omitempty"`
	ForceRefresh bool   `json:"forceRefresh"`
}

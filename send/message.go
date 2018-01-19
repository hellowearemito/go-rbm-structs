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
	Suggestions *Suggestion `json:"suggestions"`
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

// CarouselCardMessage is represents the carousel card message structure. The CarouselCardMessage is a kind of rich card.
type CarouselCardMessage struct {
	BaseMessage
	CarouselCard CarouselCard `json:"carouselCard"`
}

// StandaloneCardMessage is represents the standalone card message structure. The StandaloneCardMessage is a kind of rich card.
type StandaloneCardMessage struct {
	BaseMessage
	StandaloneCard StandaloneCard `json:"standaloneCard"`
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
	CardWidth    Width         `json:"cardWidth"`
	CardContents []CardContent `json:"cardContents"`
}

// StandaloneCard represents the standalone card structure.
type StandaloneCard struct {
	CardOrientation         Orientation `json:"cardOrientation"`
	ThumbnailImageAlignment Alignment   `json:"thumbnailImageAlignment"`
	CardContent             CardContent `json:"cardContent"`
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
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Media       Media      `json:"media"`
	Suggestions Suggestion `json:"suggestions"`
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
	FileName string `json:"fileName"`
	Height   Height `json:"height"`
}

package typeform

import (
	"net/http"
	"time"
)

// Client is the api client
type Client struct {
	httpClient *http.Client
	Config     struct {
		APIKey string
	}
	APIVersion APIVersion
}

//

// APIError type
type APIError struct {
	Error       string `json:"error"`
	Field       string `json:"field"`
	Description string `json:"description"`
}

//

type APIVersion string

const (
	Latest APIVersion = "latest"
	V0_4   APIVersion = "v0.4"
)

//

type BaseInfo struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Version       string    `json:"version"`
	Documentation string    `json:"documentation"`
	Support       string    `json:"support"`
	Time          Timestamp `json:"time"`
}

type Timestamp struct {
	time.Time
}

const TimestampFormat = "2006-01-02 15:04:05 +0000 UTC"

func (ct *Timestamp) UnmarshalJSON(b []byte) error {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	var err error
	ct.Time, err = time.Parse(TimestampFormat, string(b))
	return err
}

func (ct *Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(TimestampFormat)), nil
}

//

type Form struct {
	Title            string      `json:"title"`                        // required, The title of the typeform
	Fields           []Field     `json:"fields"`                       // required, An array of Field objects
	Tags             []string    `json:"tags,omitempty"`               // An array of Tags
	DesignID         string      `json:"design_id,omitempty"`          // The ID of the Design object you want to use
	WebhookSubmitURL string      `json:"webhook_submit_url,omitempty"` // Where you want the responses to go to when a respondent submits the typeform
	URLIDs           []string    `json:"url_ids,omitempty"`            // The IDs of the URLs you want your form to be displayed at
	Branding         bool        `json:"branding,omitempty"`           // Enables or disables the Typeform branding for the form
	LogicJumps       []LogicJump `json:"logic_jumps,omitempty"`
}

type Field struct {
	Type        FieldType `json:"type"`                  // Required, A string describing the type of the field
	Question    string    `json:"question"`              // Required, The main question text for the field
	Description string    `json:"description,omitempty"` // The description (or sub-text) that appears below the main question text (in a smaller font size).
	Required    bool      `json:"required,omitempty"`    // Decides if the field is mandatory
	Tags        []string  `json:"tags,omitempty"`        // An array of tags as strings
	Ref         string    `json:"ref,omitempty"`         // A unique reference for the field
	// Attachment // An image or video that is attached to the field (Coming soon)

	// sort_text, long_text
	MaxCharacters int `json:"max_characters,omitempty"` // The maximum number of characters the respondent can type as an answer.

	// multiple_choice
	Choices                 []Choice `json:"choices,omitempty"`                   // required, Array of choice objects with the choices that the respondent can select.
	AllowMultipleSelections bool     `json:"allow_multiple_selections,omitempty"` // Boolean to decide if the respondent can choose one or multiple choices
	Randomize               bool     `json:"randomize,omitempty"`                 // If the choice order should be randomized on every load
	VerticalAlignment       bool     `json:"vertical_alignment,omitempty"`        // If the choices should appear as one choice per row, instead of fitting as many choices as possible per row
	AddOtherChoice          bool     `json:"add_other_choice,omitempty"`          // If the field should automatically include a choice with the text "Other" which transforms into a open ended text field

	// picture_choice
	ShowLabels bool `json:"show_labels,omitempty"` // If the labels should be visible beneath the choices or not.
	// Choices
	Supersize bool `json:"supersize,omitempty"` // If the pictures choice should be set to large. (Coming soon)
	// AllowMultipleSelections
	// Randomize
	// AddOtherChoice

	// statement
	ButtonText string `json:"button_text,omitempty"` // Sets the text of the button that jumps to the next field
	HideMarks  bool   `json:"hide_marks,omitempty"`  // Boolean to indicate if the field should not be surrounded by quotation marks

	// dropdown
	// Choices
	AlphabeticalOrder bool `json:"alphabetical_order,omitempty"` // If the choices should be sorted in alphabetic order

	// number
	MinValue int `json:"min_value,omitempty"` // The minimum value your respondent can answer
	MaxValue int `json:"max_value,omitempty"` // The maximum value your respondent can answer

	// rating
	Steps int    `json:"steps,omitempty"` // The number of steps the user can chose. Is limited to a value between 1 - 10
	Shape string `json:"shape,omitempty"` // The icon to use for the steps. Use the list in Typeform.com to get the icon you want. "Stars" in Typeform.com would be used as "stars".

	// opinion_scale
	// Steps // required, The number of steps in the scale. Maximum is 11 and minimum is 5
	Labels     *Labels `json:"labels,omitempty"`       // An object, defining left, center and right labels
	StartAtOne bool    `json:"start_at_one,omitempty"` // If the scale should start at zero or one
}

//

type Choice struct {
	ImageID string `json:"image_id,omitempty"`
	Label   string `json:"label,omitempty"`
}

type Labels struct {
	Left   string `json:"left,omitempty"`
	Center string `json:"center,omitempty"`
	Right  string `json:"right,omitempty"`
}

//

type FieldType string

const (
	ShortText      FieldType = "short_text"
	LongText       FieldType = "long_text"
	MultipleChoice FieldType = "multiple_choice"
	PictureChoice  FieldType = "picture_choice"
	Statement      FieldType = "statement"
	Dropdown       FieldType = "dropdown"
	YesNo          FieldType = "yes_no"
	Number         FieldType = "number"
	Rating         FieldType = "rating"
	OpinionScale   FieldType = "opinion_scale"
	Email          FieldType = "email"
	Website        FieldType = "website"
	Legal          FieldType = "legal"
)

//

type FormInfo struct {
	Links   []Link     `json:"_links"`
	Fields  []Field    `json:"fields"`
	ID      string     `json:"id"`
	Title   string     `json:"title"`
	URLs    []URL      `json:"urls"`
	Version APIVersion `json:"version"`
}

type Link struct {
	HREF string `json:"href"`
	REL  string `json:"rel"`
}

type URL struct {
	FormID  string     `json:"form_id"`
	ID      string     `json:"id"`
	Version APIVersion `json:"version"`
}

///

type NewImage struct {
	ID          string `json:"id"`
	OriginalURL string `json:"original_url"`
	Type        string `json:"type"`
	Version     string `json:"version"`
}

//

type ImageInfo struct {
	Filename string `json:"filename"`
	Height   int    `json:"height"`
	ID       string `json:"id"`
	Type     string `json:"type"`
	URL      string `json:"url"`
	Version  string `json:"version"`
	Width    int    `json:"width"`
}

//

type Design struct {
	Colors Colors `json:"colors"`
	Font   string `json:"font"`
}

type Colors struct {
	Question   string `json:"question"`
	Button     string `json:"button"`
	Answer     string `json:"answer"`
	Background string `json:"background"`
}

type DesignInfo struct {
	ID     string `json:"id"`
	Colors struct {
		Question   string `json:"question"`
		Button     string `json:"button"`
		Answer     string `json:"answer"`
		Background string `json:"background"`
	} `json:"colors"`
	Font    string `json:"font"`
	Version string `json:"version"`
}

//

type URLInfo struct {
	ID      string `json:"id"`
	FormID  string `json:"form_id"`
	Version string `json:"version"`
	Links   []Link `json:"_links"`
}

type LogicJump struct {
	From string `json:"from"`
	To   string `json:"to"`
	If   bool   `json:"if"`
}

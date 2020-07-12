// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Item undocumented
type Item struct {
	// Entity is the base model of Item
	Entity
	// Number undocumented
	Number *string `json:"number,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// ItemCategoryID undocumented
	ItemCategoryID *UUID `json:"itemCategoryId,omitempty"`
	// ItemCategoryCode undocumented
	ItemCategoryCode *string `json:"itemCategoryCode,omitempty"`
	// Blocked undocumented
	Blocked *bool `json:"blocked,omitempty"`
	// BaseUnitOfMeasureID undocumented
	BaseUnitOfMeasureID *UUID `json:"baseUnitOfMeasureId,omitempty"`
	// Gtin undocumented
	Gtin *string `json:"gtin,omitempty"`
	// Inventory undocumented
	Inventory *int `json:"inventory,omitempty"`
	// UnitPrice undocumented
	UnitPrice *int `json:"unitPrice,omitempty"`
	// PriceIncludesTax undocumented
	PriceIncludesTax *bool `json:"priceIncludesTax,omitempty"`
	// UnitCost undocumented
	UnitCost *int `json:"unitCost,omitempty"`
	// TaxGroupID undocumented
	TaxGroupID *UUID `json:"taxGroupId,omitempty"`
	// TaxGroupCode undocumented
	TaxGroupCode *string `json:"taxGroupCode,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Picture undocumented
	Picture []Picture `json:"picture,omitempty"`
	// ItemCategory undocumented
	ItemCategory *ItemCategory `json:"itemCategory,omitempty"`
}

// ItemActionSet undocumented
type ItemActionSet struct {
	// Object is the base model of ItemActionSet
	Object
	// Comment undocumented
	Comment *CommentAction `json:"comment,omitempty"`
	// Create undocumented
	Create *CreateAction `json:"create,omitempty"`
	// Delete undocumented
	Delete *DeleteAction `json:"delete,omitempty"`
	// Edit undocumented
	Edit *EditAction `json:"edit,omitempty"`
	// Mention undocumented
	Mention *MentionAction `json:"mention,omitempty"`
	// Move undocumented
	Move *MoveAction `json:"move,omitempty"`
	// Rename undocumented
	Rename *RenameAction `json:"rename,omitempty"`
	// Restore undocumented
	Restore *RestoreAction `json:"restore,omitempty"`
	// Share undocumented
	Share *ShareAction `json:"share,omitempty"`
	// Version undocumented
	Version *VersionAction `json:"version,omitempty"`
}

// ItemActionStat undocumented
type ItemActionStat struct {
	// Object is the base model of ItemActionStat
	Object
	// ActionCount undocumented
	ActionCount *int `json:"actionCount,omitempty"`
	// ActorCount undocumented
	ActorCount *int `json:"actorCount,omitempty"`
}

// ItemActivity undocumented
type ItemActivity struct {
	// Entity is the base model of ItemActivity
	Entity
	// Access undocumented
	Access *AccessAction `json:"access,omitempty"`
	// ActivityDateTime undocumented
	ActivityDateTime *time.Time `json:"activityDateTime,omitempty"`
	// Actor undocumented
	Actor *IdentitySet `json:"actor,omitempty"`
	// DriveItem undocumented
	DriveItem *DriveItem `json:"driveItem,omitempty"`
}

// ItemActivityOLD undocumented
type ItemActivityOLD struct {
	// Entity is the base model of ItemActivityOLD
	Entity
	// Action undocumented
	Action *ItemActionSet `json:"action,omitempty"`
	// Actor undocumented
	Actor *IdentitySet `json:"actor,omitempty"`
	// Times undocumented
	Times *ItemActivityTimeSet `json:"times,omitempty"`
	// DriveItem undocumented
	DriveItem *DriveItem `json:"driveItem,omitempty"`
	// ListItem undocumented
	ListItem *ListItem `json:"listItem,omitempty"`
}

// ItemActivityStat undocumented
type ItemActivityStat struct {
	// Entity is the base model of ItemActivityStat
	Entity
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Access undocumented
	Access *ItemActionStat `json:"access,omitempty"`
	// Create undocumented
	Create *ItemActionStat `json:"create,omitempty"`
	// Delete undocumented
	Delete *ItemActionStat `json:"delete,omitempty"`
	// Edit undocumented
	Edit *ItemActionStat `json:"edit,omitempty"`
	// Move undocumented
	Move *ItemActionStat `json:"move,omitempty"`
	// IsTrending undocumented
	IsTrending *bool `json:"isTrending,omitempty"`
	// IncompleteData undocumented
	IncompleteData *IncompleteData `json:"incompleteData,omitempty"`
	// Activities undocumented
	Activities []ItemActivity `json:"activities,omitempty"`
}

// ItemActivityTimeSet undocumented
type ItemActivityTimeSet struct {
	// Object is the base model of ItemActivityTimeSet
	Object
	// LastRecordedDateTime undocumented
	LastRecordedDateTime *time.Time `json:"lastRecordedDateTime,omitempty"`
	// ObservedDateTime undocumented
	ObservedDateTime *time.Time `json:"observedDateTime,omitempty"`
	// RecordedDateTime undocumented
	RecordedDateTime *time.Time `json:"recordedDateTime,omitempty"`
}

// ItemAnalytics undocumented
type ItemAnalytics struct {
	// Entity is the base model of ItemAnalytics
	Entity
	// ItemActivityStats undocumented
	ItemActivityStats []ItemActivityStat `json:"itemActivityStats,omitempty"`
	// AllTime undocumented
	AllTime *ItemActivityStat `json:"allTime,omitempty"`
	// LastSevenDays undocumented
	LastSevenDays *ItemActivityStat `json:"lastSevenDays,omitempty"`
}

// ItemAttachment undocumented
type ItemAttachment struct {
	// Attachment is the base model of ItemAttachment
	Attachment
	// Item undocumented
	Item *OutlookItem `json:"item,omitempty"`
}

// ItemBody undocumented
type ItemBody struct {
	// Object is the base model of ItemBody
	Object
	// ContentType undocumented
	ContentType *BodyType `json:"contentType,omitempty"`
	// Content undocumented
	Content *string `json:"content,omitempty"`
}

// ItemCategory undocumented
type ItemCategory struct {
	// Entity is the base model of ItemCategory
	Entity
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

// ItemEmail undocumented
type ItemEmail struct {
	// ItemFacet is the base model of ItemEmail
	ItemFacet
	// Address undocumented
	Address *string `json:"address,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Type undocumented
	Type *EmailType `json:"type,omitempty"`
}

// ItemFacet undocumented
type ItemFacet struct {
	// Entity is the base model of ItemFacet
	Entity
	// AllowedAudiences undocumented
	AllowedAudiences *AllowedAudiences `json:"allowedAudiences,omitempty"`
	// Inference undocumented
	Inference *InferenceData `json:"inference,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
}

// ItemPhone undocumented
type ItemPhone struct {
	// ItemFacet is the base model of ItemPhone
	ItemFacet
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Type undocumented
	Type *PhoneType `json:"type,omitempty"`
	// Number undocumented
	Number *string `json:"number,omitempty"`
}

// ItemPreviewInfo undocumented
type ItemPreviewInfo struct {
	// Object is the base model of ItemPreviewInfo
	Object
	// GetURL undocumented
	GetURL *string `json:"getUrl,omitempty"`
	// PostParameters undocumented
	PostParameters *string `json:"postParameters,omitempty"`
	// PostURL undocumented
	PostURL *string `json:"postUrl,omitempty"`
}

// ItemReference undocumented
type ItemReference struct {
	// Object is the base model of ItemReference
	Object
	// DriveID undocumented
	DriveID *string `json:"driveId,omitempty"`
	// DriveType undocumented
	DriveType *string `json:"driveType,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Path undocumented
	Path *string `json:"path,omitempty"`
	// ShareID undocumented
	ShareID *string `json:"shareId,omitempty"`
	// SharepointIDs undocumented
	SharepointIDs *SharepointIDs `json:"sharepointIds,omitempty"`
	// SiteID undocumented
	SiteID *string `json:"siteId,omitempty"`
}
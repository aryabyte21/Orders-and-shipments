package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Business struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Name              string             `bson:"name"`
	Email             string             `bson:"email"`
	Phone             string             `bson:"phone"`
	Onboarded         bool               `bson:"onboarded"`
	Fqdns             []string           `bson:"fqdns"`
	TrackingPageUrl   string             `bson:"trackingPageUrl"`
	ReplaceTrackingLinks bool             `bson:"replaceTrackingLinks"`
	CurrentBillingId  primitive.ObjectID `bson:"currentBillingId" json:"-"`
	IsFreeTrialUsed   bool               `bson:"isFreeTrialUsed"`
	AuthProvider      string             `bson:"authProvider"`
	CountryCode       string             `bson:"countryCode"`
	Currency          string             `bson:"currency"`
	Category          string             `bson:"category"`
	Verified          bool               `bson:"verified"`
	EmailOptIn        bool               `bson:"emailOptIn"`
	SmsOptIn          bool               `bson:"smsOptIn"`
	ShopifyAppScopes  string             `bson:"shopifyAppScopes"`
	ShopifyAppAuthToken string            `bson:"shopifyAppAuthToken"`
	ShopifyShopDomain string             `bson:"shopifyShopDomain"`
	DropshippingMode  bool               `bson:"dropshippingMode"`
	CreatedAt         time.Time          `bson:"createdAt"`
}

package models

import "time"

// Product Описание игры.
// К сожалению не было времени подробно изучать и описовать полученный ответ.
type Product struct {
	AgeLimit      int    `json:"age_limit"`
	Bucket        string `json:"bucket"`
	ContainerType string `json:"container_type"`
	ContentOrigin int    `json:"content_origin"`
	ContentRating struct {
		Description  string `json:"description"`
		RatingSystem string `json:"rating_system"`
		URL          string `json:"url"`
	} `json:"content_rating"`
	ContentType string `json:"content_type"`
	DefaultSku  struct {
		AmortizeFlag           bool   `json:"amortizeFlag"`
		AugmentedDescription   string `json:"augmented_description"`
		BundleExclusiveFlag    bool   `json:"bundleExclusiveFlag"`
		ChargeImmediatelyFlag  bool   `json:"chargeImmediatelyFlag"`
		ChargeTypeID           int    `json:"charge_type_id"`
		CreditCardRequiredFlag int    `json:"credit_card_required_flag"`
		DefaultSku             bool   `json:"defaultSku"`
		DisplayPrice           string `json:"display_price"`
		Eligibilities          []struct {
			ID              string        `json:"id"`
			Operand         string        `json:"operand"`
			Operator        string        `json:"operator"`
			RightOperand    interface{}   `json:"rightOperand"`
			Name            string        `json:"name"`
			Description     interface{}   `json:"description"`
			EntitlementType interface{}   `json:"entitlement_type"`
			Drms            []interface{} `json:"drms"`
		} `json:"eligibilities"`
		EndDate      time.Time `json:"end_date"`
		Entitlements []struct {
			Description            interface{}   `json:"description"`
			Drms                   []interface{} `json:"drms"`
			Duration               int           `json:"duration"`
			DurationOverrideTypeID interface{}   `json:"durationOverrideTypeId"`
			ExpAfterFirstUse       int           `json:"exp_after_first_use"`
			FeatureTypeID          int           `json:"feature_type_id"`
			ID                     string        `json:"id"`
			LicenseType            int           `json:"license_type"`
			Metadata               interface{}   `json:"metadata"`
			Name                   string        `json:"name"`
			PackageType            string        `json:"packageType"`
			Packages               []struct {
				PlatformID   int    `json:"platformId"`
				PlatformName string `json:"platformName"`
				Size         int    `json:"size"`
			} `json:"packages"`
			PreorderPlaceholderFlag bool        `json:"preorder_placeholder_flag"`
			Size                    int         `json:"size"`
			SubType                 int         `json:"subType"`
			Subtitle1UageCodes      interface{} `json:"subtitle_1uage_codes"`
			Type                    int         `json:"type"`
			UseCount                int         `json:"use_count"`
			VoiceLanguageCodes      interface{} `json:"voice_language_codes"`
		} `json:"entitlements"`
		ID                          string        `json:"id"`
		IsOriginal                  bool          `json:"is_original"`
		Name                        string        `json:"name"`
		Platforms                   []int         `json:"platforms"`
		PlayabilityDate             time.Time     `json:"playability_date"`
		Price                       int           `json:"price"`
		Rewards                     []interface{} `json:"rewards"`
		SeasonPassExclusiveFlag     bool          `json:"seasonPassExclusiveFlag"`
		SkuAvailabilityOverrideFlag bool          `json:"skuAvailabilityOverrideFlag"`
		SkuType                     int           `json:"sku_type"`
		Type                        string        `json:"type"`
		WalletChargeDate            time.Time     `json:"wallet_charge_date"`
	} `json:"default_sku"`
	DobRequired          bool `json:"dob_required"`
	GameContentTypesList []struct {
		Name string `json:"name"`
		Key  string `json:"key"`
	} `json:"gameContentTypesList"`
	GameContentType string `json:"game_contentType"`
	ID              string `json:"id"`
	Images          []struct {
		Type int    `json:"type"`
		URL  string `json:"url"`
	} `json:"images"`
	Links     []interface{} `json:"links"`
	LongDesc  string        `json:"long_desc"`
	MediaList struct {
		Screenshots []struct {
			Type   string `json:"type"`
			TypeID int    `json:"typeId"`
			Source string `json:"source"`
			URL    string `json:"url"`
			Order  int    `json:"order"`
		} `json:"screenshots"`
	} `json:"mediaList"`
	MediaLayouts []struct {
		Type   string `json:"type"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"media_layouts"`
	Metadata struct {
		CnRemotePlay struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"cn_remotePlay"`
		CnVrEnabled struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"cn_vrEnabled"`
		CnPlaystationMove struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"cn_playstationMove"`
		SecondaryClassification struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"secondary_classification"`
		CnVrRequired struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"cn_vrRequired"`
		GameGenre struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"game_genre"`
		CnPsEnhanced struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"cn_psEnhanced"`
		PlayablePlatform struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"playable_platform"`
		CnNumberOfPlayers struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"cn_numberOfPlayers"`
		CnDualshockVibration struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"cn_dualshockVibration"`
		TertiaryClassification struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"tertiary_classification"`
		ContainerType struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"container_type"`
		Genre struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"genre"`
		CnInGamePurchases struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"cn_inGamePurchases"`
		CnOfflinePlayMode struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"cn_offlinePlayMode"`
		CnPsVrAimRequired struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"cn_psVrAimRequired"`
		CnPlaystationCamera struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"cn_playstationCamera"`
		CnSingstarMicrophone struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"cn_singstarMicrophone"`
		CnCrossPlatformPSVita struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"cn_crossPlatformPSVita"`
		CnPsVrAimEnabled struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"cn_psVrAimEnabled"`
		PrimaryClassification struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"primary_classification"`
	} `json:"metadata"`
	Name             string        `json:"name"`
	PageTypeID       int           `json:"pageTypeId"`
	PlayablePlatform []string      `json:"playable_platform"`
	Promomedia       []interface{} `json:"promomedia"`
	ProviderName     string        `json:"provider_name"`
	ReleaseDate      time.Time     `json:"release_date"`
	Restricted       bool          `json:"restricted"`
	Revision         int           `json:"revision"`
	ShortName        string        `json:"short_name"`
	Size             int           `json:"size"`
	SkuLinks         []interface{} `json:"sku_links"`
	Skus             []struct {
		AmortizeFlag           bool   `json:"amortizeFlag"`
		AugmentedDescription   string `json:"augmented_description"`
		BundleExclusiveFlag    bool   `json:"bundleExclusiveFlag"`
		ChargeImmediatelyFlag  bool   `json:"chargeImmediatelyFlag"`
		ChargeTypeID           int    `json:"charge_type_id"`
		CreditCardRequiredFlag int    `json:"credit_card_required_flag"`
		DefaultSku             bool   `json:"defaultSku"`
		DisplayPrice           string `json:"display_price"`
		Eligibilities          []struct {
			ID              string        `json:"id"`
			Operand         string        `json:"operand"`
			Operator        string        `json:"operator"`
			RightOperand    interface{}   `json:"rightOperand"`
			Name            string        `json:"name"`
			Description     interface{}   `json:"description"`
			EntitlementType interface{}   `json:"entitlement_type"`
			Drms            []interface{} `json:"drms"`
		} `json:"eligibilities"`
		EndDate      time.Time `json:"end_date"`
		Entitlements []struct {
			Description            interface{}   `json:"description"`
			Drms                   []interface{} `json:"drms"`
			Duration               int           `json:"duration"`
			DurationOverrideTypeID interface{}   `json:"durationOverrideTypeId"`
			ExpAfterFirstUse       int           `json:"exp_after_first_use"`
			FeatureTypeID          int           `json:"feature_type_id"`
			ID                     string        `json:"id"`
			LicenseType            int           `json:"license_type"`
			Metadata               interface{}   `json:"metadata"`
			Name                   string        `json:"name"`
			PackageType            string        `json:"packageType"`
			Packages               []struct {
				PlatformID   int    `json:"platformId"`
				PlatformName string `json:"platformName"`
				Size         int    `json:"size"`
			} `json:"packages"`
			PreorderPlaceholderFlag bool        `json:"preorder_placeholder_flag"`
			Size                    int         `json:"size"`
			SubType                 int         `json:"subType"`
			SubtitleLanguageCodes   interface{} `json:"subtitle_language_codes"`
			Type                    int         `json:"type"`
			UseCount                int         `json:"use_count"`
			VoiceLanguageCodes      interface{} `json:"voice_language_codes"`
		} `json:"entitlements"`
		ID                          string        `json:"id"`
		IsOriginal                  bool          `json:"is_original"`
		Name                        string        `json:"name"`
		Platforms                   []int         `json:"platforms"`
		PlayabilityDate             time.Time     `json:"playability_date"`
		Price                       int           `json:"price"`
		Rewards                     []interface{} `json:"rewards"`
		SeasonPassExclusiveFlag     bool          `json:"seasonPassExclusiveFlag"`
		SkuAvailabilityOverrideFlag bool          `json:"skuAvailabilityOverrideFlag"`
		SkuType                     int           `json:"sku_type"`
		Type                        string        `json:"type"`
		WalletChargeDate            time.Time     `json:"wallet_charge_date"`
	} `json:"skus"`
	StarRating struct {
		Total string `json:"total"`
		Score string `json:"score"`
		Count []struct {
			Star  int `json:"star"`
			Count int `json:"count"`
		} `json:"count"`
	} `json:"star_rating"`
	Start        int    `json:"start"`
	Timestamp    int64  `json:"timestamp"`
	TitleName    string `json:"title_name"`
	TopCategory  string `json:"top_category"`
	TotalResults int    `json:"total_results"`
}

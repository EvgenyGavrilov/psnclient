package models

import "time"

// ListGames Структура содержит список игр playstation store, служебныю иформацию.
// К сожалению не было времени подробно изучать и описовать полученный ответ от playstation store.
// TODO: В списке попадаются игры, у которых при запросе на полное описание, в ответ получаем 204 http статус.
//  Есть подазрения, что такая игра как-то связана с другой. Скорей всего это одна и та же игра, но под разные платформы.
type ListGames struct {
	AgeLimit   int `json:"age_limit"`
	Attributes struct {
		Facets struct {
			GameContentType []struct {
				Name  string `json:"name"`
				Count int    `json:"count"`
				Key   string `json:"key"`
			} `json:"game_content_type"`
			SubtitleLang []struct {
				Name  string `json:"name"`
				Count int    `json:"count"`
				Key   string `json:"key"`
			} `json:"subtitle_lang"`
			ReleaseDate []struct {
				Name  string `json:"name"`
				Count int    `json:"count"`
				Key   string `json:"key"`
			} `json:"release_date"`
			GameDemo []struct {
				Name  string `json:"name"`
				Count int    `json:"count"`
				Key   string `json:"key"`
			} `json:"game_demo"`
			Price []struct {
				Name  string `json:"name"`
				Count int    `json:"count"`
				Key   string `json:"key"`
			} `json:"price"`
			Genre []struct {
				Name  string `json:"name"`
				Count int    `json:"count"`
				Key   string `json:"key"`
			} `json:"genre"`
			TopCategory []struct {
				Name  string `json:"name"`
				Count int    `json:"count"`
				Key   string `json:"key"`
			} `json:"top_category"`
			VoiceLang []struct {
				Name  string `json:"name"`
				Count int    `json:"count"`
				Key   string `json:"key"`
			} `json:"voice_lang"`
			GameType []struct {
				Name  string `json:"name"`
				Count int    `json:"count"`
				Key   string `json:"key"`
			} `json:"game_type"`
			Relationship []struct {
				Name  string `json:"name"`
				Count int    `json:"count"`
				Key   string `json:"key"`
			} `json:"relationship"`
			Platform []struct {
				Name  string `json:"name"`
				Count int    `json:"count"`
				Key   string `json:"key"`
			} `json:"platform"`
		} `json:"facets"`
		Next []interface{} `json:"next"`
	} `json:"attributes"`
	ContainerType string          `json:"container_type"`
	ContentOrigin int             `json:"content_origin"`
	DobRequired   bool            `json:"dob_required"`
	ID            string          `json:"id"`
	Images        []interface{}   `json:"images"`
	Links         []ListGamesLink `json:"links"`
	LongDesc      string          `json:"long_desc"`
	Metadata      struct {
	} `json:"metadata"`
	Name        string        `json:"name"`
	Promomedia  []interface{} `json:"promomedia"`
	Restricted  bool          `json:"restricted"`
	Revision    int           `json:"revision"`
	SceneLayout struct {
		ID             int           `json:"id"`
		CatalogEntryID string        `json:"catalogEntryId"`
		StoreFrontID   int           `json:"storeFrontId"`
		TemplateID     int           `json:"templateId"`
		SubScenes      []interface{} `json:"subScenes"`
		StoreTypeID    int           `json:"storeTypeId"`
	} `json:"scene_layout"`
	Size                 int           `json:"size"`
	SkuLinks             []interface{} `json:"sku_links"`
	SortDefault          string        `json:"sort_default"`
	SortDefaultDirection string        `json:"sort_default_direction"`
	Start                int           `json:"start"`
	TemplateDef          struct {
		Name        string      `json:"name"`
		ID          int         `json:"id"`
		StoreTypeID int         `json:"storeTypeId"`
		ImageURL    interface{} `json:"imageUrl"`
		Locations   []struct {
			ID      int         `json:"id"`
			Name    interface{} `json:"name"`
			Width   string      `json:"width"`
			Height  string      `json:"height"`
			OffsetX string      `json:"offsetX"`
			OffsetY string      `json:"offsetY"`
			Widgets interface{} `json:"widgets"`
		} `json:"locations"`
		Extras []struct {
			TemplateExtraID int         `json:"templateExtraId"`
			Width           int         `json:"width"`
			Height          int         `json:"height"`
			Name            interface{} `json:"name"`
			Key             string      `json:"key"`
		} `json:"extras"`
	} `json:"template_def"`
	Timestamp    int64 `json:"timestamp"`
	TotalResults int   `json:"total_results"`
}

// ListGamesLink Содержит краткий набор данных об игре.
type ListGamesLink struct {
	Bucket        string `json:"bucket"`
	ContainerType string `json:"container_type"`
	ContentType   string `json:"content_type"`
	DefaultSku    struct {
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
	} `json:"default_sku,omitempty"`
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
	Name             string    `json:"name"`
	PlayablePlatform []string  `json:"playable_platform"`
	ProviderName     string    `json:"provider_name"`
	ReleaseDate      time.Time `json:"release_date"`
	Restricted       bool      `json:"restricted"`
	Revision         int       `json:"revision"`
	ShortName        string    `json:"short_name"`
	Timestamp        int64     `json:"timestamp"`
	TopCategory      string    `json:"top_category"`
	URL              string    `json:"url"`
}

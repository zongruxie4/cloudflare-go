// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// NetflowTopService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetflowTopService] method instead.
type NetflowTopService struct {
	Options []option.RequestOption
}

// NewNetflowTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetflowTopService(opts ...option.RequestOption) (r *NetflowTopService) {
	r = &NetflowTopService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems by network traffic (NetFlows).
func (r *NetflowTopService) Ases(ctx context.Context, query NetflowTopAsesParams, opts ...option.RequestOption) (res *NetflowTopAsesResponse, err error) {
	var env NetflowTopAsesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/netflows/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the top locations by network traffic (NetFlows).
func (r *NetflowTopService) Locations(ctx context.Context, query NetflowTopLocationsParams, opts ...option.RequestOption) (res *NetflowTopLocationsResponse, err error) {
	var env NetflowTopLocationsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/netflows/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NetflowTopAsesResponse struct {
	Top0 []NetflowTopAsesResponseTop0 `json:"top_0,required"`
	JSON netflowTopAsesResponseJSON   `json:"-"`
}

// netflowTopAsesResponseJSON contains the JSON metadata for the struct
// [NetflowTopAsesResponse]
type netflowTopAsesResponseJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopAsesResponseJSON) RawJSON() string {
	return r.raw
}

type NetflowTopAsesResponseTop0 struct {
	ClientASN    float64                        `json:"clientASN,required"`
	ClientAsName string                         `json:"clientASName,required"`
	Value        string                         `json:"value,required"`
	JSON         netflowTopAsesResponseTop0JSON `json:"-"`
}

// netflowTopAsesResponseTop0JSON contains the JSON metadata for the struct
// [NetflowTopAsesResponseTop0]
type netflowTopAsesResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NetflowTopAsesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopAsesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type NetflowTopLocationsResponse struct {
	Top0 []NetflowTopLocationsResponseTop0 `json:"top_0,required"`
	JSON netflowTopLocationsResponseJSON   `json:"-"`
}

// netflowTopLocationsResponseJSON contains the JSON metadata for the struct
// [NetflowTopLocationsResponse]
type netflowTopLocationsResponseJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type NetflowTopLocationsResponseTop0 struct {
	ClientCountryAlpha2 string                              `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                              `json:"clientCountryName,required"`
	Value               string                              `json:"value,required"`
	JSON                netflowTopLocationsResponseTop0JSON `json:"-"`
}

// netflowTopLocationsResponseTop0JSON contains the JSON metadata for the struct
// [NetflowTopLocationsResponseTop0]
type netflowTopLocationsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *NetflowTopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopLocationsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type NetflowTopAsesParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[NetflowTopAsesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [NetflowTopAsesParams]'s query parameters as `url.Values`.
func (r NetflowTopAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type NetflowTopAsesParamsFormat string

const (
	NetflowTopAsesParamsFormatJson NetflowTopAsesParamsFormat = "JSON"
	NetflowTopAsesParamsFormatCsv  NetflowTopAsesParamsFormat = "CSV"
)

func (r NetflowTopAsesParamsFormat) IsKnown() bool {
	switch r {
	case NetflowTopAsesParamsFormatJson, NetflowTopAsesParamsFormatCsv:
		return true
	}
	return false
}

type NetflowTopAsesResponseEnvelope struct {
	Result  NetflowTopAsesResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    netflowTopAsesResponseEnvelopeJSON `json:"-"`
}

// netflowTopAsesResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetflowTopAsesResponseEnvelope]
type netflowTopAsesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopAsesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopAsesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetflowTopLocationsParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[NetflowTopLocationsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [NetflowTopLocationsParams]'s query parameters as
// `url.Values`.
func (r NetflowTopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type NetflowTopLocationsParamsFormat string

const (
	NetflowTopLocationsParamsFormatJson NetflowTopLocationsParamsFormat = "JSON"
	NetflowTopLocationsParamsFormatCsv  NetflowTopLocationsParamsFormat = "CSV"
)

func (r NetflowTopLocationsParamsFormat) IsKnown() bool {
	switch r {
	case NetflowTopLocationsParamsFormatJson, NetflowTopLocationsParamsFormatCsv:
		return true
	}
	return false
}

type NetflowTopLocationsResponseEnvelope struct {
	Result  NetflowTopLocationsResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    netflowTopLocationsResponseEnvelopeJSON `json:"-"`
}

// netflowTopLocationsResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetflowTopLocationsResponseEnvelope]
type netflowTopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

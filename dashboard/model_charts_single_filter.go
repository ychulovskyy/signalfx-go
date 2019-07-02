/*
 * Dashboards API
 *
 * API for creating, retrieving, updating, and deleting dashboards and dashboard groups. <br> Dashboards are groups of charts. In a dashboard, all the charts that belong to the dashboard appear at the same time and follow the same filtering options. ## Dashboard layout The system lays out dashboards and the charts they contain with these dimensions: <br>   * The web UI reserves a 12x100 grid for each dashboard and assigns one or     more charts to specific locations within the grid.   * A chart associated with the dashboard can be any size from 1x1 to 12x3.   * If you assign overlapping dashboard locations for charts, the system     attempts to resize or reorganize the layout. This ensures that all     of the charts fit within the space alloted to the dashboard.  ## Dashboard access By default, all users in an organization can edit and delete dashboards and dashboard groups. If SignalFx has enabled the \"write permissions\" feature for your organization, you can limit editing or deleting of specific dashboards to specific individuals or teams, or both. Use this feature to prevent unauthorized or accidental modifications to dashboards and the charts they contain. ## Cloning dashboards Users who don't have permission to edit a dashboard can still clone it and modify the clone. ## View dashboards You can view dashboards you create using the API in the web UI by specifying their \"id\" property in a web UI URL, by following this syntax: <br> <code>https://app.signalfx.com/#/dashboard/&lt;DASHBOARD_ID&gt;</code> <br> Dashboards you create using the API also appear by name in the web UI catalog and in their dashboard group.
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dashboard

// A single filter object to apply to the charts in the dashboard. The filter specifies a default or user-defined dimension or custom property. You can either include or exclude all the data that matches the dimension or custom property.
type ChartsSingleFilter struct {
	// Flag that indicates how the filter should operate. If `true`, data that matches the criteria is _excluded_ from charts; otherwise, data that matches the criteria is included.
	NOT bool `json:"NOT,omitempty"`
	// Name of the dimension or custom property to match to the data.<br> **Note:** If the dimension or custom property doesn't exist in any of the charts for the dashboard, and `ChartsFilterObject.NOT` is `true`, the system doesn't display any data in the charts.
	Property string `json:"property"`
	// A list of values to compare to the value of the dimension or custom property specified in `ChartsFilterObject.property`. If the list contains more than one value, the filter becomes a set of queries between the value of `property` and each element of `value`. The system joins these queries with an implicit OR.
	Value         string `json:"value"`
	ApplyIfExists bool   `json:"applyIfExists,omitempty"`
}

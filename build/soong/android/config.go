package android

// Global config used by Cesium soong additions
var CesiumConfig = struct {
	// List of packages that are permitted
	// for java source overlays.
	JavaSourceOverlayModuleWhitelist []string
}{
	// JavaSourceOverlayModuleWhitelist
	[]string{
		"org.cesium.hardware",
	},
}

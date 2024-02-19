package override

var filesMap = map[string]string{
	"org.bluez.Adapter.rst":                     "adapter-api.txt",
	"org.bluez.AdminPolicySet.rst":              "admin-policy-api.txt",
	"org.bluez.AdminPolicyStatus.rst":           "admin-policy-api.txt",
	"org.bluez.AdvertisementMonitor.rst":        "advertisement-monitor-api.txt",
	"org.bluez.AdvertisementMonitorManager.rst": "advertisement-monitor-api.txt",
	"org.bluez.LEAdvertisement.rst":             "advertising-api.txt",
	"org.bluez.LEAdvertisingManager.rst":        "advertising-api.txt",
	"org.bluez.AgentManager.rst":                "agent-api.txt",
	"org.bluez.Agent.rst":                       "agent-api.txt",
	"org.bluez.Battery.rst":                     "battery-api.txt",
	"org.bluez.BatteryProviderManager.rst":      "battery-api.txt",
	"org.bluez.BatteryProvider.rst":             "battery-api.txt",
	"org.bluez.Device.rst":                      "device-api.txt",
	"org.bluez.GattService.rst":                 "gatt-api.txt",
	"org.bluez.GattCharacteristic.rst":          "gatt-api.txt",
	"org.bluez.GattDescriptor.rst":              "gatt-api.txt",
	"org.bluez.GattProfile.rst":                 "gatt-api.txt",
	"org.bluez.GattManager.rst":                 "gatt-api.txt",
	"org.bluez.HealthManager.rst":               "health-api.txt",
	"org.bluez.HealthDevice.rst":                "health-api.txt",
	"org.bluez.HealthChannel.rst":               "health-api.txt",
	"org.bluez.Input.rst":                       "input-api.txt",
	"org.bluez.Media.rst":                       "media-api.txt",
	"org.bluez.MediaControl.rst":                "media-api.txt",
	"org.bluez.MediaPlayer.rst":                 "media-api.txt",
	"org.bluez.MediaFolder.rst":                 "media-api.txt",
	"org.bluez.MediaItem.rst":                   "media-api.txt",
	"org.bluez.MediaEndpoint.rst":               "media-api.txt",
	"org.bluez.MediaTransport.rst":              "media-api.txt",
	"org.bluez.mesh.Network.rst":                "mesh-api.txt",
	"org.bluez.mesh.Node.rst":                   "mesh-api.txt",
	"org.bluez.mesh.Management.rst":             "mesh-api.txt",
	"org.bluez.mesh.Application.rst":            "mesh-api.txt",
	"org.bluez.mesh.Element.rst":                "mesh-api.txt",
	"org.bluez.mesh.Attention.rst":              "mesh-api.txt",
	"org.bluez.mesh.Provisioner.rst":            "mesh-api.txt",
	"org.bluez.mesh.ProvisionAgent.rst":         "mesh-api.txt",
	"org.bluez.Network.rst":                     "network-api.txt",
	"org.bluez.NetworkServer.rst":               "network-api.txt",
	"org.bluez.obex.AgentManager.rst":           "obex-agent-api.txt",
	"org.bluez.obex.Agent.rst":                  "obex-agent-api.txt",
	"org.bluez.obex.Client.rst":                 "obex-api.txt",
	"org.bluez.obex.Session.rst":                "obex-api.txt",
	"org.bluez.obex.Transfer.rst":               "obex-api.txt",
	"org.bluez.obex.ObjectPush.rst":             "obex-api.txt",
	"org.bluez.obex.FileTransfer.rst":           "obex-api.txt",
	"org.bluez.obex.PhonebookAccess.rst":        "obex-api.txt",
	"org.bluez.obex.Synchronization.rst":        "obex-api.txt",
	"org.bluez.obex.MessageAccess.rst":          "obex-api.txt",
	"org.bluez.obex.Message.rst":                "obex-api.txt",
	"org.bluez.ProfileManager.rst":              "profile-api.txt",
	"org.bluez.Profile.rst":                     "profile-api.txt",
	"org.bluez.SimAccess.rst":                   "sap-api.txt",
	"org.bluez.ThermometerManager.rst":          "thermometer-api.txt",
	"org.bluez.Thermometer.rst":                 "thermometer-api.txt",
	"org.bluez.ThermometerWatcher.rst":          "thermometer-api.txt",
}

func MapFile(rawfile string) (string, bool) {
	res, ok := filesMap[rawfile]
	return res, ok
}
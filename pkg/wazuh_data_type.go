package pkg

// Simple version to text JSON strings for Wazuh to ingest, might need to customize these later

import "encoding/json"

func (d *JumpCloudSystemEvent) convertToWazuhString() string {
	d.JumpCloudEventType = "system"
	b, _ := json.Marshal(d)
	return string(b)
}

func (d *JumpCloudLDAPEvent) convertToWazuhString() string {
	d.JumpCloudEventType = "ldap"
	b, _ := json.Marshal(d)
	return string(b)
}

func (d *JumpCloudDirectoryEvent) convertToWazuhString() string {
	d.JumpCloudEventType = "directory"
	b, _ := json.Marshal(d)
	return string(b)
}

func (d *JumpCloudRadiusEvent) convertToWazuhString() string {
	d.JumpCloudEventType = "radius"
	b, _ := json.Marshal(d)
	return string(b)
}

func (d *JumpCloudSSOEvent) convertToWazuhString() string {
	d.JumpCloudEventType = "sso"
	b, _ := json.Marshal(d)
	return string(b)
}

func (d *JumpCloudAdminEvent) convertToWazuhString() string {
	d.JumpCloudEventType = "admin"
	b, _ := json.Marshal(d)
	return string(b)
}

// Novos eventos adicionados
func (d *JumpCloudMDMEvent) convertToWazuhString() string {
    d.JumpCloudEventType = "mdm"
    b, _ := json.Marshal(d)
    return string(b)
}

func (d *JumpCloudNotificationChannelEvent) convertToWazuhString() string {
    d.JumpCloudEventType = "notification_channel"
    b, _ := json.Marshal(d)
    return string(b)
}

func (d *JumpCloudObjectStorageEvent) convertToWazuhString() string {
    d.JumpCloudEventType = "object_storage"
    b, _ := json.Marshal(d)
    return string(b)
}

func (d *JumpCloudAlertEvent) convertToWazuhString() string {
    d.JumpCloudEventType = "alert"
    b, _ := json.Marshal(d)
    return string(b)
}

func (d *JumpCloudSaaSManagementEvent) convertToWazuhString() string {
    d.JumpCloudEventType = "saas_management"
    b, _ := json.Marshal(d)
    return string(b)
}

func (d *JumpCloudAccessManagementEvent) convertToWazuhString() string {
    d.JumpCloudEventType = "access_management"
    b, _ := json.Marshal(d)
    return string(b)
}

func (d *JumpCloudReportsEvent) convertToWazuhString() string {
    d.JumpCloudEventType = "reports"
    b, _ := json.Marshal(d)
    return string(b)
}

func (d *JumpCloudAssetManagementEvent) convertToWazuhString() string {
    d.JumpCloudEventType = "asset_management"
    b, _ := json.Marshal(d)
    return string(b)
}


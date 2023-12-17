package models

type Network_connection struct {
	ID             string `json:"id,omitempty"`
	TypeConnection string `json:"type_connection,omitempty"`
}

type Network struct {
	ID        string             `json:"id,omitempty"`
	IpAddress string             `json:"ip_address,omitempty"`
	NetConn   Network_connection `json:"net_conn,omitempty"`
}

type ResumeNetworks struct {
	IdCampus       string `json:"id_campus,omitempty"`
	CampusName     string `json:"campus_name,omitempty"`
	NumberFloor    int    `json:"number_floor,omitempty"`
	IpAddress      string `json:"ip_address,omitempty"`
	ZoneName       string `json:"zone_name,omitempty"`
	DependendyName string `json:"dependendy_name,omitempty"`
	TypeConnection string `json:"type_connection,omitempty"`
}

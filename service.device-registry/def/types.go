// Code generated by jrpc. DO NOT EDIT.

package deviceregistrydef

// DeviceHeader is defined in the .def file
type DeviceHeader struct {
	Id             string                 `json:"id"`
	Name           string                 `json:"name"`
	Type           string                 `json:"type"`
	Kind           string                 `json:"kind"`
	Attributes     map[string]interface{} `json:"attributes"`
	RoomId         string                 `json:"room_id"`
	Room           *Room                  `json:"room"`
	ControllerName string                 `json:"controller_name"`
}

// Validate returns an error if any of the fields have bad values
func (m *DeviceHeader) Validate() error {

	if err := m.Room.Validate(); err != nil {
		return err
	}

	return nil
}

// Room is defined in the .def file
type Room struct {
	Id      string          `json:"id"`
	Name    string          `json:"name"`
	Devices []*DeviceHeader `json:"devices"`
}

// Validate returns an error if any of the fields have bad values
func (m *Room) Validate() error {

	for _, r := range m.Devices {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// GetDeviceRequest is defined in the .def file
type GetDeviceRequest struct {
	DeviceId string `json:"device_id"`
}

// Validate returns an error if any of the fields have bad values
func (m *GetDeviceRequest) Validate() error {

	return nil
}

// GetDeviceResponse is defined in the .def file
type GetDeviceResponse struct {
}

// Validate returns an error if any of the fields have bad values
func (m *GetDeviceResponse) Validate() error {
	return nil
}

// ListDevicesRequest is defined in the .def file
type ListDevicesRequest struct {
	ControllerName string `json:"controller_name"`
}

// Validate returns an error if any of the fields have bad values
func (m *ListDevicesRequest) Validate() error {

	return nil
}

// ListDevicesResponse is defined in the .def file
type ListDevicesResponse struct {
}

// Validate returns an error if any of the fields have bad values
func (m *ListDevicesResponse) Validate() error {
	return nil
}

// GetRoomRequest is defined in the .def file
type GetRoomRequest struct {
}

// Validate returns an error if any of the fields have bad values
func (m *GetRoomRequest) Validate() error {
	return nil
}

// GetRoomResponse is defined in the .def file
type GetRoomResponse struct {
}

// Validate returns an error if any of the fields have bad values
func (m *GetRoomResponse) Validate() error {
	return nil
}

// ListRoomsRequest is defined in the .def file
type ListRoomsRequest struct {
}

// Validate returns an error if any of the fields have bad values
func (m *ListRoomsRequest) Validate() error {
	return nil
}

// ListRoomsResponse is defined in the .def file
type ListRoomsResponse struct {
}

// Validate returns an error if any of the fields have bad values
func (m *ListRoomsResponse) Validate() error {
	return nil
}
